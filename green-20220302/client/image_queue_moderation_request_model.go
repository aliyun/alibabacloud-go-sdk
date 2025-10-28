// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImageQueueModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *ImageQueueModerationRequest
	GetService() *string
	SetServiceParameters(v string) *ImageQueueModerationRequest
	GetServiceParameters() *string
}

type ImageQueueModerationRequest struct {
	// example:
	//
	// baselineCheck
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// example:
	//
	// {
	//
	//         "imageUrl": "https://img.alicdn.com/xxx.jpg",
	//
	//         "dataId": "img123****"
	//
	//     }
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ImageQueueModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s ImageQueueModerationRequest) GoString() string {
	return s.String()
}

func (s *ImageQueueModerationRequest) GetService() *string {
	return s.Service
}

func (s *ImageQueueModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *ImageQueueModerationRequest) SetService(v string) *ImageQueueModerationRequest {
	s.Service = &v
	return s
}

func (s *ImageQueueModerationRequest) SetServiceParameters(v string) *ImageQueueModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *ImageQueueModerationRequest) Validate() error {
	return dara.Validate(s)
}
