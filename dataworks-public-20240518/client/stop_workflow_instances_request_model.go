// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopWorkflowInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *StopWorkflowInstancesRequest
	GetComment() *string
	SetIds(v []*int64) *StopWorkflowInstancesRequest
	GetIds() []*int64
}

type StopWorkflowInstancesRequest struct {
	// The remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The workflow instance IDs.
	//
	// This parameter is required.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s StopWorkflowInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s StopWorkflowInstancesRequest) GoString() string {
	return s.String()
}

func (s *StopWorkflowInstancesRequest) GetComment() *string {
	return s.Comment
}

func (s *StopWorkflowInstancesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *StopWorkflowInstancesRequest) SetComment(v string) *StopWorkflowInstancesRequest {
	s.Comment = &v
	return s
}

func (s *StopWorkflowInstancesRequest) SetIds(v []*int64) *StopWorkflowInstancesRequest {
	s.Ids = v
	return s
}

func (s *StopWorkflowInstancesRequest) Validate() error {
	return dara.Validate(s)
}
