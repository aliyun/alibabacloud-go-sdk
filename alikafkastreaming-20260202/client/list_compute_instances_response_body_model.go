// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *ListComputeInstancesResponseBody
	GetCode() *int64
	SetData(v []*ListComputeInstancesResponseBodyData) *ListComputeInstancesResponseBody
	GetData() []*ListComputeInstancesResponseBodyData
	SetMaxResults(v int32) *ListComputeInstancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListComputeInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListComputeInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListComputeInstancesResponseBody
	GetSuccess() *bool
}

type ListComputeInstancesResponseBody struct {
	Code       *int64                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*ListComputeInstancesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	MaxResults *int32                                  `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken  *string                                 `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListComputeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeInstancesResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *ListComputeInstancesResponseBody) GetData() []*ListComputeInstancesResponseBodyData {
	return s.Data
}

func (s *ListComputeInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListComputeInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListComputeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListComputeInstancesResponseBody) SetCode(v int64) *ListComputeInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListComputeInstancesResponseBody) SetData(v []*ListComputeInstancesResponseBodyData) *ListComputeInstancesResponseBody {
	s.Data = v
	return s
}

func (s *ListComputeInstancesResponseBody) SetMaxResults(v int32) *ListComputeInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListComputeInstancesResponseBody) SetNextToken(v string) *ListComputeInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListComputeInstancesResponseBody) SetRequestId(v string) *ListComputeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeInstancesResponseBody) SetSuccess(v bool) *ListComputeInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListComputeInstancesResponseBody) Validate() error {
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

type ListComputeInstancesResponseBodyData struct {
	// Use the UTC time format: yyyy-MM-ddTHH:mm:ssZ
	CreateTime       *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Cu               *int32  `json:"Cu,omitempty" xml:"Cu,omitempty"`
	InstanceId       *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName     *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	ServiceStatus    *string `json:"ServiceStatus,omitempty" xml:"ServiceStatus,omitempty"`
	TotalJobs        *int64  `json:"TotalJobs,omitempty" xml:"TotalJobs,omitempty"`
	TotalRunningJobs *int64  `json:"TotalRunningJobs,omitempty" xml:"TotalRunningJobs,omitempty"`
}

func (s ListComputeInstancesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListComputeInstancesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComputeInstancesResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListComputeInstancesResponseBodyData) GetCu() *int32 {
	return s.Cu
}

func (s *ListComputeInstancesResponseBodyData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeInstancesResponseBodyData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListComputeInstancesResponseBodyData) GetServiceStatus() *string {
	return s.ServiceStatus
}

func (s *ListComputeInstancesResponseBodyData) GetTotalJobs() *int64 {
	return s.TotalJobs
}

func (s *ListComputeInstancesResponseBodyData) GetTotalRunningJobs() *int64 {
	return s.TotalRunningJobs
}

func (s *ListComputeInstancesResponseBodyData) SetCreateTime(v string) *ListComputeInstancesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *ListComputeInstancesResponseBodyData) SetCu(v int32) *ListComputeInstancesResponseBodyData {
	s.Cu = &v
	return s
}

func (s *ListComputeInstancesResponseBodyData) SetInstanceId(v string) *ListComputeInstancesResponseBodyData {
	s.InstanceId = &v
	return s
}

func (s *ListComputeInstancesResponseBodyData) SetInstanceName(v string) *ListComputeInstancesResponseBodyData {
	s.InstanceName = &v
	return s
}

func (s *ListComputeInstancesResponseBodyData) SetServiceStatus(v string) *ListComputeInstancesResponseBodyData {
	s.ServiceStatus = &v
	return s
}

func (s *ListComputeInstancesResponseBodyData) SetTotalJobs(v int64) *ListComputeInstancesResponseBodyData {
	s.TotalJobs = &v
	return s
}

func (s *ListComputeInstancesResponseBodyData) SetTotalRunningJobs(v int64) *ListComputeInstancesResponseBodyData {
	s.TotalRunningJobs = &v
	return s
}

func (s *ListComputeInstancesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
