// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadFileJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateUploadFileJobResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateUploadFileJobResponseBody
	GetErrorMessage() *string
	SetJobKey(v string) *CreateUploadFileJobResponseBody
	GetJobKey() *string
	SetRequestId(v string) *CreateUploadFileJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateUploadFileJobResponseBody
	GetSuccess() *bool
}

type CreateUploadFileJobResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// InvalidParameterValid
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// Unsupported url scheme : null, scheme must be https or http
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The key of the task.
	//
	// >  You can call the [GetUserUploadFileJob](https://help.aliyun.com/document_detail/206069.html) operation to query the progress and details of the task.
	//
	// example:
	//
	// 761f18031635736380812****
	JobKey *string `json:"JobKey,omitempty" xml:"JobKey,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E103C5F9-DE47-53F2-BF34-D71DF38F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateUploadFileJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadFileJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUploadFileJobResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateUploadFileJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateUploadFileJobResponseBody) GetJobKey() *string {
	return s.JobKey
}

func (s *CreateUploadFileJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateUploadFileJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateUploadFileJobResponseBody) SetErrorCode(v string) *CreateUploadFileJobResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateUploadFileJobResponseBody) SetErrorMessage(v string) *CreateUploadFileJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateUploadFileJobResponseBody) SetJobKey(v string) *CreateUploadFileJobResponseBody {
	s.JobKey = &v
	return s
}

func (s *CreateUploadFileJobResponseBody) SetRequestId(v string) *CreateUploadFileJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateUploadFileJobResponseBody) SetSuccess(v bool) *CreateUploadFileJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateUploadFileJobResponseBody) Validate() error {
	return dara.Validate(s)
}
