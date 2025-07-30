// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendMigrationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *SuspendMigrationJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *SuspendMigrationJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *SuspendMigrationJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *SuspendMigrationJobResponseBody
	GetSuccess() *string
}

type SuspendMigrationJobResponseBody struct {
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

func (s SuspendMigrationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SuspendMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *SuspendMigrationJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *SuspendMigrationJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *SuspendMigrationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SuspendMigrationJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *SuspendMigrationJobResponseBody) SetErrCode(v string) *SuspendMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *SuspendMigrationJobResponseBody) SetErrMessage(v string) *SuspendMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *SuspendMigrationJobResponseBody) SetRequestId(v string) *SuspendMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *SuspendMigrationJobResponseBody) SetSuccess(v string) *SuspendMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *SuspendMigrationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
