// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsTemplateListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetSmsTemplateListResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetSmsTemplateListResponseBody
	GetCode() *string
	SetData(v *GetSmsTemplateListResponseBodyData) *GetSmsTemplateListResponseBody
	GetData() *GetSmsTemplateListResponseBodyData
	SetMessage(v string) *GetSmsTemplateListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSmsTemplateListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetSmsTemplateListResponseBody
	GetSuccess() *bool
}

type GetSmsTemplateListResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值示例值
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetSmsTemplateListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 示例值
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetSmsTemplateListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateListResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetSmsTemplateListResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSmsTemplateListResponseBody) GetData() *GetSmsTemplateListResponseBodyData {
	return s.Data
}

func (s *GetSmsTemplateListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSmsTemplateListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmsTemplateListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetSmsTemplateListResponseBody) SetAccessDeniedDetail(v string) *GetSmsTemplateListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetSmsTemplateListResponseBody) SetCode(v string) *GetSmsTemplateListResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmsTemplateListResponseBody) SetData(v *GetSmsTemplateListResponseBodyData) *GetSmsTemplateListResponseBody {
	s.Data = v
	return s
}

func (s *GetSmsTemplateListResponseBody) SetMessage(v string) *GetSmsTemplateListResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmsTemplateListResponseBody) SetRequestId(v string) *GetSmsTemplateListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmsTemplateListResponseBody) SetSuccess(v bool) *GetSmsTemplateListResponseBody {
	s.Success = &v
	return s
}

func (s *GetSmsTemplateListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSmsTemplateListResponseBodyData struct {
	List []*GetSmsTemplateListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 97
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 78
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 64
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetSmsTemplateListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateListResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateListResponseBodyData) GetList() []*GetSmsTemplateListResponseBodyDataList {
	return s.List
}

func (s *GetSmsTemplateListResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *GetSmsTemplateListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *GetSmsTemplateListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *GetSmsTemplateListResponseBodyData) SetList(v []*GetSmsTemplateListResponseBodyDataList) *GetSmsTemplateListResponseBodyData {
	s.List = v
	return s
}

func (s *GetSmsTemplateListResponseBodyData) SetPageNo(v int64) *GetSmsTemplateListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *GetSmsTemplateListResponseBodyData) SetPageSize(v int64) *GetSmsTemplateListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *GetSmsTemplateListResponseBodyData) SetTotal(v int64) *GetSmsTemplateListResponseBodyData {
	s.Total = &v
	return s
}

func (s *GetSmsTemplateListResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSmsTemplateListResponseBodyDataList struct {
	// 模板审核状态
	//
	// example:
	//
	// 88
	AuditStatus *int64 `json:"AuditStatus,omitempty" xml:"AuditStatus,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 示例值
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 签名
	//
	// example:
	//
	// 示例值示例值
	SignName *string `json:"SignName,omitempty" xml:"SignName,omitempty"`
	// 模板code
	//
	// example:
	//
	// 示例值示例值示例值
	TemplateCode *string `json:"TemplateCode,omitempty" xml:"TemplateCode,omitempty"`
	// 模板名称
	//
	// example:
	//
	// 示例值
	TemplateName *string `json:"TemplateName,omitempty" xml:"TemplateName,omitempty"`
	// 模板标签
	TemplateTag []*GetSmsTemplateListResponseBodyDataListTemplateTag `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty" type:"Repeated"`
	// 模板类型
	//
	// example:
	//
	// 43
	TemplateType *int64 `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
	// 模板可用状态
	UsableStateList []*string `json:"UsableStateList,omitempty" xml:"UsableStateList,omitempty" type:"Repeated"`
	// 工单号
	//
	// example:
	//
	// 示例值示例值
	WorkOrderId *string `json:"WorkOrderId,omitempty" xml:"WorkOrderId,omitempty"`
}

func (s GetSmsTemplateListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateListResponseBodyDataList) GetAuditStatus() *int64 {
	return s.AuditStatus
}

func (s *GetSmsTemplateListResponseBodyDataList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetSmsTemplateListResponseBodyDataList) GetSignName() *string {
	return s.SignName
}

func (s *GetSmsTemplateListResponseBodyDataList) GetTemplateCode() *string {
	return s.TemplateCode
}

func (s *GetSmsTemplateListResponseBodyDataList) GetTemplateName() *string {
	return s.TemplateName
}

func (s *GetSmsTemplateListResponseBodyDataList) GetTemplateTag() []*GetSmsTemplateListResponseBodyDataListTemplateTag {
	return s.TemplateTag
}

func (s *GetSmsTemplateListResponseBodyDataList) GetTemplateType() *int64 {
	return s.TemplateType
}

func (s *GetSmsTemplateListResponseBodyDataList) GetUsableStateList() []*string {
	return s.UsableStateList
}

func (s *GetSmsTemplateListResponseBodyDataList) GetWorkOrderId() *string {
	return s.WorkOrderId
}

func (s *GetSmsTemplateListResponseBodyDataList) SetAuditStatus(v int64) *GetSmsTemplateListResponseBodyDataList {
	s.AuditStatus = &v
	return s
}

func (s *GetSmsTemplateListResponseBodyDataList) SetGmtCreate(v string) *GetSmsTemplateListResponseBodyDataList {
	s.GmtCreate = &v
	return s
}

func (s *GetSmsTemplateListResponseBodyDataList) SetSignName(v string) *GetSmsTemplateListResponseBodyDataList {
	s.SignName = &v
	return s
}

func (s *GetSmsTemplateListResponseBodyDataList) SetTemplateCode(v string) *GetSmsTemplateListResponseBodyDataList {
	s.TemplateCode = &v
	return s
}

func (s *GetSmsTemplateListResponseBodyDataList) SetTemplateName(v string) *GetSmsTemplateListResponseBodyDataList {
	s.TemplateName = &v
	return s
}

func (s *GetSmsTemplateListResponseBodyDataList) SetTemplateTag(v []*GetSmsTemplateListResponseBodyDataListTemplateTag) *GetSmsTemplateListResponseBodyDataList {
	s.TemplateTag = v
	return s
}

func (s *GetSmsTemplateListResponseBodyDataList) SetTemplateType(v int64) *GetSmsTemplateListResponseBodyDataList {
	s.TemplateType = &v
	return s
}

func (s *GetSmsTemplateListResponseBodyDataList) SetUsableStateList(v []*string) *GetSmsTemplateListResponseBodyDataList {
	s.UsableStateList = v
	return s
}

func (s *GetSmsTemplateListResponseBodyDataList) SetWorkOrderId(v string) *GetSmsTemplateListResponseBodyDataList {
	s.WorkOrderId = &v
	return s
}

func (s *GetSmsTemplateListResponseBodyDataList) Validate() error {
	if s.TemplateTag != nil {
		for _, item := range s.TemplateTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetSmsTemplateListResponseBodyDataListTemplateTag struct {
	// example:
	//
	// 示例值
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// 示例值示例值
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s GetSmsTemplateListResponseBodyDataListTemplateTag) String() string {
	return dara.Prettify(s)
}

func (s GetSmsTemplateListResponseBodyDataListTemplateTag) GoString() string {
	return s.String()
}

func (s *GetSmsTemplateListResponseBodyDataListTemplateTag) GetTagKey() *string {
	return s.TagKey
}

func (s *GetSmsTemplateListResponseBodyDataListTemplateTag) GetTagValue() *string {
	return s.TagValue
}

func (s *GetSmsTemplateListResponseBodyDataListTemplateTag) SetTagKey(v string) *GetSmsTemplateListResponseBodyDataListTemplateTag {
	s.TagKey = &v
	return s
}

func (s *GetSmsTemplateListResponseBodyDataListTemplateTag) SetTagValue(v string) *GetSmsTemplateListResponseBodyDataListTemplateTag {
	s.TagValue = &v
	return s
}

func (s *GetSmsTemplateListResponseBodyDataListTemplateTag) Validate() error {
	return dara.Validate(s)
}
