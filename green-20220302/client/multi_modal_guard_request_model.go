// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalGuardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetService(v string) *MultiModalGuardRequest
	GetService() *string
	SetServiceParameters(v string) *MultiModalGuardRequest
	GetServiceParameters() *string
}

type MultiModalGuardRequest struct {
	// example:
	//
	// query_security_check
	Service           *string `json:"Service,omitempty" xml:"Service,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s MultiModalGuardRequest) String() string {
	return dara.Prettify(s)
}

func (s MultiModalGuardRequest) GoString() string {
	return s.String()
}

func (s *MultiModalGuardRequest) GetService() *string {
	return s.Service
}

func (s *MultiModalGuardRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *MultiModalGuardRequest) SetService(v string) *MultiModalGuardRequest {
	s.Service = &v
	return s
}

func (s *MultiModalGuardRequest) SetServiceParameters(v string) *MultiModalGuardRequest {
	s.ServiceParameters = &v
	return s
}

func (s *MultiModalGuardRequest) Validate() error {
	return dara.Validate(s)
}
