// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHiveTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetDoctorHiveTableRequest
	GetClusterId() *string
	SetDateTime(v string) *GetDoctorHiveTableRequest
	GetDateTime() *string
	SetRegionId(v string) *GetDoctorHiveTableRequest
	GetRegionId() *string
	SetTableName(v string) *GetDoctorHiveTableRequest
	GetTableName() *string
}

type GetDoctorHiveTableRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Specify the date in the ISO 8601 standard. For example, 2023-01-01 represents January 1, 2023.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The table name. The table name must follow the rule in Hive. A name in the {database name.table identifier} format uniquely identifies a table.
	//
	// This parameter is required.
	//
	// example:
	//
	// dw.dwd_creta_service_order_long_renew_long_da
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetDoctorHiveTableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveTableRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveTableRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorHiveTableRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *GetDoctorHiveTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorHiveTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetDoctorHiveTableRequest) SetClusterId(v string) *GetDoctorHiveTableRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorHiveTableRequest) SetDateTime(v string) *GetDoctorHiveTableRequest {
	s.DateTime = &v
	return s
}

func (s *GetDoctorHiveTableRequest) SetRegionId(v string) *GetDoctorHiveTableRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorHiveTableRequest) SetTableName(v string) *GetDoctorHiveTableRequest {
	s.TableName = &v
	return s
}

func (s *GetDoctorHiveTableRequest) Validate() error {
	return dara.Validate(s)
}
