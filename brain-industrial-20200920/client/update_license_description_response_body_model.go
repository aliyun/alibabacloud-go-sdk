// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLicenseDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *UpdateLicenseDescriptionResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *UpdateLicenseDescriptionResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *UpdateLicenseDescriptionResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateLicenseDescriptionResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateLicenseDescriptionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateLicenseDescriptionResponseBody
	GetSuccess() *string
}

type UpdateLicenseDescriptionResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// 200
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	HttpStatusCode *int32  `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 68738E75-43C1-5AE5-9F3A-AFEF576D7B5F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateLicenseDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLicenseDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLicenseDescriptionResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *UpdateLicenseDescriptionResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateLicenseDescriptionResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateLicenseDescriptionResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLicenseDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLicenseDescriptionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateLicenseDescriptionResponseBody) SetAccessDeniedDetail(v string) *UpdateLicenseDescriptionResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *UpdateLicenseDescriptionResponseBody) SetCode(v string) *UpdateLicenseDescriptionResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLicenseDescriptionResponseBody) SetHttpStatusCode(v int32) *UpdateLicenseDescriptionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateLicenseDescriptionResponseBody) SetMessage(v string) *UpdateLicenseDescriptionResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLicenseDescriptionResponseBody) SetRequestId(v string) *UpdateLicenseDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLicenseDescriptionResponseBody) SetSuccess(v string) *UpdateLicenseDescriptionResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLicenseDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
