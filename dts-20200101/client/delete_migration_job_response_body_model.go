// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMigrationJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrCode(v string) *DeleteMigrationJobResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DeleteMigrationJobResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DeleteMigrationJobResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteMigrationJobResponseBody
	GetSuccess() *string
}

type DeleteMigrationJobResponseBody struct {
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
	// F28A96B1-F897-4246-833B-310A3345****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteMigrationJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMigrationJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMigrationJobResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DeleteMigrationJobResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DeleteMigrationJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMigrationJobResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteMigrationJobResponseBody) SetErrCode(v string) *DeleteMigrationJobResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetErrMessage(v string) *DeleteMigrationJobResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetRequestId(v string) *DeleteMigrationJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) SetSuccess(v string) *DeleteMigrationJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteMigrationJobResponseBody) Validate() error {
	return dara.Validate(s)
}
