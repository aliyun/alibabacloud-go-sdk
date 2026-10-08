// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetName(v string) *GetStoryRequest
	GetDatasetName() *string
	SetObjectId(v string) *GetStoryRequest
	GetObjectId() *string
	SetProjectName(v string) *GetStoryRequest
	GetProjectName() *string
}

type GetStoryRequest struct {
	// The name of the dataset. For more information about how to obtain the dataset name, see [Create a dataset](https://help.aliyun.com/document_detail/478160.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The ID of the story object whose information you want to retrieve.
	//
	// This parameter is required.
	//
	// example:
	//
	// id1
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The name of the project. For more information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s GetStoryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetStoryRequest) GoString() string {
	return s.String()
}

func (s *GetStoryRequest) GetDatasetName() *string {
	return s.DatasetName
}

func (s *GetStoryRequest) GetObjectId() *string {
	return s.ObjectId
}

func (s *GetStoryRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GetStoryRequest) SetDatasetName(v string) *GetStoryRequest {
	s.DatasetName = &v
	return s
}

func (s *GetStoryRequest) SetObjectId(v string) *GetStoryRequest {
	s.ObjectId = &v
	return s
}

func (s *GetStoryRequest) SetProjectName(v string) *GetStoryRequest {
	s.ProjectName = &v
	return s
}

func (s *GetStoryRequest) Validate() error {
	return dara.Validate(s)
}
