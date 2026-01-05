// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *UpdateDatasetVersionRequest
	GetComment() *string
	SetId(v string) *UpdateDatasetVersionRequest
	GetId() *string
}

type UpdateDatasetVersionRequest struct {
	// The dataset version description. Maximum length: 1024 characters.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The dataset version ID. Only DataWorks datasets can be updated.
	//
	// This parameter is required.
	//
	// example:
	//
	// dataworks-datasetVersion:3pXXXb8o0ngr07njhps1
	//
	// :2
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateDatasetVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetVersionRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasetVersionRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateDatasetVersionRequest) GetId() *string {
	return s.Id
}

func (s *UpdateDatasetVersionRequest) SetComment(v string) *UpdateDatasetVersionRequest {
	s.Comment = &v
	return s
}

func (s *UpdateDatasetVersionRequest) SetId(v string) *UpdateDatasetVersionRequest {
	s.Id = &v
	return s
}

func (s *UpdateDatasetVersionRequest) Validate() error {
	return dara.Validate(s)
}
