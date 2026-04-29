// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPolarClawMCPServerShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AddPolarClawMCPServerShrinkRequest
	GetApplicationId() *string
	SetServerConfigShrink(v string) *AddPolarClawMCPServerShrinkRequest
	GetServerConfigShrink() *string
	SetServerName(v string) *AddPolarClawMCPServerShrinkRequest
	GetServerName() *string
}

type AddPolarClawMCPServerShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// {
	//
	//     "command": "node",
	//
	//     "args": [
	//
	//         "-y",
	//
	//         "@polarclaw/mcp-dev"
	//
	//     ]
	//
	// }
	ServerConfigShrink *string `json:"ServerConfig,omitempty" xml:"ServerConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-v1
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s AddPolarClawMCPServerShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPolarClawMCPServerShrinkRequest) GoString() string {
	return s.String()
}

func (s *AddPolarClawMCPServerShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AddPolarClawMCPServerShrinkRequest) GetServerConfigShrink() *string {
	return s.ServerConfigShrink
}

func (s *AddPolarClawMCPServerShrinkRequest) GetServerName() *string {
	return s.ServerName
}

func (s *AddPolarClawMCPServerShrinkRequest) SetApplicationId(v string) *AddPolarClawMCPServerShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *AddPolarClawMCPServerShrinkRequest) SetServerConfigShrink(v string) *AddPolarClawMCPServerShrinkRequest {
	s.ServerConfigShrink = &v
	return s
}

func (s *AddPolarClawMCPServerShrinkRequest) SetServerName(v string) *AddPolarClawMCPServerShrinkRequest {
	s.ServerName = &v
	return s
}

func (s *AddPolarClawMCPServerShrinkRequest) Validate() error {
	return dara.Validate(s)
}
