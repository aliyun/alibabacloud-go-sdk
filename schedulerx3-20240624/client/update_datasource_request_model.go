// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasourceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *UpdateDatasourceRequest
	GetClusterId() *string
	SetConnectionParams(v string) *UpdateDatasourceRequest
	GetConnectionParams() *string
	SetDatasourceId(v int64) *UpdateDatasourceRequest
	GetDatasourceId() *int64
	SetDescription(v string) *UpdateDatasourceRequest
	GetDescription() *string
	SetName(v string) *UpdateDatasourceRequest
	GetName() *string
	SetPassword(v string) *UpdateDatasourceRequest
	GetPassword() *string
}

type UpdateDatasourceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// {"host":"rm-bp1f03mxxxxx.mysql.rds.aliyuncs.com","port":3306,"userName":"test01","database":"test01","other":{"useSSL":"false"}}
	ConnectionParams *string `json:"ConnectionParams,omitempty" xml:"ConnectionParams,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DatasourceId *int64 `json:"DatasourceId,omitempty" xml:"DatasourceId,omitempty"`
	// example:
	//
	// None
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 2F9e9@a69c!e18b569c8
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
}

func (s UpdateDatasourceRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasourceRequest) GoString() string {
	return s.String()
}

func (s *UpdateDatasourceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *UpdateDatasourceRequest) GetConnectionParams() *string {
	return s.ConnectionParams
}

func (s *UpdateDatasourceRequest) GetDatasourceId() *int64 {
	return s.DatasourceId
}

func (s *UpdateDatasourceRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateDatasourceRequest) GetName() *string {
	return s.Name
}

func (s *UpdateDatasourceRequest) GetPassword() *string {
	return s.Password
}

func (s *UpdateDatasourceRequest) SetClusterId(v string) *UpdateDatasourceRequest {
	s.ClusterId = &v
	return s
}

func (s *UpdateDatasourceRequest) SetConnectionParams(v string) *UpdateDatasourceRequest {
	s.ConnectionParams = &v
	return s
}

func (s *UpdateDatasourceRequest) SetDatasourceId(v int64) *UpdateDatasourceRequest {
	s.DatasourceId = &v
	return s
}

func (s *UpdateDatasourceRequest) SetDescription(v string) *UpdateDatasourceRequest {
	s.Description = &v
	return s
}

func (s *UpdateDatasourceRequest) SetName(v string) *UpdateDatasourceRequest {
	s.Name = &v
	return s
}

func (s *UpdateDatasourceRequest) SetPassword(v string) *UpdateDatasourceRequest {
	s.Password = &v
	return s
}

func (s *UpdateDatasourceRequest) Validate() error {
	return dara.Validate(s)
}
