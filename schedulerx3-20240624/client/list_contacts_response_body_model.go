// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListContactsResponseBody
	GetCode() *int32
	SetData(v *ListContactsResponseBodyData) *ListContactsResponseBody
	GetData() *ListContactsResponseBodyData
	SetMessage(v string) *ListContactsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListContactsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListContactsResponseBody
	GetSuccess() *bool
}

type ListContactsResponseBody struct {
	// example:
	//
	// 200
	Code *int32                        `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListContactsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 5EF879D0-3B43-5AD1-9BF7-52418F9C5E73
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListContactsResponseBody) GoString() string {
	return s.String()
}

func (s *ListContactsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListContactsResponseBody) GetData() *ListContactsResponseBodyData {
	return s.Data
}

func (s *ListContactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListContactsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListContactsResponseBody) SetCode(v int32) *ListContactsResponseBody {
	s.Code = &v
	return s
}

func (s *ListContactsResponseBody) SetData(v *ListContactsResponseBodyData) *ListContactsResponseBody {
	s.Data = v
	return s
}

func (s *ListContactsResponseBody) SetMessage(v string) *ListContactsResponseBody {
	s.Message = &v
	return s
}

func (s *ListContactsResponseBody) SetRequestId(v string) *ListContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListContactsResponseBody) SetSuccess(v bool) *ListContactsResponseBody {
	s.Success = &v
	return s
}

func (s *ListContactsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListContactsResponseBodyData struct {
	// 当前页码
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// 每页条数
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// 联系人列表
	Records []*ListContactsResponseBodyDataRecords `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// 总记录数
	//
	// example:
	//
	// 21
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListContactsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListContactsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListContactsResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListContactsResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListContactsResponseBodyData) GetRecords() []*ListContactsResponseBodyDataRecords {
	return s.Records
}

func (s *ListContactsResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ListContactsResponseBodyData) SetPageNumber(v int32) *ListContactsResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListContactsResponseBodyData) SetPageSize(v int32) *ListContactsResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListContactsResponseBodyData) SetRecords(v []*ListContactsResponseBodyDataRecords) *ListContactsResponseBodyData {
	s.Records = v
	return s
}

func (s *ListContactsResponseBodyData) SetTotal(v int32) *ListContactsResponseBodyData {
	s.Total = &v
	return s
}

func (s *ListContactsResponseBodyData) Validate() error {
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

type ListContactsResponseBodyDataRecords struct {
	// 渠道配置 JSON（clientSecret 已脱敏为 ***）
	//
	// example:
	//
	// {"channels":[{"channelType":"dingtalk","clientId":"xxx","clientSecret":"xxx","targetType":"group","targetId":"xxx","robotCode":"xxx"}]}
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// 联系人 ID
	//
	// example:
	//
	// job-85d64bff-50b5-4d02-aa3f-19956b17449d
	ContactId *int64 `json:"ContactId,omitempty" xml:"ContactId,omitempty"`
	// 联系人名称
	//
	// example:
	//
	// Tom
	ContactName *string `json:"ContactName,omitempty" xml:"ContactName,omitempty"`
	// 是否启用
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// 创建时间
	//
	// example:
	//
	// 1783065190000
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// 最后修改时间
	//
	// example:
	//
	// 1783065190000
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// 渠道大类（IM/WEBHOOK/SMS/MAIL）
	//
	// example:
	//
	// IM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListContactsResponseBodyDataRecords) String() string {
	return dara.Prettify(s)
}

func (s ListContactsResponseBodyDataRecords) GoString() string {
	return s.String()
}

func (s *ListContactsResponseBodyDataRecords) GetConfig() *string {
	return s.Config
}

func (s *ListContactsResponseBodyDataRecords) GetContactId() *int64 {
	return s.ContactId
}

func (s *ListContactsResponseBodyDataRecords) GetContactName() *string {
	return s.ContactName
}

func (s *ListContactsResponseBodyDataRecords) GetEnabled() *bool {
	return s.Enabled
}

func (s *ListContactsResponseBodyDataRecords) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListContactsResponseBodyDataRecords) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListContactsResponseBodyDataRecords) GetType() *string {
	return s.Type
}

func (s *ListContactsResponseBodyDataRecords) SetConfig(v string) *ListContactsResponseBodyDataRecords {
	s.Config = &v
	return s
}

func (s *ListContactsResponseBodyDataRecords) SetContactId(v int64) *ListContactsResponseBodyDataRecords {
	s.ContactId = &v
	return s
}

func (s *ListContactsResponseBodyDataRecords) SetContactName(v string) *ListContactsResponseBodyDataRecords {
	s.ContactName = &v
	return s
}

func (s *ListContactsResponseBodyDataRecords) SetEnabled(v bool) *ListContactsResponseBodyDataRecords {
	s.Enabled = &v
	return s
}

func (s *ListContactsResponseBodyDataRecords) SetGmtCreate(v string) *ListContactsResponseBodyDataRecords {
	s.GmtCreate = &v
	return s
}

func (s *ListContactsResponseBodyDataRecords) SetGmtModified(v string) *ListContactsResponseBodyDataRecords {
	s.GmtModified = &v
	return s
}

func (s *ListContactsResponseBodyDataRecords) SetType(v string) *ListContactsResponseBodyDataRecords {
	s.Type = &v
	return s
}

func (s *ListContactsResponseBodyDataRecords) Validate() error {
	return dara.Validate(s)
}
