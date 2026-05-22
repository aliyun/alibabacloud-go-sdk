// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardAsyncResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *MultiModalGuardAsyncResultRequest
	GetService() *string
	SetServiceParameters(v string) *MultiModalGuardAsyncResultRequest
	GetServiceParameters() *string
}

type MultiModalGuardAsyncResultRequest struct {
	// example:
	//
	// audio_security_check
	Service *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// example:
	//
	// {
	//
	//   "url": "https://xxx.mp4"
	//
	// }
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s MultiModalGuardAsyncResultRequest) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardAsyncResultRequest) GoString() string {
	return s.String()
}

func (s *MultiModalGuardAsyncResultRequest) GetService() *string {
	return s.Service
}

func (s *MultiModalGuardAsyncResultRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *MultiModalGuardAsyncResultRequest) SetService(v string) *MultiModalGuardAsyncResultRequest {
	s.Service = &v
	return s
}

func (s *MultiModalGuardAsyncResultRequest) SetServiceParameters(v string) *MultiModalGuardAsyncResultRequest {
	s.ServiceParameters = &v
	return s
}

func (s *MultiModalGuardAsyncResultRequest) Validate() error {
	return dara.Validate(s)
}
