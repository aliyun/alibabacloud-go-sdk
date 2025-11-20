// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActivateTemplateMCPRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEnabledTools(v []*string) *ActivateTemplateMCPRequest
	GetEnabledTools() []*string
	SetTransport(v string) *ActivateTemplateMCPRequest
	GetTransport() *string
}

type ActivateTemplateMCPRequest struct {
	EnabledTools []*string `json:"enabledTools,omitempty" xml:"enabledTools,omitempty" type:"Repeated"`
	// example:
	//
	// streamable-http
	Transport *string `json:"transport,omitempty" xml:"transport,omitempty"`
}

func (s ActivateTemplateMCPRequest) String() string {
	return dara.Prettify(s)
}

func (s ActivateTemplateMCPRequest) GoString() string {
	return s.String()
}

func (s *ActivateTemplateMCPRequest) GetEnabledTools() []*string {
	return s.EnabledTools
}

func (s *ActivateTemplateMCPRequest) GetTransport() *string {
	return s.Transport
}

func (s *ActivateTemplateMCPRequest) SetEnabledTools(v []*string) *ActivateTemplateMCPRequest {
	s.EnabledTools = v
	return s
}

func (s *ActivateTemplateMCPRequest) SetTransport(v string) *ActivateTemplateMCPRequest {
	s.Transport = &v
	return s
}

func (s *ActivateTemplateMCPRequest) Validate() error {
	return dara.Validate(s)
}
