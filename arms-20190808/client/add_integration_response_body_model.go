// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *AddIntegrationResponseBody
	GetCode() *int32
	SetData(v string) *AddIntegrationResponseBody
	GetData() *string
	SetMessage(v string) *AddIntegrationResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddIntegrationResponseBody
	GetRequestId() *string
}

type AddIntegrationResponseBody struct {
	// Status code. 200 means success, other status codes are exceptions.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// success
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The prompt information of the returned result.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *AddIntegrationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *AddIntegrationResponseBody) GetData() *string {
	return s.Data
}

func (s *AddIntegrationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddIntegrationResponseBody) SetCode(v int32) *AddIntegrationResponseBody {
	s.Code = &v
	return s
}

func (s *AddIntegrationResponseBody) SetData(v string) *AddIntegrationResponseBody {
	s.Data = &v
	return s
}

func (s *AddIntegrationResponseBody) SetMessage(v string) *AddIntegrationResponseBody {
	s.Message = &v
	return s
}

func (s *AddIntegrationResponseBody) SetRequestId(v string) *AddIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddIntegrationResponseBody) Validate() error {
	return dara.Validate(s)
}
