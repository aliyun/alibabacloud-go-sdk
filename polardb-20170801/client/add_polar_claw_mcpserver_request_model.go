// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPolarClawMCPServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *AddPolarClawMCPServerRequest
	GetApplicationId() *string
	SetServerConfig(v map[string]interface{}) *AddPolarClawMCPServerRequest
	GetServerConfig() map[string]interface{}
	SetServerName(v string) *AddPolarClawMCPServerRequest
	GetServerName() *string
}

type AddPolarClawMCPServerRequest struct {
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
	ServerConfig map[string]interface{} `json:"ServerConfig,omitempty" xml:"ServerConfig,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test-v1
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s AddPolarClawMCPServerRequest) String() string {
	return dara.Prettify(s)
}

func (s AddPolarClawMCPServerRequest) GoString() string {
	return s.String()
}

func (s *AddPolarClawMCPServerRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *AddPolarClawMCPServerRequest) GetServerConfig() map[string]interface{} {
	return s.ServerConfig
}

func (s *AddPolarClawMCPServerRequest) GetServerName() *string {
	return s.ServerName
}

func (s *AddPolarClawMCPServerRequest) SetApplicationId(v string) *AddPolarClawMCPServerRequest {
	s.ApplicationId = &v
	return s
}

func (s *AddPolarClawMCPServerRequest) SetServerConfig(v map[string]interface{}) *AddPolarClawMCPServerRequest {
	s.ServerConfig = v
	return s
}

func (s *AddPolarClawMCPServerRequest) SetServerName(v string) *AddPolarClawMCPServerRequest {
	s.ServerName = &v
	return s
}

func (s *AddPolarClawMCPServerRequest) Validate() error {
	return dara.Validate(s)
}
