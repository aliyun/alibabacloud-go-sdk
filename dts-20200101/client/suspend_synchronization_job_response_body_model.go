// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendSynchronizationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *SuspendSynchronizationJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SuspendSynchronizationJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *SuspendSynchronizationJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SuspendSynchronizationJobResponseBody
	GetSuccess() *string
}

type SuspendSynchronizationJobResponseBody struct {
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
	// C306C198-7807-409D-930A-D6CE6C32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SuspendSynchronizationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendSynchronizationJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SuspendSynchronizationJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SuspendSynchronizationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendSynchronizationJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SuspendSynchronizationJobResponseBody) SetErrCode(v string) *SuspendSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SuspendSynchronizationJobResponseBody) SetErrMessage(v string) *SuspendSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SuspendSynchronizationJobResponseBody) SetRequestId(v string) *SuspendSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendSynchronizationJobResponseBody) SetSuccess(v string) *SuspendSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendSynchronizationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
