// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *SuspendTaskInstancesRequest
	GetComment() *string
	SetIds(v []*int64) *SuspendTaskInstancesRequest
	GetIds() []*int64
}

type SuspendTaskInstancesRequest struct {
	// Remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID list of the task instance.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s SuspendTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s SuspendTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *SuspendTaskInstancesRequest) GetComment() *string {
	return s.Comment
}

func (s *SuspendTaskInstancesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *SuspendTaskInstancesRequest) SetComment(v string) *SuspendTaskInstancesRequest {
	s.Comment = &v
	return s
}

func (s *SuspendTaskInstancesRequest) SetIds(v []*int64) *SuspendTaskInstancesRequest {
	s.Ids = v
	return s
}

func (s *SuspendTaskInstancesRequest) Validate() error {
	return dara.Validate(s)
}
