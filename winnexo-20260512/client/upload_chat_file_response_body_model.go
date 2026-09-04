// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadChatFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadChatFileResponseBody
	GetCode() *string
	SetContentType(v string) *UploadChatFileResponseBody
	GetContentType() *string
	SetFileName(v string) *UploadChatFileResponseBody
	GetFileName() *string
	SetFilePublicUrl(v string) *UploadChatFileResponseBody
	GetFilePublicUrl() *string
	SetFileRecordId(v string) *UploadChatFileResponseBody
	GetFileRecordId() *string
	SetFileUrl(v string) *UploadChatFileResponseBody
	GetFileUrl() *string
	SetMessage(v string) *UploadChatFileResponseBody
	GetMessage() *string
	SetObjectName(v string) *UploadChatFileResponseBody
	GetObjectName() *string
	SetRequestId(v string) *UploadChatFileResponseBody
	GetRequestId() *string
	SetUploadSignatureUrl(v string) *UploadChatFileResponseBody
	GetUploadSignatureUrl() *string
}

type UploadChatFileResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The content type of the file. Valid values:
	//
	// - **image**: Image.
	//
	// - **document**: General document.
	//
	// - **alidoc**: Alibaba document.
	//
	// - **text**: Text.
	//
	// - **video**: Video.
	//
	// - **audio**: Audio.
	//
	// - **archive**: Archive.
	//
	// - **app**: Application.
	//
	// - **link**: Shortcut.
	//
	// - **other**: Other.
	//
	// example:
	//
	// string_value
	ContentType *string `json:"contentType,omitempty" xml:"contentType,omitempty"`
	// The full path name of the file.
	//
	// example:
	//
	// example.pdf
	FileName *string `json:"fileName,omitempty" xml:"fileName,omitempty"`
	// The publicly accessible URL of the AliDing online document.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FilePublicUrl *string `json:"filePublicUrl,omitempty" xml:"filePublicUrl,omitempty"`
	// The file record ID. This parameter is optional and corresponds to settings.file_record_id.
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
	// The Yida attachment address.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	FileUrl *string `json:"fileUrl,omitempty" xml:"fileUrl,omitempty"`
	// The description of the status code.
	//
	// example:
	//
	// ok
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The object name.
	//
	// example:
	//
	// string_value
	ObjectName *string `json:"objectName,omitempty" xml:"objectName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The signature URL.
	//
	// example:
	//
	// https://example.com/winnexo/resource
	UploadSignatureUrl *string `json:"uploadSignatureUrl,omitempty" xml:"uploadSignatureUrl,omitempty"`
}

func (s UploadChatFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadChatFileResponseBody) GoString() string {
	return s.String()
}

func (s *UploadChatFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadChatFileResponseBody) GetContentType() *string {
	return s.ContentType
}

func (s *UploadChatFileResponseBody) GetFileName() *string {
	return s.FileName
}

func (s *UploadChatFileResponseBody) GetFilePublicUrl() *string {
	return s.FilePublicUrl
}

func (s *UploadChatFileResponseBody) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *UploadChatFileResponseBody) GetFileUrl() *string {
	return s.FileUrl
}

func (s *UploadChatFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadChatFileResponseBody) GetObjectName() *string {
	return s.ObjectName
}

func (s *UploadChatFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadChatFileResponseBody) GetUploadSignatureUrl() *string {
	return s.UploadSignatureUrl
}

func (s *UploadChatFileResponseBody) SetCode(v string) *UploadChatFileResponseBody {
	s.Code = &v
	return s
}

func (s *UploadChatFileResponseBody) SetContentType(v string) *UploadChatFileResponseBody {
	s.ContentType = &v
	return s
}

func (s *UploadChatFileResponseBody) SetFileName(v string) *UploadChatFileResponseBody {
	s.FileName = &v
	return s
}

func (s *UploadChatFileResponseBody) SetFilePublicUrl(v string) *UploadChatFileResponseBody {
	s.FilePublicUrl = &v
	return s
}

func (s *UploadChatFileResponseBody) SetFileRecordId(v string) *UploadChatFileResponseBody {
	s.FileRecordId = &v
	return s
}

func (s *UploadChatFileResponseBody) SetFileUrl(v string) *UploadChatFileResponseBody {
	s.FileUrl = &v
	return s
}

func (s *UploadChatFileResponseBody) SetMessage(v string) *UploadChatFileResponseBody {
	s.Message = &v
	return s
}

func (s *UploadChatFileResponseBody) SetObjectName(v string) *UploadChatFileResponseBody {
	s.ObjectName = &v
	return s
}

func (s *UploadChatFileResponseBody) SetRequestId(v string) *UploadChatFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadChatFileResponseBody) SetUploadSignatureUrl(v string) *UploadChatFileResponseBody {
	s.UploadSignatureUrl = &v
	return s
}

func (s *UploadChatFileResponseBody) Validate() error {
	return dara.Validate(s)
}
