// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstancePortResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *UpdateInstancePortResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *UpdateInstancePortResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *UpdateInstancePortResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *UpdateInstancePortResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateInstancePortResponseBody
	GetSuccess() *string
	SetData(v bool) *UpdateInstancePortResponseBody
	GetData() *bool
}

type UpdateInstancePortResponseBody struct {
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
	// D1303CD4-AA70-5998-8025-F55B22C50840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// true
	Data *bool `json:"data,omitempty" xml:"data,omitempty"`
}

func (s UpdateInstancePortResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstancePortResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateInstancePortResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpdateInstancePortResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *UpdateInstancePortResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *UpdateInstancePortResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateInstancePortResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateInstancePortResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateInstancePortResponseBody) SetErrorCode(v string) *UpdateInstancePortResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *UpdateInstancePortResponseBody) SetErrorMessage(v string) *UpdateInstancePortResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateInstancePortResponseBody) SetHttpStatusCode(v string) *UpdateInstancePortResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateInstancePortResponseBody) SetRequestId(v string) *UpdateInstancePortResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateInstancePortResponseBody) SetSuccess(v string) *UpdateInstancePortResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateInstancePortResponseBody) SetData(v bool) *UpdateInstancePortResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateInstancePortResponseBody) Validate() error {
	return dara.Validate(s)
}
