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
	// The parameters that are configured for authentication.
	AuthParametersShrink *string `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty"`
	// The name of the connection that you want to update. The name must be 2 to 127 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// connection-name
	ConnectionName *string `json:"ConnectionName,omitempty" xml:"ConnectionName,omitempty"`
	// The description of the connection. The description can be up to 255 characters in length.
	//
	// example:
	//
	// The description of the connection.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the network.
	//
	// This parameter is required.
	NetworkParametersShrink *string `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty"`
	// 数据源连接参数（JSON 对象）。具体字段定义请调用 GetConnectionType 接口，参考返回结果中的 ParamsSchema
	//
	// example:
	//
	// {"HostName":"xxx.mysql.rds.aliyuncs.com","Port":"3306","User":"root","Password":"xxx","DatabaseName":"demo_db"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// 连接类型。可选值：MySQL、PostgreSQL、Elasticsearch、Http
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
