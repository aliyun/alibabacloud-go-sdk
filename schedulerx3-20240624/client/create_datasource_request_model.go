// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateDatasourceRequest
	GetClusterId() *string
	SetConnectionParams(v string) *CreateDatasourceRequest
	GetConnectionParams() *string
	SetDescription(v string) *CreateDatasourceRequest
	GetDescription() *string
	SetName(v string) *CreateDatasourceRequest
	GetName() *string
	SetPassword(v string) *CreateDatasourceRequest
	GetPassword() *string
	SetType(v int32) *CreateDatasourceRequest
	GetType() *int32
}

type CreateDatasourceRequest struct {
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
	// example:
	//
	// my first datasource
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// myDatasouce
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
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
	// 0
	Type *int32 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateDatasourceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasourceRequest) GoString() string {
	return s.String()
}

func (s *CreateDatasourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateDatasourceRequest) GetConnectionParams() *string {
	return s.ConnectionParams
}

func (s *CreateDatasourceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDatasourceRequest) GetName() *string {
	return s.Name
}

func (s *CreateDatasourceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateDatasourceRequest) GetType() *int32 {
	return s.Type
}

func (s *CreateDatasourceRequest) SetClusterId(v string) *CreateDatasourceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateDatasourceRequest) SetConnectionParams(v string) *CreateDatasourceRequest {
	s.ConnectionParams = &v
	return s
}

func (s *CreateDatasourceRequest) SetDescription(v string) *CreateDatasourceRequest {
	s.Description = &v
	return s
}

func (s *CreateDatasourceRequest) SetName(v string) *CreateDatasourceRequest {
	s.Name = &v
	return s
}

func (s *CreateDatasourceRequest) SetPassword(v string) *CreateDatasourceRequest {
	s.Password = &v
	return s
}

func (s *CreateDatasourceRequest) SetType(v int32) *CreateDatasourceRequest {
	s.Type = &v
	return s
}

func (s *CreateDatasourceRequest) Validate() error {
	return dara.Validate(s)
}
