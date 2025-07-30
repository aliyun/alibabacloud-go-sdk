// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSynchronizationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *CreateSynchronizationJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateSynchronizationJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *CreateSynchronizationJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateSynchronizationJobResponseBody
	GetSuccess() *string
	SetSynchronizationJobId(v string) *CreateSynchronizationJobResponseBody
	GetSynchronizationJobId() *string
}

type CreateSynchronizationJobResponseBody struct {
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 9056C8B0-5799-493A-9655-70F607B8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The ID of the data synchronization instance.
	//
	// example:
	//
	// dtshvj11k25255****
	SynchronizationJobId *string `json:"SynchronizationJobId,omitempty" xml:"SynchronizationJobId,omitempty"`
}

func (s CreateSynchronizationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSynchronizationJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateSynchronizationJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateSynchronizationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSynchronizationJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateSynchronizationJobResponseBody) GetSynchronizationJobId() *string {
	return s.SynchronizationJobId
}

func (s *CreateSynchronizationJobResponseBody) SetErrCode(v string) *CreateSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetErrMessage(v string) *CreateSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetRequestId(v string) *CreateSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetSuccess(v string) *CreateSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) SetSynchronizationJobId(v string) *CreateSynchronizationJobResponseBody {
	s.SynchronizationJobId = &v
	return s
}

func (s *CreateSynchronizationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
