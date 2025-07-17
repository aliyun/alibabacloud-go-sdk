// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *StopTaskInstancesRequest
	GetComment() *string
	SetIds(v []*int64) *StopTaskInstancesRequest
	GetIds() []*int64
}

type StopTaskInstancesRequest struct {
	// Remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID list of the task instance.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s StopTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s StopTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *StopTaskInstancesRequest) GetComment() *string {
	return s.Comment
}

func (s *StopTaskInstancesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *StopTaskInstancesRequest) SetComment(v string) *StopTaskInstancesRequest {
	s.Comment = &v
	return s
}

func (s *StopTaskInstancesRequest) SetIds(v []*int64) *StopTaskInstancesRequest {
	s.Ids = v
	return s
}

func (s *StopTaskInstancesRequest) Validate() error {
	return dara.Validate(s)
}
