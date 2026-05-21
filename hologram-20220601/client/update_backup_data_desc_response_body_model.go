// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupDataDescResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateBackupDataDescResponseBody
	GetData() *bool
	SetErrorCode(v string) *UpdateBackupDataDescResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateBackupDataDescResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *UpdateBackupDataDescResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UpdateBackupDataDescResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateBackupDataDescResponseBody
	GetSuccess() *bool
}

type UpdateBackupDataDescResponseBody struct {
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
	// 87060182-4B69-5C3D-B3A1-A944D20138A7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateBackupDataDescResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupDataDescResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateBackupDataDescResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateBackupDataDescResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateBackupDataDescResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateBackupDataDescResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateBackupDataDescResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateBackupDataDescResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateBackupDataDescResponseBody) SetData(v bool) *UpdateBackupDataDescResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateBackupDataDescResponseBody) SetErrorCode(v string) *UpdateBackupDataDescResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateBackupDataDescResponseBody) SetErrorMessage(v string) *UpdateBackupDataDescResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateBackupDataDescResponseBody) SetHttpStatusCode(v string) *UpdateBackupDataDescResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateBackupDataDescResponseBody) SetRequestId(v string) *UpdateBackupDataDescResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateBackupDataDescResponseBody) SetSuccess(v bool) *UpdateBackupDataDescResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateBackupDataDescResponseBody) Validate() error {
	return dara.Validate(s)
}
