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
	SetHttpCode(v int32) *SumStorageMetricsByDateResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *SumStorageMetricsByDateResponseBody
	GetRequestId() *string
}

type SumStorageMetricsByDateResponseBody struct {
	// The returned data.
	Data []*SumStorageMetricsByDateResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The HTTP status code.
	//
	// - 1xx: Informational. The request is received and the process is continuing.
	//
	// - 2xx: Success. The request is successfully received, understood, and accepted.
	//
	// - 3xx: Redirection. Further action needs to be taken to complete the request.
	//
	// - 4xx: Client Error. The request contains bad syntax or cannot be fulfilled.
	//
	// - 5xx: Server Error. The server fails to fulfill an apparently valid request.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
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

func (s *SumStorageMetricsByDateResponseBody) SetHttpCode(v int32) *SumStorageMetricsByDateResponseBody {
	s.HttpCode = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBody) SetRequestId(v string) *SumStorageMetricsByDateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SumStorageMetricsByDateResponseBody) Validate() error {
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

type SumStorageMetricsByDateResponseBodyData struct {
	// The date of the statistics. The format is yyyyMMdd.
	//
	// example:
	//
	// 20250719
	DateTime *string `json:"dateTime,omitempty" xml:"dateTime,omitempty"`
	// The list of storage usage of a specified type.
	ItemStorageMetrics []*SumStorageMetricsByDateResponseBodyDataItemStorageMetrics `json:"itemStorageMetrics,omitempty" xml:"itemStorageMetrics,omitempty" type:"Repeated"`
	// The storage type. Valid values:
	//
	// - Storage: Standard.
	//
	// - LowFreqStorage: Infrequent Access (IA).
	//
	// - ColdStorage: Archive.
	//
	// example:
	//
	// Storage
	StorageType *string `json:"storageType,omitempty" xml:"storageType,omitempty"`
	// The unit of the total storage.
	//
	// example:
	//
	// GB
	Unit *string `json:"unit,omitempty" xml:"unit,omitempty"`
	// The total storage.
	//
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
	if s.ItemStorageMetrics != nil {
		for _, item := range s.ItemStorageMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SumStorageMetricsByDateResponseBodyDataItemStorageMetrics struct {
	// If the type is PROJECT, this parameter indicates the project name. If the type is STORAGE_TYPE, this parameter indicates the storage type.
	//
	// example:
	//
	// prj
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// The percentage of the storage usage.
	//
	// example:
	//
	// 91.22
	Percentage *float64 `json:"percentage,omitempty" xml:"percentage,omitempty"`
	// The storage usage.
	//
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
