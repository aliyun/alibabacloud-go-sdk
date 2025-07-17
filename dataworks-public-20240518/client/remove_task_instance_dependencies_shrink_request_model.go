// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveTaskInstanceDependenciesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *RemoveTaskInstanceDependenciesShrinkRequest
	GetComment() *string
	SetId(v int64) *RemoveTaskInstanceDependenciesShrinkRequest
	GetId() *int64
	SetUpstreamTaskInstanceIdsShrink(v string) *RemoveTaskInstanceDependenciesShrinkRequest
	GetUpstreamTaskInstanceIdsShrink() *string
}

type RemoveTaskInstanceDependenciesShrinkRequest struct {
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
	UpstreamTaskInstanceIdsShrink *string `json:"UpstreamTaskInstanceIds,omitempty" xml:"UpstreamTaskInstanceIds,omitempty"`
}

func (s RemoveTaskInstanceDependenciesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveTaskInstanceDependenciesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RemoveTaskInstanceDependenciesShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *RemoveTaskInstanceDependenciesShrinkRequest) GetId() *int64 {
	return s.Id
}

func (s *RemoveTaskInstanceDependenciesShrinkRequest) GetUpstreamTaskInstanceIdsShrink() *string {
	return s.UpstreamTaskInstanceIdsShrink
}

func (s *RemoveTaskInstanceDependenciesShrinkRequest) SetComment(v string) *RemoveTaskInstanceDependenciesShrinkRequest {
	s.Comment = &v
	return s
}

func (s *RemoveTaskInstanceDependenciesShrinkRequest) SetId(v int64) *RemoveTaskInstanceDependenciesShrinkRequest {
	s.Id = &v
	return s
}

func (s *RemoveTaskInstanceDependenciesShrinkRequest) SetUpstreamTaskInstanceIdsShrink(v string) *RemoveTaskInstanceDependenciesShrinkRequest {
	s.UpstreamTaskInstanceIdsShrink = &v
	return s
}

func (s *RemoveTaskInstanceDependenciesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
