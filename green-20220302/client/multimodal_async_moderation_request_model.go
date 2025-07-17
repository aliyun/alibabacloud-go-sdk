// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultimodalAsyncModerationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *MultimodalAsyncModerationRequest
	GetService() *string
	SetServiceParameters(v string) *MultimodalAsyncModerationRequest
	GetServiceParameters() *string
}

type MultimodalAsyncModerationRequest struct {
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s MultimodalAsyncModerationRequest) String() string {
	return dara.Prettify(s)
}

func (s MultimodalAsyncModerationRequest) GoString() string {
	return s.String()
}

func (s *MultimodalAsyncModerationRequest) GetService() *string {
	return s.Service
}

func (s *MultimodalAsyncModerationRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *MultimodalAsyncModerationRequest) SetService(v string) *MultimodalAsyncModerationRequest {
	s.Service = &v
	return s
}

func (s *MultimodalAsyncModerationRequest) SetServiceParameters(v string) *MultimodalAsyncModerationRequest {
	s.ServiceParameters = &v
	return s
}

func (s *MultimodalAsyncModerationRequest) Validate() error {
	return dara.Validate(s)
}
