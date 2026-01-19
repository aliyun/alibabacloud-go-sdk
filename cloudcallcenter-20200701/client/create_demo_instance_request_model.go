// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDemoInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOutboundCallWhitelist(v string) *CreateDemoInstanceRequest
	GetOutboundCallWhitelist() *string
}

type CreateDemoInstanceRequest struct {
	// This parameter is required.
	OutboundCallWhitelist *string `json:"OutboundCallWhitelist,omitempty" xml:"OutboundCallWhitelist,omitempty"`
}

func (s CreateDemoInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDemoInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateDemoInstanceRequest) GetOutboundCallWhitelist() *string {
	return s.OutboundCallWhitelist
}

func (s *CreateDemoInstanceRequest) SetOutboundCallWhitelist(v string) *CreateDemoInstanceRequest {
	s.OutboundCallWhitelist = &v
	return s
}

func (s *CreateDemoInstanceRequest) Validate() error {
	return dara.Validate(s)
}
