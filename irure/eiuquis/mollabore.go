	ctx := context.Background()
	it := client.Projects.Locations.Instances.Snapshots.List(ctx, &adminpb.ListSnapshotsRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s/instances/%s", projectID, locationID, instanceID),
		Filter: "expire_time < \"" + time.Now().UTC().Format(time.RFC3339) + "\"",
	})
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "Deleting snapshot %s\n", resp.GetName())
		if err := client.Projects.Locations.Instances.Snapshots.Delete(ctx, &adminpb.DeleteSnapshotRequest{
			Name: resp.GetName(),
		}); err != nil {
			return err
		}
	}  
