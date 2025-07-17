// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartWorkflowInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *StartWorkflowInstancesRequest
	GetComment() *string
	SetIds(v []*int64) *StartWorkflowInstancesRequest
	GetIds() []*int64
}

type StartWorkflowInstancesRequest struct {
	// The remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The IDs of workflow instances.
	//
	// This parameter is required.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s StartWorkflowInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s StartWorkflowInstancesRequest) GoString() string {
	return s.String()
}

func (s *StartWorkflowInstancesRequest) GetComment() *string {
	return s.Comment
}

func (s *StartWorkflowInstancesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *StartWorkflowInstancesRequest) SetComment(v string) *StartWorkflowInstancesRequest {
	s.Comment = &v
	return s
}

func (s *StartWorkflowInstancesRequest) SetIds(v []*int64) *StartWorkflowInstancesRequest {
	s.Ids = v
	return s
}

func (s *StartWorkflowInstancesRequest) Validate() error {
	return dara.Validate(s)
}
