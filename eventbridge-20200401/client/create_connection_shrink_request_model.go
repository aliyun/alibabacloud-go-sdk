// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthParametersShrink(v string) *CreateConnectionShrinkRequest
	GetAuthParametersShrink() *string
	SetConnectionName(v string) *CreateConnectionShrinkRequest
	GetConnectionName() *string
	SetDescription(v string) *CreateConnectionShrinkRequest
	GetDescription() *string
	SetNetworkParametersShrink(v string) *CreateConnectionShrinkRequest
	GetNetworkParametersShrink() *string
	SetParametersShrink(v string) *CreateConnectionShrinkRequest
	GetParametersShrink() *string
	SetType(v string) *CreateConnectionShrinkRequest
	GetType() *string
}

type CreateConnectionShrinkRequest struct {
	// The authentication configuration.
	AuthParametersShrink *string `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty"`
	// The connection configuration name. Maximum length: 127 characters. Minimum length: 2 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the connection configuration. Maximum length: 255 characters.
	//
	// example:
	//
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The network configuration.
	//
	// This parameter is required.
	NetworkParametersShrink *string `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty"`
	// The data source connection parameters (JSON object). This parameter is required when Type is set to a data source type. This parameter is not required for the Http type. For specific field definitions, call the GetConnectionType operation and refer to ParamsSchema in the response.
	//
	// example:
	//
	// {"HostName":"xxx.mysql.rds.aliyuncs.com","Port":"3306","User":"root","Password":"xxx","DatabaseName":"demo_db"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The connection type. Valid values: MySQL, PostgreSQL, Elasticsearch, and Http. This parameter is required for data source connections. If this parameter is not specified, the default value Http is used. The Http type is used for HTTP protocol targets such as API Destination. Data source types are used for data connections in the integration marketplace.
	//
	// example:
	//
	// Http
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateConnectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateConnectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateConnectionShrinkRequest) GetAuthParametersShrink() *string {
	return s.AuthParametersShrink
}

func (s *CreateConnectionShrinkRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *CreateConnectionShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateConnectionShrinkRequest) GetNetworkParametersShrink() *string {
	return s.NetworkParametersShrink
}

func (s *CreateConnectionShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *CreateConnectionShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateConnectionShrinkRequest) SetAuthParametersShrink(v string) *CreateConnectionShrinkRequest {
	s.AuthParametersShrink = &v
	return s
}

func (s *CreateConnectionShrinkRequest) SetConnectionName(v string) *CreateConnectionShrinkRequest {
	s.ConnectionName = &v
	return s
}

func (s *CreateConnectionShrinkRequest) SetDescription(v string) *CreateConnectionShrinkRequest {
	s.Description = &v
	return s
}

func (s *CreateConnectionShrinkRequest) SetNetworkParametersShrink(v string) *CreateConnectionShrinkRequest {
	s.NetworkParametersShrink = &v
	return s
}

func (s *CreateConnectionShrinkRequest) SetParametersShrink(v string) *CreateConnectionShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *CreateConnectionShrinkRequest) SetType(v string) *CreateConnectionShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateConnectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
