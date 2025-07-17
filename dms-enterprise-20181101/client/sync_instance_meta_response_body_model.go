// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSyncInstanceMetaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *SyncInstanceMetaResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *SyncInstanceMetaResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *SyncInstanceMetaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SyncInstanceMetaResponseBody
	GetSuccess() *bool
}

type SyncInstanceMetaResponseBody struct {
	// Details about the topology of the data table.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error code returned.
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
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SyncInstanceMetaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SyncInstanceMetaResponseBody) GoString() string {
	return s.String()
}

func (s *SyncInstanceMetaResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *SyncInstanceMetaResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *SyncInstanceMetaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SyncInstanceMetaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SyncInstanceMetaResponseBody) SetErrorCode(v string) *SyncInstanceMetaResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *SyncInstanceMetaResponseBody) SetErrorMessage(v string) *SyncInstanceMetaResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *SyncInstanceMetaResponseBody) SetRequestId(v string) *SyncInstanceMetaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SyncInstanceMetaResponseBody) SetSuccess(v bool) *SyncInstanceMetaResponseBody {
	s.Success = &v
	return s
}

func (s *SyncInstanceMetaResponseBody) Validate() error {
	return dara.Validate(s)
}
