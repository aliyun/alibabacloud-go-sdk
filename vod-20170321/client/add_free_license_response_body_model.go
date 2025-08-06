// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFreeLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddFreeLicenseResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *AddFreeLicenseResponseBody
	GetHttpStatusCode() *int32
	SetLicenseList(v []*AddFreeLicenseResponseBodyLicenseList) *AddFreeLicenseResponseBody
	GetLicenseList() []*AddFreeLicenseResponseBodyLicenseList
	SetMessage(v string) *AddFreeLicenseResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddFreeLicenseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *AddFreeLicenseResponseBody
	GetSuccess() *bool
}

type AddFreeLicenseResponseBody struct {
	Code           *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32                                   `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	LicenseList    []*AddFreeLicenseResponseBodyLicenseList `json:"LicenseList,omitempty" xml:"LicenseList,omitempty" type:"Repeated"`
	Message        *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddFreeLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddFreeLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *AddFreeLicenseResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddFreeLicenseResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AddFreeLicenseResponseBody) GetLicenseList() []*AddFreeLicenseResponseBodyLicenseList {
	return s.LicenseList
}

func (s *AddFreeLicenseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddFreeLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddFreeLicenseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddFreeLicenseResponseBody) SetCode(v string) *AddFreeLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *AddFreeLicenseResponseBody) SetHttpStatusCode(v int32) *AddFreeLicenseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AddFreeLicenseResponseBody) SetLicenseList(v []*AddFreeLicenseResponseBodyLicenseList) *AddFreeLicenseResponseBody {
	s.LicenseList = v
	return s
}

func (s *AddFreeLicenseResponseBody) SetMessage(v string) *AddFreeLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *AddFreeLicenseResponseBody) SetRequestId(v string) *AddFreeLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddFreeLicenseResponseBody) SetSuccess(v bool) *AddFreeLicenseResponseBody {
	s.Success = &v
	return s
}

func (s *AddFreeLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}

type AddFreeLicenseResponseBodyLicenseList struct {
	AppItemId     *string `json:"AppItemId,omitempty" xml:"AppItemId,omitempty"`
	BusinessType  *string `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	LicenseId     *int64  `json:"LicenseId,omitempty" xml:"LicenseId,omitempty"`
	LicenseItemId *string `json:"LicenseItemId,omitempty" xml:"LicenseItemId,omitempty"`
	SdkType       *string `json:"SdkType,omitempty" xml:"SdkType,omitempty"`
}

func (s AddFreeLicenseResponseBodyLicenseList) String() string {
	return dara.Prettify(s)
}

func (s AddFreeLicenseResponseBodyLicenseList) GoString() string {
	return s.String()
}

func (s *AddFreeLicenseResponseBodyLicenseList) GetAppItemId() *string {
	return s.AppItemId
}

func (s *AddFreeLicenseResponseBodyLicenseList) GetBusinessType() *string {
	return s.BusinessType
}

func (s *AddFreeLicenseResponseBodyLicenseList) GetLicenseId() *int64 {
	return s.LicenseId
}

func (s *AddFreeLicenseResponseBodyLicenseList) GetLicenseItemId() *string {
	return s.LicenseItemId
}

func (s *AddFreeLicenseResponseBodyLicenseList) GetSdkType() *string {
	return s.SdkType
}

func (s *AddFreeLicenseResponseBodyLicenseList) SetAppItemId(v string) *AddFreeLicenseResponseBodyLicenseList {
	s.AppItemId = &v
	return s
}

func (s *AddFreeLicenseResponseBodyLicenseList) SetBusinessType(v string) *AddFreeLicenseResponseBodyLicenseList {
	s.BusinessType = &v
	return s
}

func (s *AddFreeLicenseResponseBodyLicenseList) SetLicenseId(v int64) *AddFreeLicenseResponseBodyLicenseList {
	s.LicenseId = &v
	return s
}

func (s *AddFreeLicenseResponseBodyLicenseList) SetLicenseItemId(v string) *AddFreeLicenseResponseBodyLicenseList {
	s.LicenseItemId = &v
	return s
}

func (s *AddFreeLicenseResponseBodyLicenseList) SetSdkType(v string) *AddFreeLicenseResponseBodyLicenseList {
	s.SdkType = &v
	return s
}

func (s *AddFreeLicenseResponseBodyLicenseList) Validate() error {
	return dara.Validate(s)
}
