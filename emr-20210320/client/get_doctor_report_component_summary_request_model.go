// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoctorReportComponentSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *GetDoctorReportComponentSummaryRequest
	GetClusterId() *string
	SetComponentType(v string) *GetDoctorReportComponentSummaryRequest
	GetComponentType() *string
	SetDateTime(v string) *GetDoctorReportComponentSummaryRequest
	GetDateTime() *string
	SetRegionId(v string) *GetDoctorReportComponentSummaryRequest
	GetRegionId() *string
}

type GetDoctorReportComponentSummaryRequest struct {
	// Cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-b933c5aac8fe****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Select component filter type. Values:
	//
	// - compute
	//
	// - hive
	//
	// - hdfs
	//
	// - yarn
	//
	// - oss
	//
	// - hbase
	//
	// This parameter is required.
	//
	// example:
	//
	// compute
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// Report date.
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
}

func (s GetDoctorReportComponentSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoctorReportComponentSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetDoctorReportComponentSummaryRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDoctorReportComponentSummaryRequest) GetComponentType() *string {
	return s.ComponentType
}

func (s *GetDoctorReportComponentSummaryRequest) GetDateTime() *string {
	return s.DateTime
}

func (s *GetDoctorReportComponentSummaryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetDoctorReportComponentSummaryRequest) SetClusterId(v string) *GetDoctorReportComponentSummaryRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDoctorReportComponentSummaryRequest) SetComponentType(v string) *GetDoctorReportComponentSummaryRequest {
	s.ComponentType = &v
	return s
}

func (s *GetDoctorReportComponentSummaryRequest) SetDateTime(v string) *GetDoctorReportComponentSummaryRequest {
	s.DateTime = &v
	return s
}

func (s *GetDoctorReportComponentSummaryRequest) SetRegionId(v string) *GetDoctorReportComponentSummaryRequest {
	s.RegionId = &v
	return s
}

func (s *GetDoctorReportComponentSummaryRequest) Validate() error {
	return dara.Validate(s)
}
