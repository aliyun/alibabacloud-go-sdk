// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRerunTaskInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *RerunTaskInstancesRequest
	GetComment() *string
	SetIds(v []*int64) *RerunTaskInstancesRequest
	GetIds() []*int64
}

type RerunTaskInstancesRequest struct {
	// Remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID list of the task instance.
	Ids []*int64 `json:"Ids,omitempty" xml:"Ids,omitempty" type:"Repeated"`
}

func (s RerunTaskInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RerunTaskInstancesRequest) GoString() string {
	return s.String()
}

func (s *RerunTaskInstancesRequest) GetComment() *string {
	return s.Comment
}

func (s *RerunTaskInstancesRequest) GetIds() []*int64 {
	return s.Ids
}

func (s *RerunTaskInstancesRequest) SetComment(v string) *RerunTaskInstancesRequest {
	s.Comment = &v
	return s
}

func (s *RerunTaskInstancesRequest) SetIds(v []*int64) *RerunTaskInstancesRequest {
	s.Ids = v
	return s
}

func (s *RerunTaskInstancesRequest) Validate() error {
	return dara.Validate(s)
}
