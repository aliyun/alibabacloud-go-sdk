// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFreeLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteFreeLicenseResponseBody
	GetCode() *string
	SetFailedLicenseItemIdList(v []*string) *DeleteFreeLicenseResponseBody
	GetFailedLicenseItemIdList() []*string
	SetHttpStatusCode(v int32) *DeleteFreeLicenseResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteFreeLicenseResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteFreeLicenseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFreeLicenseResponseBody
	GetSuccess() *bool
}

type DeleteFreeLicenseResponseBody struct {
	Code                    *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	FailedLicenseItemIdList []*string `json:"FailedLicenseItemIdList,omitempty" xml:"FailedLicenseItemIdList,omitempty" type:"Repeated"`
	HttpStatusCode          *int32    `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	Message                 *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId               *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success                 *bool     `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFreeLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFreeLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFreeLicenseResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteFreeLicenseResponseBody) GetFailedLicenseItemIdList() []*string {
	return s.FailedLicenseItemIdList
}

func (s *DeleteFreeLicenseResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteFreeLicenseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteFreeLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFreeLicenseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFreeLicenseResponseBody) SetCode(v string) *DeleteFreeLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteFreeLicenseResponseBody) SetFailedLicenseItemIdList(v []*string) *DeleteFreeLicenseResponseBody {
	s.FailedLicenseItemIdList = v
	return s
}

func (s *DeleteFreeLicenseResponseBody) SetHttpStatusCode(v int32) *DeleteFreeLicenseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteFreeLicenseResponseBody) SetMessage(v string) *DeleteFreeLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFreeLicenseResponseBody) SetRequestId(v string) *DeleteFreeLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFreeLicenseResponseBody) SetSuccess(v bool) *DeleteFreeLicenseResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFreeLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
