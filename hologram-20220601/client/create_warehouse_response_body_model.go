// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWarehouseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *CreateWarehouseResponseBody
	GetData() *bool
	SetErrorMessage(v string) *CreateWarehouseResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *CreateWarehouseResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *CreateWarehouseResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateWarehouseResponseBody
	GetSuccess() *string
}

type CreateWarehouseResponseBody struct {
	// Indicates whether the virtual warehouse was created. Valid values: true and false.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// null
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2A8DEF6E-067E-5DB0-BAE1-2894266E6C6A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. The request result is irrelevant to the business.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateWarehouseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateWarehouseResponseBody) GoString() string {
	return s.String()
}

func (s *CreateWarehouseResponseBody) GetData() *bool {
	return s.Data
}

func (s *CreateWarehouseResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateWarehouseResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateWarehouseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateWarehouseResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateWarehouseResponseBody) SetData(v bool) *CreateWarehouseResponseBody {
	s.Data = &v
	return s
}

func (s *CreateWarehouseResponseBody) SetErrorMessage(v string) *CreateWarehouseResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateWarehouseResponseBody) SetHttpStatusCode(v string) *CreateWarehouseResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateWarehouseResponseBody) SetRequestId(v string) *CreateWarehouseResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateWarehouseResponseBody) SetSuccess(v string) *CreateWarehouseResponseBody {
	s.Success = &v
	return s
}

func (s *CreateWarehouseResponseBody) Validate() error {
	return dara.Validate(s)
}
