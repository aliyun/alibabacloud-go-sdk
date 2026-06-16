// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageAsyncModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *ImageAsyncModerationRequest
	GetService() *string
	SetServiceParameters(v string) *ImageAsyncModerationRequest
	GetServiceParameters() *string
}

type ImageAsyncModerationRequest struct {
	// The detection service supported by the enhanced image moderation feature. Valid values:
	//
	// - baselineCheck: common baseline moderation
	//
	// - baselineCheck_pro: common baseline moderation Professional Edition
	//
	// - baselineCheck_cb: common baseline moderation for services outside China
	//
	// - tonalityImprove: content administration moderation
	//
	// - aigcCheck: AIGC image moderation
	//
	// - profilePhotoCheck: profile picture moderation
	//
	// - advertisingCheck: ad material moderation
	//
	// - liveStreamCheck: video or live stream screenshot moderation
	//
	// example:
	//
	// baselineCheck
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// A set of parameters related to the content to be moderated. The value must be a JSON string.
	//
	// example:
	//
	// {"imageUrl":"https://img.alicdn.com/tfs/TB1U4r9AeH2gK0jSZJnXXaT1FXa-2880-480.png","dataId":"img123****"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ImageAsyncModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageAsyncModerationRequest) GoString() string {
	return s.String()
}

func (s *ImageAsyncModerationRequest) GetService() *string {
	return s.Service
}

func (s *ImageAsyncModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *ImageAsyncModerationRequest) SetService(v string) *ImageAsyncModerationRequest {
	s.Service = &v
	return s
}

func (s *ImageAsyncModerationRequest) SetServiceParameters(v string) *ImageAsyncModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *ImageAsyncModerationRequest) Validate() error {
	return dara.Validate(s)
}
