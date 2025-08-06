// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDelFreeLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DelFreeLicenseResponseBody
	GetCode() *string
	SetFailedLicenseItemIdList(v []*string) *DelFreeLicenseResponseBody
	GetFailedLicenseItemIdList() []*string
	SetHttpStatusCode(v int32) *DelFreeLicenseResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DelFreeLicenseResponseBody
	GetMessage() *string
	SetRequestId(v string) *DelFreeLicenseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DelFreeLicenseResponseBody
	GetSuccess() *bool
}

type DelFreeLicenseResponseBody struct {
	Code                    *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	FailedLicenseItemIdList []*string `json:"FailedLicenseItemIdList,omitempty" xml:"FailedLicenseItemIdList,omitempty" type:"Repeated"`
	HttpStatusCode          *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message                 *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId               *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                 *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DelFreeLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DelFreeLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *DelFreeLicenseResponseBody) GetCode() *string {
	return s.Code
}

func (s *DelFreeLicenseResponseBody) GetFailedLicenseItemIdList() []*string {
	return s.FailedLicenseItemIdList
}

func (s *DelFreeLicenseResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DelFreeLicenseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DelFreeLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DelFreeLicenseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DelFreeLicenseResponseBody) SetCode(v string) *DelFreeLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *DelFreeLicenseResponseBody) SetFailedLicenseItemIdList(v []*string) *DelFreeLicenseResponseBody {
	s.FailedLicenseItemIdList = v
	return s
}

func (s *DelFreeLicenseResponseBody) SetHttpStatusCode(v int32) *DelFreeLicenseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DelFreeLicenseResponseBody) SetMessage(v string) *DelFreeLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *DelFreeLicenseResponseBody) SetRequestId(v string) *DelFreeLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DelFreeLicenseResponseBody) SetSuccess(v bool) *DelFreeLicenseResponseBody {
	s.Success = &v
	return s
}

func (s *DelFreeLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
