// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSourceUploadSignatureResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSourceUploadSignatureResponseBody
	GetCode() *string
	SetContentType(v string) *GetSourceUploadSignatureResponseBody
	GetContentType() *string
	SetExpiresIn(v int64) *GetSourceUploadSignatureResponseBody
	GetExpiresIn() *int64
	SetFilePublicUrl(v string) *GetSourceUploadSignatureResponseBody
	GetFilePublicUrl() *string
	SetFileRecordId(v string) *GetSourceUploadSignatureResponseBody
	GetFileRecordId() *string
	SetFileUrl(v string) *GetSourceUploadSignatureResponseBody
	GetFileUrl() *string
	SetMessage(v string) *GetSourceUploadSignatureResponseBody
	GetMessage() *string
	SetMethod(v string) *GetSourceUploadSignatureResponseBody
	GetMethod() *string
	SetObjectName(v string) *GetSourceUploadSignatureResponseBody
	GetObjectName() *string
	SetRequestId(v string) *GetSourceUploadSignatureResponseBody
	GetRequestId() *string
	SetUploadSignatureUrl(v string) *GetSourceUploadSignatureResponseBody
	GetUploadSignatureUrl() *string
}

type GetSourceUploadSignatureResponseBody struct {
	// 业务状态码：成功为 200，失败为后端错误码（ERR.	- / InvalidParameter.*）
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// 文件 Content-Type
	//
	// example:
	//
	// string_value
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// 签名 URL 有效时长（秒）
	//
	// example:
	//
	// 1
	ExpiresIn *int64 `json:"expiresIn,omitempty" xml:"expiresIn,omitempty"`
	// 文件公开访问 URL
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// 文件记录 ID
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// 文件 OSS 内部 URL
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// 错误描述，成功时为空
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// 上传 HTTP 方法（固定为 PUT）
	//
	// example:
	//
	// string_value
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// OSS 对象名
	//
	// example:
	//
	// string_value
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	// 请求追踪 ID
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// 签名上传 URL（PUT 方式）
	//
	// example:
	//
	// https://example.com/winnexo/resource
	UploadSignatureUrl *string `json:"uploadSignatureUrl,omitempty" xml:"uploadSignatureUrl,omitempty"`
}

func (s GetSourceUploadSignatureResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSourceUploadSignatureResponseBody) GoString() string {
	return s.String()
}

func (s *GetSourceUploadSignatureResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSourceUploadSignatureResponseBody) GetContentType() *string {
	return s.ContentType
}

func (s *GetSourceUploadSignatureResponseBody) GetExpiresIn() *int64 {
	return s.ExpiresIn
}

func (s *GetSourceUploadSignatureResponseBody) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *GetSourceUploadSignatureResponseBody) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *GetSourceUploadSignatureResponseBody) GetFileUrl() *string {
	return s.FileUrl
}

func (s *GetSourceUploadSignatureResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSourceUploadSignatureResponseBody) GetMethod() *string {
	return s.Method
}

func (s *GetSourceUploadSignatureResponseBody) GetObjectName() *string {
	return s.ObjectName
}

func (s *GetSourceUploadSignatureResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSourceUploadSignatureResponseBody) GetUploadSignatureUrl() *string {
	return s.UploadSignatureUrl
}

func (s *GetSourceUploadSignatureResponseBody) SetCode(v string) *GetSourceUploadSignatureResponseBody {
	s.Code = &v
	return s
}

func (s *GetSourceUploadSignatureResponseBody) SetContentType(v string) *GetSourceUploadSignatureResponseBody {
	s.ContentType = &v
	return s
}

func (s *GetSourceUploadSignatureResponseBody) SetExpiresIn(v int64) *GetSourceUploadSignatureResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *GetSourceUploadSignatureResponseBody) SetFilePublicUrl(v string) *GetSourceUploadSignatureResponseBody {
	s.FilePublicUrl = &v
	return s
}

func (s *GetSourceUploadSignatureResponseBody) SetFileRecordId(v string) *GetSourceUploadSignatureResponseBody {
	s.FileRecordId = &v
	return s
}

func (s *GetSourceUploadSignatureResponseBody) SetFileUrl(v string) *GetSourceUploadSignatureResponseBody {
	s.FileUrl = &v
	return s
}

func (s *GetSourceUploadSignatureResponseBody) SetMessage(v string) *GetSourceUploadSignatureResponseBody {
	s.Message = &v
	return s
}

func (s *GetSourceUploadSignatureResponseBody) SetMethod(v string) *GetSourceUploadSignatureResponseBody {
	s.Method = &v
	return s
}

func (s *GetSourceUploadSignatureResponseBody) SetObjectName(v string) *GetSourceUploadSignatureResponseBody {
	s.ObjectName = &v
	return s
}

func (s *GetSourceUploadSignatureResponseBody) SetRequestId(v string) *GetSourceUploadSignatureResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSourceUploadSignatureResponseBody) SetUploadSignatureUrl(v string) *GetSourceUploadSignatureResponseBody {
	s.UploadSignatureUrl = &v
	return s
}

func (s *GetSourceUploadSignatureResponseBody) Validate() error {
	return dara.Validate(s)
}
