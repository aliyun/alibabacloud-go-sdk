// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAgentGroupPageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *AgentGroupPageResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *AgentGroupPageResponseBody
	GetCode() *string
	SetMessage(v string) *AgentGroupPageResponseBody
	GetMessage() *string
	SetModel(v *AgentGroupPageResponseBodyModel) *AgentGroupPageResponseBody
	GetModel() *AgentGroupPageResponseBodyModel
	SetRequestId(v string) *AgentGroupPageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AgentGroupPageResponseBody
	GetSuccess() *bool
	SetTimestamp(v int64) *AgentGroupPageResponseBody
	GetTimestamp() *int64
}

type AgentGroupPageResponseBody struct {
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 示例值
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// a
	Message *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Model   *AgentGroupPageResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 12345
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 71
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
}

func (s AgentGroupPageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AgentGroupPageResponseBody) GoString() string {
	return s.String()
}

func (s *AgentGroupPageResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *AgentGroupPageResponseBody) GetCode() *string {
	return s.Code
}

func (s *AgentGroupPageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AgentGroupPageResponseBody) GetModel() *AgentGroupPageResponseBodyModel {
	return s.Model
}

func (s *AgentGroupPageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AgentGroupPageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AgentGroupPageResponseBody) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *AgentGroupPageResponseBody) SetAccessDeniedDetail(v string) *AgentGroupPageResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *AgentGroupPageResponseBody) SetCode(v string) *AgentGroupPageResponseBody {
	s.Code = &v
	return s
}

func (s *AgentGroupPageResponseBody) SetMessage(v string) *AgentGroupPageResponseBody {
	s.Message = &v
	return s
}

func (s *AgentGroupPageResponseBody) SetModel(v *AgentGroupPageResponseBodyModel) *AgentGroupPageResponseBody {
	s.Model = v
	return s
}

func (s *AgentGroupPageResponseBody) SetRequestId(v string) *AgentGroupPageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AgentGroupPageResponseBody) SetSuccess(v bool) *AgentGroupPageResponseBody {
	s.Success = &v
	return s
}

func (s *AgentGroupPageResponseBody) SetTimestamp(v int64) *AgentGroupPageResponseBody {
	s.Timestamp = &v
	return s
}

func (s *AgentGroupPageResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AgentGroupPageResponseBodyModel struct {
	// example:
	//
	// 62
	Current *int64 `json:"Current,omitempty" xml:"Current,omitempty"`
	// example:
	//
	// 82
	Pages   *int64                                    `json:"Pages,omitempty" xml:"Pages,omitempty"`
	Records []*AgentGroupPageResponseBodyModelRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 15
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// 12
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s AgentGroupPageResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s AgentGroupPageResponseBodyModel) GoString() string {
	return s.String()
}

func (s *AgentGroupPageResponseBodyModel) GetCurrent() *int64 {
	return s.Current
}

func (s *AgentGroupPageResponseBodyModel) GetPages() *int64 {
	return s.Pages
}

func (s *AgentGroupPageResponseBodyModel) GetRecords() []*AgentGroupPageResponseBodyModelRecords {
	return s.Records
}

func (s *AgentGroupPageResponseBodyModel) GetSize() *int64 {
	return s.Size
}

func (s *AgentGroupPageResponseBodyModel) GetTotal() *int64 {
	return s.Total
}

func (s *AgentGroupPageResponseBodyModel) SetCurrent(v int64) *AgentGroupPageResponseBodyModel {
	s.Current = &v
	return s
}

func (s *AgentGroupPageResponseBodyModel) SetPages(v int64) *AgentGroupPageResponseBodyModel {
	s.Pages = &v
	return s
}

func (s *AgentGroupPageResponseBodyModel) SetRecords(v []*AgentGroupPageResponseBodyModelRecords) *AgentGroupPageResponseBodyModel {
	s.Records = v
	return s
}

func (s *AgentGroupPageResponseBodyModel) SetSize(v int64) *AgentGroupPageResponseBodyModel {
	s.Size = &v
	return s
}

func (s *AgentGroupPageResponseBodyModel) SetTotal(v int64) *AgentGroupPageResponseBodyModel {
	s.Total = &v
	return s
}

func (s *AgentGroupPageResponseBodyModel) Validate() error {
	if s.Records != nil {
		for _, item := range s.Records {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AgentGroupPageResponseBodyModelRecords struct {
	// 坐席组ID
	//
	// example:
	//
	// 1
	AgentGroupId *int64 `json:"AgentGroupId,omitempty" xml:"AgentGroupId,omitempty"`
	// 坐席组名称
	//
	// example:
	//
	// a
	AgentGroupName *string `json:"AgentGroupName,omitempty" xml:"AgentGroupName,omitempty"`
	// 坐席组下的坐席id列表
	AgentIds []*int64 `json:"AgentIds,omitempty" xml:"AgentIds,omitempty" type:"Repeated"`
	// 创建坐席组的时间
	//
	// example:
	//
	// 2026-01-01 11:11:11
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
}

func (s AgentGroupPageResponseBodyModelRecords) String() string {
	return dara.Prettify(s)
}

func (s AgentGroupPageResponseBodyModelRecords) GoString() string {
	return s.String()
}

func (s *AgentGroupPageResponseBodyModelRecords) GetAgentGroupId() *int64 {
	return s.AgentGroupId
}

func (s *AgentGroupPageResponseBodyModelRecords) GetAgentGroupName() *string {
	return s.AgentGroupName
}

func (s *AgentGroupPageResponseBodyModelRecords) GetAgentIds() []*int64 {
	return s.AgentIds
}

func (s *AgentGroupPageResponseBodyModelRecords) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AgentGroupPageResponseBodyModelRecords) SetAgentGroupId(v int64) *AgentGroupPageResponseBodyModelRecords {
	s.AgentGroupId = &v
	return s
}

func (s *AgentGroupPageResponseBodyModelRecords) SetAgentGroupName(v string) *AgentGroupPageResponseBodyModelRecords {
	s.AgentGroupName = &v
	return s
}

func (s *AgentGroupPageResponseBodyModelRecords) SetAgentIds(v []*int64) *AgentGroupPageResponseBodyModelRecords {
	s.AgentIds = v
	return s
}

func (s *AgentGroupPageResponseBodyModelRecords) SetCreateTime(v string) *AgentGroupPageResponseBodyModelRecords {
	s.CreateTime = &v
	return s
}

func (s *AgentGroupPageResponseBodyModelRecords) Validate() error {
	return dara.Validate(s)
}
