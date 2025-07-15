// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitCustomHotTopicBroadcastJobShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHotTopicBroadcastConfigShrink(v string) *SubmitCustomHotTopicBroadcastJobShrinkRequest
	GetHotTopicBroadcastConfigShrink() *string
	SetHotTopicVersion(v string) *SubmitCustomHotTopicBroadcastJobShrinkRequest
	GetHotTopicVersion() *string
	SetTopicsShrink(v string) *SubmitCustomHotTopicBroadcastJobShrinkRequest
	GetTopicsShrink() *string
	SetWorkspaceId(v string) *SubmitCustomHotTopicBroadcastJobShrinkRequest
	GetWorkspaceId() *string
}

type SubmitCustomHotTopicBroadcastJobShrinkRequest struct {
	// This parameter is required.
	HotTopicBroadcastConfigShrink *string `json:"HotTopicBroadcastConfig,omitempty" xml:"HotTopicBroadcastConfig,omitempty"`
	// example:
	//
	// 热点版本
	HotTopicVersion *string `json:"HotTopicVersion,omitempty" xml:"HotTopicVersion,omitempty"`
	TopicsShrink    *string `json:"Topics,omitempty" xml:"Topics,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s SubmitCustomHotTopicBroadcastJobShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitCustomHotTopicBroadcastJobShrinkRequest) GoString() string {
	return s.String()
}

func (s *SubmitCustomHotTopicBroadcastJobShrinkRequest) GetHotTopicBroadcastConfigShrink() *string {
	return s.HotTopicBroadcastConfigShrink
}

func (s *SubmitCustomHotTopicBroadcastJobShrinkRequest) GetHotTopicVersion() *string {
	return s.HotTopicVersion
}

func (s *SubmitCustomHotTopicBroadcastJobShrinkRequest) GetTopicsShrink() *string {
	return s.TopicsShrink
}

func (s *SubmitCustomHotTopicBroadcastJobShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *SubmitCustomHotTopicBroadcastJobShrinkRequest) SetHotTopicBroadcastConfigShrink(v string) *SubmitCustomHotTopicBroadcastJobShrinkRequest {
	s.HotTopicBroadcastConfigShrink = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobShrinkRequest) SetHotTopicVersion(v string) *SubmitCustomHotTopicBroadcastJobShrinkRequest {
	s.HotTopicVersion = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobShrinkRequest) SetTopicsShrink(v string) *SubmitCustomHotTopicBroadcastJobShrinkRequest {
	s.TopicsShrink = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobShrinkRequest) SetWorkspaceId(v string) *SubmitCustomHotTopicBroadcastJobShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *SubmitCustomHotTopicBroadcastJobShrinkRequest) Validate() error {
	return dara.Validate(s)
}
