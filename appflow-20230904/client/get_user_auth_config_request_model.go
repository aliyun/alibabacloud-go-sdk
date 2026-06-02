// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAuthConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfigId(v string) *GetUserAuthConfigRequest
	GetAuthConfigId() *string
	SetConnectorId(v string) *GetUserAuthConfigRequest
	GetConnectorId() *string
	SetConnectorVersion(v string) *GetUserAuthConfigRequest
	GetConnectorVersion() *string
}

type GetUserAuthConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// uac-ac11e0f1db7647ce8469
	AuthConfigId *string `json:"AuthConfigId,omitempty" xml:"AuthConfigId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// connector-5e43ef26b53e4a158529
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// example:
	//
	// 1
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" xml:"ConnectorVersion,omitempty"`
}

func (s GetUserAuthConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetUserAuthConfigRequest) GoString() string {
	return s.String()
}

func (s *GetUserAuthConfigRequest) GetAuthConfigId() *string {
	return s.AuthConfigId
}

func (s *GetUserAuthConfigRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *GetUserAuthConfigRequest) GetConnectorVersion() *string {
	return s.ConnectorVersion
}

func (s *GetUserAuthConfigRequest) SetAuthConfigId(v string) *GetUserAuthConfigRequest {
	s.AuthConfigId = &v
	return s
}

func (s *GetUserAuthConfigRequest) SetConnectorId(v string) *GetUserAuthConfigRequest {
	s.ConnectorId = &v
	return s
}

func (s *GetUserAuthConfigRequest) SetConnectorVersion(v string) *GetUserAuthConfigRequest {
	s.ConnectorVersion = &v
	return s
}

func (s *GetUserAuthConfigRequest) Validate() error {
	return dara.Validate(s)
}
