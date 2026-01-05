// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFileUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DeleteFileUploadResponseBodyData) *DeleteFileUploadResponseBody
	GetData() *DeleteFileUploadResponseBodyData
	SetErrorCode(v string) *DeleteFileUploadResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteFileUploadResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *DeleteFileUploadResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteFileUploadResponseBody
	GetSuccess() *bool
}

type DeleteFileUploadResponseBody struct {
	Data *DeleteFileUploadResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
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
	// 67E910F2-4B62-5B0C-ACA3-7547695C****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteFileUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileUploadResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileUploadResponseBody) GetData() *DeleteFileUploadResponseBodyData {
	return s.Data
}

func (s *DeleteFileUploadResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteFileUploadResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteFileUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFileUploadResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteFileUploadResponseBody) SetData(v *DeleteFileUploadResponseBodyData) *DeleteFileUploadResponseBody {
	s.Data = v
	return s
}

func (s *DeleteFileUploadResponseBody) SetErrorCode(v string) *DeleteFileUploadResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteFileUploadResponseBody) SetErrorMessage(v string) *DeleteFileUploadResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteFileUploadResponseBody) SetRequestId(v string) *DeleteFileUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFileUploadResponseBody) SetSuccess(v bool) *DeleteFileUploadResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteFileUploadResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteFileUploadResponseBodyData struct {
	// example:
	//
	// f-8*******01m
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s DeleteFileUploadResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DeleteFileUploadResponseBodyData) GoString() string {
	return s.String()
}

func (s *DeleteFileUploadResponseBodyData) GetFileId() *string {
	return s.FileId
}

func (s *DeleteFileUploadResponseBodyData) SetFileId(v string) *DeleteFileUploadResponseBodyData {
	s.FileId = &v
	return s
}

func (s *DeleteFileUploadResponseBodyData) Validate() error {
	return dara.Validate(s)
}
