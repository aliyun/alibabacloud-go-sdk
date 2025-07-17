// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *ResumeTaskInstancesRequest
	GetComment() *string
	SetIds(v []*int64) *ResumeTaskInstancesRequest
	GetIds() []*int64
}

type ResumeTaskInstancesRequest struct {
	// Remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID list of the task instance.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s ResumeTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ResumeTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *ResumeTaskInstancesRequest) GetComment() *string {
	return s.Comment
}

func (s *ResumeTaskInstancesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *ResumeTaskInstancesRequest) SetComment(v string) *ResumeTaskInstancesRequest {
	s.Comment = &v
	return s
}

func (s *ResumeTaskInstancesRequest) SetIds(v []*int64) *ResumeTaskInstancesRequest {
	s.Ids = v
	return s
}

func (s *ResumeTaskInstancesRequest) Validate() error {
	return dara.Validate(s)
}
