// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPipelineJobsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ListPipelineJobsRequest
	GetCategory() *string
}

type ListPipelineJobsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// DEPLOY
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
}

func (s ListPipelineJobsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPipelineJobsRequest) GoString() string {
	return s.String()
}

func (s *ListPipelineJobsRequest) GetCategory() *string {
	return s.Category
}

func (s *ListPipelineJobsRequest) SetCategory(v string) *ListPipelineJobsRequest {
	s.Category = &v
	return s
}

func (s *ListPipelineJobsRequest) Validate() error {
	return dara.Validate(s)
}
