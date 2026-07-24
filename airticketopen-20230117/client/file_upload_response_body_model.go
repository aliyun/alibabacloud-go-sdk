// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileUploadResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *FileUploadResponseBody
	GetRequestId() *string
	SetData(v *FileUploadResponseBodyData) *FileUploadResponseBody
	GetData() *FileUploadResponseBodyData
	SetErrorCode(v string) *FileUploadResponseBody
	GetErrorCode() *string
	SetErrorData(v interface{}) *FileUploadResponseBody
	GetErrorData() interface{}
	SetErrorMsg(v string) *FileUploadResponseBody
	GetErrorMsg() *string
	SetStatus(v int32) *FileUploadResponseBody
	GetStatus() *int32
	SetSuccess(v bool) *FileUploadResponseBody
	GetSuccess() *bool
}

type FileUploadResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 51593418-8C73-5E47-8BA8-3F1D4A00CC0B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The data returned for a successful request.
	Data *FileUploadResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The business error code.
	//
	// example:
	//
	// null
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// The data returned with an error response.
	//
	// example:
	//
	// null
	ErrorData interface{} `json:"error_data,omitempty" xml:"error_data,omitempty"`
	// The error message.
	//
	// example:
	//
	// null
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// The HTTP status code. The value is always 200 for successful requests.
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s FileUploadResponseBody) String() string {
	return dara.Prettify(s)
}

func (s FileUploadResponseBody) GoString() string {
	return s.String()
}

func (s *FileUploadResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *FileUploadResponseBody) GetData() *FileUploadResponseBodyData {
	return s.Data
}

func (s *FileUploadResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *FileUploadResponseBody) GetErrorData() interface{} {
	return s.ErrorData
}

func (s *FileUploadResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *FileUploadResponseBody) GetStatus() *int32 {
	return s.Status
}

func (s *FileUploadResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *FileUploadResponseBody) SetRequestId(v string) *FileUploadResponseBody {
	s.RequestId = &v
	return s
}

func (s *FileUploadResponseBody) SetData(v *FileUploadResponseBodyData) *FileUploadResponseBody {
	s.Data = v
	return s
}

func (s *FileUploadResponseBody) SetErrorCode(v string) *FileUploadResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *FileUploadResponseBody) SetErrorData(v interface{}) *FileUploadResponseBody {
	s.ErrorData = v
	return s
}

func (s *FileUploadResponseBody) SetErrorMsg(v string) *FileUploadResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *FileUploadResponseBody) SetStatus(v int32) *FileUploadResponseBody {
	s.Status = &v
	return s
}

func (s *FileUploadResponseBody) SetSuccess(v bool) *FileUploadResponseBody {
	s.Success = &v
	return s
}

func (s *FileUploadResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type FileUploadResponseBodyData struct {
	// The file name after upload.
	//
	// example:
	//
	// https://fliggy-flight-jinghang-bucket.oss-cn-zhangjiakou.aliyuncs.com/suez/flight_suez_9a634376****47.jpeg
	UploadedFileUrl *string `json:"uploaded_file_url,omitempty" xml:"uploaded_file_url,omitempty"`
}

func (s FileUploadResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s FileUploadResponseBodyData) GoString() string {
	return s.String()
}

func (s *FileUploadResponseBodyData) GetUploadedFileUrl() *string {
	return s.UploadedFileUrl
}

func (s *FileUploadResponseBodyData) SetUploadedFileUrl(v string) *FileUploadResponseBodyData {
	s.UploadedFileUrl = &v
	return s
}

func (s *FileUploadResponseBodyData) Validate() error {
	return dara.Validate(s)
}
