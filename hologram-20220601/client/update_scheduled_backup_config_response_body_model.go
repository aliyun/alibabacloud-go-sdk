// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledBackupConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateScheduledBackupConfigResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateScheduledBackupConfigResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateScheduledBackupConfigResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *UpdateScheduledBackupConfigResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UpdateScheduledBackupConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateScheduledBackupConfigResponseBody
	GetSuccess() *bool
}

type UpdateScheduledBackupConfigResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// null
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 69221013-D303-5046-91A7-A0BE2BC2E7D7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateScheduledBackupConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledBackupConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScheduledBackupConfigResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateScheduledBackupConfigResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateScheduledBackupConfigResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateScheduledBackupConfigResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateScheduledBackupConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScheduledBackupConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateScheduledBackupConfigResponseBody) SetData(v bool) *UpdateScheduledBackupConfigResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateScheduledBackupConfigResponseBody) SetErrorCode(v string) *UpdateScheduledBackupConfigResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateScheduledBackupConfigResponseBody) SetErrorMessage(v string) *UpdateScheduledBackupConfigResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateScheduledBackupConfigResponseBody) SetHttpStatusCode(v string) *UpdateScheduledBackupConfigResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateScheduledBackupConfigResponseBody) SetRequestId(v string) *UpdateScheduledBackupConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScheduledBackupConfigResponseBody) SetSuccess(v bool) *UpdateScheduledBackupConfigResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateScheduledBackupConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
