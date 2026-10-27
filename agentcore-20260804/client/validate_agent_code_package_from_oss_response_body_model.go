// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateAgentCodePackageFromOssResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ValidateAgentCodePackageFromOssResponseBody
	GetCode() *string
	SetData(v *ValidateAgentCodePackageFromOssResponseBodyData) *ValidateAgentCodePackageFromOssResponseBody
	GetData() *ValidateAgentCodePackageFromOssResponseBodyData
	SetHttpStatusCode(v int32) *ValidateAgentCodePackageFromOssResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ValidateAgentCodePackageFromOssResponseBody
	GetMessage() *string
	SetRequestId(v string) *ValidateAgentCodePackageFromOssResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ValidateAgentCodePackageFromOssResponseBody
	GetSuccess() *bool
}

type ValidateAgentCodePackageFromOssResponseBody struct {
	// The business status code.
	//
	// example:
	//
	// SUCCESS
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The returned result.
	Data *ValidateAgentCodePackageFromOssResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
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

func (s ValidateAgentCodePackageFromOssResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ValidateAgentCodePackageFromOssResponseBody) GoString() string {
	return s.String()
}

func (s *ValidateAgentCodePackageFromOssResponseBody) GetCode() *string {
	return s.Code
}

func (s *ValidateAgentCodePackageFromOssResponseBody) GetData() *ValidateAgentCodePackageFromOssResponseBodyData {
	return s.Data
}

func (s *ValidateAgentCodePackageFromOssResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ValidateAgentCodePackageFromOssResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ValidateAgentCodePackageFromOssResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ValidateAgentCodePackageFromOssResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ValidateAgentCodePackageFromOssResponseBody) SetCode(v string) *ValidateAgentCodePackageFromOssResponseBody {
	s.Code = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponseBody) SetData(v *ValidateAgentCodePackageFromOssResponseBodyData) *ValidateAgentCodePackageFromOssResponseBody {
	s.Data = v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponseBody) SetHttpStatusCode(v int32) *ValidateAgentCodePackageFromOssResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponseBody) SetMessage(v string) *ValidateAgentCodePackageFromOssResponseBody {
	s.Message = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponseBody) SetRequestId(v string) *ValidateAgentCodePackageFromOssResponseBody {
	s.RequestId = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponseBody) SetSuccess(v bool) *ValidateAgentCodePackageFromOssResponseBody {
	s.Success = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ValidateAgentCodePackageFromOssResponseBodyData struct {
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
	// The OSS URI of the customer source code package.
	//
	// example:
	//
	// oss://my-agent-bucket/packages/main.zip
	SourceUri *string `json:"sourceUri,omitempty" xml:"sourceUri,omitempty"`
}

func (s ValidateAgentCodePackageFromOssResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ValidateAgentCodePackageFromOssResponseBodyData) GoString() string {
	return s.String()
}

func (s *ValidateAgentCodePackageFromOssResponseBodyData) GetFileName() *string {
	return s.FileName
}

func (s *ValidateAgentCodePackageFromOssResponseBodyData) GetFileSize() *int64 {
	return s.FileSize
}

func (s *ValidateAgentCodePackageFromOssResponseBodyData) GetFormat() *string {
	return s.Format
}

func (s *ValidateAgentCodePackageFromOssResponseBodyData) GetSourceUri() *string {
	return s.SourceUri
}

func (s *ValidateAgentCodePackageFromOssResponseBodyData) SetFileName(v string) *ValidateAgentCodePackageFromOssResponseBodyData {
	s.FileName = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponseBodyData) SetFileSize(v int64) *ValidateAgentCodePackageFromOssResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponseBodyData) SetFormat(v string) *ValidateAgentCodePackageFromOssResponseBodyData {
	s.Format = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponseBodyData) SetSourceUri(v string) *ValidateAgentCodePackageFromOssResponseBodyData {
	s.SourceUri = &v
	return s
}

func (s *ValidateAgentCodePackageFromOssResponseBodyData) Validate() error {
	return dara.Validate(s)
}
