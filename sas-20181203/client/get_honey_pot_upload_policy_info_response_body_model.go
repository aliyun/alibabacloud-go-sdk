// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneyPotUploadPolicyInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHoneyPotUploadPolicyInfoResponseBody
	GetCode() *string
	SetData(v *GetHoneyPotUploadPolicyInfoResponseBodyData) *GetHoneyPotUploadPolicyInfoResponseBody
	GetData() *GetHoneyPotUploadPolicyInfoResponseBodyData
	SetHttpStatusCode(v int32) *GetHoneyPotUploadPolicyInfoResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetHoneyPotUploadPolicyInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHoneyPotUploadPolicyInfoResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHoneyPotUploadPolicyInfoResponseBody
	GetSuccess() *bool
}

type GetHoneyPotUploadPolicyInfoResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *GetHoneyPotUploadPolicyInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code that is returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8C177095-A734-59B2-9409-7D4F26FF****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHoneyPotUploadPolicyInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHoneyPotUploadPolicyInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) GetData() *GetHoneyPotUploadPolicyInfoResponseBodyData {
	return s.Data
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) SetCode(v string) *GetHoneyPotUploadPolicyInfoResponseBody {
	s.Code = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) SetData(v *GetHoneyPotUploadPolicyInfoResponseBodyData) *GetHoneyPotUploadPolicyInfoResponseBody {
	s.Data = v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) SetHttpStatusCode(v int32) *GetHoneyPotUploadPolicyInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) SetMessage(v string) *GetHoneyPotUploadPolicyInfoResponseBody {
	s.Message = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) SetRequestId(v string) *GetHoneyPotUploadPolicyInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) SetSuccess(v bool) *GetHoneyPotUploadPolicyInfoResponseBody {
	s.Success = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetHoneyPotUploadPolicyInfoResponseBodyData struct {
	// The key ID that is required for the file upload.
	//
	// example:
	//
	// yourAccessKeyID
	Accessid *string `json:"Accessid,omitempty" xml:"Accessid,omitempty"`
	// The expiration time of the URL. The value is a timestamp. You can use the value to determine whether the URL expires. If the expiration time arrives, you can no longer use the URL to upload files.
	//
	// example:
	//
	// 1661443200000
	Expire *string `json:"Expire,omitempty" xml:"Expire,omitempty"`
	// The request URL during the upload.
	//
	// example:
	//
	// https://aegis-update-static-file.oss-cn-hangzhou.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The full path of the file in OSS. The file is uploaded by calling the OSS PostObject operation.
	//
	// example:
	//
	// HONEYPOT_FILE/1766185894104675_169********
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The limits that are imposed on the file upload. The limits include the file size.
	//
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMjAyMy0wOS0wMVQwMzoyNTozNS44MzZaIiwiY29uZGl0aW9ucyI6W1siY29udGVudC1sZW5ndGgtcmFuZ2UiLDAsMTA0ODU3NjAwXSxbInN0YXJ0cy13aXRoIiwiJGtleSIsIkhPTkVZUE9UX0ZJTEUvMTc2NjE4NTg5NDEwNDY3NV8xNjkzNTM4NDM1N*****************
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The security token.
	//
	// example:
	//
	// ***
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The signature that is calculated based on **AccessKeySecret*	- and **Policy**. When you call an Object Storage Service (OSS) API operation, OSS uses the signature information to verify the POST request.
	//
	// example:
	//
	// wKPqlFneNTZPn52k2Rz9GTY*****
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetHoneyPotUploadPolicyInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHoneyPotUploadPolicyInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) GetAccessid() *string {
	return s.Accessid
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) GetExpire() *string {
	return s.Expire
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) GetHost() *string {
	return s.Host
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) GetKey() *string {
	return s.Key
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) GetPolicy() *string {
	return s.Policy
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) GetSignature() *string {
	return s.Signature
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) SetAccessid(v string) *GetHoneyPotUploadPolicyInfoResponseBodyData {
	s.Accessid = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) SetExpire(v string) *GetHoneyPotUploadPolicyInfoResponseBodyData {
	s.Expire = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) SetHost(v string) *GetHoneyPotUploadPolicyInfoResponseBodyData {
	s.Host = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) SetKey(v string) *GetHoneyPotUploadPolicyInfoResponseBodyData {
	s.Key = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) SetPolicy(v string) *GetHoneyPotUploadPolicyInfoResponseBodyData {
	s.Policy = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) SetSecurityToken(v string) *GetHoneyPotUploadPolicyInfoResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) SetSignature(v string) *GetHoneyPotUploadPolicyInfoResponseBodyData {
	s.Signature = &v
	return s
}

func (s *GetHoneyPotUploadPolicyInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
