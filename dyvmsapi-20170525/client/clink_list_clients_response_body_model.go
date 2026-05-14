// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClinkListClientsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ClinkListClientsResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *ClinkListClientsResponseBody
	GetCode() *string
	SetData(v *ClinkListClientsResponseBodyData) *ClinkListClientsResponseBody
	GetData() *ClinkListClientsResponseBodyData
	SetMessage(v string) *ClinkListClientsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ClinkListClientsResponseBody
	GetRequestId() *string
}

type ClinkListClientsResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ClinkListClientsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D9CB3933-9FE3-4870-BA8E-2BEE91B69D23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ClinkListClientsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ClinkListClientsResponseBody) GoString() string {
	return s.String()
}

func (s *ClinkListClientsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ClinkListClientsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ClinkListClientsResponseBody) GetData() *ClinkListClientsResponseBodyData {
	return s.Data
}

func (s *ClinkListClientsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ClinkListClientsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ClinkListClientsResponseBody) SetAccessDeniedDetail(v string) *ClinkListClientsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ClinkListClientsResponseBody) SetCode(v string) *ClinkListClientsResponseBody {
	s.Code = &v
	return s
}

func (s *ClinkListClientsResponseBody) SetData(v *ClinkListClientsResponseBodyData) *ClinkListClientsResponseBody {
	s.Data = v
	return s
}

func (s *ClinkListClientsResponseBody) SetMessage(v string) *ClinkListClientsResponseBody {
	s.Message = &v
	return s
}

func (s *ClinkListClientsResponseBody) SetRequestId(v string) *ClinkListClientsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ClinkListClientsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ClinkListClientsResponseBodyData struct {
	// 座席对象列表
	Clients []*ClinkListClientsResponseBodyDataClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// 请求 id
	//
	// example:
	//
	// xxx
	ClinkRequestId *string `json:"ClinkRequestId,omitempty" xml:"ClinkRequestId,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 22
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ClinkListClientsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ClinkListClientsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ClinkListClientsResponseBodyData) GetClients() []*ClinkListClientsResponseBodyDataClients {
	return s.Clients
}

func (s *ClinkListClientsResponseBodyData) GetClinkRequestId() *string {
	return s.ClinkRequestId
}

func (s *ClinkListClientsResponseBodyData) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ClinkListClientsResponseBodyData) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ClinkListClientsResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ClinkListClientsResponseBodyData) SetClients(v []*ClinkListClientsResponseBodyDataClients) *ClinkListClientsResponseBodyData {
	s.Clients = v
	return s
}

func (s *ClinkListClientsResponseBodyData) SetClinkRequestId(v string) *ClinkListClientsResponseBodyData {
	s.ClinkRequestId = &v
	return s
}

func (s *ClinkListClientsResponseBodyData) SetPageNumber(v int64) *ClinkListClientsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ClinkListClientsResponseBodyData) SetPageSize(v int64) *ClinkListClientsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ClinkListClientsResponseBodyData) SetTotalCount(v int64) *ClinkListClientsResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *ClinkListClientsResponseBodyData) Validate() error {
	if s.Clients != nil {
		for _, item := range s.Clients {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ClinkListClientsResponseBodyDataClients struct {
	// 是否激活，0: 否；1: 是
	//
	// example:
	//
	// 1
	Active *int64 `json:"Active,omitempty" xml:"Active,omitempty"`
	// 区号
	//
	// example:
	//
	// 010
	AreaCode *string `json:"AreaCode,omitempty" xml:"AreaCode,omitempty"`
	// 座席绑定电话
	//
	// example:
	//
	// xxx
	BindTel *string `json:"BindTel,omitempty" xml:"BindTel,omitempty"`
	// 座席号
	//
	// example:
	//
	// 1111
	Cno *string `json:"Cno,omitempty" xml:"Cno,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 1775024775
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// 座席crm id
	//
	// example:
	//
	// xxx
	CrmId *string `json:"CrmId,omitempty" xml:"CrmId,omitempty"`
	// 号码隐藏类型，0: 不隐藏；1: 隐藏规则与全局设置保持一致
	//
	// example:
	//
	// 0
	HiddenTel *int64 `json:"HiddenTel,omitempty" xml:"HiddenTel,omitempty"`
	// 座席 Id
	//
	// example:
	//
	// 34
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// 座席名称
	//
	// example:
	//
	// 2222
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// 座席所属队列的队列号数组
	Qnos []*string `json:"Qnos,omitempty" xml:"Qnos,omitempty" type:"Repeated"`
	// 座席角色，0: 普通座席；1: 班长座席
	//
	// example:
	//
	// 0
	Role *int64 `json:"Role,omitempty" xml:"Role,omitempty"`
	// 座席状态，0: 离线；1: 在线
	//
	// example:
	//
	// 1
	Status *int64 `json:"Status,omitempty" xml:"Status,omitempty"`
	// 座席类型，1：全渠道、2：呼叫中心、3：在线客服
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
	// 修改时间
	//
	// example:
	//
	// 1775024775
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ClinkListClientsResponseBodyDataClients) String() string {
	return dara.Prettify(s)
}

func (s ClinkListClientsResponseBodyDataClients) GoString() string {
	return s.String()
}

func (s *ClinkListClientsResponseBodyDataClients) GetActive() *int64 {
	return s.Active
}

func (s *ClinkListClientsResponseBodyDataClients) GetAreaCode() *string {
	return s.AreaCode
}

func (s *ClinkListClientsResponseBodyDataClients) GetBindTel() *string {
	return s.BindTel
}

func (s *ClinkListClientsResponseBodyDataClients) GetCno() *string {
	return s.Cno
}

func (s *ClinkListClientsResponseBodyDataClients) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ClinkListClientsResponseBodyDataClients) GetCrmId() *string {
	return s.CrmId
}

func (s *ClinkListClientsResponseBodyDataClients) GetHiddenTel() *int64 {
	return s.HiddenTel
}

func (s *ClinkListClientsResponseBodyDataClients) GetId() *int64 {
	return s.Id
}

func (s *ClinkListClientsResponseBodyDataClients) GetName() *string {
	return s.Name
}

func (s *ClinkListClientsResponseBodyDataClients) GetQnos() []*string {
	return s.Qnos
}

func (s *ClinkListClientsResponseBodyDataClients) GetRole() *int64 {
	return s.Role
}

func (s *ClinkListClientsResponseBodyDataClients) GetStatus() *int64 {
	return s.Status
}

func (s *ClinkListClientsResponseBodyDataClients) GetType() *int64 {
	return s.Type
}

func (s *ClinkListClientsResponseBodyDataClients) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *ClinkListClientsResponseBodyDataClients) SetActive(v int64) *ClinkListClientsResponseBodyDataClients {
	s.Active = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetAreaCode(v string) *ClinkListClientsResponseBodyDataClients {
	s.AreaCode = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetBindTel(v string) *ClinkListClientsResponseBodyDataClients {
	s.BindTel = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetCno(v string) *ClinkListClientsResponseBodyDataClients {
	s.Cno = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetCreateTime(v int64) *ClinkListClientsResponseBodyDataClients {
	s.CreateTime = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetCrmId(v string) *ClinkListClientsResponseBodyDataClients {
	s.CrmId = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetHiddenTel(v int64) *ClinkListClientsResponseBodyDataClients {
	s.HiddenTel = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetId(v int64) *ClinkListClientsResponseBodyDataClients {
	s.Id = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetName(v string) *ClinkListClientsResponseBodyDataClients {
	s.Name = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetQnos(v []*string) *ClinkListClientsResponseBodyDataClients {
	s.Qnos = v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetRole(v int64) *ClinkListClientsResponseBodyDataClients {
	s.Role = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetStatus(v int64) *ClinkListClientsResponseBodyDataClients {
	s.Status = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetType(v int64) *ClinkListClientsResponseBodyDataClients {
	s.Type = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) SetUpdateTime(v int64) *ClinkListClientsResponseBodyDataClients {
	s.UpdateTime = &v
	return s
}

func (s *ClinkListClientsResponseBodyDataClients) Validate() error {
	return dara.Validate(s)
}
