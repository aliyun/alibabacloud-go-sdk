// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeMetricsBySignatureRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndDate(v int64) *ListComputeMetricsBySignatureRequest
	GetEndDate() *int64
	SetInstanceId(v string) *ListComputeMetricsBySignatureRequest
	GetInstanceId() *string
	SetJobOwner(v string) *ListComputeMetricsBySignatureRequest
	GetJobOwner() *string
	SetPageNumber(v int64) *ListComputeMetricsBySignatureRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListComputeMetricsBySignatureRequest
	GetPageSize() *int64
	SetProjectNames(v []*string) *ListComputeMetricsBySignatureRequest
	GetProjectNames() []*string
	SetSignature(v string) *ListComputeMetricsBySignatureRequest
	GetSignature() *string
	SetStartDate(v int64) *ListComputeMetricsBySignatureRequest
	GetStartDate() *int64
	SetTypes(v []*string) *ListComputeMetricsBySignatureRequest
	GetTypes() []*string
}

type ListComputeMetricsBySignatureRequest struct {
	// example:
	//
	// 1718590596556
	EndDate *int64 `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// example:
	//
	// 20240730****ddlr
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// ALIYUN$7632***@aliyun.com
	JobOwner *string `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize     *int64    `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	ProjectNames []*string `json:"projectNames,omitempty" xml:"projectNames,omitempty" type:"Repeated"`
	// example:
	//
	// ghijkl789012
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// example:
	//
	// 1715393576201
	StartDate *int64    `json:"startDate,omitempty" xml:"startDate,omitempty"`
	Types     []*string `json:"types,omitempty" xml:"types,omitempty" type:"Repeated"`
}

func (s ListComputeMetricsBySignatureRequest) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsBySignatureRequest) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsBySignatureRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *ListComputeMetricsBySignatureRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeMetricsBySignatureRequest) GetJobOwner() *string {
	return s.JobOwner
}

func (s *ListComputeMetricsBySignatureRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListComputeMetricsBySignatureRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListComputeMetricsBySignatureRequest) GetProjectNames() []*string {
	return s.ProjectNames
}

func (s *ListComputeMetricsBySignatureRequest) GetSignature() *string {
	return s.Signature
}

func (s *ListComputeMetricsBySignatureRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *ListComputeMetricsBySignatureRequest) GetTypes() []*string {
	return s.Types
}

func (s *ListComputeMetricsBySignatureRequest) SetEndDate(v int64) *ListComputeMetricsBySignatureRequest {
	s.EndDate = &v
	return s
}

func (s *ListComputeMetricsBySignatureRequest) SetInstanceId(v string) *ListComputeMetricsBySignatureRequest {
	s.InstanceId = &v
	return s
}

func (s *ListComputeMetricsBySignatureRequest) SetJobOwner(v string) *ListComputeMetricsBySignatureRequest {
	s.JobOwner = &v
	return s
}

func (s *ListComputeMetricsBySignatureRequest) SetPageNumber(v int64) *ListComputeMetricsBySignatureRequest {
	s.PageNumber = &v
	return s
}

func (s *ListComputeMetricsBySignatureRequest) SetPageSize(v int64) *ListComputeMetricsBySignatureRequest {
	s.PageSize = &v
	return s
}

func (s *ListComputeMetricsBySignatureRequest) SetProjectNames(v []*string) *ListComputeMetricsBySignatureRequest {
	s.ProjectNames = v
	return s
}

func (s *ListComputeMetricsBySignatureRequest) SetSignature(v string) *ListComputeMetricsBySignatureRequest {
	s.Signature = &v
	return s
}

func (s *ListComputeMetricsBySignatureRequest) SetStartDate(v int64) *ListComputeMetricsBySignatureRequest {
	s.StartDate = &v
	return s
}

func (s *ListComputeMetricsBySignatureRequest) SetTypes(v []*string) *ListComputeMetricsBySignatureRequest {
	s.Types = v
	return s
}

func (s *ListComputeMetricsBySignatureRequest) Validate() error {
	return dara.Validate(s)
}
