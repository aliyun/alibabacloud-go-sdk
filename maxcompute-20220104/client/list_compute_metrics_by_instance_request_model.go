// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeMetricsByInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *ListComputeMetricsByInstanceRequest
	GetEndDate() *int64
	SetInstanceId(v string) *ListComputeMetricsByInstanceRequest
	GetInstanceId() *string
	SetJobOwner(v string) *ListComputeMetricsByInstanceRequest
	GetJobOwner() *string
	SetPageNumber(v int64) *ListComputeMetricsByInstanceRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListComputeMetricsByInstanceRequest
	GetPageSize() *int64
	SetProjectNames(v []*string) *ListComputeMetricsByInstanceRequest
	GetProjectNames() []*string
	SetRegion(v string) *ListComputeMetricsByInstanceRequest
	GetRegion() *string
	SetSignature(v string) *ListComputeMetricsByInstanceRequest
	GetSignature() *string
	SetSpecCodes(v []*string) *ListComputeMetricsByInstanceRequest
	GetSpecCodes() []*string
	SetStartDate(v int64) *ListComputeMetricsByInstanceRequest
	GetStartDate() *int64
	SetTypes(v []*string) *ListComputeMetricsByInstanceRequest
	GetTypes() []*string
}

type ListComputeMetricsByInstanceRequest struct {
	// The end time for the period.
	//
	// example:
	//
	// 1718590596556
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// The job(instance) ID.
	//
	// example:
	//
	// 20240730****ddlr
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// The Alibaba Cloud account that is used to run the MaxCompute job.
	//
	// example:
	//
	// ALIYUN$7632***@aliyun.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// The name of MaxCompute project.
	ProjectNames []*string `json:"projectNames,omitempty" xml:"projectNames,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"region,omitempty" xml:"region,omitempty"`
	// The signature of the SQL job.
	//
	// example:
	//
	// ghijkl789012
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// Specification types.
	SpecCodes []*string `json:"specCodes,omitempty" xml:"specCodes,omitempty" type:"Repeated"`
	// The start time for the period.
	//
	// example:
	//
	// 1715393576201
	StartDate *int64 `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// Metering types.
	Types []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s ListComputeMetricsByInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsByInstanceRequest) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *ListComputeMetricsByInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeMetricsByInstanceRequest) GetJobOwner() *string {
	return s.JobOwner
}

func (s *ListComputeMetricsByInstanceRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListComputeMetricsByInstanceRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListComputeMetricsByInstanceRequest) GetProjectNames() []*string {
	return s.ProjectNames
}

func (s *ListComputeMetricsByInstanceRequest) GetRegion() *string {
	return s.Region
}

func (s *ListComputeMetricsByInstanceRequest) GetSignature() *string {
	return s.Signature
}

func (s *ListComputeMetricsByInstanceRequest) GetSpecCodes() []*string {
	return s.SpecCodes
}

func (s *ListComputeMetricsByInstanceRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *ListComputeMetricsByInstanceRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListComputeMetricsByInstanceRequest) SetEndDate(v int64) *ListComputeMetricsByInstanceRequest {
	s.EndDate = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetInstanceId(v string) *ListComputeMetricsByInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetJobOwner(v string) *ListComputeMetricsByInstanceRequest {
	s.JobOwner = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetPageNumber(v int64) *ListComputeMetricsByInstanceRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetPageSize(v int64) *ListComputeMetricsByInstanceRequest {
	s.PageSize = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetProjectNames(v []*string) *ListComputeMetricsByInstanceRequest {
	s.ProjectNames = v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetRegion(v string) *ListComputeMetricsByInstanceRequest {
	s.Region = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetSignature(v string) *ListComputeMetricsByInstanceRequest {
	s.Signature = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetSpecCodes(v []*string) *ListComputeMetricsByInstanceRequest {
	s.SpecCodes = v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetStartDate(v int64) *ListComputeMetricsByInstanceRequest {
	s.StartDate = &v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) SetTypes(v []*string) *ListComputeMetricsByInstanceRequest {
	s.Types = v
	return s
}

func (s *ListComputeMetricsByInstanceRequest) Validate() error {
	return dara.Validate(s)
}
