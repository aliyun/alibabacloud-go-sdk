// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataCollectionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListDataCollectionsResponseBody
	GetRequestId() *string
	SetResult(v []*ListDataCollectionsResponseBodyResult) *ListDataCollectionsResponseBody
	GetResult() []*ListDataCollectionsResponseBodyResult
	SetTotalCount(v int32) *ListDataCollectionsResponseBody
	GetTotalCount() *int32
}

type ListDataCollectionsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 959D8782-B130-95EB-86CC-1F6ED447981F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the data collection task.
	//
	// For more information, see [DataCollection](https://help.aliyun.com/document_detail/173605.html).
	Result []*ListDataCollectionsResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ListDataCollectionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataCollectionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataCollectionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataCollectionsResponseBody) GetResult() []*ListDataCollectionsResponseBodyResult {
	return s.Result
}

func (s *ListDataCollectionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListDataCollectionsResponseBody) SetRequestId(v string) *ListDataCollectionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataCollectionsResponseBody) SetResult(v []*ListDataCollectionsResponseBodyResult) *ListDataCollectionsResponseBody {
	s.Result = v
	return s
}

func (s *ListDataCollectionsResponseBody) SetTotalCount(v int32) *ListDataCollectionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListDataCollectionsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataCollectionsResponseBodyResult struct {
	// The time when the task was created.
	//
	// example:
	//
	// 1581065837
	Created *int32 `json:"created,omitempty" xml:"created,omitempty"`
	// The type of data collected. Valid values:
	//
	// 	- behavior: behavioral data.
	//
	// 	- item_info: project information.
	//
	// 	- industry_specific: industry-specific data.
	//
	// example:
	//
	// BEHAVIOR
	DataCollectionType *string `json:"dataCollectionType,omitempty" xml:"dataCollectionType,omitempty"`
	// The data collection ID.
	//
	// example:
	//
	// 286
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The industry name. Valid values:
	//
	// 	- general
	//
	// 	- ecommerce
	//
	// example:
	//
	// GENERAL
	IndustryName *string `json:"industryName,omitempty" xml:"industryName,omitempty"`
	// The name of the data collection task.
	//
	// example:
	//
	// os_function_test_v1
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The status of the data collection feature. Valid values:
	//
	// 	- 0: The feature is disabled.
	//
	// 	- 1: The feature is being enabled.
	//
	// 	- 2: The feature is enabled.
	//
	// 	- 3: The feature failed to be enabled.
	//
	// example:
	//
	// 2
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// The sundial ID.
	//
	// example:
	//
	// 1755
	SundialId *string `json:"sundialId,omitempty" xml:"sundialId,omitempty"`
	// The type of the source from which data was collected. Valid values:
	//
	// 	- server
	//
	// 	- web
	//
	// 	- app
	//
	// Only server is supported.
	//
	// example:
	//
	// server
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The time when the data collection task was updated.
	//
	// example:
	//
	// 1581065904
	Updated *int32 `json:"updated,omitempty" xml:"updated,omitempty"`
}

func (s ListDataCollectionsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListDataCollectionsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListDataCollectionsResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *ListDataCollectionsResponseBodyResult) GetDataCollectionType() *string {
	return s.DataCollectionType
}

func (s *ListDataCollectionsResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *ListDataCollectionsResponseBodyResult) GetIndustryName() *string {
	return s.IndustryName
}

func (s *ListDataCollectionsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *ListDataCollectionsResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *ListDataCollectionsResponseBodyResult) GetSundialId() *string {
	return s.SundialId
}

func (s *ListDataCollectionsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *ListDataCollectionsResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *ListDataCollectionsResponseBodyResult) SetCreated(v int32) *ListDataCollectionsResponseBodyResult {
	s.Created = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetDataCollectionType(v string) *ListDataCollectionsResponseBodyResult {
	s.DataCollectionType = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetId(v string) *ListDataCollectionsResponseBodyResult {
	s.Id = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetIndustryName(v string) *ListDataCollectionsResponseBodyResult {
	s.IndustryName = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetName(v string) *ListDataCollectionsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetStatus(v int32) *ListDataCollectionsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetSundialId(v string) *ListDataCollectionsResponseBodyResult {
	s.SundialId = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetType(v string) *ListDataCollectionsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) SetUpdated(v int32) *ListDataCollectionsResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *ListDataCollectionsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
