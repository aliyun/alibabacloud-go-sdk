// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnvironmentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateEnvironmentResponseBody
	GetCode() *string
	SetData(v *CreateEnvironmentResponseBodyData) *CreateEnvironmentResponseBody
	GetData() *CreateEnvironmentResponseBodyData
	SetMessage(v string) *CreateEnvironmentResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateEnvironmentResponseBody
	GetRequestId() *string
}

type CreateEnvironmentResponseBody struct {
	// Response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// Response data.
	Data *CreateEnvironmentResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// Response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Request ID, used for tracing the API call chain.
	//
	// example:
	//
	// 3C3B9A12-3868-5EB9-8BEA-F99E03DD125C
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateEnvironmentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvironmentResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateEnvironmentResponseBody) GetData() *CreateEnvironmentResponseBodyData {
	return s.Data
}

func (s *CreateEnvironmentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateEnvironmentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnvironmentResponseBody) SetCode(v string) *CreateEnvironmentResponseBody {
	s.Code = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetData(v *CreateEnvironmentResponseBodyData) *CreateEnvironmentResponseBody {
	s.Data = v
	return s
}

func (s *CreateEnvironmentResponseBody) SetMessage(v string) *CreateEnvironmentResponseBody {
	s.Message = &v
	return s
}

func (s *CreateEnvironmentResponseBody) SetRequestId(v string) *CreateEnvironmentResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnvironmentResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateEnvironmentResponseBodyData struct {
	// Environment ID.
	//
	// example:
	//
	// env-cq7l5s5lhtgi6qasrdc0
	EnvironmentId *string `json:"environmentId,omitempty" xml:"environmentId,omitempty"`
}

func (s CreateEnvironmentResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateEnvironmentResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateEnvironmentResponseBodyData) GetEnvironmentId() *string {
	return s.EnvironmentId
}

func (s *CreateEnvironmentResponseBodyData) SetEnvironmentId(v string) *CreateEnvironmentResponseBodyData {
	s.EnvironmentId = &v
	return s
}

func (s *CreateEnvironmentResponseBodyData) Validate() error {
	return dara.Validate(s)
}
