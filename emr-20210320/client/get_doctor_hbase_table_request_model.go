// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHBaseTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetDoctorHBaseTableRequest
	GetClusterId() *string
	SetDateTime(v string) *GetDoctorHBaseTableRequest
	GetDateTime() *string
	SetRegionId(v string) *GetDoctorHBaseTableRequest
	GetRegionId() *string
	SetTableName(v string) *GetDoctorHBaseTableRequest
	GetTableName() *string
}

type GetDoctorHBaseTableRequest struct {
	// Cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Date.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2023-01-01
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Table name.
	//
	// This parameter is required.
	//
	// example:
	//
	// namespace1:tb_item
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s GetDoctorHBaseTableRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHBaseTableRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHBaseTableRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorHBaseTableRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *GetDoctorHBaseTableRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorHBaseTableRequest) GetTableName() *string {
	return s.TableName
}

func (s *GetDoctorHBaseTableRequest) SetClusterId(v string) *GetDoctorHBaseTableRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorHBaseTableRequest) SetDateTime(v string) *GetDoctorHBaseTableRequest {
	s.DateTime = &v
	return s
}

func (s *GetDoctorHBaseTableRequest) SetRegionId(v string) *GetDoctorHBaseTableRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorHBaseTableRequest) SetTableName(v string) *GetDoctorHBaseTableRequest {
	s.TableName = &v
	return s
}

func (s *GetDoctorHBaseTableRequest) Validate() error {
	return dara.Validate(s)
}
