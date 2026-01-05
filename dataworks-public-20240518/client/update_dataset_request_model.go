// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *UpdateDatasetRequest
	GetComment() *string
	SetId(v string) *UpdateDatasetRequest
	GetId() *string
	SetName(v string) *UpdateDatasetRequest
	GetName() *string
	SetReadme(v string) *UpdateDatasetRequest
	GetReadme() *string
}

type UpdateDatasetRequest struct {
	// The dataset description. Length not exceeding 1024.
	//
	// example:
	//
	// new comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The dataset ID. Only DataWorks datasets are supported for update.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataworks-dataset:3pXXXb8o0ngr07njhps1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The dataset name. A non-empty string, length not exceeding 128.
	//
	// example:
	//
	// test_oss_dataset_new
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The user guide, supports Markdown formatted rich text.
	//
	// example:
	//
	// ## introduction
	Readme *string `json:"Readme,omitempty" xml:"Readme,omitempty"`
}

func (s UpdateDatasetRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateDatasetRequest) GetId() *string {
	return s.Id
}

func (s *UpdateDatasetRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDatasetRequest) GetReadme() *string {
	return s.Readme
}

func (s *UpdateDatasetRequest) SetComment(v string) *UpdateDatasetRequest {
	s.Comment = &v
	return s
}

func (s *UpdateDatasetRequest) SetId(v string) *UpdateDatasetRequest {
	s.Id = &v
	return s
}

func (s *UpdateDatasetRequest) SetName(v string) *UpdateDatasetRequest {
	s.Name = &v
	return s
}

func (s *UpdateDatasetRequest) SetReadme(v string) *UpdateDatasetRequest {
	s.Readme = &v
	return s
}

func (s *UpdateDatasetRequest) Validate() error {
	return dara.Validate(s)
}
