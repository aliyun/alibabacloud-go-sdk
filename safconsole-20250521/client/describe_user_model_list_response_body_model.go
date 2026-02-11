// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserModelListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeUserModelListResponseBody
	GetCode() *int32
	SetCurrentPage(v int64) *DescribeUserModelListResponseBody
	GetCurrentPage() *int64
	SetHttpStatusCode(v int64) *DescribeUserModelListResponseBody
	GetHttpStatusCode() *int64
	SetPageSize(v int64) *DescribeUserModelListResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeUserModelListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeUserModelListResponseBodyResultObject) *DescribeUserModelListResponseBody
	GetResultObject() []*DescribeUserModelListResponseBodyResultObject
	SetSuccess(v bool) *DescribeUserModelListResponseBody
	GetSuccess() *bool
	SetTotalItem(v int64) *DescribeUserModelListResponseBody
	GetTotalItem() *int64
	SetTotalPage(v int64) *DescribeUserModelListResponseBody
	GetTotalPage() *int64
}

type DescribeUserModelListResponseBody struct {
	// `code`
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Current page.
	//
	// example:
	//
	// 1
	CurrentPage *int64 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Pagination parameter: number of items per page, default value is 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	ResultObject []*DescribeUserModelListResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Repeated"`
	// Indicates whether the request was successful, with values as follows:
	//
	// - true, request succeeded
	//
	// - false, request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// Total number of records.
	//
	// example:
	//
	// 1
	TotalItem *int64 `json:"TotalItem,omitempty" xml:"TotalItem,omitempty"`
	// Total number of pages.
	//
	// example:
	//
	// 1
	TotalPage *int64 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s DescribeUserModelListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserModelListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserModelListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeUserModelListResponseBody) GetCurrentPage() *int64 {
	return s.CurrentPage
}

func (s *DescribeUserModelListResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeUserModelListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeUserModelListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserModelListResponseBody) GetResultObject() []*DescribeUserModelListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeUserModelListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeUserModelListResponseBody) GetTotalItem() *int64 {
	return s.TotalItem
}

func (s *DescribeUserModelListResponseBody) GetTotalPage() *int64 {
	return s.TotalPage
}

func (s *DescribeUserModelListResponseBody) SetCode(v int32) *DescribeUserModelListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUserModelListResponseBody) SetCurrentPage(v int64) *DescribeUserModelListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeUserModelListResponseBody) SetHttpStatusCode(v int64) *DescribeUserModelListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeUserModelListResponseBody) SetPageSize(v int64) *DescribeUserModelListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeUserModelListResponseBody) SetRequestId(v string) *DescribeUserModelListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserModelListResponseBody) SetResultObject(v []*DescribeUserModelListResponseBodyResultObject) *DescribeUserModelListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeUserModelListResponseBody) SetSuccess(v bool) *DescribeUserModelListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeUserModelListResponseBody) SetTotalItem(v int64) *DescribeUserModelListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeUserModelListResponseBody) SetTotalPage(v int64) *DescribeUserModelListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeUserModelListResponseBody) Validate() error {
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

type DescribeUserModelListResponseBodyResultObject struct {
	// Authorization type.
	//
	// example:
	//
	// READ
	AuthType *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// Model name.
	//
	// example:
	//
	// Model_A
	CustomerModuleName *string `json:"CustomerModuleName,omitempty" xml:"CustomerModuleName,omitempty"`
	// Remarks.
	//
	// example:
	//
	// 备注。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Associated features.
	//
	// example:
	//
	// template_a
	FeatureTemplate *string `json:"FeatureTemplate,omitempty" xml:"FeatureTemplate,omitempty"`
	// Modification time.
	//
	// example:
	//
	// 1673578656000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Primary key ID of the model.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Publication status.
	//
	// example:
	//
	// ONLINE
	InnerDefineStatus *string `json:"InnerDefineStatus,omitempty" xml:"InnerDefineStatus,omitempty"`
	// Model identifier.
	//
	// example:
	//
	// inner_model_a
	InnerModuleName *string `json:"InnerModuleName,omitempty" xml:"InnerModuleName,omitempty"`
	// Whether iteration is allowed.
	//
	// example:
	//
	// true
	IsAllowIterate *bool `json:"IsAllowIterate,omitempty" xml:"IsAllowIterate,omitempty"`
	// Whether rollback is allowed.
	//
	// example:
	//
	// true
	IsAllowRollback *bool `json:"IsAllowRollback,omitempty" xml:"IsAllowRollback,omitempty"`
	// Iteration version.
	//
	// example:
	//
	// 1
	IterationVersion *int64 `json:"IterationVersion,omitempty" xml:"IterationVersion,omitempty"`
	// Root model ID.
	//
	// example:
	//
	// 10
	RootModuleId *int64 `json:"RootModuleId,omitempty" xml:"RootModuleId,omitempty"`
}

func (s DescribeUserModelListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserModelListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeUserModelListResponseBodyResultObject) GetAuthType() *string {
	return s.AuthType
}

func (s *DescribeUserModelListResponseBodyResultObject) GetCustomerModuleName() *string {
	return s.CustomerModuleName
}

func (s *DescribeUserModelListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeUserModelListResponseBodyResultObject) GetFeatureTemplate() *string {
	return s.FeatureTemplate
}

func (s *DescribeUserModelListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeUserModelListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeUserModelListResponseBodyResultObject) GetInnerDefineStatus() *string {
	return s.InnerDefineStatus
}

func (s *DescribeUserModelListResponseBodyResultObject) GetInnerModuleName() *string {
	return s.InnerModuleName
}

func (s *DescribeUserModelListResponseBodyResultObject) GetIsAllowIterate() *bool {
	return s.IsAllowIterate
}

func (s *DescribeUserModelListResponseBodyResultObject) GetIsAllowRollback() *bool {
	return s.IsAllowRollback
}

func (s *DescribeUserModelListResponseBodyResultObject) GetIterationVersion() *int64 {
	return s.IterationVersion
}

func (s *DescribeUserModelListResponseBodyResultObject) GetRootModuleId() *int64 {
	return s.RootModuleId
}

func (s *DescribeUserModelListResponseBodyResultObject) SetAuthType(v string) *DescribeUserModelListResponseBodyResultObject {
	s.AuthType = &v
	return s
}

func (s *DescribeUserModelListResponseBodyResultObject) SetCustomerModuleName(v string) *DescribeUserModelListResponseBodyResultObject {
	s.CustomerModuleName = &v
	return s
}

func (s *DescribeUserModelListResponseBodyResultObject) SetDescription(v string) *DescribeUserModelListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeUserModelListResponseBodyResultObject) SetFeatureTemplate(v string) *DescribeUserModelListResponseBodyResultObject {
	s.FeatureTemplate = &v
	return s
}

func (s *DescribeUserModelListResponseBodyResultObject) SetGmtModified(v int64) *DescribeUserModelListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeUserModelListResponseBodyResultObject) SetId(v int64) *DescribeUserModelListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeUserModelListResponseBodyResultObject) SetInnerDefineStatus(v string) *DescribeUserModelListResponseBodyResultObject {
	s.InnerDefineStatus = &v
	return s
}

func (s *DescribeUserModelListResponseBodyResultObject) SetInnerModuleName(v string) *DescribeUserModelListResponseBodyResultObject {
	s.InnerModuleName = &v
	return s
}

func (s *DescribeUserModelListResponseBodyResultObject) SetIsAllowIterate(v bool) *DescribeUserModelListResponseBodyResultObject {
	s.IsAllowIterate = &v
	return s
}

func (s *DescribeUserModelListResponseBodyResultObject) SetIsAllowRollback(v bool) *DescribeUserModelListResponseBodyResultObject {
	s.IsAllowRollback = &v
	return s
}

func (s *DescribeUserModelListResponseBodyResultObject) SetIterationVersion(v int64) *DescribeUserModelListResponseBodyResultObject {
	s.IterationVersion = &v
	return s
}

func (s *DescribeUserModelListResponseBodyResultObject) SetRootModuleId(v int64) *DescribeUserModelListResponseBodyResultObject {
	s.RootModuleId = &v
	return s
}

func (s *DescribeUserModelListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
