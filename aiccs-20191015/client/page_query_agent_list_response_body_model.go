// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageQueryAgentListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *PageQueryAgentListResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *PageQueryAgentListResponseBody
	GetCode() *string
	SetData(v *PageQueryAgentListResponseBodyData) *PageQueryAgentListResponseBody
	GetData() *PageQueryAgentListResponseBodyData
	SetMessage(v string) *PageQueryAgentListResponseBody
	GetMessage() *string
	SetRequestId(v string) *PageQueryAgentListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *PageQueryAgentListResponseBody
	GetSuccess() *bool
}

type PageQueryAgentListResponseBody struct {
	// example:
	//
	// Access Denied
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *PageQueryAgentListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 742C9243-2870-B8D6-0C68-C60BEB2DF09A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s PageQueryAgentListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PageQueryAgentListResponseBody) GoString() string {
	return s.String()
}

func (s *PageQueryAgentListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *PageQueryAgentListResponseBody) GetCode() *string {
	return s.Code
}

func (s *PageQueryAgentListResponseBody) GetData() *PageQueryAgentListResponseBodyData {
	return s.Data
}

func (s *PageQueryAgentListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PageQueryAgentListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PageQueryAgentListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *PageQueryAgentListResponseBody) SetAccessDeniedDetail(v string) *PageQueryAgentListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *PageQueryAgentListResponseBody) SetCode(v string) *PageQueryAgentListResponseBody {
	s.Code = &v
	return s
}

func (s *PageQueryAgentListResponseBody) SetData(v *PageQueryAgentListResponseBodyData) *PageQueryAgentListResponseBody {
	s.Data = v
	return s
}

func (s *PageQueryAgentListResponseBody) SetMessage(v string) *PageQueryAgentListResponseBody {
	s.Message = &v
	return s
}

func (s *PageQueryAgentListResponseBody) SetRequestId(v string) *PageQueryAgentListResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageQueryAgentListResponseBody) SetSuccess(v bool) *PageQueryAgentListResponseBody {
	s.Success = &v
	return s
}

func (s *PageQueryAgentListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PageQueryAgentListResponseBodyData struct {
	List []*PageQueryAgentListResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNo *int64 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 14
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s PageQueryAgentListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PageQueryAgentListResponseBodyData) GoString() string {
	return s.String()
}

func (s *PageQueryAgentListResponseBodyData) GetList() []*PageQueryAgentListResponseBodyDataList {
	return s.List
}

func (s *PageQueryAgentListResponseBodyData) GetPageNo() *int64 {
	return s.PageNo
}

func (s *PageQueryAgentListResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *PageQueryAgentListResponseBodyData) GetTotal() *int64 {
	return s.Total
}

func (s *PageQueryAgentListResponseBodyData) SetList(v []*PageQueryAgentListResponseBodyDataList) *PageQueryAgentListResponseBodyData {
	s.List = v
	return s
}

func (s *PageQueryAgentListResponseBodyData) SetPageNo(v int64) *PageQueryAgentListResponseBodyData {
	s.PageNo = &v
	return s
}

func (s *PageQueryAgentListResponseBodyData) SetPageSize(v int64) *PageQueryAgentListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *PageQueryAgentListResponseBodyData) SetTotal(v int64) *PageQueryAgentListResponseBodyData {
	s.Total = &v
	return s
}

func (s *PageQueryAgentListResponseBodyData) Validate() error {
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

type PageQueryAgentListResponseBodyDataList struct {
	// example:
	//
	// 121312*******
	AgentId *string `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	// example:
	//
	// 示例值
	AgentName *string `json:"AgentName,omitempty" xml:"AgentName,omitempty"`
	// example:
	//
	// DFAS*****
	ApplicationCode *string `json:"ApplicationCode,omitempty" xml:"ApplicationCode,omitempty"`
	// example:
	//
	// 示例值
	AuditReason *string `json:"AuditReason,omitempty" xml:"AuditReason,omitempty"`
	// example:
	//
	// 构建失败原因
	BuildFailReason *string `json:"BuildFailReason,omitempty" xml:"BuildFailReason,omitempty"`
	// example:
	//
	// 示例值示例值示例值
	BusinessTypeName *string `json:"BusinessTypeName,omitempty" xml:"BusinessTypeName,omitempty"`
	// example:
	//
	// 2025-10-28 17:10:17
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 示例值示例值
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2025-10-28 14:38:15
	LastOnlineTime *string `json:"LastOnlineTime,omitempty" xml:"LastOnlineTime,omitempty"`
	// example:
	//
	// 2025-10-28 17:10:17
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 7
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// false
	WithActivePrompt *bool `json:"WithActivePrompt,omitempty" xml:"WithActivePrompt,omitempty"`
	// example:
	//
	// false
	WithConfig *bool `json:"WithConfig,omitempty" xml:"WithConfig,omitempty"`
}

func (s PageQueryAgentListResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s PageQueryAgentListResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *PageQueryAgentListResponseBodyDataList) GetAgentId() *string {
	return s.AgentId
}

func (s *PageQueryAgentListResponseBodyDataList) GetAgentName() *string {
	return s.AgentName
}

func (s *PageQueryAgentListResponseBodyDataList) GetApplicationCode() *string {
	return s.ApplicationCode
}

func (s *PageQueryAgentListResponseBodyDataList) GetAuditReason() *string {
	return s.AuditReason
}

func (s *PageQueryAgentListResponseBodyDataList) GetBuildFailReason() *string {
	return s.BuildFailReason
}

func (s *PageQueryAgentListResponseBodyDataList) GetBusinessTypeName() *string {
	return s.BusinessTypeName
}

func (s *PageQueryAgentListResponseBodyDataList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *PageQueryAgentListResponseBodyDataList) GetDescription() *string {
	return s.Description
}

func (s *PageQueryAgentListResponseBodyDataList) GetLastOnlineTime() *string {
	return s.LastOnlineTime
}

func (s *PageQueryAgentListResponseBodyDataList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *PageQueryAgentListResponseBodyDataList) GetStatus() *int64 {
	return s.Status
}

func (s *PageQueryAgentListResponseBodyDataList) GetWithActivePrompt() *bool {
	return s.WithActivePrompt
}

func (s *PageQueryAgentListResponseBodyDataList) GetWithConfig() *bool {
	return s.WithConfig
}

func (s *PageQueryAgentListResponseBodyDataList) SetAgentId(v string) *PageQueryAgentListResponseBodyDataList {
	s.AgentId = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) SetAgentName(v string) *PageQueryAgentListResponseBodyDataList {
	s.AgentName = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) SetApplicationCode(v string) *PageQueryAgentListResponseBodyDataList {
	s.ApplicationCode = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) SetAuditReason(v string) *PageQueryAgentListResponseBodyDataList {
	s.AuditReason = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) SetBuildFailReason(v string) *PageQueryAgentListResponseBodyDataList {
	s.BuildFailReason = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) SetBusinessTypeName(v string) *PageQueryAgentListResponseBodyDataList {
	s.BusinessTypeName = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) SetCreateTime(v string) *PageQueryAgentListResponseBodyDataList {
	s.CreateTime = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) SetDescription(v string) *PageQueryAgentListResponseBodyDataList {
	s.Description = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) SetLastOnlineTime(v string) *PageQueryAgentListResponseBodyDataList {
	s.LastOnlineTime = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) SetModifyTime(v string) *PageQueryAgentListResponseBodyDataList {
	s.ModifyTime = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) SetStatus(v int64) *PageQueryAgentListResponseBodyDataList {
	s.Status = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) SetWithActivePrompt(v bool) *PageQueryAgentListResponseBodyDataList {
	s.WithActivePrompt = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) SetWithConfig(v bool) *PageQueryAgentListResponseBodyDataList {
	s.WithConfig = &v
	return s
}

func (s *PageQueryAgentListResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
