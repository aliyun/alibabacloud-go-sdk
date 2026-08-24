// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgeBasePreSignedUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetKnowledgeBasePreSignedUrlResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v int32) *GetKnowledgeBasePreSignedUrlResponseBody
	GetCode() *int32
	SetData(v *GetKnowledgeBasePreSignedUrlResponseBodyData) *GetKnowledgeBasePreSignedUrlResponseBody
	GetData() *GetKnowledgeBasePreSignedUrlResponseBodyData
	SetHttpStatusCode(v int32) *GetKnowledgeBasePreSignedUrlResponseBody
	GetHttpStatusCode() *int32
	SetSuccess(v bool) *GetKnowledgeBasePreSignedUrlResponseBody
	GetSuccess() *bool
}

type GetKnowledgeBasePreSignedUrlResponseBody struct {
	// The details of the permission verification failure.
	//
	// example:
	//
	// {"PolicyType":"AccountLevelIdentityBasedPolicy","AuthPrincipalOwnerId":"1234567890123456","AuthPrincipalType":"SubUser","AuthPrincipalDisplayName":"1234567890123456","NoPermissionType":"ImplicitDeny","AuthAction":"milvusknowledgebase:ListDatasets"}
	AccessDeniedDetail *string `json:"accessDeniedDetail,omitempty" xml:"accessDeniedDetail,omitempty"`
	// The business status code.
	//
	// example:
	//
	// 0
	Code *int32 `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetKnowledgeBasePreSignedUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 403
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetKnowledgeBasePreSignedUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBasePreSignedUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBasePreSignedUrlResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetKnowledgeBasePreSignedUrlResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetKnowledgeBasePreSignedUrlResponseBody) GetData() *GetKnowledgeBasePreSignedUrlResponseBodyData {
	return s.Data
}

func (s *GetKnowledgeBasePreSignedUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetKnowledgeBasePreSignedUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKnowledgeBasePreSignedUrlResponseBody) SetAccessDeniedDetail(v string) *GetKnowledgeBasePreSignedUrlResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlResponseBody) SetCode(v int32) *GetKnowledgeBasePreSignedUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlResponseBody) SetData(v *GetKnowledgeBasePreSignedUrlResponseBodyData) *GetKnowledgeBasePreSignedUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlResponseBody) SetHttpStatusCode(v int32) *GetKnowledgeBasePreSignedUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlResponseBody) SetSuccess(v bool) *GetKnowledgeBasePreSignedUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKnowledgeBasePreSignedUrlResponseBodyData struct {
	// The bucket name.
	//
	// example:
	//
	// knowledgebase-fileupload
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// The validity period of the pre-signed URL in seconds.
	//
	// example:
	//
	// 3600
	ExpiresIn *int32 `json:"expiresIn,omitempty" xml:"expiresIn,omitempty"`
	// The list of pre-signed PUT URLs. **The order corresponds one-to-one with the `Documents` in the request.**
	PreSignedUrls []*string `json:"preSignedUrls,omitempty" xml:"preSignedUrls,omitempty" type:"Repeated"`
}

func (s GetKnowledgeBasePreSignedUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgeBasePreSignedUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKnowledgeBasePreSignedUrlResponseBodyData) GetBucketName() *string {
	return s.BucketName
}

func (s *GetKnowledgeBasePreSignedUrlResponseBodyData) GetExpiresIn() *int32 {
	return s.ExpiresIn
}

func (s *GetKnowledgeBasePreSignedUrlResponseBodyData) GetPreSignedUrls() []*string {
	return s.PreSignedUrls
}

func (s *GetKnowledgeBasePreSignedUrlResponseBodyData) SetBucketName(v string) *GetKnowledgeBasePreSignedUrlResponseBodyData {
	s.BucketName = &v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlResponseBodyData) SetExpiresIn(v int32) *GetKnowledgeBasePreSignedUrlResponseBodyData {
	s.ExpiresIn = &v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlResponseBodyData) SetPreSignedUrls(v []*string) *GetKnowledgeBasePreSignedUrlResponseBodyData {
	s.PreSignedUrls = v
	return s
}

func (s *GetKnowledgeBasePreSignedUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
