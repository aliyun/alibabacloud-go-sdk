// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNumbersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteNumbersResponseBody
	GetCode() *string
	SetData(v string) *DeleteNumbersResponseBody
	GetData() *string
	SetHttpStatusCode(v int32) *DeleteNumbersResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *DeleteNumbersResponseBody
	GetMessage() *string
	SetParams(v []*string) *DeleteNumbersResponseBody
	GetParams() []*string
	SetRequestId(v string) *DeleteNumbersResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteNumbersResponseBody
	GetSuccess() *bool
}

type DeleteNumbersResponseBody struct {
	// The response code.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The scenario version ID.
	//
	// example:
	//
	// 4f9a8e2b-6c1d-4a7e-9b3f-2d5c8a1e7b15
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned by the API.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of variable values in the error message.
	Params []*string `json:"Params,omitempty" xml:"Params,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 254EB995-DEDF-48A4-9101-9CA5B72FFBCC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteNumbersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNumbersResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNumbersResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteNumbersResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteNumbersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DeleteNumbersResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteNumbersResponseBody) GetParams() []*string {
	return s.Params
}

func (s *DeleteNumbersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNumbersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteNumbersResponseBody) SetCode(v string) *DeleteNumbersResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteNumbersResponseBody) SetData(v string) *DeleteNumbersResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteNumbersResponseBody) SetHttpStatusCode(v int32) *DeleteNumbersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DeleteNumbersResponseBody) SetMessage(v string) *DeleteNumbersResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteNumbersResponseBody) SetParams(v []*string) *DeleteNumbersResponseBody {
	s.Params = v
	return s
}

func (s *DeleteNumbersResponseBody) SetRequestId(v string) *DeleteNumbersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNumbersResponseBody) SetSuccess(v bool) *DeleteNumbersResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteNumbersResponseBody) Validate() error {
	return dara.Validate(s)
}
