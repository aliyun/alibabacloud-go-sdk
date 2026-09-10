// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaOssTempKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetMetaOssTempKeyResponseBodyData) *GetMetaOssTempKeyResponseBody
	GetData() *GetMetaOssTempKeyResponseBodyData
	SetErrCode(v string) *GetMetaOssTempKeyResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *GetMetaOssTempKeyResponseBody
	GetErrMessage() *string
	SetSuccess(v bool) *GetMetaOssTempKeyResponseBody
	GetSuccess() *bool
}

type GetMetaOssTempKeyResponseBody struct {
	// The response body. For more information about the fields, see the child field descriptions.
	Data *GetMetaOssTempKeyResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code. An empty string is returned if the call is successful.
	//
	// example:
	//
	// Success
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message. An empty string is returned if the call is successful.
	//
	// example:
	//
	// success
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// - true: The call is successful.
	//
	// - false: The call failed. Check errCode and errMessage for troubleshooting.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMetaOssTempKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetaOssTempKeyResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetaOssTempKeyResponseBody) GetData() *GetMetaOssTempKeyResponseBodyData {
	return s.Data
}

func (s *GetMetaOssTempKeyResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *GetMetaOssTempKeyResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *GetMetaOssTempKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetaOssTempKeyResponseBody) SetData(v *GetMetaOssTempKeyResponseBodyData) *GetMetaOssTempKeyResponseBody {
	s.Data = v
	return s
}

func (s *GetMetaOssTempKeyResponseBody) SetErrCode(v string) *GetMetaOssTempKeyResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetMetaOssTempKeyResponseBody) SetErrMessage(v string) *GetMetaOssTempKeyResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *GetMetaOssTempKeyResponseBody) SetSuccess(v bool) *GetMetaOssTempKeyResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetaOssTempKeyResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMetaOssTempKeyResponseBodyData struct {
	// The temporary AccessKey ID (STS token). This value is used together with securityToken to authenticate direct uploads to OSS. This is a sensitive credential. Do not hard-code it in your code or print it to logs.
	//
	// example:
	//
	// STS.NY6bbCNqNPpt5GcSTEzB6Lahn
	Ak *string `json:"ak,omitempty" xml:"ak,omitempty"`
	// The name of the OSS bucket.
	//
	// example:
	//
	// lhm-pre-cn-hangzhou
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// The allowed OSS upload directory prefix. The value must end with a forward slash (/). The key of the uploaded object must start with this prefix. Otherwise, the request is rejected by OSS.
	//
	// example:
	//
	// teleport/meta/1063934947625635/
	Dir *string `json:"dir,omitempty" xml:"dir,omitempty"`
	// The endpoint of the region where the OSS bucket resides.
	//
	// example:
	//
	// oss-cn-hangzhou.aliyuncs.com
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// The credential expiration timestamp in Unix seconds. Before use, verify whether the current time has exceeded this value. If the credential has expired, obtain new credentials.
	//
	// example:
	//
	// 1779966540
	Expire *int64 `json:"expire,omitempty" xml:"expire,omitempty"`
	// The Base64-encoded upload policy that defines constraints such as file size and path prefix. The decoded value is a JSON string.
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyNi0wNS0yOFQxMToxMjo1OC43MzZaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwXSxbInN0YXJ0cy13aXRoIiwiJGtleSIsInRlbGVwb3J0L21ldGEvMTA2MzkzNDk0NzYyNTYzNS8iXV19
	Policy *string `json:"policy,omitempty" xml:"policy,omitempty"`
	// The STS temporary security token. This value is used together with ak for authentication and is returned only in STS authentication mode. This is a sensitive credential. Do not hard-code it in your code or print it to logs.
	//
	// example:
	//
	// CAIS3QJ1q6Ft5B2yfSjIr5rsAOjugKcY9YqlSRPBlWEFZN1V3fD6gzz2IHhMfHFvA
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// The signature calculated based on the policy. The OSS server uses this signature to verify the validity of upload requests.
	//
	// example:
	//
	// ydDYrWUzfKNM6slVhjPhUx83qUo=
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s GetMetaOssTempKeyResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMetaOssTempKeyResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMetaOssTempKeyResponseBodyData) GetAk() *string {
	return s.Ak
}

func (s *GetMetaOssTempKeyResponseBodyData) GetBucket() *string {
	return s.Bucket
}

func (s *GetMetaOssTempKeyResponseBodyData) GetDir() *string {
	return s.Dir
}

func (s *GetMetaOssTempKeyResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetMetaOssTempKeyResponseBodyData) GetExpire() *int64 {
	return s.Expire
}

func (s *GetMetaOssTempKeyResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetMetaOssTempKeyResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetMetaOssTempKeyResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetMetaOssTempKeyResponseBodyData) SetAk(v string) *GetMetaOssTempKeyResponseBodyData {
	s.Ak = &v
	return s
}

func (s *GetMetaOssTempKeyResponseBodyData) SetBucket(v string) *GetMetaOssTempKeyResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GetMetaOssTempKeyResponseBodyData) SetDir(v string) *GetMetaOssTempKeyResponseBodyData {
	s.Dir = &v
	return s
}

func (s *GetMetaOssTempKeyResponseBodyData) SetEndpoint(v string) *GetMetaOssTempKeyResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetMetaOssTempKeyResponseBodyData) SetExpire(v int64) *GetMetaOssTempKeyResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetMetaOssTempKeyResponseBodyData) SetPolicy(v string) *GetMetaOssTempKeyResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetMetaOssTempKeyResponseBodyData) SetSecurityToken(v string) *GetMetaOssTempKeyResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetMetaOssTempKeyResponseBodyData) SetSignature(v string) *GetMetaOssTempKeyResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetMetaOssTempKeyResponseBodyData) Validate() error {
	return dara.Validate(s)
}
