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
	// The detection types supported by Image Moderation Enhanced Edition. Valid values:
	//
	// - baselineCheck: general baseline check
	//
	// - baselineCheck_pro: general baseline check (Professional Edition)
	//
	// - baselineCheck_cb: general baseline check (Overseas Edition)
	//
	// - tonalityImprove: content governance detection
	//
	// - aigcCheck: AIGC image detection
	//
	// - aigcViolationDetection: AIGC image infringement detection
	//
	// - aigcDetector: AIGC image generation determination
	//
	// - profilePhotoCheck: profile picture detection
	//
	// - postImageCheck: post and comment image detection
	//
	// - advertisingCheck: marketing material detection
	//
	// - liveStreamCheck: video or live stream screenshot detection
	//
	// - generalOcr: general image and text OCR
	//
	// - generalRecognition: universal image recognition
	//
	// - postImageCheckByVL: image moderation service with large and small model fusion
	//
	// - postImageCheckByVL_cb: image moderation service with large and small model fusion (Overseas Edition)
	//
	// - baselineCheckByVL: general image moderation large model service
	//
	// example:
	//
	// baselineCheck
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameter set for the content moderation object. The value is a JSON string.
	//
	// - imageUrl: the URL of the object to be moderated. Required.
	//
	// - dataId: the data ID corresponding to the moderation object. Optional.
	//
	// - referer: the Referer request header, used for scenarios such as hotlink protection. Optional.
	//
	// example:
	//
	// {"imageUrl":"https://img.alicdn.com/tfs/TB1U4r9AeH2gK0jSZJnXXaT1FXa-2880-480.png","dataId":"img1234567"}
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
