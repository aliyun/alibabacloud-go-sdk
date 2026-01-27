// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryDownloadTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetryDownloadTaskResponseBody
	GetCode() *string
	SetData(v string) *RetryDownloadTaskResponseBody
	GetData() *string
	SetErrCode(v string) *RetryDownloadTaskResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *RetryDownloadTaskResponseBody
	GetErrMessage() *string
	SetMessage(v string) *RetryDownloadTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetryDownloadTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *RetryDownloadTaskResponseBody
	GetSuccess() *string
}

type RetryDownloadTaskResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Success
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// instanceName can not be empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 49FE4E8E-39B9-56DE-BC07-5AEBFAXXXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetryDownloadTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetryDownloadTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RetryDownloadTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetryDownloadTaskResponseBody) GetData() *string {
	return s.Data
}

func (s *RetryDownloadTaskResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *RetryDownloadTaskResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *RetryDownloadTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetryDownloadTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetryDownloadTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *RetryDownloadTaskResponseBody) SetCode(v string) *RetryDownloadTaskResponseBody {
	s.Code = &v
	return s
}

func (s *RetryDownloadTaskResponseBody) SetData(v string) *RetryDownloadTaskResponseBody {
	s.Data = &v
	return s
}

func (s *RetryDownloadTaskResponseBody) SetErrCode(v string) *RetryDownloadTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *RetryDownloadTaskResponseBody) SetErrMessage(v string) *RetryDownloadTaskResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *RetryDownloadTaskResponseBody) SetMessage(v string) *RetryDownloadTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RetryDownloadTaskResponseBody) SetRequestId(v string) *RetryDownloadTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryDownloadTaskResponseBody) SetSuccess(v string) *RetryDownloadTaskResponseBody {
	s.Success = &v
	return s
}

func (s *RetryDownloadTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
