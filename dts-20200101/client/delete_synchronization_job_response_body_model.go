// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSynchronizationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DeleteSynchronizationJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteSynchronizationJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DeleteSynchronizationJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteSynchronizationJobResponseBody
	GetSuccess() *string
}

type DeleteSynchronizationJobResponseBody struct {
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

func (s DeleteSynchronizationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSynchronizationJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteSynchronizationJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteSynchronizationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSynchronizationJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteSynchronizationJobResponseBody) SetErrCode(v string) *DeleteSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteSynchronizationJobResponseBody) SetErrMessage(v string) *DeleteSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteSynchronizationJobResponseBody) SetRequestId(v string) *DeleteSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSynchronizationJobResponseBody) SetSuccess(v string) *DeleteSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSynchronizationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
