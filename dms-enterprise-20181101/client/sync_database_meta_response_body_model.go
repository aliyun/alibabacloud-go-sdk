// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncDatabaseMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SyncDatabaseMetaResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SyncDatabaseMetaResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SyncDatabaseMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncDatabaseMetaResponseBody
	GetSuccess() *bool
}

type SyncDatabaseMetaResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A99CD576-1E18-4E86-931E-C3CCE56DC030
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncDatabaseMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncDatabaseMetaResponseBody) GoString() string {
	return s.String()
}

func (s *SyncDatabaseMetaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SyncDatabaseMetaResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SyncDatabaseMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncDatabaseMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncDatabaseMetaResponseBody) SetErrorCode(v string) *SyncDatabaseMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SyncDatabaseMetaResponseBody) SetErrorMessage(v string) *SyncDatabaseMetaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SyncDatabaseMetaResponseBody) SetRequestId(v string) *SyncDatabaseMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncDatabaseMetaResponseBody) SetSuccess(v bool) *SyncDatabaseMetaResponseBody {
	s.Success = &v
	return s
}

func (s *SyncDatabaseMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
