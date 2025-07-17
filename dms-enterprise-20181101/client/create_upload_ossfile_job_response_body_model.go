// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadOSSFileJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateUploadOSSFileJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateUploadOSSFileJobResponseBody
	GetErrorMessage() *string
	SetJobKey(v string) *CreateUploadOSSFileJobResponseBody
	GetJobKey() *string
	SetRequestId(v string) *CreateUploadOSSFileJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateUploadOSSFileJobResponseBody
	GetSuccess() *bool
}

type CreateUploadOSSFileJobResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The key of the file upload task. You can query the upload progress and task details. For more information, see [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html).
	//
	// example:
	//
	// 65254a4c1614235217749100e
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 4E1D2B4D-3E53-4ABC-999D-1D2520B3471A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateUploadOSSFileJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadOSSFileJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadOSSFileJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateUploadOSSFileJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateUploadOSSFileJobResponseBody) GetJobKey() *string {
	return s.JobKey
}

func (s *CreateUploadOSSFileJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUploadOSSFileJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUploadOSSFileJobResponseBody) SetErrorCode(v string) *CreateUploadOSSFileJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUploadOSSFileJobResponseBody) SetErrorMessage(v string) *CreateUploadOSSFileJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateUploadOSSFileJobResponseBody) SetJobKey(v string) *CreateUploadOSSFileJobResponseBody {
	s.JobKey = &v
	return s
}

func (s *CreateUploadOSSFileJobResponseBody) SetRequestId(v string) *CreateUploadOSSFileJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadOSSFileJobResponseBody) SetSuccess(v bool) *CreateUploadOSSFileJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUploadOSSFileJobResponseBody) Validate() error {
	return dara.Validate(s)
}
