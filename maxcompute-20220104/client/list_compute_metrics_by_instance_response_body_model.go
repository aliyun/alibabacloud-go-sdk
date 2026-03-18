// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeMetricsByInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListComputeMetricsByInstanceResponseBodyData) *ListComputeMetricsByInstanceResponseBody
	GetData() *ListComputeMetricsByInstanceResponseBodyData
	SetErrorCode(v string) *ListComputeMetricsByInstanceResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *ListComputeMetricsByInstanceResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *ListComputeMetricsByInstanceResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListComputeMetricsByInstanceResponseBody
	GetRequestId() *string
}

type ListComputeMetricsByInstanceResponseBody struct {
	Data      *ListComputeMetricsByInstanceResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode *string                                       `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                                       `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                                        `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListComputeMetricsByInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsByInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceResponseBody) GetData() *ListComputeMetricsByInstanceResponseBodyData {
	return s.Data
}

func (s *ListComputeMetricsByInstanceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListComputeMetricsByInstanceResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *ListComputeMetricsByInstanceResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListComputeMetricsByInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeMetricsByInstanceResponseBody) SetData(v *ListComputeMetricsByInstanceResponseBodyData) *ListComputeMetricsByInstanceResponseBody {
	s.Data = v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) SetErrorCode(v string) *ListComputeMetricsByInstanceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) SetErrorMsg(v string) *ListComputeMetricsByInstanceResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) SetHttpCode(v int32) *ListComputeMetricsByInstanceResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) SetRequestId(v string) *ListComputeMetricsByInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComputeMetricsByInstanceResponseBodyData struct {
	InstanceComputeMetrics []*ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics `json:"instanceComputeMetrics,omitempty" xml:"instanceComputeMetrics,omitempty" type:"Repeated"`
	PageNumber             *int64                                                                `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	PageSize               *int64                                                                `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	TotalCount             *int64                                                                `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListComputeMetricsByInstanceResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsByInstanceResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceResponseBodyData) GetInstanceComputeMetrics() []*ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	return s.InstanceComputeMetrics
}

func (s *ListComputeMetricsByInstanceResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListComputeMetricsByInstanceResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListComputeMetricsByInstanceResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListComputeMetricsByInstanceResponseBodyData) SetInstanceComputeMetrics(v []*ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) *ListComputeMetricsByInstanceResponseBodyData {
	s.InstanceComputeMetrics = v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyData) SetPageNumber(v int64) *ListComputeMetricsByInstanceResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyData) SetPageSize(v int64) *ListComputeMetricsByInstanceResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyData) SetTotalCount(v int64) *ListComputeMetricsByInstanceResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyData) Validate() error {
	if s.InstanceComputeMetrics != nil {
		for _, item := range s.InstanceComputeMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics struct {
	EndTime     *int64   `json:"endTime,omitempty" xml:"endTime,omitempty"`
	InstanceId  *string  `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	JobOwner    *string  `json:"jobOwner,omitempty" xml:"jobOwner,omitempty"`
	ProjectName *string  `json:"projectName,omitempty" xml:"projectName,omitempty"`
	Signature   *string  `json:"signature,omitempty" xml:"signature,omitempty"`
	SpecCode    *string  `json:"specCode,omitempty" xml:"specCode,omitempty"`
	SubmitTime  *int64   `json:"submitTime,omitempty" xml:"submitTime,omitempty"`
	Type        *string  `json:"type,omitempty" xml:"type,omitempty"`
	Unit        *string  `json:"unit,omitempty" xml:"unit,omitempty"`
	Usage       *float64 `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetJobOwner() *string {
	return s.JobOwner
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetProjectName() *string {
	return s.ProjectName
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetSignature() *string {
	return s.Signature
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetSpecCode() *string {
	return s.SpecCode
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetSubmitTime() *int64 {
	return s.SubmitTime
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetType() *string {
	return s.Type
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetUnit() *string {
	return s.Unit
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) GetUsage() *float64 {
	return s.Usage
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetEndTime(v int64) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.EndTime = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetInstanceId(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.InstanceId = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetJobOwner(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.JobOwner = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetProjectName(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.ProjectName = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetSignature(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.Signature = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetSpecCode(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.SpecCode = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetSubmitTime(v int64) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.SubmitTime = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetType(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.Type = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetUnit(v string) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.Unit = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) SetUsage(v float64) *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics {
	s.Usage = &v
	return s
}

func (s *ListComputeMetricsByInstanceResponseBodyDataInstanceComputeMetrics) Validate() error {
	return dara.Validate(s)
}
