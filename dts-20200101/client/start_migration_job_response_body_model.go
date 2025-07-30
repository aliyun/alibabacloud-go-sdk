// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartMigrationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *StartMigrationJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *StartMigrationJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *StartMigrationJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *StartMigrationJobResponseBody
	GetSuccess() *string
}

type StartMigrationJobResponseBody struct {
	// The error code returned if the request failed.
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
	// Request ID.
	//
	// example:
	//
	// FDC111B1-ACBF-457D-9656-247FDEE9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartMigrationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *StartMigrationJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *StartMigrationJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *StartMigrationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartMigrationJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *StartMigrationJobResponseBody) SetErrCode(v string) *StartMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *StartMigrationJobResponseBody) SetErrMessage(v string) *StartMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *StartMigrationJobResponseBody) SetRequestId(v string) *StartMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartMigrationJobResponseBody) SetSuccess(v string) *StartMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *StartMigrationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
