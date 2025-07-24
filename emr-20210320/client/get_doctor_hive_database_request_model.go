// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorHiveDatabaseRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetDoctorHiveDatabaseRequest
	GetClusterId() *string
	SetDatabaseName(v string) *GetDoctorHiveDatabaseRequest
	GetDatabaseName() *string
	SetDateTime(v string) *GetDoctorHiveDatabaseRequest
	GetDateTime() *string
	SetRegionId(v string) *GetDoctorHiveDatabaseRequest
	GetRegionId() *string
}

type GetDoctorHiveDatabaseRequest struct {
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The database name.
	//
	// This parameter is required.
	//
	// example:
	//
	// db1
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The query date.
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
}

func (s GetDoctorHiveDatabaseRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorHiveDatabaseRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorHiveDatabaseRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorHiveDatabaseRequest) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *GetDoctorHiveDatabaseRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *GetDoctorHiveDatabaseRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorHiveDatabaseRequest) SetClusterId(v string) *GetDoctorHiveDatabaseRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorHiveDatabaseRequest) SetDatabaseName(v string) *GetDoctorHiveDatabaseRequest {
	s.DatabaseName = &v
	return s
}

func (s *GetDoctorHiveDatabaseRequest) SetDateTime(v string) *GetDoctorHiveDatabaseRequest {
	s.DateTime = &v
	return s
}

func (s *GetDoctorHiveDatabaseRequest) SetRegionId(v string) *GetDoctorHiveDatabaseRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorHiveDatabaseRequest) Validate() error {
	return dara.Validate(s)
}
