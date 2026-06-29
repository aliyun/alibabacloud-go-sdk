// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataDomainsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDataDomainsResponseBody
	GetCode() *string
	SetData(v *ListDataDomainsResponseBodyData) *ListDataDomainsResponseBody
	GetData() *ListDataDomainsResponseBodyData
	SetHttpStatusCode(v int32) *ListDataDomainsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListDataDomainsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListDataDomainsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListDataDomainsResponseBody
	GetSuccess() *bool
}

type ListDataDomainsResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data *ListDataDomainsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned by the backend.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDataDomainsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataDomainsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDataDomainsResponseBody) GetData() *ListDataDomainsResponseBodyData {
	return s.Data
}

func (s *ListDataDomainsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDataDomainsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDataDomainsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataDomainsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDataDomainsResponseBody) SetCode(v string) *ListDataDomainsResponseBody {
	s.Code = &v
	return s
}

func (s *ListDataDomainsResponseBody) SetData(v *ListDataDomainsResponseBodyData) *ListDataDomainsResponseBody {
	s.Data = v
	return s
}

func (s *ListDataDomainsResponseBody) SetHttpStatusCode(v int32) *ListDataDomainsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDataDomainsResponseBody) SetMessage(v string) *ListDataDomainsResponseBody {
	s.Message = &v
	return s
}

func (s *ListDataDomainsResponseBody) SetRequestId(v string) *ListDataDomainsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataDomainsResponseBody) SetSuccess(v bool) *ListDataDomainsResponseBody {
	s.Success = &v
	return s
}

func (s *ListDataDomainsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataDomainsResponseBodyData struct {
	// The details of the data domains.
	DataDomainList []*ListDataDomainsResponseBodyDataDataDomainList `json:"DataDomainList,omitempty" xml:"DataDomainList,omitempty" type:"Repeated"`
}

func (s ListDataDomainsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataDomainsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataDomainsResponseBodyData) GetDataDomainList() []*ListDataDomainsResponseBodyDataDataDomainList {
	return s.DataDomainList
}

func (s *ListDataDomainsResponseBodyData) SetDataDomainList(v []*ListDataDomainsResponseBodyDataDataDomainList) *ListDataDomainsResponseBodyData {
	s.DataDomainList = v
	return s
}

func (s *ListDataDomainsResponseBodyData) Validate() error {
	if s.DataDomainList != nil {
		for _, item := range s.DataDomainList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataDomainsResponseBodyDataDataDomainList struct {
	// The abbreviation of the data domain.
	//
	// example:
	//
	// dm_code_name
	Abbreviation *string `json:"Abbreviation,omitempty" xml:"Abbreviation,omitempty"`
	// The ID of the business unit to which the data domain belongs.
	//
	// example:
	//
	// 545844456
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The description of the business object.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name of the data domain.
	//
	// example:
	//
	// 测试
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-10-10 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2024-10-10 10:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the data domain.
	//
	// example:
	//
	// 1241844456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the user who last modified the data domain.
	//
	// example:
	//
	// 30010010
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// The name of the user who last modified the data domain.
	//
	// example:
	//
	// 张三
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// The name of the data domain.
	//
	// example:
	//
	// dm_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the owner of the business object.
	//
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The ID of the owner of the business object.
	//
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The IDs of the parent data domains.
	//
	// example:
	//
	// 10232311
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s ListDataDomainsResponseBodyDataDataDomainList) String() string {
	return dara.Prettify(s)
}

func (s ListDataDomainsResponseBodyDataDataDomainList) GoString() string {
	return s.String()
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetAbbreviation() *string {
	return s.Abbreviation
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetDescription() *string {
	return s.Description
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetId() *int64 {
	return s.Id
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetLastModifier() *string {
	return s.LastModifier
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetName() *string {
	return s.Name
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) GetParentId() *int64 {
	return s.ParentId
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetAbbreviation(v string) *ListDataDomainsResponseBodyDataDataDomainList {
	s.Abbreviation = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetBizUnitId(v int64) *ListDataDomainsResponseBodyDataDataDomainList {
	s.BizUnitId = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetDescription(v string) *ListDataDomainsResponseBodyDataDataDomainList {
	s.Description = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetDisplayName(v string) *ListDataDomainsResponseBodyDataDataDomainList {
	s.DisplayName = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetGmtCreate(v string) *ListDataDomainsResponseBodyDataDataDomainList {
	s.GmtCreate = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetGmtModified(v string) *ListDataDomainsResponseBodyDataDataDomainList {
	s.GmtModified = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetId(v int64) *ListDataDomainsResponseBodyDataDataDomainList {
	s.Id = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetLastModifier(v string) *ListDataDomainsResponseBodyDataDataDomainList {
	s.LastModifier = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetLastModifierName(v string) *ListDataDomainsResponseBodyDataDataDomainList {
	s.LastModifierName = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetName(v string) *ListDataDomainsResponseBodyDataDataDomainList {
	s.Name = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetOwnerName(v string) *ListDataDomainsResponseBodyDataDataDomainList {
	s.OwnerName = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetOwnerUserId(v string) *ListDataDomainsResponseBodyDataDataDomainList {
	s.OwnerUserId = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) SetParentId(v int64) *ListDataDomainsResponseBodyDataDataDomainList {
	s.ParentId = &v
	return s
}

func (s *ListDataDomainsResponseBodyDataDataDomainList) Validate() error {
	return dara.Validate(s)
}
