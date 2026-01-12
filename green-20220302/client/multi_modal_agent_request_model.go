// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMultiModalAgentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppID(v string) *MultiModalAgentRequest
	GetAppID() *string
	SetServiceParameters(v string) *MultiModalAgentRequest
	GetServiceParameters() *string
}

type MultiModalAgentRequest struct {
	// example:
	//
	// txt_check_agent_01
	AppID             *string `json:"AppID,omitempty" xml:"AppID,omitempty"`
	ServiceParameters *string `json:"ServiceParameters,omitempty" xml:"ServiceParameters,omitempty"`
}

func (s MultiModalAgentRequest) String() string {
	return dara.Prettify(s)
}

func (s MultiModalAgentRequest) GoString() string {
	return s.String()
}

func (s *MultiModalAgentRequest) GetAppID() *string {
	return s.AppID
}

func (s *MultiModalAgentRequest) GetServiceParameters() *string {
	return s.ServiceParameters
}

func (s *MultiModalAgentRequest) SetAppID(v string) *MultiModalAgentRequest {
	s.AppID = &v
	return s
}

func (s *MultiModalAgentRequest) SetServiceParameters(v string) *MultiModalAgentRequest {
	s.ServiceParameters = &v
	return s
}

func (s *MultiModalAgentRequest) Validate() error {
	return dara.Validate(s)
}
