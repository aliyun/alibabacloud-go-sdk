// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveStoryFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *RemoveStoryFilesRequest
	GetDatasetName() *string
	SetFiles(v []*RemoveStoryFilesRequestFiles) *RemoveStoryFilesRequest
	GetFiles() []*RemoveStoryFilesRequestFiles
	SetObjectId(v string) *RemoveStoryFilesRequest
	GetObjectId() *string
	SetProjectName(v string) *RemoveStoryFilesRequest
	GetProjectName() *string
}

type RemoveStoryFilesRequest struct {
	// The name of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// testdataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The files that you want to delete.
	//
	// This parameter is required.
	Files []*RemoveStoryFilesRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The ID of the story.
	//
	// This parameter is required.
	//
	// example:
	//
	// testid
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the project.
	//
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s RemoveStoryFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveStoryFilesRequest) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *RemoveStoryFilesRequest) GetFiles() []*RemoveStoryFilesRequestFiles {
	return s.Files
}

func (s *RemoveStoryFilesRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *RemoveStoryFilesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *RemoveStoryFilesRequest) SetDatasetName(v string) *RemoveStoryFilesRequest {
	s.DatasetName = &v
	return s
}

func (s *RemoveStoryFilesRequest) SetFiles(v []*RemoveStoryFilesRequestFiles) *RemoveStoryFilesRequest {
	s.Files = v
	return s
}

func (s *RemoveStoryFilesRequest) SetObjectId(v string) *RemoveStoryFilesRequest {
	s.ObjectId = &v
	return s
}

func (s *RemoveStoryFilesRequest) SetProjectName(v string) *RemoveStoryFilesRequest {
	s.ProjectName = &v
	return s
}

func (s *RemoveStoryFilesRequest) Validate() error {
	return dara.Validate(s)
}

type RemoveStoryFilesRequestFiles struct {
	// The URI of the Object Storage Service (OSS) bucket where you store the files that you want to delete.
	//
	// Specify the value in the oss://${Bucket}/${Object} format. `${Bucket}` specifies the name of the OSS bucket that resides in the same region as the current project. `${Object}` specifies the complete path to the files that have an extension.
	//
	// example:
	//
	// oss://bucket1/object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s RemoveStoryFilesRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s RemoveStoryFilesRequestFiles) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesRequestFiles) GetURI() *string {
	return s.URI
}

func (s *RemoveStoryFilesRequestFiles) SetURI(v string) *RemoveStoryFilesRequestFiles {
	s.URI = &v
	return s
}

func (s *RemoveStoryFilesRequestFiles) Validate() error {
	return dara.Validate(s)
}
