// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVariableSceneListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVariableSceneListResponseBody
	GetRequestId() *string
	SetCurrentPage(v int32) *DescribeVariableSceneListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeVariableSceneListResponseBody
	GetPageSize() *int32
	SetResultObject(v []*DescribeVariableSceneListResponseBodyResultObject) *DescribeVariableSceneListResponseBody
	GetResultObject() []*DescribeVariableSceneListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeVariableSceneListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeVariableSceneListResponseBody
	GetTotalPage() *int32
}

type DescribeVariableSceneListResponseBody struct {
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
	ResultObject []*DescribeVariableSceneListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
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

func (s DescribeVariableSceneListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableSceneListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVariableSceneListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVariableSceneListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeVariableSceneListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVariableSceneListResponseBody) GetResultObject() []*DescribeVariableSceneListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeVariableSceneListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeVariableSceneListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeVariableSceneListResponseBody) SetRequestId(v string) *DescribeVariableSceneListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVariableSceneListResponseBody) SetCurrentPage(v int32) *DescribeVariableSceneListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeVariableSceneListResponseBody) SetPageSize(v int32) *DescribeVariableSceneListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVariableSceneListResponseBody) SetResultObject(v []*DescribeVariableSceneListResponseBodyResultObject) *DescribeVariableSceneListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeVariableSceneListResponseBody) SetTotalItem(v int32) *DescribeVariableSceneListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeVariableSceneListResponseBody) SetTotalPage(v int32) *DescribeVariableSceneListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeVariableSceneListResponseBody) Validate() error {
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

type DescribeVariableSceneListResponseBodyResultObject struct {
	// Business category identifier.
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
	// Created by.
	//
	// example:
	//
	// 1519714049632764
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// Description information.
	//
	// example:
	//
	// 配置描述信息
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
	// Primary key ID of the configuration.
	//
	// example:
	//
	// 497
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

func (s DescribeVariableSceneListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeVariableSceneListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeVariableSceneListResponseBodyResultObject) GetBizType() *string {
	return s.BizType
}

func (s *DescribeVariableSceneListResponseBodyResultObject) GetConfigKey() *string {
	return s.ConfigKey
}

func (s *DescribeVariableSceneListResponseBodyResultObject) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *DescribeVariableSceneListResponseBodyResultObject) GetCreator() *string {
	return s.Creator
}

func (s *DescribeVariableSceneListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeVariableSceneListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeVariableSceneListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeVariableSceneListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeVariableSceneListResponseBodyResultObject) GetLastModifiedOperator() *string {
	return s.LastModifiedOperator
}

func (s *DescribeVariableSceneListResponseBodyResultObject) GetStatus() *string {
	return s.Status
}

func (s *DescribeVariableSceneListResponseBodyResultObject) SetBizType(v string) *DescribeVariableSceneListResponseBodyResultObject {
	s.BizType = &v
	return s
}

func (s *DescribeVariableSceneListResponseBodyResultObject) SetConfigKey(v string) *DescribeVariableSceneListResponseBodyResultObject {
	s.ConfigKey = &v
	return s
}

func (s *DescribeVariableSceneListResponseBodyResultObject) SetConfigValue(v string) *DescribeVariableSceneListResponseBodyResultObject {
	s.ConfigValue = &v
	return s
}

func (s *DescribeVariableSceneListResponseBodyResultObject) SetCreator(v string) *DescribeVariableSceneListResponseBodyResultObject {
	s.Creator = &v
	return s
}

func (s *DescribeVariableSceneListResponseBodyResultObject) SetDescription(v string) *DescribeVariableSceneListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeVariableSceneListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeVariableSceneListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeVariableSceneListResponseBodyResultObject) SetGmtModified(v int64) *DescribeVariableSceneListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeVariableSceneListResponseBodyResultObject) SetId(v int64) *DescribeVariableSceneListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeVariableSceneListResponseBodyResultObject) SetLastModifiedOperator(v string) *DescribeVariableSceneListResponseBodyResultObject {
	s.LastModifiedOperator = &v
	return s
}

func (s *DescribeVariableSceneListResponseBodyResultObject) SetStatus(v string) *DescribeVariableSceneListResponseBodyResultObject {
	s.Status = &v
	return s
}

func (s *DescribeVariableSceneListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
