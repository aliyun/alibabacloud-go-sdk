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
	// example:
	//
	// OK
	Code *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListDataDomainsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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
	// example:
	//
	// dm_code_name
	Abbreviation *string `json:"Abbreviation,omitempty" xml:"Abbreviation,omitempty"`
	// example:
	//
	// 545844456
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 测试
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// example:
	//
	// 2024-10-10 10:00:00
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2024-10-10 10:00:00
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// 1241844456
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 30010010
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// example:
	//
	// 张三
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// example:
	//
	// dm_code_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
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
