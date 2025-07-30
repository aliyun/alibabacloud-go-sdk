// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetSynchronizationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *ResetSynchronizationJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ResetSynchronizationJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *ResetSynchronizationJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ResetSynchronizationJobResponseBody
	GetSuccess() *string
}

type ResetSynchronizationJobResponseBody struct {
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
	// FDC111B1-ACBF-457D-9656-247FDEE9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ResetSynchronizationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *ResetSynchronizationJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ResetSynchronizationJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ResetSynchronizationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetSynchronizationJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ResetSynchronizationJobResponseBody) SetErrCode(v string) *ResetSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ResetSynchronizationJobResponseBody) SetErrMessage(v string) *ResetSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ResetSynchronizationJobResponseBody) SetRequestId(v string) *ResetSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetSynchronizationJobResponseBody) SetSuccess(v string) *ResetSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *ResetSynchronizationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
