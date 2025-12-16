// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataCollctionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDataCollctionResponseBody
	GetRequestId() *string
	SetResult(v *DescribeDataCollctionResponseBodyResult) *DescribeDataCollctionResponseBody
	GetResult() *DescribeDataCollctionResponseBodyResult
}

type DescribeDataCollctionResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 72FAD77B-83F9-F393-BA8E-5834E2427BF8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The details of the data collection task.
	Result *DescribeDataCollctionResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s DescribeDataCollctionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCollctionResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataCollctionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataCollctionResponseBody) GetResult() *DescribeDataCollctionResponseBodyResult {
	return s.Result
}

func (s *DescribeDataCollctionResponseBody) SetRequestId(v string) *DescribeDataCollctionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataCollctionResponseBody) SetResult(v *DescribeDataCollctionResponseBodyResult) *DescribeDataCollctionResponseBody {
	s.Result = v
	return s
}

func (s *DescribeDataCollctionResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataCollctionResponseBodyResult struct {
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
	// The ID of the data collection task.
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
	// 	- app Note: Only server is supported.
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

func (s DescribeDataCollctionResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataCollctionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDataCollctionResponseBodyResult) GetCreated() *int32 {
	return s.Created
}

func (s *DescribeDataCollctionResponseBodyResult) GetDataCollectionType() *string {
	return s.DataCollectionType
}

func (s *DescribeDataCollctionResponseBodyResult) GetId() *string {
	return s.Id
}

func (s *DescribeDataCollctionResponseBodyResult) GetIndustryName() *string {
	return s.IndustryName
}

func (s *DescribeDataCollctionResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *DescribeDataCollctionResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeDataCollctionResponseBodyResult) GetSundialId() *string {
	return s.SundialId
}

func (s *DescribeDataCollctionResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *DescribeDataCollctionResponseBodyResult) GetUpdated() *int32 {
	return s.Updated
}

func (s *DescribeDataCollctionResponseBodyResult) SetCreated(v int32) *DescribeDataCollctionResponseBodyResult {
	s.Created = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetDataCollectionType(v string) *DescribeDataCollctionResponseBodyResult {
	s.DataCollectionType = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetId(v string) *DescribeDataCollctionResponseBodyResult {
	s.Id = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetIndustryName(v string) *DescribeDataCollctionResponseBodyResult {
	s.IndustryName = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetName(v string) *DescribeDataCollctionResponseBodyResult {
	s.Name = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetStatus(v int32) *DescribeDataCollctionResponseBodyResult {
	s.Status = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetSundialId(v string) *DescribeDataCollctionResponseBodyResult {
	s.SundialId = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetType(v string) *DescribeDataCollctionResponseBodyResult {
	s.Type = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) SetUpdated(v int32) *DescribeDataCollctionResponseBodyResult {
	s.Updated = &v
	return s
}

func (s *DescribeDataCollctionResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
