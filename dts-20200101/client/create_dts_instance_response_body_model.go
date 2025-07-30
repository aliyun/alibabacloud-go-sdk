// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDtsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *CreateDtsInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateDtsInstanceResponseBody
	GetErrMessage() *string
	SetInstanceId(v string) *CreateDtsInstanceResponseBody
	GetInstanceId() *string
	SetJobId(v string) *CreateDtsInstanceResponseBody
	GetJobId() *string
	SetRequestId(v string) *CreateDtsInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateDtsInstanceResponseBody
	GetSuccess() *string
}

type CreateDtsInstanceResponseBody struct {
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the DTS instance.
	//
	// example:
	//
	// dtsbi6e22ay243****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the task.
	//
	// example:
	//
	// bi6e22ay243****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C166D79D-436B-45F0-B5A5-25E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDtsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDtsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDtsInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateDtsInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateDtsInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateDtsInstanceResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *CreateDtsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDtsInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateDtsInstanceResponseBody) SetErrCode(v string) *CreateDtsInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateDtsInstanceResponseBody) SetErrMessage(v string) *CreateDtsInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateDtsInstanceResponseBody) SetInstanceId(v string) *CreateDtsInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateDtsInstanceResponseBody) SetJobId(v string) *CreateDtsInstanceResponseBody {
	s.JobId = &v
	return s
}

func (s *CreateDtsInstanceResponseBody) SetRequestId(v string) *CreateDtsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDtsInstanceResponseBody) SetSuccess(v string) *CreateDtsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDtsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
