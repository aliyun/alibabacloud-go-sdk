// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumStorageMetricsByTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SumStorageMetricsByTypeResponseBodyData) *SumStorageMetricsByTypeResponseBody
	GetData() []*SumStorageMetricsByTypeResponseBodyData
	SetHttpCode(v int32) *SumStorageMetricsByTypeResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *SumStorageMetricsByTypeResponseBody
	GetRequestId() *string
}

type SumStorageMetricsByTypeResponseBody struct {
	Data []*SumStorageMetricsByTypeResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 0a06dc0917476202205161986edbbc
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SumStorageMetricsByTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SumStorageMetricsByTypeResponseBody) GoString() string {
	return s.String()
}

func (s *SumStorageMetricsByTypeResponseBody) GetData() []*SumStorageMetricsByTypeResponseBodyData {
	return s.Data
}

func (s *SumStorageMetricsByTypeResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *SumStorageMetricsByTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SumStorageMetricsByTypeResponseBody) SetData(v []*SumStorageMetricsByTypeResponseBodyData) *SumStorageMetricsByTypeResponseBody {
	s.Data = v
	return s
}

func (s *SumStorageMetricsByTypeResponseBody) SetHttpCode(v int32) *SumStorageMetricsByTypeResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SumStorageMetricsByTypeResponseBody) SetRequestId(v string) *SumStorageMetricsByTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SumStorageMetricsByTypeResponseBody) Validate() error {
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

type SumStorageMetricsByTypeResponseBodyData struct {
	DailyStorageMetrics []*SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics `json:"dailyStorageMetrics,omitempty" xml:"dailyStorageMetrics,omitempty" type:"Repeated"`
	// example:
	//
	// Storage
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
	// example:
	//
	// GB
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// 329.503338
	Usage *float64 `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s SumStorageMetricsByTypeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SumStorageMetricsByTypeResponseBodyData) GoString() string {
	return s.String()
}

func (s *SumStorageMetricsByTypeResponseBodyData) GetDailyStorageMetrics() []*SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics {
	return s.DailyStorageMetrics
}

func (s *SumStorageMetricsByTypeResponseBodyData) GetStorageType() *string {
	return s.StorageType
}

func (s *SumStorageMetricsByTypeResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *SumStorageMetricsByTypeResponseBodyData) GetUsage() *float64 {
	return s.Usage
}

func (s *SumStorageMetricsByTypeResponseBodyData) SetDailyStorageMetrics(v []*SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) *SumStorageMetricsByTypeResponseBodyData {
	s.DailyStorageMetrics = v
	return s
}

func (s *SumStorageMetricsByTypeResponseBodyData) SetStorageType(v string) *SumStorageMetricsByTypeResponseBodyData {
	s.StorageType = &v
	return s
}

func (s *SumStorageMetricsByTypeResponseBodyData) SetUnit(v string) *SumStorageMetricsByTypeResponseBodyData {
	s.Unit = &v
	return s
}

func (s *SumStorageMetricsByTypeResponseBodyData) SetUsage(v float64) *SumStorageMetricsByTypeResponseBodyData {
	s.Usage = &v
	return s
}

func (s *SumStorageMetricsByTypeResponseBodyData) Validate() error {
	if s.DailyStorageMetrics != nil {
		for _, item := range s.DailyStorageMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics struct {
	// example:
	//
	// 20260410
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// example:
	//
	// 50
	Percentage *float64 `json:"percentage,omitempty" xml:"percentage,omitempty"`
	// example:
	//
	// Storage
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
	// example:
	//
	// GB
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// example:
	//
	// 30
	Usage *float64 `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) String() string {
	return dara.Prettify(s)
}

func (s SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) GoString() string {
	return s.String()
}

func (s *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) GetDateTime() *string {
	return s.DateTime
}

func (s *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) GetPercentage() *float64 {
	return s.Percentage
}

func (s *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) GetStorageType() *string {
	return s.StorageType
}

func (s *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) GetUnit() *string {
	return s.Unit
}

func (s *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) GetUsage() *float64 {
	return s.Usage
}

func (s *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) SetDateTime(v string) *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics {
	s.DateTime = &v
	return s
}

func (s *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) SetPercentage(v float64) *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics {
	s.Percentage = &v
	return s
}

func (s *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) SetStorageType(v string) *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics {
	s.StorageType = &v
	return s
}

func (s *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) SetUnit(v string) *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics {
	s.Unit = &v
	return s
}

func (s *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) SetUsage(v float64) *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics {
	s.Usage = &v
	return s
}

func (s *SumStorageMetricsByTypeResponseBodyDataDailyStorageMetrics) Validate() error {
	return dara.Validate(s)
}
