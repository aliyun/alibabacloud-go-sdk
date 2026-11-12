// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcpCodePackageUploadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMcpCodePackageUploadUrlResponseBody
	GetCode() *string
	SetData(v *GetMcpCodePackageUploadUrlResponseBodyData) *GetMcpCodePackageUploadUrlResponseBody
	GetData() *GetMcpCodePackageUploadUrlResponseBodyData
	SetHttpStatusCode(v int32) *GetMcpCodePackageUploadUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetMcpCodePackageUploadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMcpCodePackageUploadUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMcpCodePackageUploadUrlResponseBody
	GetSuccess() *bool
}

type GetMcpCodePackageUploadUrlResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetMcpCodePackageUploadUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID, which is used to locate and troubleshoot the request.
	//
	// example:
	//
	// request-123456
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetMcpCodePackageUploadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMcpCodePackageUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetMcpCodePackageUploadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMcpCodePackageUploadUrlResponseBody) GetData() *GetMcpCodePackageUploadUrlResponseBodyData {
	return s.Data
}

func (s *GetMcpCodePackageUploadUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetMcpCodePackageUploadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMcpCodePackageUploadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMcpCodePackageUploadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMcpCodePackageUploadUrlResponseBody) SetCode(v string) *GetMcpCodePackageUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBody) SetData(v *GetMcpCodePackageUploadUrlResponseBodyData) *GetMcpCodePackageUploadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBody) SetHttpStatusCode(v int32) *GetMcpCodePackageUploadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBody) SetMessage(v string) *GetMcpCodePackageUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBody) SetRequestId(v string) *GetMcpCodePackageUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBody) SetSuccess(v bool) *GetMcpCodePackageUploadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMcpCodePackageUploadUrlResponseBodyData struct {
	// The expiration time of the temporary code package artifact, in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-08-24T00:00:00Z
	ArtifactExpiresAt *string `json:"artifactExpiresAt,omitempty" xml:"artifactExpiresAt,omitempty"`
	// The code package upload token.
	//
	// example:
	//
	// upload-token
	CodePackageToken *string `json:"codePackageToken,omitempty" xml:"codePackageToken,omitempty"`
	// The expiration time of the presigned upload URL, in ISO 8601 UTC format.
	//
	// example:
	//
	// 2026-08-23T00:15:00Z
	ExpiresAt *string `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	// The HTTP request headers. The key is the header name, and the value is the header content.
	//
	// example:
	//
	// {"Content-Type":"application/zip"}
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The maximum allowed size of the code package to be uploaded, in bytes.
	//
	// example:
	//
	// 52428800
	MaxSize *int64 `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	// The HTTP method used to upload the code package.
	//
	// example:
	//
	// PUT
	UploadMethod *string `json:"uploadMethod,omitempty" xml:"uploadMethod,omitempty"`
	// The presigned upload URL of the code package.
	//
	// example:
	//
	// https://example-bucket.oss-cn-hangzhou.aliyuncs.com/server.zip?Signature=example-signature
	UploadUrl *string `json:"uploadUrl,omitempty" xml:"uploadUrl,omitempty"`
}

func (s GetMcpCodePackageUploadUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetMcpCodePackageUploadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) GetArtifactExpiresAt() *string {
	return s.ArtifactExpiresAt
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) GetCodePackageToken() *string {
	return s.CodePackageToken
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) GetMaxSize() *int64 {
	return s.MaxSize
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) GetUploadMethod() *string {
	return s.UploadMethod
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) SetArtifactExpiresAt(v string) *GetMcpCodePackageUploadUrlResponseBodyData {
	s.ArtifactExpiresAt = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) SetCodePackageToken(v string) *GetMcpCodePackageUploadUrlResponseBodyData {
	s.CodePackageToken = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) SetExpiresAt(v string) *GetMcpCodePackageUploadUrlResponseBodyData {
	s.ExpiresAt = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) SetHeaders(v map[string]*string) *GetMcpCodePackageUploadUrlResponseBodyData {
	s.Headers = v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) SetMaxSize(v int64) *GetMcpCodePackageUploadUrlResponseBodyData {
	s.MaxSize = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) SetUploadMethod(v string) *GetMcpCodePackageUploadUrlResponseBodyData {
	s.UploadMethod = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) SetUploadUrl(v string) *GetMcpCodePackageUploadUrlResponseBodyData {
	s.UploadUrl = &v
	return s
}

func (s *GetMcpCodePackageUploadUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
