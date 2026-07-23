// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateConnectionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthParametersShrink(v string) *UpdateConnectionShrinkRequest
	GetAuthParametersShrink() *string
	SetConnectionName(v string) *UpdateConnectionShrinkRequest
	GetConnectionName() *string
	SetDescription(v string) *UpdateConnectionShrinkRequest
	GetDescription() *string
	SetNetworkParametersShrink(v string) *UpdateConnectionShrinkRequest
	GetNetworkParametersShrink() *string
	SetParametersShrink(v string) *UpdateConnectionShrinkRequest
	GetParametersShrink() *string
	SetType(v string) *UpdateConnectionShrinkRequest
	GetType() *string
}

type UpdateConnectionShrinkRequest struct {
	// The data structure of the authentication parameters.
	AuthParametersShrink *string `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty"`
	// The name of the connection to be updated. The maximum length is 127 characters. The minimum length is 2 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description. The maximum length is 255 characters.
	//
	// example:
	//
	// Description of the connection configuration
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The data structure of the network configuration.
	//
	// This parameter is required.
	NetworkParametersShrink *string `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty"`
	// The data source connection parameters (JSON object). For specific field definitions, call the GetConnectionType API and refer to the ParamsSchema in the response.
	//
	// example:
	//
	// {"HostName":"xxx.mysql.rds.aliyuncs.com","Port":"3306","User":"root","Password":"xxx","DatabaseName":"demo_db"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The connection type. Valid values: MySQL, PostgreSQL, Elasticsearch, and Http.
	//
	// example:
	//
	// Http
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateConnectionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateConnectionShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateConnectionShrinkRequest) GetAuthParametersShrink() *string {
	return s.AuthParametersShrink
}

func (s *UpdateConnectionShrinkRequest) GetConnectionName() *string {
	return s.ConnectionName
}

func (s *UpdateConnectionShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateConnectionShrinkRequest) GetNetworkParametersShrink() *string {
	return s.NetworkParametersShrink
}

func (s *UpdateConnectionShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *UpdateConnectionShrinkRequest) GetType() *string {
	return s.Type
}

func (s *UpdateConnectionShrinkRequest) SetAuthParametersShrink(v string) *UpdateConnectionShrinkRequest {
	s.AuthParametersShrink = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) SetConnectionName(v string) *UpdateConnectionShrinkRequest {
	s.ConnectionName = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) SetDescription(v string) *UpdateConnectionShrinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) SetNetworkParametersShrink(v string) *UpdateConnectionShrinkRequest {
	s.NetworkParametersShrink = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) SetParametersShrink(v string) *UpdateConnectionShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) SetType(v string) *UpdateConnectionShrinkRequest {
	s.Type = &v
	return s
}

func (s *UpdateConnectionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
