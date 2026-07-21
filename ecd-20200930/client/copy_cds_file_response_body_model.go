// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCdsFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CopyCdsFileResponseBody
	GetCode() *string
	SetCopyCdsFileModel(v *CopyCdsFileResponseBodyCopyCdsFileModel) *CopyCdsFileResponseBody
	GetCopyCdsFileModel() *CopyCdsFileResponseBodyCopyCdsFileModel
	SetMessage(v string) *CopyCdsFileResponseBody
	GetMessage() *string
	SetRequestId(v string) *CopyCdsFileResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CopyCdsFileResponseBody
	GetSuccess() *string
}

type CopyCdsFileResponseBody struct {
	// The execution result. A value of `success` indicates success. Otherwise, an error message is returned.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of copying the file.
	CopyCdsFileModel *CopyCdsFileResponseBodyCopyCdsFileModel `json:"CopyCdsFileModel,omitempty" xml:"CopyCdsFileModel,omitempty" type:"Struct"`
	// The error message. This parameter is not returned if Code is `success`.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 93AD30C1-16B8-5C54-AD23-A51FF53F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CopyCdsFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyCdsFileResponseBody) GoString() string {
	return s.String()
}

func (s *CopyCdsFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *CopyCdsFileResponseBody) GetCopyCdsFileModel() *CopyCdsFileResponseBodyCopyCdsFileModel {
	return s.CopyCdsFileModel
}

func (s *CopyCdsFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CopyCdsFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyCdsFileResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CopyCdsFileResponseBody) SetCode(v string) *CopyCdsFileResponseBody {
	s.Code = &v
	return s
}

func (s *CopyCdsFileResponseBody) SetCopyCdsFileModel(v *CopyCdsFileResponseBodyCopyCdsFileModel) *CopyCdsFileResponseBody {
	s.CopyCdsFileModel = v
	return s
}

func (s *CopyCdsFileResponseBody) SetMessage(v string) *CopyCdsFileResponseBody {
	s.Message = &v
	return s
}

func (s *CopyCdsFileResponseBody) SetRequestId(v string) *CopyCdsFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyCdsFileResponseBody) SetSuccess(v string) *CopyCdsFileResponseBody {
	s.Success = &v
	return s
}

func (s *CopyCdsFileResponseBody) Validate() error {
	if s.CopyCdsFileModel != nil {
		if err := s.CopyCdsFileModel.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CopyCdsFileResponseBodyCopyCdsFileModel struct {
	// The asynchronous task ID. This field is not returned when a file is copied. When a folder is copied, the copy operation is performed asynchronously in the background, so this field is returned. You can call [GetAsyncTask](https://help.aliyun.com/document_detail/2357404.html) and pass in this asynchronous task ID to obtain the task details.
	//
	// example:
	//
	// 4221bf6e6ab43a255edc4463bffa6f5f5d31****
	AsyncTaskId *string `json:"AsyncTaskId,omitempty" xml:"AsyncTaskId,omitempty"`
	// The ID of the new file or folder after the copy operation.
	//
	// example:
	//
	// 6400727cb878821bcb414615a609b4072463****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s CopyCdsFileResponseBodyCopyCdsFileModel) String() string {
	return dara.Prettify(s)
}

func (s CopyCdsFileResponseBodyCopyCdsFileModel) GoString() string {
	return s.String()
}

func (s *CopyCdsFileResponseBodyCopyCdsFileModel) GetAsyncTaskId() *string {
	return s.AsyncTaskId
}

func (s *CopyCdsFileResponseBodyCopyCdsFileModel) GetFileId() *string {
	return s.FileId
}

func (s *CopyCdsFileResponseBodyCopyCdsFileModel) SetAsyncTaskId(v string) *CopyCdsFileResponseBodyCopyCdsFileModel {
	s.AsyncTaskId = &v
	return s
}

func (s *CopyCdsFileResponseBodyCopyCdsFileModel) SetFileId(v string) *CopyCdsFileResponseBodyCopyCdsFileModel {
	s.FileId = &v
	return s
}

func (s *CopyCdsFileResponseBodyCopyCdsFileModel) Validate() error {
	return dara.Validate(s)
}
