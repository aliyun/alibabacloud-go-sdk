// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSynchronizationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *StartSynchronizationJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StartSynchronizationJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *StartSynchronizationJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *StartSynchronizationJobResponseBody
	GetSuccess() *string
}

type StartSynchronizationJobResponseBody struct {
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
	// The request ID.
	//
	// example:
	//
	// FDC111B1-ACBF-457D-9656-247FDEE9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartSynchronizationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSynchronizationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartSynchronizationJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StartSynchronizationJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StartSynchronizationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSynchronizationJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *StartSynchronizationJobResponseBody) SetErrCode(v string) *StartSynchronizationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartSynchronizationJobResponseBody) SetErrMessage(v string) *StartSynchronizationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartSynchronizationJobResponseBody) SetRequestId(v string) *StartSynchronizationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSynchronizationJobResponseBody) SetSuccess(v string) *StartSynchronizationJobResponseBody {
	s.Success = &v
	return s
}

func (s *StartSynchronizationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
