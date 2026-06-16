// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *MultiModalGuardAsyncRequest
	GetService() *string
	SetServiceParameters(v string) *MultiModalGuardAsyncRequest
	GetServiceParameters() *string
}

type MultiModalGuardAsyncRequest struct {
	// The moderation service type. Valid values: `audio_security_check` and `video_security_check`.
	//
	// example:
	//
	// audio_security_check
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// The parameter set required for the moderation service.
	//
	// example:
	//
	// {"url": "https://testxxx.oss-cn-shanghai.aliyuncs.com/xxx.mp4", "dataId": "data1234"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s MultiModalGuardAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncRequest) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncRequest) GetService() *string {
	return s.Service
}

func (s *MultiModalGuardAsyncRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *MultiModalGuardAsyncRequest) SetService(v string) *MultiModalGuardAsyncRequest {
	s.Service = &v
	return s
}

func (s *MultiModalGuardAsyncRequest) SetServiceParameters(v string) *MultiModalGuardAsyncRequest {
	s.ServiceParameters = &v
	return s
}

func (s *MultiModalGuardAsyncRequest) Validate() error {
	return dara.Validate(s)
}
