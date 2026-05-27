// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListComputeMetricsBySignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListComputeMetricsBySignatureResponseBodyData) *ListComputeMetricsBySignatureResponseBody
	GetData() *ListComputeMetricsBySignatureResponseBodyData
	SetHttpCode(v int32) *ListComputeMetricsBySignatureResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *ListComputeMetricsBySignatureResponseBody
	GetRequestId() *string
}

type ListComputeMetricsBySignatureResponseBody struct {
	Data *ListComputeMetricsBySignatureResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 0adbef0f17600626304764284d0001
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListComputeMetricsBySignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsBySignatureResponseBody) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsBySignatureResponseBody) GetData() *ListComputeMetricsBySignatureResponseBodyData {
	return s.Data
}

func (s *ListComputeMetricsBySignatureResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *ListComputeMetricsBySignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListComputeMetricsBySignatureResponseBody) SetData(v *ListComputeMetricsBySignatureResponseBodyData) *ListComputeMetricsBySignatureResponseBody {
	s.Data = v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBody) SetHttpCode(v int32) *ListComputeMetricsBySignatureResponseBody {
	s.HttpCode = &v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBody) SetRequestId(v string) *ListComputeMetricsBySignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListComputeMetricsBySignatureResponseBodyData struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize                *int64                                                                  `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	SignatureComputeMetrics []*ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics `json:"signatureComputeMetrics,omitempty" xml:"signatureComputeMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// 57
	TotalCount *int64 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListComputeMetricsBySignatureResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsBySignatureResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsBySignatureResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListComputeMetricsBySignatureResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListComputeMetricsBySignatureResponseBodyData) GetSignatureComputeMetrics() []*ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics {
	return s.SignatureComputeMetrics
}

func (s *ListComputeMetricsBySignatureResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListComputeMetricsBySignatureResponseBodyData) SetPageNumber(v int64) *ListComputeMetricsBySignatureResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBodyData) SetPageSize(v int64) *ListComputeMetricsBySignatureResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBodyData) SetSignatureComputeMetrics(v []*ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) *ListComputeMetricsBySignatureResponseBodyData {
	s.SignatureComputeMetrics = v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBodyData) SetTotalCount(v int64) *ListComputeMetricsBySignatureResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBodyData) Validate() error {
	if s.SignatureComputeMetrics != nil {
		for _, item := range s.SignatureComputeMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics struct {
	Instances    []*ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances `json:"instances,omitempty" xml:"instances,omitempty" type:"Repeated"`
	ProjectNames []*string                                                                        `json:"projectNames,omitempty" xml:"projectNames,omitempty" type:"Repeated"`
	// example:
	//
	// oC0HkG5aTRscH3BDrttrLwHf8XA=
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
	// example:
	//
	// GBCplx
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// 100
	Usage *float64 `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) GetInstances() []*ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances {
	return s.Instances
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) GetProjectNames() []*string {
	return s.ProjectNames
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) GetSignature() *string {
	return s.Signature
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) GetUnit() *string {
	return s.Unit
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) GetUsage() *float64 {
	return s.Usage
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) SetInstances(v []*ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances) *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics {
	s.Instances = v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) SetProjectNames(v []*string) *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics {
	s.ProjectNames = v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) SetSignature(v string) *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics {
	s.Signature = &v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) SetUnit(v string) *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics {
	s.Unit = &v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) SetUsage(v float64) *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics {
	s.Usage = &v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetrics) Validate() error {
	if s.Instances != nil {
		for _, item := range s.Instances {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances struct {
	// example:
	//
	// 1770886999000
	EndTime *int64 `json:"endTime,omitempty" xml:"endTime,omitempty"`
	// example:
	//
	// 20250910185415772gu7vl8vwk22
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// example:
	//
	// 1762946698000
	StartTime *int64 `json:"startTime,omitempty" xml:"startTime,omitempty"`
}

func (s ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances) String() string {
	return dara.Prettify(s)
}

func (s ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances) GoString() string {
	return s.String()
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances) GetEndTime() *int64 {
	return s.EndTime
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances) GetStartTime() *int64 {
	return s.StartTime
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances) SetEndTime(v int64) *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances {
	s.EndTime = &v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances) SetInstanceId(v string) *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances {
	s.InstanceId = &v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances) SetStartTime(v int64) *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances {
	s.StartTime = &v
	return s
}

func (s *ListComputeMetricsBySignatureResponseBodyDataSignatureComputeMetricsInstances) Validate() error {
	return dara.Validate(s)
}
