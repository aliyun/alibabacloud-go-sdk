// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTemplatePageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPage(v int32) *DescribeTemplatePageListResponseBody
	GetCurrentPage() *int32
	SetPageSize(v int32) *DescribeTemplatePageListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeTemplatePageListResponseBody
	GetRequestId() *string
	SetResultObject(v []*DescribeTemplatePageListResponseBodyResultObject) *DescribeTemplatePageListResponseBody
	GetResultObject() []*DescribeTemplatePageListResponseBodyResultObject
	SetTotalItem(v int32) *DescribeTemplatePageListResponseBody
	GetTotalItem() *int32
	SetTotalPage(v int32) *DescribeTemplatePageListResponseBody
	GetTotalPage() *int32
}

type DescribeTemplatePageListResponseBody struct {
	// Current page number.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"currentPage,omitempty" xml:"currentPage,omitempty"`
	// Number of items per page, default value is 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// Request ID
	//
	// example:
	//
	// AE7E6105-7DEB-5125-9B24-DCBC139F6CD2
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Returned object
	ResultObject []*DescribeTemplatePageListResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Repeated"`
	// Total number of items
	//
	// example:
	//
	// 6
	TotalItem *int32 `json:"totalItem,omitempty" xml:"totalItem,omitempty"`
	// Total number of pages
	//
	// example:
	//
	// 1
	TotalPage *int32 `json:"totalPage,omitempty" xml:"totalPage,omitempty"`
}

func (s DescribeTemplatePageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatePageListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTemplatePageListResponseBody) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *DescribeTemplatePageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeTemplatePageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTemplatePageListResponseBody) GetResultObject() []*DescribeTemplatePageListResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeTemplatePageListResponseBody) GetTotalItem() *int32 {
	return s.TotalItem
}

func (s *DescribeTemplatePageListResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *DescribeTemplatePageListResponseBody) SetCurrentPage(v int32) *DescribeTemplatePageListResponseBody {
	s.CurrentPage = &v
	return s
}

func (s *DescribeTemplatePageListResponseBody) SetPageSize(v int32) *DescribeTemplatePageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeTemplatePageListResponseBody) SetRequestId(v string) *DescribeTemplatePageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTemplatePageListResponseBody) SetResultObject(v []*DescribeTemplatePageListResponseBodyResultObject) *DescribeTemplatePageListResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeTemplatePageListResponseBody) SetTotalItem(v int32) *DescribeTemplatePageListResponseBody {
	s.TotalItem = &v
	return s
}

func (s *DescribeTemplatePageListResponseBody) SetTotalPage(v int32) *DescribeTemplatePageListResponseBody {
	s.TotalPage = &v
	return s
}

func (s *DescribeTemplatePageListResponseBody) Validate() error {
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

type DescribeTemplatePageListResponseBodyResultObject struct {
	// Description information.
	//
	// example:
	//
	// 描述
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Event code
	//
	// example:
	//
	// de_aszbjb7236
	EventCode *string `json:"eventCode,omitempty" xml:"eventCode,omitempty"`
	// Event name.
	//
	// example:
	//
	// 注册风险
	EventName *string `json:"eventName,omitempty" xml:"eventName,omitempty"`
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
	// Template ID.
	//
	// example:
	//
	// 497
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// Number of policies
	//
	// example:
	//
	// 3
	RuleCount *int32 `json:"ruleCount,omitempty" xml:"ruleCount,omitempty"`
	// Template code.
	//
	// example:
	//
	// register
	TemplateCode *string `json:"templateCode,omitempty" xml:"templateCode,omitempty"`
	// Template name.
	//
	// example:
	//
	// 注册事件
	TemplateName *string `json:"templateName,omitempty" xml:"templateName,omitempty"`
	// Event status
	//
	// example:
	//
	// ONLINE
	TemplateStatus *string `json:"templateStatus,omitempty" xml:"templateStatus,omitempty"`
	// Template type
	//
	// example:
	//
	// PUB_SERVICE
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
	// Number of customer authorizations
	//
	// example:
	//
	// 1
	UserCount *int32 `json:"userCount,omitempty" xml:"userCount,omitempty"`
	// Version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"version,omitempty" xml:"version,omitempty"`
}

func (s DescribeTemplatePageListResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeTemplatePageListResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetEventCode() *string {
	return s.EventCode
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetEventName() *string {
	return s.EventName
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetId() *int64 {
	return s.Id
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetTemplateName() *string {
	return s.TemplateName
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetTemplateStatus() *string {
	return s.TemplateStatus
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetTemplateType() *string {
	return s.TemplateType
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetUserCount() *int32 {
	return s.UserCount
}

func (s *DescribeTemplatePageListResponseBodyResultObject) GetVersion() *int32 {
	return s.Version
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetDescription(v string) *DescribeTemplatePageListResponseBodyResultObject {
	s.Description = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetEventCode(v string) *DescribeTemplatePageListResponseBodyResultObject {
	s.EventCode = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetEventName(v string) *DescribeTemplatePageListResponseBodyResultObject {
	s.EventName = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetGmtCreate(v int64) *DescribeTemplatePageListResponseBodyResultObject {
	s.GmtCreate = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetGmtModified(v int64) *DescribeTemplatePageListResponseBodyResultObject {
	s.GmtModified = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetId(v int64) *DescribeTemplatePageListResponseBodyResultObject {
	s.Id = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetRuleCount(v int32) *DescribeTemplatePageListResponseBodyResultObject {
	s.RuleCount = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetTemplateCode(v string) *DescribeTemplatePageListResponseBodyResultObject {
	s.TemplateCode = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetTemplateName(v string) *DescribeTemplatePageListResponseBodyResultObject {
	s.TemplateName = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetTemplateStatus(v string) *DescribeTemplatePageListResponseBodyResultObject {
	s.TemplateStatus = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetTemplateType(v string) *DescribeTemplatePageListResponseBodyResultObject {
	s.TemplateType = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetUserCount(v int32) *DescribeTemplatePageListResponseBodyResultObject {
	s.UserCount = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) SetVersion(v int32) *DescribeTemplatePageListResponseBodyResultObject {
	s.Version = &v
	return s
}

func (s *DescribeTemplatePageListResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
