// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSemanticJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DeleteSemanticJobRequest
	GetName() *string
}

type DeleteSemanticJobRequest struct {
	// The name of the job to delete. Use the Data.Name value from the CreateSemanticJob response or the Name value from a ListSemanticJobs list item.
	//
	// This parameter is required.
	//
	// example:
	//
	// semantic-job-demo
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteSemanticJobRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSemanticJobRequest) GoString() string {
	return s.String()
}

func (s *DeleteSemanticJobRequest) GetName() *string {
	return s.Name
}

func (s *DeleteSemanticJobRequest) SetName(v string) *DeleteSemanticJobRequest {
	s.Name = &v
	return s
}

func (s *DeleteSemanticJobRequest) Validate() error {
	return dara.Validate(s)
}
