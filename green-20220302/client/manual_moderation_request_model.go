// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManualModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *ManualModerationRequest
	GetService() *string
	SetServiceParameters(v string) *ManualModerationRequest
	GetServiceParameters() *string
}

type ManualModerationRequest struct {
	// Service.
	//
	// example:
	//
	// imageManualCheck
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// Parameters required by the moderation service, in JSON string format.
	//
	// - url: The URL of the object to be inspected. Make sure the URL is accessible through the public network.
	//
	// - dataId: Optional. The data ID corresponding to the object being inspected.
	//
	// example:
	//
	// {"url": "https://talesofai.oss-cn-shanghai.aliyuncs.com/xxx.mp4", "dataId": "data1234"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s ManualModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s ManualModerationRequest) GoString() string {
	return s.String()
}

func (s *ManualModerationRequest) GetService() *string {
	return s.Service
}

func (s *ManualModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *ManualModerationRequest) SetService(v string) *ManualModerationRequest {
	s.Service = &v
	return s
}

func (s *ManualModerationRequest) SetServiceParameters(v string) *ManualModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *ManualModerationRequest) Validate() error {
	return dara.Validate(s)
}
