// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateManualBackupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateManualBackupResponseBody
	GetData() *bool
	SetErrorCode(v string) *CreateManualBackupResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateManualBackupResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *CreateManualBackupResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *CreateManualBackupResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateManualBackupResponseBody
	GetSuccess() *bool
}

type CreateManualBackupResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 5C11C8CB-9851-504D-8A41-3CF5F4FC0358
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateManualBackupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateManualBackupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateManualBackupResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateManualBackupResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateManualBackupResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateManualBackupResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateManualBackupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateManualBackupResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateManualBackupResponseBody) SetData(v bool) *CreateManualBackupResponseBody {
	s.Data = &v
	return s
}

func (s *CreateManualBackupResponseBody) SetErrorCode(v string) *CreateManualBackupResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateManualBackupResponseBody) SetErrorMessage(v string) *CreateManualBackupResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateManualBackupResponseBody) SetHttpStatusCode(v string) *CreateManualBackupResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateManualBackupResponseBody) SetRequestId(v string) *CreateManualBackupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateManualBackupResponseBody) SetSuccess(v bool) *CreateManualBackupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateManualBackupResponseBody) Validate() error {
	return dara.Validate(s)
}
