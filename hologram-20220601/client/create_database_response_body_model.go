// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateDatabaseResponseBody
	GetData() *bool
	SetErrorCode(v string) *CreateDatabaseResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDatabaseResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *CreateDatabaseResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *CreateDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateDatabaseResponseBody
	GetSuccess() *string
}

type CreateDatabaseResponseBody struct {
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
	// 92F00FA1-CE4E-55A1-8BA9-F2C61534B518
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatabaseResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateDatabaseResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDatabaseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDatabaseResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatabaseResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateDatabaseResponseBody) SetData(v bool) *CreateDatabaseResponseBody {
	s.Data = &v
	return s
}

func (s *CreateDatabaseResponseBody) SetErrorCode(v string) *CreateDatabaseResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDatabaseResponseBody) SetErrorMessage(v string) *CreateDatabaseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDatabaseResponseBody) SetHttpStatusCode(v string) *CreateDatabaseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateDatabaseResponseBody) SetRequestId(v string) *CreateDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatabaseResponseBody) SetSuccess(v string) *CreateDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
