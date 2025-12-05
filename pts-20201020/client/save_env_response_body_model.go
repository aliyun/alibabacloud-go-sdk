// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveEnvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SaveEnvResponseBody
	GetCode() *string
	SetEnvId(v string) *SaveEnvResponseBody
	GetEnvId() *string
	SetHttpStatusCode(v int32) *SaveEnvResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SaveEnvResponseBody
	GetMessage() *string
	SetRequestId(v string) *SaveEnvResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveEnvResponseBody
	GetSuccess() *bool
}

type SaveEnvResponseBody struct {
	// The system status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the environment.
	//
	// example:
	//
	// 10YPA8H
	EnvId *string `json:"EnvId,omitempty" xml:"EnvId,omitempty"`
	// The HTTP status code. If the operation is successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The error message. If the operation is successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SaveEnvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveEnvResponseBody) GoString() string {
	return s.String()
}

func (s *SaveEnvResponseBody) GetCode() *string {
	return s.Code
}

func (s *SaveEnvResponseBody) GetEnvId() *string {
	return s.EnvId
}

func (s *SaveEnvResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SaveEnvResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SaveEnvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveEnvResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveEnvResponseBody) SetCode(v string) *SaveEnvResponseBody {
	s.Code = &v
	return s
}

func (s *SaveEnvResponseBody) SetEnvId(v string) *SaveEnvResponseBody {
	s.EnvId = &v
	return s
}

func (s *SaveEnvResponseBody) SetHttpStatusCode(v int32) *SaveEnvResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SaveEnvResponseBody) SetMessage(v string) *SaveEnvResponseBody {
	s.Message = &v
	return s
}

func (s *SaveEnvResponseBody) SetRequestId(v string) *SaveEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveEnvResponseBody) SetSuccess(v bool) *SaveEnvResponseBody {
	s.Success = &v
	return s
}

func (s *SaveEnvResponseBody) Validate() error {
	return dara.Validate(s)
}
