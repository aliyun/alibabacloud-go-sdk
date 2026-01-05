// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileUploadCallbackResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *FileUploadCallbackResponseBodyData) *FileUploadCallbackResponseBody
	GetData() *FileUploadCallbackResponseBodyData
	SetErrorCode(v string) *FileUploadCallbackResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *FileUploadCallbackResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *FileUploadCallbackResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *FileUploadCallbackResponseBody
	GetSuccess() *bool
}

type FileUploadCallbackResponseBody struct {
	Data *FileUploadCallbackResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Specified parameter Tid is not valid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 841BC14F-8E21-56B0-A7D6-593C5841AC84
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s FileUploadCallbackResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FileUploadCallbackResponseBody) GoString() string {
	return s.String()
}

func (s *FileUploadCallbackResponseBody) GetData() *FileUploadCallbackResponseBodyData {
	return s.Data
}

func (s *FileUploadCallbackResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *FileUploadCallbackResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *FileUploadCallbackResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FileUploadCallbackResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FileUploadCallbackResponseBody) SetData(v *FileUploadCallbackResponseBodyData) *FileUploadCallbackResponseBody {
	s.Data = v
	return s
}

func (s *FileUploadCallbackResponseBody) SetErrorCode(v string) *FileUploadCallbackResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FileUploadCallbackResponseBody) SetErrorMessage(v string) *FileUploadCallbackResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *FileUploadCallbackResponseBody) SetRequestId(v string) *FileUploadCallbackResponseBody {
	s.RequestId = &v
	return s
}

func (s *FileUploadCallbackResponseBody) SetSuccess(v bool) *FileUploadCallbackResponseBody {
	s.Success = &v
	return s
}

func (s *FileUploadCallbackResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FileUploadCallbackResponseBodyData struct {
	// example:
	//
	// f-8*******01m
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s FileUploadCallbackResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FileUploadCallbackResponseBodyData) GoString() string {
	return s.String()
}

func (s *FileUploadCallbackResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *FileUploadCallbackResponseBodyData) SetFileId(v string) *FileUploadCallbackResponseBodyData {
	s.FileId = &v
	return s
}

func (s *FileUploadCallbackResponseBodyData) Validate() error {
	return dara.Validate(s)
}
