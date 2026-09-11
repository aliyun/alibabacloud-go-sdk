// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateFileUploadParamsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GenerateFileUploadParamsResponseBody
	GetCode() *string
	SetData(v *GenerateFileUploadParamsResponseBodyData) *GenerateFileUploadParamsResponseBody
	GetData() *GenerateFileUploadParamsResponseBodyData
	SetHttpStatusCode(v int32) *GenerateFileUploadParamsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GenerateFileUploadParamsResponseBody
	GetMessage() *string
	SetParams(v []*string) *GenerateFileUploadParamsResponseBody
	GetParams() []*string
	SetRequestId(v string) *GenerateFileUploadParamsResponseBody
	GetRequestId() *string
}

type GenerateFileUploadParamsResponseBody struct {
	// The internal error code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The file upload parameters.
	Data *GenerateFileUploadParamsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// Instance llm-xdne77rxe14ziszr
	//
	//  does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of dynamic error parameters.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// D771A1B6-3D5F-174A-BEE1-98CE1000D337
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GenerateFileUploadParamsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GenerateFileUploadParamsResponseBody) GoString() string {
	return s.String()
}

func (s *GenerateFileUploadParamsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GenerateFileUploadParamsResponseBody) GetData() *GenerateFileUploadParamsResponseBodyData {
	return s.Data
}

func (s *GenerateFileUploadParamsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GenerateFileUploadParamsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GenerateFileUploadParamsResponseBody) GetParams() []*string {
	return s.Params
}

func (s *GenerateFileUploadParamsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GenerateFileUploadParamsResponseBody) SetCode(v string) *GenerateFileUploadParamsResponseBody {
	s.Code = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBody) SetData(v *GenerateFileUploadParamsResponseBodyData) *GenerateFileUploadParamsResponseBody {
	s.Data = v
	return s
}

func (s *GenerateFileUploadParamsResponseBody) SetHttpStatusCode(v int32) *GenerateFileUploadParamsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBody) SetMessage(v string) *GenerateFileUploadParamsResponseBody {
	s.Message = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBody) SetParams(v []*string) *GenerateFileUploadParamsResponseBody {
	s.Params = v
	return s
}

func (s *GenerateFileUploadParamsResponseBody) SetRequestId(v string) *GenerateFileUploadParamsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GenerateFileUploadParamsResponseBodyData struct {
	// The AccessKey ID used for signing.
	//
	// example:
	//
	// STS.NYGg9ejEjYqySx3EsuRutagbd
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The OSS secret used for authorized file upload.
	//
	// example:
	//
	// DGhwedF4SsbsqUMfzNBCjZFLJZSAdhiSE4hFPbKMm6JE
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The name of the OSS bucket where files are stored.
	//
	// example:
	//
	// cab
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The upload validity period.
	//
	// example:
	//
	// 1774794266093
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The upload file path.
	//
	// example:
	//
	// vocabulary/B678CA67-C8CB-150C-AD7F-6FA7F0A811BA_热词导入模版 (7).zip
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// The access domain name of OSS.
	//
	// example:
	//
	// http://cab.oss-cn-hangzhou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The policy that OSS uses to verify the validity of the request form fields.
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNi0wMy0yOVQxMzoyNDoyNi4yMDNaIiwiY29uZGl0aW9ucyI6W239
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The region.
	//
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The security token.
	//
	// example:
	//
	// CAISzwJ1q6Ft5B2yfSjIr5ryLIjRh5pL7rOSUV6CoXMgXvpYjqLJhjz2IHhMfnlvB+gYsfU2m2xR5/Yclrp6SJtIXleCZtF94oxN9h2gb4fb42Jqag+/08/LI3OaLjKm9u2wCryLYbGwU/OpbE++5U0X6LDmdDKkckW4OJmS8/BOZcgWWQ/KBlgvRq0hRG1YpdQdKGHaONu0LxfumRCwNkdzvRdmgm4NgsbWgO/ks0GG3ASmlrFF+9mufMb5M/MBZskvD42Hu8VtbbfE3SJq7BxHybx7lqQs+02c5onHUwYPu0vZYrOLroQ+fFFjHKMzDdtPq/7ylPI9ofDamIXxxAarin3kufQeLmrJ4LwneIvBXr5RHd5wa2rbWAEsmLNBEhL2EJMKtT476hcbIAuUI3bC5F+kxOHp9i6ErImtRWbLssUUla4R5TGOWbLJWzkTH93xuRqAAapuIRuRt0d2Odr1hsaYukMd42UkNapdTrehzmXeR6lyv1jlLmkAHve9Cbl9N5bO3A96FSlEfjHksQBWG0CEXRm3jLW41bpR00dgnM6gpOj7lRW2z33L0dTtaRw79X3+Uqz3gv9md5QvoaVi1jnr/cFRNxbjl7DI39pdcGlTI2lqIAA=
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The signature information calculated based on the AccessKey secret and the policy. When you call an OSS API operation, OSS verifies this signature information to confirm the validity of the POST request.
	//
	// example:
	//
	// 6oETypl+gbYHwbgcwnQiyDYoQbA=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GenerateFileUploadParamsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GenerateFileUploadParamsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GenerateFileUploadParamsResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GenerateFileUploadParamsResponseBodyData) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GenerateFileUploadParamsResponseBodyData) GetBucket() *string {
	return s.Bucket
}

func (s *GenerateFileUploadParamsResponseBodyData) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GenerateFileUploadParamsResponseBodyData) GetFileKey() *string {
	return s.FileKey
}

func (s *GenerateFileUploadParamsResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GenerateFileUploadParamsResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GenerateFileUploadParamsResponseBodyData) GetRegion() *string {
	return s.Region
}

func (s *GenerateFileUploadParamsResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GenerateFileUploadParamsResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GenerateFileUploadParamsResponseBodyData) SetAccessKeyId(v string) *GenerateFileUploadParamsResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBodyData) SetAccessKeySecret(v string) *GenerateFileUploadParamsResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBodyData) SetBucket(v string) *GenerateFileUploadParamsResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBodyData) SetExpirationTime(v int64) *GenerateFileUploadParamsResponseBodyData {
	s.ExpirationTime = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBodyData) SetFileKey(v string) *GenerateFileUploadParamsResponseBodyData {
	s.FileKey = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBodyData) SetHost(v string) *GenerateFileUploadParamsResponseBodyData {
	s.Host = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBodyData) SetPolicy(v string) *GenerateFileUploadParamsResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBodyData) SetRegion(v string) *GenerateFileUploadParamsResponseBodyData {
	s.Region = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBodyData) SetSecurityToken(v string) *GenerateFileUploadParamsResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBodyData) SetSignature(v string) *GenerateFileUploadParamsResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GenerateFileUploadParamsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
