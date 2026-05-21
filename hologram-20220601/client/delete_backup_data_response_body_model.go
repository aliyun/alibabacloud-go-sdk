// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteBackupDataResponseBody
	GetData() *bool
	SetErrorCode(v string) *DeleteBackupDataResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DeleteBackupDataResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *DeleteBackupDataResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *DeleteBackupDataResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteBackupDataResponseBody
	GetSuccess() *bool
}

type DeleteBackupDataResponseBody struct {
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
	// 1907BA1B-B443-5B39-8113-7F3E260DC20D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteBackupDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupDataResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteBackupDataResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteBackupDataResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DeleteBackupDataResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DeleteBackupDataResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DeleteBackupDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteBackupDataResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteBackupDataResponseBody) SetData(v bool) *DeleteBackupDataResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteBackupDataResponseBody) SetErrorCode(v string) *DeleteBackupDataResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DeleteBackupDataResponseBody) SetErrorMessage(v string) *DeleteBackupDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteBackupDataResponseBody) SetHttpStatusCode(v string) *DeleteBackupDataResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteBackupDataResponseBody) SetRequestId(v string) *DeleteBackupDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteBackupDataResponseBody) SetSuccess(v bool) *DeleteBackupDataResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteBackupDataResponseBody) Validate() error {
	return dara.Validate(s)
}
