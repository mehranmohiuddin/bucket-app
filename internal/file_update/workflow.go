package file_update

type Workflow interface {
	Process() error
}

type WorkflowImpl struct {
}

func (self *WorkflowImpl) Process() error {
	// retrieve file from local or bucket

	// update file

	// re-upload file to local or bucket

	return nil
}
