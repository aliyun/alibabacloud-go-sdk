// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSumStorageMetricsByDateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SumStorageMetricsByDateResponseBodyData) *SumStorageMetricsByDateResponseBody
	GetData() []*SumStorageMetricsByDateResponseBodyData
	SetErrorCode(v string) *SumStorageMetricsByDateResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *SumStorageMetricsByDateResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *SumStorageMetricsByDateResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *SumStorageMetricsByDateResponseBody
	GetRequestId() *string
}

type SumStorageMetricsByDateResponseBody struct {
	Data []*SumStorageMetricsByDateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// this quota is not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// example:
	//
	// 0abb781a17411408145995819e0dae
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s SumStorageMetricsByDateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SumStorageMetricsByDateResponseBody) GoString() string {
	return s.String()
}

func (s *SumStorageMetricsByDateResponseBody) GetData() []*SumStorageMetricsByDateResponseBodyData {
	return s.Data
}

func (s *SumStorageMetricsByDateResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SumStorageMetricsByDateResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *SumStorageMetricsByDateResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *SumStorageMetricsByDateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SumStorageMetricsByDateResponseBody) SetData(v []*SumStorageMetricsByDateResponseBodyData) *SumStorageMetricsByDateResponseBody {
	s.Data = v
	return s
}

func (s *SumStorageMetricsByDateResponseBody) SetErrorCode(v string) *SumStorageMetricsByDateResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBody) SetErrorMsg(v string) *SumStorageMetricsByDateResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBody) SetHttpCode(v int32) *SumStorageMetricsByDateResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBody) SetRequestId(v string) *SumStorageMetricsByDateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBody) Validate() error {
	return dara.Validate(s)
}

type SumStorageMetricsByDateResponseBodyData struct {
	// example:
	//
	// 20250719
	DateTime           *string                                                      `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	ItemStorageMetrics []*SumStorageMetricsByDateResponseBodyDataItemStorageMetrics `json:"itemStorageMetrics,omitempty" xml:"itemStorageMetrics,omitempty" type:"Repeated"`
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
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s SumStorageMetricsByDateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SumStorageMetricsByDateResponseBodyData) GoString() string {
	return s.String()
}

func (s *SumStorageMetricsByDateResponseBodyData) GetDateTime() *string {
	return s.DateTime
}

func (s *SumStorageMetricsByDateResponseBodyData) GetItemStorageMetrics() []*SumStorageMetricsByDateResponseBodyDataItemStorageMetrics {
	return s.ItemStorageMetrics
}

func (s *SumStorageMetricsByDateResponseBodyData) GetStorageType() *string {
	return s.StorageType
}

func (s *SumStorageMetricsByDateResponseBodyData) GetUnit() *string {
	return s.Unit
}

func (s *SumStorageMetricsByDateResponseBodyData) GetUsage() *string {
	return s.Usage
}

func (s *SumStorageMetricsByDateResponseBodyData) SetDateTime(v string) *SumStorageMetricsByDateResponseBodyData {
	s.DateTime = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBodyData) SetItemStorageMetrics(v []*SumStorageMetricsByDateResponseBodyDataItemStorageMetrics) *SumStorageMetricsByDateResponseBodyData {
	s.ItemStorageMetrics = v
	return s
}

func (s *SumStorageMetricsByDateResponseBodyData) SetStorageType(v string) *SumStorageMetricsByDateResponseBodyData {
	s.StorageType = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBodyData) SetUnit(v string) *SumStorageMetricsByDateResponseBodyData {
	s.Unit = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBodyData) SetUsage(v string) *SumStorageMetricsByDateResponseBodyData {
	s.Usage = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type SumStorageMetricsByDateResponseBodyDataItemStorageMetrics struct {
	// example:
	//
	// prj
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// example:
	//
	// 91.22
	Percentage *float64 `json:"percentage,omitempty" xml:"percentage,omitempty"`
	// example:
	//
	// 300.560392
	Usage *string `json:"usage,omitempty" xml:"usage,omitempty"`
}

func (s SumStorageMetricsByDateResponseBodyDataItemStorageMetrics) String() string {
	return dara.Prettify(s)
}

func (s SumStorageMetricsByDateResponseBodyDataItemStorageMetrics) GoString() string {
	return s.String()
}

func (s *SumStorageMetricsByDateResponseBodyDataItemStorageMetrics) GetItemName() *string {
	return s.ItemName
}

func (s *SumStorageMetricsByDateResponseBodyDataItemStorageMetrics) GetPercentage() *float64 {
	return s.Percentage
}

func (s *SumStorageMetricsByDateResponseBodyDataItemStorageMetrics) GetUsage() *string {
	return s.Usage
}

func (s *SumStorageMetricsByDateResponseBodyDataItemStorageMetrics) SetItemName(v string) *SumStorageMetricsByDateResponseBodyDataItemStorageMetrics {
	s.ItemName = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBodyDataItemStorageMetrics) SetPercentage(v float64) *SumStorageMetricsByDateResponseBodyDataItemStorageMetrics {
	s.Percentage = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBodyDataItemStorageMetrics) SetUsage(v string) *SumStorageMetricsByDateResponseBodyDataItemStorageMetrics {
	s.Usage = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBodyDataItemStorageMetrics) Validate() error {
	return dara.Validate(s)
}
