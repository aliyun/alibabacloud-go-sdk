// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *ImageModerationRequest
	GetService() *string
	SetServiceParameters(v string) *ImageModerationRequest
	GetServiceParameters() *string
}

type ImageModerationRequest struct {
	// The moderation services supported by Image Moderation 2.0. Valid values:
	//
	// 	- baselineCheck: common baseline moderation
	//
	// 	- baselineCheck_pro: common baseline moderation_Professional
	//
	// 	- baselineCheck_cb: common baseline moderation_For regions outside the Chinese mainland
	//
	// 	- tonalityImprove: content governance moderation
	//
	// 	- aigcCheck: AI-generated image identification
	//
	// 	- profilePhotoCheck: avatar image moderation
	//
	// 	- advertisingCheck: marketing material identification
	//
	// 	- liveStreamCheck: moderation of screenshots of videos and live streams
	//
	// Valid values:
	//
	// 	- liveStreamCheck: moderation of screenshots of videos and live streams
	//
	// 	- baselineCheck: common baseline moderation
	//
	// 	- aigcCheck: AI-generated image identification
	//
	// 	- baselineCheck_pro: common baseline moderation_Professional
	//
	// 	- advertisingCheck: marketing material identification
	//
	// 	- baselineCheck_cb: common baseline moderation_For regions outside the Chinese mainland
	//
	// 	- tonalityImprove: content governance moderation
	//
	// 	- profilePhotoCheck: avatar image moderation
	//
	// example:
	//
	// baselineCheck
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// 	- imageUrl: the URL of the object that you want to moderate. This parameter is required.
	//
	// 	- dataId: the ID of the object that you want to moderate. This parameter is optional.
	//
	// example:
	//
	// {"imageUrl":"https://www.aliyun.com/test.jpg","dataId":"img1234567"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ImageModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageModerationRequest) GoString() string {
	return s.String()
}

func (s *ImageModerationRequest) GetService() *string {
	return s.Service
}

func (s *ImageModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *ImageModerationRequest) SetService(v string) *ImageModerationRequest {
	s.Service = &v
	return s
}

func (s *ImageModerationRequest) SetServiceParameters(v string) *ImageModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *ImageModerationRequest) Validate() error {
	return dara.Validate(s)
}
