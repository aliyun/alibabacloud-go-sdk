// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewAppLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RenewAppLicenseResponseBody
	GetCode() *string
	SetData(v *RenewAppLicenseResponseBodyData) *RenewAppLicenseResponseBody
	GetData() *RenewAppLicenseResponseBodyData
	SetHttpStatusCode(v int32) *RenewAppLicenseResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RenewAppLicenseResponseBody
	GetMessage() *string
	SetRequestId(v string) *RenewAppLicenseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenewAppLicenseResponseBody
	GetSuccess() *bool
}

type RenewAppLicenseResponseBody struct {
	Code           *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Data           *RenewAppLicenseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	HttpStatusCode *int32                           `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message        *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success        *bool                            `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewAppLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewAppLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RenewAppLicenseResponseBody) GetCode() *string {
	return s.Code
}

func (s *RenewAppLicenseResponseBody) GetData() *RenewAppLicenseResponseBodyData {
	return s.Data
}

func (s *RenewAppLicenseResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RenewAppLicenseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RenewAppLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewAppLicenseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenewAppLicenseResponseBody) SetCode(v string) *RenewAppLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RenewAppLicenseResponseBody) SetData(v *RenewAppLicenseResponseBodyData) *RenewAppLicenseResponseBody {
	s.Data = v
	return s
}

func (s *RenewAppLicenseResponseBody) SetHttpStatusCode(v int32) *RenewAppLicenseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RenewAppLicenseResponseBody) SetMessage(v string) *RenewAppLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RenewAppLicenseResponseBody) SetRequestId(v string) *RenewAppLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewAppLicenseResponseBody) SetSuccess(v bool) *RenewAppLicenseResponseBody {
	s.Success = &v
	return s
}

func (s *RenewAppLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}

type RenewAppLicenseResponseBodyData struct {
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s RenewAppLicenseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RenewAppLicenseResponseBodyData) GoString() string {
	return s.String()
}

func (s *RenewAppLicenseResponseBodyData) GetResult() *bool {
	return s.Result
}

func (s *RenewAppLicenseResponseBodyData) SetResult(v bool) *RenewAppLicenseResponseBodyData {
	s.Result = &v
	return s
}

func (s *RenewAppLicenseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
