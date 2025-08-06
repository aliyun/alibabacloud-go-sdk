// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewFreeLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RenewFreeLicenseResponseBody
	GetCode() *string
	SetFailedLicenseItemIdList(v []*string) *RenewFreeLicenseResponseBody
	GetFailedLicenseItemIdList() []*string
	SetHttpStatusCode(v int32) *RenewFreeLicenseResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RenewFreeLicenseResponseBody
	GetMessage() *string
	SetRequestId(v string) *RenewFreeLicenseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RenewFreeLicenseResponseBody
	GetSuccess() *bool
}

type RenewFreeLicenseResponseBody struct {
	Code                    *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	FailedLicenseItemIdList []*string `json:"FailedLicenseItemIdList,omitempty" xml:"FailedLicenseItemIdList,omitempty" type:"Repeated"`
	HttpStatusCode          *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message                 *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId               *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                 *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RenewFreeLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewFreeLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *RenewFreeLicenseResponseBody) GetCode() *string {
	return s.Code
}

func (s *RenewFreeLicenseResponseBody) GetFailedLicenseItemIdList() []*string {
	return s.FailedLicenseItemIdList
}

func (s *RenewFreeLicenseResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RenewFreeLicenseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RenewFreeLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewFreeLicenseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RenewFreeLicenseResponseBody) SetCode(v string) *RenewFreeLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *RenewFreeLicenseResponseBody) SetFailedLicenseItemIdList(v []*string) *RenewFreeLicenseResponseBody {
	s.FailedLicenseItemIdList = v
	return s
}

func (s *RenewFreeLicenseResponseBody) SetHttpStatusCode(v int32) *RenewFreeLicenseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RenewFreeLicenseResponseBody) SetMessage(v string) *RenewFreeLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *RenewFreeLicenseResponseBody) SetRequestId(v string) *RenewFreeLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewFreeLicenseResponseBody) SetSuccess(v bool) *RenewFreeLicenseResponseBody {
	s.Success = &v
	return s
}

func (s *RenewFreeLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
