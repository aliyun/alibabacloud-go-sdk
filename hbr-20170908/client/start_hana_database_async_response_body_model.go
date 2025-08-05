// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartHanaDatabaseAsyncResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *StartHanaDatabaseAsyncResponseBody
	GetCode() *string
	SetMessage(v string) *StartHanaDatabaseAsyncResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartHanaDatabaseAsyncResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartHanaDatabaseAsyncResponseBody
	GetSuccess() *bool
	SetTaskId(v string) *StartHanaDatabaseAsyncResponseBody
	GetTaskId() *string
}

type StartHanaDatabaseAsyncResponseBody struct {
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, "successful" is returned. If the request failed, an error message is returned.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the job that is used to initialize the backup vault. You can call the DescribeTask operation to query the job status.
	//
	// example:
	//
	// t-000bjt479yefheij1o0x
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StartHanaDatabaseAsyncResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartHanaDatabaseAsyncResponseBody) GoString() string {
	return s.String()
}

func (s *StartHanaDatabaseAsyncResponseBody) GetCode() *string {
	return s.Code
}

func (s *StartHanaDatabaseAsyncResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartHanaDatabaseAsyncResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartHanaDatabaseAsyncResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartHanaDatabaseAsyncResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StartHanaDatabaseAsyncResponseBody) SetCode(v string) *StartHanaDatabaseAsyncResponseBody {
	s.Code = &v
	return s
}

func (s *StartHanaDatabaseAsyncResponseBody) SetMessage(v string) *StartHanaDatabaseAsyncResponseBody {
	s.Message = &v
	return s
}

func (s *StartHanaDatabaseAsyncResponseBody) SetRequestId(v string) *StartHanaDatabaseAsyncResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartHanaDatabaseAsyncResponseBody) SetSuccess(v bool) *StartHanaDatabaseAsyncResponseBody {
	s.Success = &v
	return s
}

func (s *StartHanaDatabaseAsyncResponseBody) SetTaskId(v string) *StartHanaDatabaseAsyncResponseBody {
	s.TaskId = &v
	return s
}

func (s *StartHanaDatabaseAsyncResponseBody) Validate() error {
	return dara.Validate(s)
}
