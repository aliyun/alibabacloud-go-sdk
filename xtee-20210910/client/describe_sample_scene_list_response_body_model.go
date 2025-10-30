// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleSceneListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSampleSceneListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeSampleSceneListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeSampleSceneListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeSampleSceneListResponseBodyResultObject) *DescribeSampleSceneListResponseBody
	GetResultObject() []*DescribeSampleSceneListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeSampleSceneListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeSampleSceneListResponseBody
	GetTotalPage() *int32
}

type DescribeSampleSceneListResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Page size, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Returned object.
	ResultObject []*DescribeSampleSceneListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items.
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeSampleSceneListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleSceneListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleSceneListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSampleSceneListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeSampleSceneListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSampleSceneListResponseBody) GetResultObject() []*DescribeSampleSceneListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSampleSceneListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeSampleSceneListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeSampleSceneListResponseBody) SetRequestId(v string) *DescribeSampleSceneListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSampleSceneListResponseBody) SetCurrentPage(v int32) *DescribeSampleSceneListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeSampleSceneListResponseBody) SetPageSize(v int32) *DescribeSampleSceneListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSampleSceneListResponseBody) SetResultObject(v []*DescribeSampleSceneListResponseBodyResultObject) *DescribeSampleSceneListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSampleSceneListResponseBody) SetTotalItem(v int32) *DescribeSampleSceneListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeSampleSceneListResponseBody) SetTotalPage(v int32) *DescribeSampleSceneListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeSampleSceneListResponseBody) Validate() error {
	if s.ResultObject != nil {
		for _, item := range s.ResultObject {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSampleSceneListResponseBodyResultObject struct {
	// Business type.
	//
	// example:
	//
	// variable_scene
	BizType *string `json:"bizType,omitempty" xml:"bizType,omitempty"`
	// Configuration key.
	//
	// example:
	//
	// account_abuse_detection
	ConfigKey *string `json:"configKey,omitempty" xml:"configKey,omitempty"`
	// Configuration value.
	//
	// example:
	//
	// 1
	ConfigValue *string `json:"configValue,omitempty" xml:"configValue,omitempty"`
	// Creator.
	//
	// example:
	//
	// 1519714049632764
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// Description information.
	//
	// example:
	//
	// 变量描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Creation time.
	//
	// example:
	//
	// 1621578648000
	GmtCreate *int64 `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1565701886000
	GmtModified *int64 `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// Task ID.
	//
	// example:
	//
	// 3144
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Last modified by.
	//
	// example:
	//
	// 1519714049632764
	LastModifiedOperator *string `json:"lastModifiedOperator,omitempty" xml:"lastModifiedOperator,omitempty"`
	// Status.
	//
	// example:
	//
	// ENABLE
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s DescribeSampleSceneListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleSceneListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSampleSceneListResponseBodyResultObject) GetBizType() *string {
	return s.BizType
}

func (s *DescribeSampleSceneListResponseBodyResultObject) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DescribeSampleSceneListResponseBodyResultObject) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *DescribeSampleSceneListResponseBodyResultObject) GetCreator() *string {
	return s.Creator
}

func (s *DescribeSampleSceneListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeSampleSceneListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeSampleSceneListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeSampleSceneListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeSampleSceneListResponseBodyResultObject) GetLastModifiedOperator() *string {
	return s.LastModifiedOperator
}

func (s *DescribeSampleSceneListResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeSampleSceneListResponseBodyResultObject) SetBizType(v string) *DescribeSampleSceneListResponseBodyResultObject {
	s.BizType = &v
	return s
}

func (s *DescribeSampleSceneListResponseBodyResultObject) SetConfigKey(v string) *DescribeSampleSceneListResponseBodyResultObject {
	s.ConfigKey = &v
	return s
}

func (s *DescribeSampleSceneListResponseBodyResultObject) SetConfigValue(v string) *DescribeSampleSceneListResponseBodyResultObject {
	s.ConfigValue = &v
	return s
}

func (s *DescribeSampleSceneListResponseBodyResultObject) SetCreator(v string) *DescribeSampleSceneListResponseBodyResultObject {
	s.Creator = &v
	return s
}

func (s *DescribeSampleSceneListResponseBodyResultObject) SetDescription(v string) *DescribeSampleSceneListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeSampleSceneListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeSampleSceneListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeSampleSceneListResponseBodyResultObject) SetGmtModified(v int64) *DescribeSampleSceneListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeSampleSceneListResponseBodyResultObject) SetId(v int64) *DescribeSampleSceneListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeSampleSceneListResponseBodyResultObject) SetLastModifiedOperator(v string) *DescribeSampleSceneListResponseBodyResultObject {
	s.LastModifiedOperator = &v
	return s
}

func (s *DescribeSampleSceneListResponseBodyResultObject) SetStatus(v string) *DescribeSampleSceneListResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeSampleSceneListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
