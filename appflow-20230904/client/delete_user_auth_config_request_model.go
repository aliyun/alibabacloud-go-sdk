// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserAuthConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfigId(v string) *DeleteUserAuthConfigRequest
	GetAuthConfigId() *string
	SetConnectorId(v string) *DeleteUserAuthConfigRequest
	GetConnectorId() *string
	SetConnectorVersion(v string) *DeleteUserAuthConfigRequest
	GetConnectorVersion() *string
}

type DeleteUserAuthConfigRequest struct {
	// The credential ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// uac-8f4bfe29041441d8bdd9
	AuthConfigId *string `json:"AuthConfigId,omitempty" xml:"AuthConfigId,omitempty"`
	// The ID of the connector.
	//
	// This parameter is required.
	//
	// example:
	//
	// connector-5e43ef26b53e4a158529
	ConnectorId *string `json:"ConnectorId,omitempty" xml:"ConnectorId,omitempty"`
	// The connector name.
	//
	// example:
	//
	// 2
	ConnectorVersion *string `json:"ConnectorVersion,omitempty" xml:"ConnectorVersion,omitempty"`
}

func (s DeleteUserAuthConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserAuthConfigRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserAuthConfigRequest) GetAuthConfigId() *string {
	return s.AuthConfigId
}

func (s *DeleteUserAuthConfigRequest) GetConnectorId() *string {
	return s.ConnectorId
}

func (s *DeleteUserAuthConfigRequest) GetConnectorVersion() *string {
	return s.ConnectorVersion
}

func (s *DeleteUserAuthConfigRequest) SetAuthConfigId(v string) *DeleteUserAuthConfigRequest {
	s.AuthConfigId = &v
	return s
}

func (s *DeleteUserAuthConfigRequest) SetConnectorId(v string) *DeleteUserAuthConfigRequest {
	s.ConnectorId = &v
	return s
}

func (s *DeleteUserAuthConfigRequest) SetConnectorVersion(v string) *DeleteUserAuthConfigRequest {
	s.ConnectorVersion = &v
	return s
}

func (s *DeleteUserAuthConfigRequest) Validate() error {
	return dara.Validate(s)
}
