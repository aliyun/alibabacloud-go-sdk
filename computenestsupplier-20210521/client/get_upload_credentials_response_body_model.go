// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadCredentialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUploadCredentialsResponseBody
	GetCode() *string
	SetData(v *GetUploadCredentialsResponseBodyData) *GetUploadCredentialsResponseBody
	GetData() *GetUploadCredentialsResponseBodyData
	SetHttpStatusCode(v int32) *GetUploadCredentialsResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUploadCredentialsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetUploadCredentialsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUploadCredentialsResponseBody
	GetSuccess() *bool
}

type GetUploadCredentialsResponseBody struct {
	// The response code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	Data *GetUploadCredentialsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// FCC3321E-D518-1BC4-861E-588E9D4DAFB7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. A value of true indicates the request was successful. A value of false indicates the request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetUploadCredentialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadCredentialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadCredentialsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUploadCredentialsResponseBody) GetData() *GetUploadCredentialsResponseBodyData {
	return s.Data
}

func (s *GetUploadCredentialsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUploadCredentialsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUploadCredentialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUploadCredentialsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUploadCredentialsResponseBody) SetCode(v string) *GetUploadCredentialsResponseBody {
	s.Code = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetData(v *GetUploadCredentialsResponseBodyData) *GetUploadCredentialsResponseBody {
	s.Data = v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetHttpStatusCode(v int32) *GetUploadCredentialsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetMessage(v string) *GetUploadCredentialsResponseBody {
	s.Message = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetRequestId(v string) *GetUploadCredentialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) SetSuccess(v bool) *GetUploadCredentialsResponseBody {
	s.Success = &v
	return s
}

func (s *GetUploadCredentialsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetUploadCredentialsResponseBodyData struct {
	// The AccessKey ID.
	//
	// example:
	//
	// STS.NUCe19W1FKaHAYAhe********
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// 8LQGp59mY23pcXeTdcvSA1cUQZBeD92sFrXi********
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The bucket name.
	//
	// example:
	//
	// service-info-private
	BucketName *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	// The time when the AccessKey pair expires.
	//
	// example:
	//
	// 2023-05-18T12:27:59Z
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The name of the key.
	//
	// example:
	//
	// 221514575922756034/cn-hangzhou/d57c62fbd508xxxxxxxx.json
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The security token.
	//
	// example:
	//
	// CAISzQN1q6Ft5B2yfSjIr5b2LouNuu5n/KOjQ3/wjGUHYdlagYGdmzz2IH1Le3NrBO8esfgymGFU6v8dlo1dYLQeHhadQI5cs80HtFqLSNaE65LswPlZ2M2ISETPJzfV9pCK
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
}

func (s GetUploadCredentialsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetUploadCredentialsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetUploadCredentialsResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetUploadCredentialsResponseBodyData) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GetUploadCredentialsResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *GetUploadCredentialsResponseBodyData) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *GetUploadCredentialsResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *GetUploadCredentialsResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetUploadCredentialsResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetUploadCredentialsResponseBodyData) SetAccessKeyId(v string) *GetUploadCredentialsResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetAccessKeySecret(v string) *GetUploadCredentialsResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetBucketName(v string) *GetUploadCredentialsResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetExpireDate(v string) *GetUploadCredentialsResponseBodyData {
	s.ExpireDate = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetKey(v string) *GetUploadCredentialsResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetRegionId(v string) *GetUploadCredentialsResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) SetSecurityToken(v string) *GetUploadCredentialsResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetUploadCredentialsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
