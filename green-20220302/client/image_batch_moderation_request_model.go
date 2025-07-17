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
	// The types of detection supported by the enhanced image review, separated by English commas. Values:
	//
	// - baselineCheck：General Baseline Detection
	//
	// - baselineCheck_pro：General Baseline Detection_Pro Edition
	//
	// - tonalityImprove：Content governance monitoring
	//
	// - aigcCheck：AIGC image detection
	//
	// example:
	//
	// baselineCheck,tonalityImprove
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The set of relevant parameters for content detection objects.
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
