// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntegrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteIntegrationResponseBody
	GetCode() *int32
	SetData(v string) *DeleteIntegrationResponseBody
	GetData() *string
	SetMessage(v string) *DeleteIntegrationResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteIntegrationResponseBody
	GetRequestId() *string
}

type DeleteIntegrationResponseBody struct {
	// The HTTP status code. The status code 200 indicates that the request was successful. Other status codes indicate that the request failed.
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
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1A9C645C-C83F-4C9D-8CCB-29BEC9E1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIntegrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntegrationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIntegrationResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteIntegrationResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteIntegrationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteIntegrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIntegrationResponseBody) SetCode(v int32) *DeleteIntegrationResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteIntegrationResponseBody) SetData(v string) *DeleteIntegrationResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteIntegrationResponseBody) SetMessage(v string) *DeleteIntegrationResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteIntegrationResponseBody) SetRequestId(v string) *DeleteIntegrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIntegrationResponseBody) Validate() error {
	return dara.Validate(s)
}
