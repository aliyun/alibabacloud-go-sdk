// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBizUnitsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBizUnitsResponseBody
	GetCode() *string
	SetData(v *ListBizUnitsResponseBodyData) *ListBizUnitsResponseBody
	GetData() *ListBizUnitsResponseBodyData
	SetHttpStatusCode(v int32) *ListBizUnitsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBizUnitsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBizUnitsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBizUnitsResponseBody
	GetSuccess() *bool
}

type ListBizUnitsResponseBody struct {
	// The error code. A value of OK indicates that the request was successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The query result.
	Data *ListBizUnitsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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

func (s ListBizUnitsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBizUnitsResponseBody) GoString() string {
	return s.String()
}

func (s *ListBizUnitsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBizUnitsResponseBody) GetData() *ListBizUnitsResponseBodyData {
	return s.Data
}

func (s *ListBizUnitsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBizUnitsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBizUnitsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBizUnitsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBizUnitsResponseBody) SetCode(v string) *ListBizUnitsResponseBody {
	s.Code = &v
	return s
}

func (s *ListBizUnitsResponseBody) SetData(v *ListBizUnitsResponseBodyData) *ListBizUnitsResponseBody {
	s.Data = v
	return s
}

func (s *ListBizUnitsResponseBody) SetHttpStatusCode(v int32) *ListBizUnitsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBizUnitsResponseBody) SetMessage(v string) *ListBizUnitsResponseBody {
	s.Message = &v
	return s
}

func (s *ListBizUnitsResponseBody) SetRequestId(v string) *ListBizUnitsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBizUnitsResponseBody) SetSuccess(v bool) *ListBizUnitsResponseBody {
	s.Success = &v
	return s
}

func (s *ListBizUnitsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListBizUnitsResponseBodyData struct {
	// The business unit details.
	BizUnitList []*ListBizUnitsResponseBodyDataBizUnitList `json:"BizUnitList,omitempty" xml:"BizUnitList,omitempty" type:"Repeated"`
}

func (s ListBizUnitsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListBizUnitsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListBizUnitsResponseBodyData) GetBizUnitList() []*ListBizUnitsResponseBodyDataBizUnitList {
	return s.BizUnitList
}

func (s *ListBizUnitsResponseBodyData) SetBizUnitList(v []*ListBizUnitsResponseBodyDataBizUnitList) *ListBizUnitsResponseBodyData {
	s.BizUnitList = v
	return s
}

func (s *ListBizUnitsResponseBodyData) Validate() error {
	if s.BizUnitList != nil {
		for _, item := range s.BizUnitList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBizUnitsResponseBodyDataBizUnitList struct {
	// The business unit architects.
	AccountList []*ListBizUnitsResponseBodyDataBizUnitListAccountList `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Repeated"`
	// The description of the business object.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The display name.
	//
	// example:
	//
	// 测试
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The environment identifier. Valid values:
	//
	// - DEV: development environment.
	//
	// - PROD: production environment.
	//
	// example:
	//
	// DEV
	Env *string `json:"Env,omitempty" xml:"Env,omitempty"`
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
	// The business unit icon.
	//
	// example:
	//
	// icon-e-commerce
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// The business unit ID.
	//
	// example:
	//
	// 101001201
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the user who last modified the business unit.
	//
	// example:
	//
	// 30010010
	LastModifier *string `json:"LastModifier,omitempty" xml:"LastModifier,omitempty"`
	// The name of the user who last modified the business unit.
	//
	// example:
	//
	// 张三
	LastModifierName *string `json:"LastModifierName,omitempty" xml:"LastModifierName,omitempty"`
	// The production mode. Valid values:
	//
	// - BASIC: single-environment mode.
	//
	// - DEV_PROD: development/production dual-environment mode.
	//
	// example:
	//
	// DEV_PROD
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name.
	//
	// example:
	//
	// test01
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner of the business object.
	//
	// example:
	//
	// 张三
	OwnerName *string `json:"OwnerName,omitempty" xml:"OwnerName,omitempty"`
	// The owner of the business object.
	//
	// example:
	//
	// 30010010
	OwnerUserId *string `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
}

func (s ListBizUnitsResponseBodyDataBizUnitList) String() string {
	return dara.Prettify(s)
}

func (s ListBizUnitsResponseBodyDataBizUnitList) GoString() string {
	return s.String()
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetAccountList() []*ListBizUnitsResponseBodyDataBizUnitListAccountList {
	return s.AccountList
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetDescription() *string {
	return s.Description
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetEnv() *string {
	return s.Env
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetIcon() *string {
	return s.Icon
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetId() *int64 {
	return s.Id
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetLastModifier() *string {
	return s.LastModifier
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetLastModifierName() *string {
	return s.LastModifierName
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetMode() *string {
	return s.Mode
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetName() *string {
	return s.Name
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetOwnerName() *string {
	return s.OwnerName
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) GetOwnerUserId() *string {
	return s.OwnerUserId
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetAccountList(v []*ListBizUnitsResponseBodyDataBizUnitListAccountList) *ListBizUnitsResponseBodyDataBizUnitList {
	s.AccountList = v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetDescription(v string) *ListBizUnitsResponseBodyDataBizUnitList {
	s.Description = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetDisplayName(v string) *ListBizUnitsResponseBodyDataBizUnitList {
	s.DisplayName = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetEnv(v string) *ListBizUnitsResponseBodyDataBizUnitList {
	s.Env = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetGmtCreate(v string) *ListBizUnitsResponseBodyDataBizUnitList {
	s.GmtCreate = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetGmtModified(v string) *ListBizUnitsResponseBodyDataBizUnitList {
	s.GmtModified = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetIcon(v string) *ListBizUnitsResponseBodyDataBizUnitList {
	s.Icon = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetId(v int64) *ListBizUnitsResponseBodyDataBizUnitList {
	s.Id = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetLastModifier(v string) *ListBizUnitsResponseBodyDataBizUnitList {
	s.LastModifier = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetLastModifierName(v string) *ListBizUnitsResponseBodyDataBizUnitList {
	s.LastModifierName = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetMode(v string) *ListBizUnitsResponseBodyDataBizUnitList {
	s.Mode = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetName(v string) *ListBizUnitsResponseBodyDataBizUnitList {
	s.Name = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetOwnerName(v string) *ListBizUnitsResponseBodyDataBizUnitList {
	s.OwnerName = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) SetOwnerUserId(v string) *ListBizUnitsResponseBodyDataBizUnitList {
	s.OwnerUserId = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitList) Validate() error {
	if s.AccountList != nil {
		for _, item := range s.AccountList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListBizUnitsResponseBodyDataBizUnitListAccountList struct {
	// The user ID.
	//
	// example:
	//
	// 20001201
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListBizUnitsResponseBodyDataBizUnitListAccountList) String() string {
	return dara.Prettify(s)
}

func (s ListBizUnitsResponseBodyDataBizUnitListAccountList) GoString() string {
	return s.String()
}

func (s *ListBizUnitsResponseBodyDataBizUnitListAccountList) GetId() *string {
	return s.Id
}

func (s *ListBizUnitsResponseBodyDataBizUnitListAccountList) SetId(v string) *ListBizUnitsResponseBodyDataBizUnitListAccountList {
	s.Id = &v
	return s
}

func (s *ListBizUnitsResponseBodyDataBizUnitListAccountList) Validate() error {
	return dara.Validate(s)
}
