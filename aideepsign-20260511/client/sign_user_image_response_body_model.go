// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSignUserImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithm(v string) *SignUserImageResponseBody
	GetAlgorithm() *string
	SetCertificateSubject(v string) *SignUserImageResponseBody
	GetCertificateSubject() *string
	SetCode(v string) *SignUserImageResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SignUserImageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SignUserImageResponseBody
	GetMessage() *string
	SetObjectKey(v string) *SignUserImageResponseBody
	GetObjectKey() *string
	SetRequestId(v string) *SignUserImageResponseBody
	GetRequestId() *string
	SetSignTime(v string) *SignUserImageResponseBody
	GetSignTime() *string
	SetSignedImageUrl(v string) *SignUserImageResponseBody
	GetSignedImageUrl() *string
	SetSuccess(v bool) *SignUserImageResponseBody
	GetSuccess() *bool
}

type SignUserImageResponseBody struct {
	// The algorithm used for signing, such as ps256 or es256.
	//
	// example:
	//
	// ps256
	Algorithm *string `json:"Algorithm,omitempty" xml:"Algorithm,omitempty"`
	// The subject information of the signing certificate.
	//
	// example:
	//
	// CN=AIDeepSign User Certificate,O=Alibaba Cloud
	CertificateSubject *string `json:"CertificateSubject,omitempty" xml:"CertificateSubject,omitempty"`
	// The business error code. The value "OK" is returned if the request succeeds.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. The value 200 is returned if the request succeeds.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The additional information. The value "success" is returned if the request succeeds.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ObjectKey of the signed image in OSS. You can use this value for subsequent API calls.
	//
	// example:
	//
	// deepsign/123456789/signed/abc12345-def6-7890-abcd-ef1234567890.png
	ObjectKey *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-7890-ABCD-EF1234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signing time in ISO 8601 format, such as `2026-01-15T08:30:00Z`.
	//
	// example:
	//
	// 2026-06-18T10:30:00Z
	SignTime *string `json:"SignTime,omitempty" xml:"SignTime,omitempty"`
	// The pre-signed download URL of the signed image.
	//
	// example:
	//
	// https://bucket.oss-cn-hangzhou.aliyuncs.com/deepsign/123456789/signed/abc12345.png?Expires=1718700000&OSSAccessKeyId=...
	SignedImageUrl *string `json:"SignedImageUrl,omitempty" xml:"SignedImageUrl,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SignUserImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SignUserImageResponseBody) GoString() string {
	return s.String()
}

func (s *SignUserImageResponseBody) GetAlgorithm() *string {
	return s.Algorithm
}

func (s *SignUserImageResponseBody) GetCertificateSubject() *string {
	return s.CertificateSubject
}

func (s *SignUserImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *SignUserImageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SignUserImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SignUserImageResponseBody) GetObjectKey() *string {
	return s.ObjectKey
}

func (s *SignUserImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SignUserImageResponseBody) GetSignTime() *string {
	return s.SignTime
}

func (s *SignUserImageResponseBody) GetSignedImageUrl() *string {
	return s.SignedImageUrl
}

func (s *SignUserImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SignUserImageResponseBody) SetAlgorithm(v string) *SignUserImageResponseBody {
	s.Algorithm = &v
	return s
}

func (s *SignUserImageResponseBody) SetCertificateSubject(v string) *SignUserImageResponseBody {
	s.CertificateSubject = &v
	return s
}

func (s *SignUserImageResponseBody) SetCode(v string) *SignUserImageResponseBody {
	s.Code = &v
	return s
}

func (s *SignUserImageResponseBody) SetHttpStatusCode(v int32) *SignUserImageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SignUserImageResponseBody) SetMessage(v string) *SignUserImageResponseBody {
	s.Message = &v
	return s
}

func (s *SignUserImageResponseBody) SetObjectKey(v string) *SignUserImageResponseBody {
	s.ObjectKey = &v
	return s
}

func (s *SignUserImageResponseBody) SetRequestId(v string) *SignUserImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *SignUserImageResponseBody) SetSignTime(v string) *SignUserImageResponseBody {
	s.SignTime = &v
	return s
}

func (s *SignUserImageResponseBody) SetSignedImageUrl(v string) *SignUserImageResponseBody {
	s.SignedImageUrl = &v
	return s
}

func (s *SignUserImageResponseBody) SetSuccess(v bool) *SignUserImageResponseBody {
	s.Success = &v
	return s
}

func (s *SignUserImageResponseBody) Validate() error {
	return dara.Validate(s)
}
