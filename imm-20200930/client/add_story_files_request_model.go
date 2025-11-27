// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStoryFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *AddStoryFilesRequest
	GetDatasetName() *string
	SetFiles(v []*AddStoryFilesRequestFiles) *AddStoryFilesRequest
	GetFiles() []*AddStoryFilesRequestFiles
	SetObjectId(v string) *AddStoryFilesRequest
	GetObjectId() *string
	SetProjectName(v string) *AddStoryFilesRequest
	GetProjectName() *string
}

type AddStoryFilesRequest struct {
	// The name of the dataset.
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The objects that you want to add.
	//
	// This parameter is required.
	Files []*AddStoryFilesRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
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
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s AddStoryFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s AddStoryFilesRequest) GoString() string {
	return s.String()
}

func (s *AddStoryFilesRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *AddStoryFilesRequest) GetFiles() []*AddStoryFilesRequestFiles {
	return s.Files
}

func (s *AddStoryFilesRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *AddStoryFilesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *AddStoryFilesRequest) SetDatasetName(v string) *AddStoryFilesRequest {
	s.DatasetName = &v
	return s
}

func (s *AddStoryFilesRequest) SetFiles(v []*AddStoryFilesRequestFiles) *AddStoryFilesRequest {
	s.Files = v
	return s
}

func (s *AddStoryFilesRequest) SetObjectId(v string) *AddStoryFilesRequest {
	s.ObjectId = &v
	return s
}

func (s *AddStoryFilesRequest) SetProjectName(v string) *AddStoryFilesRequest {
	s.ProjectName = &v
	return s
}

func (s *AddStoryFilesRequest) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddStoryFilesRequestFiles struct {
	// The URI of the object.
	//
	// Specify the OSS URI in the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// example:
	//
	// oss://test-bucket/test-object
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
}

func (s AddStoryFilesRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s AddStoryFilesRequestFiles) GoString() string {
	return s.String()
}

func (s *AddStoryFilesRequestFiles) GetURI() *string {
	return s.URI
}

func (s *AddStoryFilesRequestFiles) SetURI(v string) *AddStoryFilesRequestFiles {
	s.URI = &v
	return s
}

func (s *AddStoryFilesRequestFiles) Validate() error {
	return dara.Validate(s)
}
