// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeJobsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListComputeJobsResponseBody
	GetCode() *int64
	SetData(v []*ListComputeJobsResponseBodyData) *ListComputeJobsResponseBody
	GetData() []*ListComputeJobsResponseBodyData
	SetMaxResults(v int32) *ListComputeJobsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListComputeJobsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListComputeJobsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListComputeJobsResponseBody
	GetSuccess() *bool
	SetTotal(v int64) *ListComputeJobsResponseBody
	GetTotal() *int64
}

type ListComputeJobsResponseBody struct {
	Code       *int64                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListComputeJobsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	MaxResults *int32                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
	Total      *int64                             `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListComputeJobsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeJobsResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeJobsResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListComputeJobsResponseBody) GetData() []*ListComputeJobsResponseBodyData {
	return s.Data
}

func (s *ListComputeJobsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComputeJobsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComputeJobsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeJobsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListComputeJobsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListComputeJobsResponseBody) SetCode(v int64) *ListComputeJobsResponseBody {
	s.Code = &v
	return s
}

func (s *ListComputeJobsResponseBody) SetData(v []*ListComputeJobsResponseBodyData) *ListComputeJobsResponseBody {
	s.Data = v
	return s
}

func (s *ListComputeJobsResponseBody) SetMaxResults(v int32) *ListComputeJobsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListComputeJobsResponseBody) SetNextToken(v string) *ListComputeJobsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListComputeJobsResponseBody) SetRequestId(v string) *ListComputeJobsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeJobsResponseBody) SetSuccess(v bool) *ListComputeJobsResponseBody {
	s.Success = &v
	return s
}

func (s *ListComputeJobsResponseBody) SetTotal(v int64) *ListComputeJobsResponseBody {
	s.Total = &v
	return s
}

func (s *ListComputeJobsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComputeJobsResponseBodyData struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	CreateTime *string  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CuLimit    *float64 `json:"CuLimit,omitempty" xml:"CuLimit,omitempty"`
	CuReserved *float64 `json:"CuReserved,omitempty" xml:"CuReserved,omitempty"`
	CuUsed     *float64 `json:"CuUsed,omitempty" xml:"CuUsed,omitempty"`
	DebugMode  *int32   `json:"DebugMode,omitempty" xml:"DebugMode,omitempty"`
	InstanceId *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	JobName    *string  `json:"JobName,omitempty" xml:"JobName,omitempty"`
	Owner      *string  `json:"Owner,omitempty" xml:"Owner,omitempty"`
	RegionId   *string  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Remark     *string  `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Status     *string  `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListComputeJobsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListComputeJobsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComputeJobsResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListComputeJobsResponseBodyData) GetCuLimit() *float64 {
	return s.CuLimit
}

func (s *ListComputeJobsResponseBodyData) GetCuReserved() *float64 {
	return s.CuReserved
}

func (s *ListComputeJobsResponseBodyData) GetCuUsed() *float64 {
	return s.CuUsed
}

func (s *ListComputeJobsResponseBodyData) GetDebugMode() *int32 {
	return s.DebugMode
}

func (s *ListComputeJobsResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeJobsResponseBodyData) GetJobName() *string {
	return s.JobName
}

func (s *ListComputeJobsResponseBodyData) GetOwner() *string {
	return s.Owner
}

func (s *ListComputeJobsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *ListComputeJobsResponseBodyData) GetRemark() *string {
	return s.Remark
}

func (s *ListComputeJobsResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *ListComputeJobsResponseBodyData) SetCreateTime(v string) *ListComputeJobsResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListComputeJobsResponseBodyData) SetCuLimit(v float64) *ListComputeJobsResponseBodyData {
	s.CuLimit = &v
	return s
}

func (s *ListComputeJobsResponseBodyData) SetCuReserved(v float64) *ListComputeJobsResponseBodyData {
	s.CuReserved = &v
	return s
}

func (s *ListComputeJobsResponseBodyData) SetCuUsed(v float64) *ListComputeJobsResponseBodyData {
	s.CuUsed = &v
	return s
}

func (s *ListComputeJobsResponseBodyData) SetDebugMode(v int32) *ListComputeJobsResponseBodyData {
	s.DebugMode = &v
	return s
}

func (s *ListComputeJobsResponseBodyData) SetInstanceId(v string) *ListComputeJobsResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListComputeJobsResponseBodyData) SetJobName(v string) *ListComputeJobsResponseBodyData {
	s.JobName = &v
	return s
}

func (s *ListComputeJobsResponseBodyData) SetOwner(v string) *ListComputeJobsResponseBodyData {
	s.Owner = &v
	return s
}

func (s *ListComputeJobsResponseBodyData) SetRegionId(v string) *ListComputeJobsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *ListComputeJobsResponseBodyData) SetRemark(v string) *ListComputeJobsResponseBodyData {
	s.Remark = &v
	return s
}

func (s *ListComputeJobsResponseBodyData) SetStatus(v string) *ListComputeJobsResponseBodyData {
	s.Status = &v
	return s
}

func (s *ListComputeJobsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
