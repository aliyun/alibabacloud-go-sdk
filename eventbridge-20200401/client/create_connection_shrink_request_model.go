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
	// The parameters that are configured for authentication.
	AuthParametersShrink *string `json:"AuthParameters,omitempty" xml:"AuthParameters,omitempty"`
	// The name of the connection. The name must be 2 to 127 characters in length.
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
	// demo
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameters that are configured for the network.
	//
	// This parameter is required.
	NetworkParametersShrink *string `json:"NetworkParameters,omitempty" xml:"NetworkParameters,omitempty"`
	// 数据源连接参数（JSON 对象），Type 为数据源类型时必填，Http 类型不需要。具体字段定义请调用 GetConnectionType 接口，参考返回结果中的 ParamsSchema
	//
	// example:
	//
	// {"HostName":"xxx.mysql.rds.aliyuncs.com","Port":"3306","User":"root","Password":"xxx","DatabaseName":"demo_db"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// 连接类型。可选值：MySQL、PostgreSQL、Elasticsearch、Http。数据源类型连接必填，不传默认为 Http。Http 类型用于 API Destination 等 HTTP 协议目标；数据源类型用于集成广场的数据连接
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
