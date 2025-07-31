// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataDomainInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetDataDomainInfoResponseBody
	GetCode() *string
	SetDataDomainInfo(v *GetDataDomainInfoResponseBodyDataDomainInfo) *GetDataDomainInfoResponseBody
	GetDataDomainInfo() *GetDataDomainInfoResponseBodyDataDomainInfo
	SetHttpStatusCode(v int32) *GetDataDomainInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetDataDomainInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetDataDomainInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataDomainInfoResponseBody
	GetSuccess() *bool
}

type GetDataDomainInfoResponseBody struct {
	// example:
	//
	// OK
	Code           *string                                      `json:"Code,omitempty" xml:"Code,omitempty"`
	DataDomainInfo *GetDataDomainInfoResponseBodyDataDomainInfo `json:"DataDomainInfo,omitempty" xml:"DataDomainInfo,omitempty" type:"Struct"`
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

func (s GetDataDomainInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataDomainInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataDomainInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetDataDomainInfoResponseBody) GetDataDomainInfo() *GetDataDomainInfoResponseBodyDataDomainInfo {
	return s.DataDomainInfo
}

func (s *GetDataDomainInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetDataDomainInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetDataDomainInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataDomainInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataDomainInfoResponseBody) SetCode(v string) *GetDataDomainInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetDataDomainInfoResponseBody) SetDataDomainInfo(v *GetDataDomainInfoResponseBodyDataDomainInfo) *GetDataDomainInfoResponseBody {
	s.DataDomainInfo = v
	return s
}

func (s *GetDataDomainInfoResponseBody) SetHttpStatusCode(v int32) *GetDataDomainInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetDataDomainInfoResponseBody) SetMessage(v string) *GetDataDomainInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetDataDomainInfoResponseBody) SetRequestId(v string) *GetDataDomainInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataDomainInfoResponseBody) SetSuccess(v bool) *GetDataDomainInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataDomainInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetDataDomainInfoResponseBodyDataDomainInfo struct {
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

func (s GetDataDomainInfoResponseBodyDataDomainInfo) String() string {
	return dara.Prettify(s)
}

func (s GetDataDomainInfoResponseBodyDataDomainInfo) GoString() string {
	return s.String()
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetAbbreviation() *string {
	return s.Abbreviation
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetDescription() *string {
	return s.Description
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetDisplayName() *string {
	return s.DisplayName
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetId() *int64 {
	return s.Id
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetLastModifier() *string {
	return s.LastModifier
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetName() *string {
	return s.Name
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetOwnerName() *string {
	return s.OwnerName
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetAbbreviation(v string) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.Abbreviation = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetBizUnitId(v int64) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.BizUnitId = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetDescription(v string) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.Description = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetDisplayName(v string) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.DisplayName = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetGmtCreate(v string) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.GmtCreate = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetGmtModified(v string) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.GmtModified = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetId(v int64) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.Id = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetLastModifier(v string) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.LastModifier = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetLastModifierName(v string) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.LastModifierName = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetName(v string) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.Name = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetOwnerName(v string) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.OwnerName = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetOwnerUserId(v string) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.OwnerUserId = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) SetParentId(v int64) *GetDataDomainInfoResponseBodyDataDomainInfo {
	s.ParentId = &v
	return s
}

func (s *GetDataDomainInfoResponseBodyDataDomainInfo) Validate() error {
	return dara.Validate(s)
}
