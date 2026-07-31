// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunSemanticJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *RunSemanticJobRequest
	GetName() *string
}

type RunSemanticJobRequest struct {
	// The name of the job to run. Use the Data.Name value returned by CreateSemanticJob or the Name field from the ListSemanticJobs response. The Source, ResourceGroupId, and reference files of the job are determined by the definition saved at creation time.
	//
	// This parameter is required.
	//
	// example:
	//
	// semantic-job-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s RunSemanticJobRequest) String() string {
	return dara.Prettify(s)
}

func (s RunSemanticJobRequest) GoString() string {
	return s.String()
}

func (s *RunSemanticJobRequest) GetName() *string {
	return s.Name
}

func (s *RunSemanticJobRequest) SetName(v string) *RunSemanticJobRequest {
	s.Name = &v
	return s
}

func (s *RunSemanticJobRequest) Validate() error {
	return dara.Validate(s)
}
