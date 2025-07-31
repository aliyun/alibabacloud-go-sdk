// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFileStorageCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetFileStorageCredentialResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetFileStorageCredentialResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetFileStorageCredentialResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFileStorageCredentialResponseBody
	GetRequestId() *string
	SetStorageCredential(v *GetFileStorageCredentialResponseBodyStorageCredential) *GetFileStorageCredentialResponseBody
	GetStorageCredential() *GetFileStorageCredentialResponseBodyStorageCredential
	SetSuccess(v bool) *GetFileStorageCredentialResponseBody
	GetSuccess() *bool
}

type GetFileStorageCredentialResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// internal error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 82E78D6B-AA8F-1FEF-8AA3-5C9DA2A79140
	RequestId         *string                                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageCredential *GetFileStorageCredentialResponseBodyStorageCredential `json:"StorageCredential,omitempty" xml:"StorageCredential,omitempty" type:"Struct"`
	Success           *bool                                                  `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFileStorageCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFileStorageCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *GetFileStorageCredentialResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFileStorageCredentialResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetFileStorageCredentialResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFileStorageCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFileStorageCredentialResponseBody) GetStorageCredential() *GetFileStorageCredentialResponseBodyStorageCredential {
	return s.StorageCredential
}

func (s *GetFileStorageCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFileStorageCredentialResponseBody) SetCode(v string) *GetFileStorageCredentialResponseBody {
	s.Code = &v
	return s
}

func (s *GetFileStorageCredentialResponseBody) SetHttpStatusCode(v int32) *GetFileStorageCredentialResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetFileStorageCredentialResponseBody) SetMessage(v string) *GetFileStorageCredentialResponseBody {
	s.Message = &v
	return s
}

func (s *GetFileStorageCredentialResponseBody) SetRequestId(v string) *GetFileStorageCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFileStorageCredentialResponseBody) SetStorageCredential(v *GetFileStorageCredentialResponseBodyStorageCredential) *GetFileStorageCredentialResponseBody {
	s.StorageCredential = v
	return s
}

func (s *GetFileStorageCredentialResponseBody) SetSuccess(v bool) *GetFileStorageCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *GetFileStorageCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetFileStorageCredentialResponseBodyStorageCredential struct {
	// example:
	//
	// temp.akId
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// temp.akKey
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// example:
	//
	// dataphin
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// example:
	//
	// oss-cn-shanghai.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// example:
	//
	// 17343434343434
	Expiration *int64 `json:"Expiration,omitempty" xml:"Expiration,omitempty"`
	// example:
	//
	// 1023231111/abc
	ObjectName *string `json:"ObjectName,omitempty" xml:"ObjectName,omitempty"`
	// region
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// temp.token
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// example:
	//
	// oss
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s GetFileStorageCredentialResponseBodyStorageCredential) String() string {
	return dara.Prettify(s)
}

func (s GetFileStorageCredentialResponseBodyStorageCredential) GoString() string {
	return s.String()
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) GetAccessId() *string {
	return s.AccessId
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) GetAccessKey() *string {
	return s.AccessKey
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) GetBucket() *string {
	return s.Bucket
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) GetExpiration() *int64 {
	return s.Expiration
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) GetObjectName() *string {
	return s.ObjectName
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) GetRegion() *string {
	return s.Region
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) GetStorageType() *string {
	return s.StorageType
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) SetAccessId(v string) *GetFileStorageCredentialResponseBodyStorageCredential {
	s.AccessId = &v
	return s
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) SetAccessKey(v string) *GetFileStorageCredentialResponseBodyStorageCredential {
	s.AccessKey = &v
	return s
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) SetBucket(v string) *GetFileStorageCredentialResponseBodyStorageCredential {
	s.Bucket = &v
	return s
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) SetEndpoint(v string) *GetFileStorageCredentialResponseBodyStorageCredential {
	s.Endpoint = &v
	return s
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) SetExpiration(v int64) *GetFileStorageCredentialResponseBodyStorageCredential {
	s.Expiration = &v
	return s
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) SetObjectName(v string) *GetFileStorageCredentialResponseBodyStorageCredential {
	s.ObjectName = &v
	return s
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) SetRegion(v string) *GetFileStorageCredentialResponseBodyStorageCredential {
	s.Region = &v
	return s
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) SetSecurityToken(v string) *GetFileStorageCredentialResponseBodyStorageCredential {
	s.SecurityToken = &v
	return s
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) SetStorageType(v string) *GetFileStorageCredentialResponseBodyStorageCredential {
	s.StorageType = &v
	return s
}

func (s *GetFileStorageCredentialResponseBodyStorageCredential) Validate() error {
	return dara.Validate(s)
}
