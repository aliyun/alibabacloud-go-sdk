// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveEnvResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RemoveEnvResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *RemoveEnvResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *RemoveEnvResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveEnvResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RemoveEnvResponseBody
	GetSuccess() *bool
}

type RemoveEnvResponseBody struct {
	// The system status code. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. If the request was successful, this parameter is not returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, this parameter is not returned.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
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

func (s RemoveEnvResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveEnvResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveEnvResponseBody) GetCode() *string {
	return s.Code
}

func (s *RemoveEnvResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *RemoveEnvResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveEnvResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveEnvResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RemoveEnvResponseBody) SetCode(v string) *RemoveEnvResponseBody {
	s.Code = &v
	return s
}

func (s *RemoveEnvResponseBody) SetHttpStatusCode(v int32) *RemoveEnvResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *RemoveEnvResponseBody) SetMessage(v string) *RemoveEnvResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveEnvResponseBody) SetRequestId(v string) *RemoveEnvResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveEnvResponseBody) SetSuccess(v bool) *RemoveEnvResponseBody {
	s.Success = &v
	return s
}

func (s *RemoveEnvResponseBody) Validate() error {
	return dara.Validate(s)
}
