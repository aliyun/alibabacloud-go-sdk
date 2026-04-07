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
	// example:
	//
	// OK
	Code *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GenerateFileUploadParamsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// Instance af81a389-91f0-4157-8d82-720edd02b66a
	//
	//  does not exist.
	Message *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	Params  []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// example:
	//
	// F132DDBA-66C4-5BD3-BF7E-9642FE877158
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
	// example:
	//
	// STS.NYGg9ejEjYqySx3EsuRutagbd
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// example:
	//
	// DGhwedF4SsbsqUMfzNBCjZFLJZSAdhiSE4hFPbKMm6JE
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// example:
	//
	// cab
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// 1774794266093
	ExpirationTime *int64  `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	FileKey        *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// example:
	//
	// http://cab.oss-cn-hangzhou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNi0wMy0yOVQxMzoyNDoyNi4yMDNaIiwiY29uZGl0aW9ucyI6W239
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// cn-hangzhou
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// CAISzwJ1q6Ft5B2yfSjIr5ryLIjRh5pL7rOSUV6CoXMgXvpYjqLJhjz2IHhMfnlvB+gYsfU2m2xR5/Yclrp6SJtIXleCZtF94oxN9h2gb4fb42Jqag+/08/LI3OaLjKm9u2wCryLYbGwU/OpbE++5U0X6LDmdDKkckW4OJmS8/BOZcgWWQ/KBlgvRq0hRG1YpdQdKGHaONu0LxfumRCwNkdzvRdmgm4NgsbWgO/ks0GG3ASmlrFF+9mufMb5M/MBZskvD42Hu8VtbbfE3SJq7BxHybx7lqQs+02c5onHUwYPu0vZYrOLroQ+fFFjHKMzDdtPq/7ylPI9ofDamIXxxAarin3kufQeLmrJ4LwneIvBXr5RHd5wa2rbWAEsmLNBEhL2EJMKtT476hcbIAuUI3bC5F+kxOHp9i6ErImtRWbLssUUla4R5TGOWbLJWzkTH93xuRqAAapuIRuRt0d2Odr1hsaYukMd42UkNapdTrehzmXeR6lyv1jlLmkAHve9Cbl9N5bO3A96FSlEfjHksQBWG0CEXRm3jLW41bpR00dgnM6gpOj7lRW2z33L0dTtaRw79X3+Uqz3gv9md5QvoaVi1jnr/cFRNxbjl7DI39pdcGlTI2lqIAA=
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
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
