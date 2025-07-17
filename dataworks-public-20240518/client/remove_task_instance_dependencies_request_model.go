// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTaskInstanceDependenciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *RemoveTaskInstanceDependenciesRequest
	GetComment() *string
	SetId(v int64) *RemoveTaskInstanceDependenciesRequest
	GetId() *int64
	SetUpstreamTaskInstanceIds(v []*int64) *RemoveTaskInstanceDependenciesRequest
	GetUpstreamTaskInstanceIds() []*int64
}

type RemoveTaskInstanceDependenciesRequest struct {
	// The remarks.
	//
	// example:
	//
	// this is a comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The IDs of ancestor instances of the instance
	UpstreamTaskInstanceIds []*int64 `json:"UpstreamTaskInstanceIds,omitempty" xml:"UpstreamTaskInstanceIds,omitempty" type:"Repeated"`
}

func (s RemoveTaskInstanceDependenciesRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTaskInstanceDependenciesRequest) GoString() string {
	return s.String()
}

func (s *RemoveTaskInstanceDependenciesRequest) GetComment() *string {
	return s.Comment
}

func (s *RemoveTaskInstanceDependenciesRequest) GetId() *int64 {
	return s.Id
}

func (s *RemoveTaskInstanceDependenciesRequest) GetUpstreamTaskInstanceIds() []*int64 {
	return s.UpstreamTaskInstanceIds
}

func (s *RemoveTaskInstanceDependenciesRequest) SetComment(v string) *RemoveTaskInstanceDependenciesRequest {
	s.Comment = &v
	return s
}

func (s *RemoveTaskInstanceDependenciesRequest) SetId(v int64) *RemoveTaskInstanceDependenciesRequest {
	s.Id = &v
	return s
}

func (s *RemoveTaskInstanceDependenciesRequest) SetUpstreamTaskInstanceIds(v []*int64) *RemoveTaskInstanceDependenciesRequest {
	s.UpstreamTaskInstanceIds = v
	return s
}

func (s *RemoveTaskInstanceDependenciesRequest) Validate() error {
	return dara.Validate(s)
}
