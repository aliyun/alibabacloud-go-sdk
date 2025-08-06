// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddLicenseResponseBody
	GetCode() *string
	SetData(v []*AddLicenseResponseBodyData) *AddLicenseResponseBody
	GetData() []*AddLicenseResponseBodyData
	SetHttpStatusCode(v int32) *AddLicenseResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *AddLicenseResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddLicenseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddLicenseResponseBody
	GetSuccess() *bool
}

type AddLicenseResponseBody struct {
	Code           *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           []*AddLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	HttpStatusCode *int32                        `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *AddLicenseResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddLicenseResponseBody) GetData() []*AddLicenseResponseBodyData {
	return s.Data
}

func (s *AddLicenseResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddLicenseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLicenseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddLicenseResponseBody) SetCode(v string) *AddLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *AddLicenseResponseBody) SetData(v []*AddLicenseResponseBodyData) *AddLicenseResponseBody {
	s.Data = v
	return s
}

func (s *AddLicenseResponseBody) SetHttpStatusCode(v int32) *AddLicenseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddLicenseResponseBody) SetMessage(v string) *AddLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *AddLicenseResponseBody) SetRequestId(v string) *AddLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLicenseResponseBody) SetSuccess(v bool) *AddLicenseResponseBody {
	s.Success = &v
	return s
}

func (s *AddLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddLicenseResponseBodyData struct {
	AppItemId    *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
	BusinessType *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	ItemId       *string `json:"ItemId,omitempty" xml:"ItemId,omitempty"`
	LicenseId    *int64  `json:"LicenseId,omitempty" xml:"LicenseId,omitempty"`
	SdkType      *string `json:"SdkType,omitempty" xml:"SdkType,omitempty"`
}

func (s AddLicenseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s AddLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *AddLicenseResponseBodyData) GetAppItemId() *string {
	return s.AppItemId
}

func (s *AddLicenseResponseBodyData) GetBusinessType() *string {
	return s.BusinessType
}

func (s *AddLicenseResponseBodyData) GetItemId() *string {
	return s.ItemId
}

func (s *AddLicenseResponseBodyData) GetLicenseId() *int64 {
	return s.LicenseId
}

func (s *AddLicenseResponseBodyData) GetSdkType() *string {
	return s.SdkType
}

func (s *AddLicenseResponseBodyData) SetAppItemId(v string) *AddLicenseResponseBodyData {
	s.AppItemId = &v
	return s
}

func (s *AddLicenseResponseBodyData) SetBusinessType(v string) *AddLicenseResponseBodyData {
	s.BusinessType = &v
	return s
}

func (s *AddLicenseResponseBodyData) SetItemId(v string) *AddLicenseResponseBodyData {
	s.ItemId = &v
	return s
}

func (s *AddLicenseResponseBodyData) SetLicenseId(v int64) *AddLicenseResponseBodyData {
	s.LicenseId = &v
	return s
}

func (s *AddLicenseResponseBodyData) SetSdkType(v string) *AddLicenseResponseBodyData {
	s.SdkType = &v
	return s
}

func (s *AddLicenseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
