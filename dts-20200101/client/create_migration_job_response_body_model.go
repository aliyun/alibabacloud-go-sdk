// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMigrationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *CreateMigrationJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *CreateMigrationJobResponseBody
	GetErrMessage() *string
	SetMigrationJobId(v string) *CreateMigrationJobResponseBody
	GetMigrationJobId() *string
	SetRequestId(v string) *CreateMigrationJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateMigrationJobResponseBody
	GetSuccess() *string
}

type CreateMigrationJobResponseBody struct {
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
	// The ID of the data migration instance.
	//
	// example:
	//
	// dtsi8911td9233****
	MigrationJobId *string `json:"MigrationJobId,omitempty" xml:"MigrationJobId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C166D79D-436B-45F0-B5A5-25E1959F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMigrationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMigrationJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *CreateMigrationJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *CreateMigrationJobResponseBody) GetMigrationJobId() *string {
	return s.MigrationJobId
}

func (s *CreateMigrationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMigrationJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateMigrationJobResponseBody) SetErrCode(v string) *CreateMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetErrMessage(v string) *CreateMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetMigrationJobId(v string) *CreateMigrationJobResponseBody {
	s.MigrationJobId = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetRequestId(v string) *CreateMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMigrationJobResponseBody) SetSuccess(v string) *CreateMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMigrationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
