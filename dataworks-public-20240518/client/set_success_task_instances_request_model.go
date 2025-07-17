// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSuccessTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *SetSuccessTaskInstancesRequest
	GetComment() *string
	SetIds(v []*int64) *SetSuccessTaskInstancesRequest
	GetIds() []*int64
}

type SetSuccessTaskInstancesRequest struct {
	// Remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID list of the task instance.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s SetSuccessTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetSuccessTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *SetSuccessTaskInstancesRequest) GetComment() *string {
	return s.Comment
}

func (s *SetSuccessTaskInstancesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *SetSuccessTaskInstancesRequest) SetComment(v string) *SetSuccessTaskInstancesRequest {
	s.Comment = &v
	return s
}

func (s *SetSuccessTaskInstancesRequest) SetIds(v []*int64) *SetSuccessTaskInstancesRequest {
	s.Ids = v
	return s
}

func (s *SetSuccessTaskInstancesRequest) Validate() error {
	return dara.Validate(s)
}
