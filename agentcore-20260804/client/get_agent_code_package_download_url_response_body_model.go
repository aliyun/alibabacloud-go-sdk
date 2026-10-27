// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentCodePackageDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentCodePackageDownloadUrlResponseBody
	GetCode() *string
	SetData(v *GetAgentCodePackageDownloadUrlResponseBodyData) *GetAgentCodePackageDownloadUrlResponseBody
	GetData() *GetAgentCodePackageDownloadUrlResponseBodyData
	SetHttpStatusCode(v int32) *GetAgentCodePackageDownloadUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetAgentCodePackageDownloadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentCodePackageDownloadUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentCodePackageDownloadUrlResponseBody
	GetSuccess() *bool
}

type GetAgentCodePackageDownloadUrlResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data *GetAgentCodePackageDownloadUrlResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s GetAgentCodePackageDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) GetData() *GetAgentCodePackageDownloadUrlResponseBodyData {
	return s.Data
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) SetCode(v string) *GetAgentCodePackageDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) SetData(v *GetAgentCodePackageDownloadUrlResponseBodyData) *GetAgentCodePackageDownloadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) SetHttpStatusCode(v int32) *GetAgentCodePackageDownloadUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) SetMessage(v string) *GetAgentCodePackageDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) SetRequestId(v string) *GetAgentCodePackageDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) SetSuccess(v bool) *GetAgentCodePackageDownloadUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentCodePackageDownloadUrlResponseBodyData struct {
	// The CRC64 value of the code package object.
	//
	// example:
	//
	// 456
	Crc64 *string `json:"crc64,omitempty" xml:"crc64,omitempty"`
	// The time-limited, read-only download URL for the code package.
	//
	// example:
	//
	// https://agentcore-platform-example.oss-cn-hangzhou.aliyuncs.com/code-packages/workspaces/ws-a1b2c3d4e5f6/550e8400-e29b-41d4-a716-446655440000/package.zip?OSSAccessKeyId=LTAI5tExample****&Expires=1792627800&Signature=ExampleSignature****
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// The ETag of the code package object.
	//
	// example:
	//
	// fixed-etag
	Etag *string `json:"etag,omitempty" xml:"etag,omitempty"`
	// The expiration time of the download URL.
	//
	// example:
	//
	// 2026-10-22T00:10:00Z
	ExpiresAt *string `json:"expiresAt,omitempty" xml:"expiresAt,omitempty"`
	// The original file name of the code package.
	//
	// example:
	//
	// main.zip
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The compressed size of the code package.
	//
	// example:
	//
	// 123
	FileSize *int64 `json:"fileSize,omitempty" xml:"fileSize,omitempty"`
	// The format of the code package.
	//
	// example:
	//
	// zip
	Format *string `json:"format,omitempty" xml:"format,omitempty"`
	// The resource version ID corresponding to the currently saved configuration.
	//
	// example:
	//
	// 42
	ResourceVersionId *string `json:"resourceVersionId,omitempty" xml:"resourceVersionId,omitempty"`
}

func (s GetAgentCodePackageDownloadUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentCodePackageDownloadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) GetCrc64() *string {
	return s.Crc64
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) GetEtag() *string {
	return s.Etag
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) GetExpiresAt() *string {
	return s.ExpiresAt
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) GetFileSize() *int64 {
	return s.FileSize
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) GetFormat() *string {
	return s.Format
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) GetResourceVersionId() *string {
	return s.ResourceVersionId
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) SetCrc64(v string) *GetAgentCodePackageDownloadUrlResponseBodyData {
	s.Crc64 = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) SetDownloadUrl(v string) *GetAgentCodePackageDownloadUrlResponseBodyData {
	s.DownloadUrl = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) SetEtag(v string) *GetAgentCodePackageDownloadUrlResponseBodyData {
	s.Etag = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) SetExpiresAt(v string) *GetAgentCodePackageDownloadUrlResponseBodyData {
	s.ExpiresAt = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) SetFileName(v string) *GetAgentCodePackageDownloadUrlResponseBodyData {
	s.FileName = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) SetFileSize(v int64) *GetAgentCodePackageDownloadUrlResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) SetFormat(v string) *GetAgentCodePackageDownloadUrlResponseBodyData {
	s.Format = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) SetResourceVersionId(v string) *GetAgentCodePackageDownloadUrlResponseBodyData {
	s.ResourceVersionId = &v
	return s
}

func (s *GetAgentCodePackageDownloadUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
