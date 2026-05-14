// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExternalDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateExternalDatabaseResponseBody
	GetData() *bool
	SetErrorCode(v string) *CreateExternalDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateExternalDatabaseResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *CreateExternalDatabaseResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *CreateExternalDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateExternalDatabaseResponseBody
	GetSuccess() *string
}

type CreateExternalDatabaseResponseBody struct {
	// example:
	//
	// true
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
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateExternalDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateExternalDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateExternalDatabaseResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateExternalDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateExternalDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateExternalDatabaseResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateExternalDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateExternalDatabaseResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateExternalDatabaseResponseBody) SetData(v bool) *CreateExternalDatabaseResponseBody {
	s.Data = &v
	return s
}

func (s *CreateExternalDatabaseResponseBody) SetErrorCode(v string) *CreateExternalDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateExternalDatabaseResponseBody) SetErrorMessage(v string) *CreateExternalDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateExternalDatabaseResponseBody) SetHttpStatusCode(v string) *CreateExternalDatabaseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateExternalDatabaseResponseBody) SetRequestId(v string) *CreateExternalDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateExternalDatabaseResponseBody) SetSuccess(v string) *CreateExternalDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *CreateExternalDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
