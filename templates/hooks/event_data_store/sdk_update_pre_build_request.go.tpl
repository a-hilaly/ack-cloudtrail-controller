	if delta.DifferentAt("Spec.Tags") {
		err = rm.syncTags(ctx, latest, desired)
		if err != nil {
			return nil, err
		}
	} else if !delta.DifferentExcept("Spec.Tags") {
        return desired, nil
    }