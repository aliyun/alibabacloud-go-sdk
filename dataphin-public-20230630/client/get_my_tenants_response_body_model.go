// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMyTenantsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMyTenantsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetMyTenantsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMyTenantsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMyTenantsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMyTenantsResponseBody
	GetSuccess() *bool
	SetTenantList(v []*GetMyTenantsResponseBodyTenantList) *GetMyTenantsResponseBody
	GetTenantList() []*GetMyTenantsResponseBodyTenantList
}

type GetMyTenantsResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
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
	// example:
	//
	// true
	Success    *bool                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TenantList []*GetMyTenantsResponseBodyTenantList `json:"TenantList,omitempty" xml:"TenantList,omitempty" type:"Repeated"`
}

func (s GetMyTenantsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMyTenantsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMyTenantsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMyTenantsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMyTenantsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMyTenantsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMyTenantsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMyTenantsResponseBody) GetTenantList() []*GetMyTenantsResponseBodyTenantList {
	return s.TenantList
}

func (s *GetMyTenantsResponseBody) SetCode(v string) *GetMyTenantsResponseBody {
	s.Code = &v
	return s
}

func (s *GetMyTenantsResponseBody) SetHttpStatusCode(v int32) *GetMyTenantsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMyTenantsResponseBody) SetMessage(v string) *GetMyTenantsResponseBody {
	s.Message = &v
	return s
}

func (s *GetMyTenantsResponseBody) SetRequestId(v string) *GetMyTenantsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMyTenantsResponseBody) SetSuccess(v bool) *GetMyTenantsResponseBody {
	s.Success = &v
	return s
}

func (s *GetMyTenantsResponseBody) SetTenantList(v []*GetMyTenantsResponseBodyTenantList) *GetMyTenantsResponseBody {
	s.TenantList = v
	return s
}

func (s *GetMyTenantsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMyTenantsResponseBodyTenantList struct {
	// example:
	//
	// 1717343597000
	DeleteTime *int64 `json:"DeleteTime,omitempty" xml:"DeleteTime,omitempty"`
	// example:
	//
	// false
	Deleted     *bool   `json:"Deleted,omitempty" xml:"Deleted,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 132311
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// false
	OpsTenant *bool `json:"OpsTenant,omitempty" xml:"OpsTenant,omitempty"`
	// example:
	//
	// 21323231
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// false
	ResourceLimited *bool     `json:"ResourceLimited,omitempty" xml:"ResourceLimited,omitempty"`
	TenantTypeList  []*string `json:"TenantTypeList,omitempty" xml:"TenantTypeList,omitempty" type:"Repeated"`
	// example:
	//
	// icon
	TitleType *string `json:"TitleType,omitempty" xml:"TitleType,omitempty"`
	// example:
	//
	// true
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s GetMyTenantsResponseBodyTenantList) String() string {
	return dara.Prettify(s)
}

func (s GetMyTenantsResponseBodyTenantList) GoString() string {
	return s.String()
}

func (s *GetMyTenantsResponseBodyTenantList) GetDeleteTime() *int64 {
	return s.DeleteTime
}

func (s *GetMyTenantsResponseBodyTenantList) GetDeleted() *bool {
	return s.Deleted
}

func (s *GetMyTenantsResponseBodyTenantList) GetDescription() *string {
	return s.Description
}

func (s *GetMyTenantsResponseBodyTenantList) GetId() *int64 {
	return s.Id
}

func (s *GetMyTenantsResponseBodyTenantList) GetName() *string {
	return s.Name
}

func (s *GetMyTenantsResponseBodyTenantList) GetOpsTenant() *bool {
	return s.OpsTenant
}

func (s *GetMyTenantsResponseBodyTenantList) GetOwnerId() *string {
	return s.OwnerId
}

func (s *GetMyTenantsResponseBodyTenantList) GetResourceLimited() *bool {
	return s.ResourceLimited
}

func (s *GetMyTenantsResponseBodyTenantList) GetTenantTypeList() []*string {
	return s.TenantTypeList
}

func (s *GetMyTenantsResponseBodyTenantList) GetTitleType() *string {
	return s.TitleType
}

func (s *GetMyTenantsResponseBodyTenantList) GetVisible() *bool {
	return s.Visible
}

func (s *GetMyTenantsResponseBodyTenantList) SetDeleteTime(v int64) *GetMyTenantsResponseBodyTenantList {
	s.DeleteTime = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetDeleted(v bool) *GetMyTenantsResponseBodyTenantList {
	s.Deleted = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetDescription(v string) *GetMyTenantsResponseBodyTenantList {
	s.Description = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetId(v int64) *GetMyTenantsResponseBodyTenantList {
	s.Id = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetName(v string) *GetMyTenantsResponseBodyTenantList {
	s.Name = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetOpsTenant(v bool) *GetMyTenantsResponseBodyTenantList {
	s.OpsTenant = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetOwnerId(v string) *GetMyTenantsResponseBodyTenantList {
	s.OwnerId = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetResourceLimited(v bool) *GetMyTenantsResponseBodyTenantList {
	s.ResourceLimited = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetTenantTypeList(v []*string) *GetMyTenantsResponseBodyTenantList {
	s.TenantTypeList = v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetTitleType(v string) *GetMyTenantsResponseBodyTenantList {
	s.TitleType = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) SetVisible(v bool) *GetMyTenantsResponseBodyTenantList {
	s.Visible = &v
	return s
}

func (s *GetMyTenantsResponseBodyTenantList) Validate() error {
	return dara.Validate(s)
}
