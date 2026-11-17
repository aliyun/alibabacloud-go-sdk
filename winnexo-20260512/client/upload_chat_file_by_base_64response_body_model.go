// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadChatFileByBase64ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UploadChatFileByBase64ResponseBody
	GetCode() *string
	SetFileRecordId(v string) *UploadChatFileByBase64ResponseBody
	GetFileRecordId() *string
	SetMessage(v string) *UploadChatFileByBase64ResponseBody
	GetMessage() *string
	SetRequestId(v string) *UploadChatFileByBase64ResponseBody
	GetRequestId() *string
}

type UploadChatFileByBase64ResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The file record ID.
	//
	// example:
	//
	// exampleFileRecordId
	FileRecordId *string `json:"fileRecordId,omitempty" xml:"fileRecordId,omitempty"`
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
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UploadChatFileByBase64ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UploadChatFileByBase64ResponseBody) GoString() string {
	return s.String()
}

func (s *UploadChatFileByBase64ResponseBody) GetCode() *string {
	return s.Code
}

func (s *UploadChatFileByBase64ResponseBody) GetFileRecordId() *string {
	return s.FileRecordId
}

func (s *UploadChatFileByBase64ResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UploadChatFileByBase64ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UploadChatFileByBase64ResponseBody) SetCode(v string) *UploadChatFileByBase64ResponseBody {
	s.Code = &v
	return s
}

func (s *UploadChatFileByBase64ResponseBody) SetFileRecordId(v string) *UploadChatFileByBase64ResponseBody {
	s.FileRecordId = &v
	return s
}

func (s *UploadChatFileByBase64ResponseBody) SetMessage(v string) *UploadChatFileByBase64ResponseBody {
	s.Message = &v
	return s
}

func (s *UploadChatFileByBase64ResponseBody) SetRequestId(v string) *UploadChatFileByBase64ResponseBody {
	s.RequestId = &v
	return s
}

func (s *UploadChatFileByBase64ResponseBody) Validate() error {
	return dara.Validate(s)
}
