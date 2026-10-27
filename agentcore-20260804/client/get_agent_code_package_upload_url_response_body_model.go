// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentCodePackageUploadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentCodePackageUploadUrlResponseBody
	GetCode() *string
	SetData(v *GetAgentCodePackageUploadUrlResponseBodyData) *GetAgentCodePackageUploadUrlResponseBody
	GetData() *GetAgentCodePackageUploadUrlResponseBodyData
	SetHttpStatusCode(v int32) *GetAgentCodePackageUploadUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAgentCodePackageUploadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentCodePackageUploadUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentCodePackageUploadUrlResponseBody
	GetSuccess() *bool
}

type GetAgentCodePackageUploadUrlResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetAgentCodePackageUploadUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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
	// The request ID.
	//
	// example:
	//
	// 550e8400-e29b-41d4-a716-446655440000
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s GetAgentCodePackageUploadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageUploadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentCodePackageUploadUrlResponseBody) GetData() *GetAgentCodePackageUploadUrlResponseBodyData {
	return s.Data
}

func (s *GetAgentCodePackageUploadUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAgentCodePackageUploadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentCodePackageUploadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentCodePackageUploadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentCodePackageUploadUrlResponseBody) SetCode(v string) *GetAgentCodePackageUploadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponseBody) SetData(v *GetAgentCodePackageUploadUrlResponseBodyData) *GetAgentCodePackageUploadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponseBody) SetHttpStatusCode(v int32) *GetAgentCodePackageUploadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponseBody) SetMessage(v string) *GetAgentCodePackageUploadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponseBody) SetRequestId(v string) *GetAgentCodePackageUploadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponseBody) SetSuccess(v bool) *GetAgentCodePackageUploadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentCodePackageUploadUrlResponseBodyData struct {
	// The expiration time of the upload URL.
	//
	// example:
	//
	// 2026-10-22T00:10:00Z
	ExpiresAt *string `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	// The HTTP headers that must be included in the direct upload request.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The maximum compressed size of the code package.
	//
	// example:
	//
	// 500000000
	MaxSize *int64 `json:"maxSize,omitempty" xml:"maxSize,omitempty"`
	// The temporary Object Storage Service (OSS) URI of the code package.
	//
	// example:
	//
	// oss://agentcore-platform-example/code-package-uploads/workspaces/ws-a1b2c3d4e5f6/550e8400-e29b-41d4-a716-446655440000/package.zip
	SourceUri *string `json:"sourceUri,omitempty" xml:"sourceUri,omitempty"`
	// The HTTP method for direct upload.
	//
	// example:
	//
	// PUT
	UploadMethod *string `json:"uploadMethod,omitempty" xml:"uploadMethod,omitempty"`
	// The time-limited presigned upload URL.
	//
	// example:
	//
	// https://agentcore-platform-example.oss-cn-hangzhou.aliyuncs.com/code-package-uploads/workspaces/ws-a1b2c3d4e5f6/550e8400-e29b-41d4-a716-446655440000/package.zip?OSSAccessKeyId=LTAI5tExample****&Expires=1792627800&Signature=ExampleSignature****
	UploadUrl *string `json:"uploadUrl,omitempty" xml:"uploadUrl,omitempty"`
}

func (s GetAgentCodePackageUploadUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageUploadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) GetMaxSize() *int64 {
	return s.MaxSize
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) GetSourceUri() *string {
	return s.SourceUri
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) GetUploadMethod() *string {
	return s.UploadMethod
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) GetUploadUrl() *string {
	return s.UploadUrl
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) SetExpiresAt(v string) *GetAgentCodePackageUploadUrlResponseBodyData {
	s.ExpiresAt = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) SetHeaders(v map[string]*string) *GetAgentCodePackageUploadUrlResponseBodyData {
	s.Headers = v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) SetMaxSize(v int64) *GetAgentCodePackageUploadUrlResponseBodyData {
	s.MaxSize = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) SetSourceUri(v string) *GetAgentCodePackageUploadUrlResponseBodyData {
	s.SourceUri = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) SetUploadMethod(v string) *GetAgentCodePackageUploadUrlResponseBodyData {
	s.UploadMethod = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) SetUploadUrl(v string) *GetAgentCodePackageUploadUrlResponseBodyData {
	s.UploadUrl = &v
	return s
}

func (s *GetAgentCodePackageUploadUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
