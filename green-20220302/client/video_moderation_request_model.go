// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVideoModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *VideoModerationRequest
	GetService() *string
	SetServiceParameters(v string) *VideoModerationRequest
	GetServiceParameters() *string
}

type VideoModerationRequest struct {
	// The type of the moderation service.
	//
	// example:
	//
	// videoDetection
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameters required by the moderation service. The value is a JSON string.
	//
	// example:
	//
	// {\\"url\\": \\"https://talesofai.oss-cn-shanghai.aliyuncs.com/xxx.mp4\\", \\"dataId\\": \\"94db0b88-f521-11ed-806e-fae21c1f239c\\"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s VideoModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s VideoModerationRequest) GoString() string {
	return s.String()
}

func (s *VideoModerationRequest) GetService() *string {
	return s.Service
}

func (s *VideoModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *VideoModerationRequest) SetService(v string) *VideoModerationRequest {
	s.Service = &v
	return s
}

func (s *VideoModerationRequest) SetServiceParameters(v string) *VideoModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *VideoModerationRequest) Validate() error {
	return dara.Validate(s)
}
