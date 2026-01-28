// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateConnectDatasourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *OperateConnectDatasourceRequest
	GetClusterId() *string
	SetConnectionParams(v string) *OperateConnectDatasourceRequest
	GetConnectionParams() *string
	SetPassword(v string) *OperateConnectDatasourceRequest
	GetPassword() *string
	SetType(v int32) *OperateConnectDatasourceRequest
	GetType() *int32
}

type OperateConnectDatasourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// {"host":"rm-bp1f03mxxxxx.mysql.rds.aliyuncs.com","port":3306,"userName":"test01","database":"test01","other":{"useSSL":"false"}}
	ConnectionParams *string `json:"ConnectionParams,omitempty" xml:"ConnectionParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ***
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OperateConnectDatasourceRequest) String() string {
	return dara.Prettify(s)
}

func (s OperateConnectDatasourceRequest) GoString() string {
	return s.String()
}

func (s *OperateConnectDatasourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *OperateConnectDatasourceRequest) GetConnectionParams() *string {
	return s.ConnectionParams
}

func (s *OperateConnectDatasourceRequest) GetPassword() *string {
	return s.Password
}

func (s *OperateConnectDatasourceRequest) GetType() *int32 {
	return s.Type
}

func (s *OperateConnectDatasourceRequest) SetClusterId(v string) *OperateConnectDatasourceRequest {
	s.ClusterId = &v
	return s
}

func (s *OperateConnectDatasourceRequest) SetConnectionParams(v string) *OperateConnectDatasourceRequest {
	s.ConnectionParams = &v
	return s
}

func (s *OperateConnectDatasourceRequest) SetPassword(v string) *OperateConnectDatasourceRequest {
	s.Password = &v
	return s
}

func (s *OperateConnectDatasourceRequest) SetType(v int32) *OperateConnectDatasourceRequest {
	s.Type = &v
	return s
}

func (s *OperateConnectDatasourceRequest) Validate() error {
	return dara.Validate(s)
}
