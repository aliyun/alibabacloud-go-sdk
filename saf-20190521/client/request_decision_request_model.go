// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRequestDecisionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventCode(v string) *RequestDecisionRequest
	GetEventCode() *string
	SetServiceParameters(v string) *RequestDecisionRequest
	GetServiceParameters() *string
}

type RequestDecisionRequest struct {
	// The event code.
	//
	// example:
	//
	// de_*
	EventCode *string `json:"EventCode,omitempty" xml:"EventCode,omitempty"`
	// The service input parameters. The value is a complete JSON string.
	//
	// example:
	//
	// {"accountId": "1233***"}
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s RequestDecisionRequest) String() string {
	return dara.Prettify(s)
}

func (s RequestDecisionRequest) GoString() string {
	return s.String()
}

func (s *RequestDecisionRequest) GetEventCode() *string {
	return s.EventCode
}

func (s *RequestDecisionRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *RequestDecisionRequest) SetEventCode(v string) *RequestDecisionRequest {
	s.EventCode = &v
	return s
}

func (s *RequestDecisionRequest) SetServiceParameters(v string) *RequestDecisionRequest {
	s.ServiceParameters = &v
	return s
}

func (s *RequestDecisionRequest) Validate() error {
	return dara.Validate(s)
}
