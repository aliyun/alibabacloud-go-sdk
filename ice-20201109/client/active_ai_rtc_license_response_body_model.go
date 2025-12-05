// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iActiveAiRtcLicenseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ActiveAiRtcLicenseResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ActiveAiRtcLicenseResponseBody
	GetHttpStatusCode() *int32
	SetLicense(v string) *ActiveAiRtcLicenseResponseBody
	GetLicense() *string
	SetMessage(v string) *ActiveAiRtcLicenseResponseBody
	GetMessage() *string
	SetRequestId(v string) *ActiveAiRtcLicenseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ActiveAiRtcLicenseResponseBody
	GetSuccess() *bool
}

type ActiveAiRtcLicenseResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The license information.
	//
	// example:
	//
	// a659a06659a***
	License *string `json:"License,omitempty" xml:"License,omitempty"`
	// The error message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4F9C14FE-1147-15AC-8EDF-A590FF12***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ActiveAiRtcLicenseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ActiveAiRtcLicenseResponseBody) GoString() string {
	return s.String()
}

func (s *ActiveAiRtcLicenseResponseBody) GetCode() *string {
	return s.Code
}

func (s *ActiveAiRtcLicenseResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ActiveAiRtcLicenseResponseBody) GetLicense() *string {
	return s.License
}

func (s *ActiveAiRtcLicenseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ActiveAiRtcLicenseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ActiveAiRtcLicenseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ActiveAiRtcLicenseResponseBody) SetCode(v string) *ActiveAiRtcLicenseResponseBody {
	s.Code = &v
	return s
}

func (s *ActiveAiRtcLicenseResponseBody) SetHttpStatusCode(v int32) *ActiveAiRtcLicenseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ActiveAiRtcLicenseResponseBody) SetLicense(v string) *ActiveAiRtcLicenseResponseBody {
	s.License = &v
	return s
}

func (s *ActiveAiRtcLicenseResponseBody) SetMessage(v string) *ActiveAiRtcLicenseResponseBody {
	s.Message = &v
	return s
}

func (s *ActiveAiRtcLicenseResponseBody) SetRequestId(v string) *ActiveAiRtcLicenseResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActiveAiRtcLicenseResponseBody) SetSuccess(v bool) *ActiveAiRtcLicenseResponseBody {
	s.Success = &v
	return s
}

func (s *ActiveAiRtcLicenseResponseBody) Validate() error {
	return dara.Validate(s)
}
