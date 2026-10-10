// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageBatchModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *ImageBatchModerationRequest
	GetService() *string
	SetServiceParameters(v string) *ImageBatchModerationRequest
	GetServiceParameters() *string
}

type ImageBatchModerationRequest struct {
	// The detection types supported by Image Moderation Enhanced Edition. Separate multiple values with commas. Valid values:
	//
	// - baselineCheck: general baseline check
	//
	// - baselineCheck_pro: general baseline check professional edition
	//
	// - tonalityImprove: content governance detection
	//
	// - aigcCheck: AIGC image detection
	//
	// example:
	//
	// baselineCheck,tonalityImprove
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameter set for the content moderation object.
	//
	// example:
	//
	// {
	//
	//         "imageUrl": "https://img.alicdn.com/tfs/TB1U4r9AeH2gK0jSZJnXXaT1FXa-2880-480.png",
	//
	//         "dataId": "img123****"
	//
	//     }
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ImageBatchModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageBatchModerationRequest) GoString() string {
	return s.String()
}

func (s *ImageBatchModerationRequest) GetService() *string {
	return s.Service
}

func (s *ImageBatchModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *ImageBatchModerationRequest) SetService(v string) *ImageBatchModerationRequest {
	s.Service = &v
	return s
}

func (s *ImageBatchModerationRequest) SetServiceParameters(v string) *ImageBatchModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *ImageBatchModerationRequest) Validate() error {
	return dara.Validate(s)
}
