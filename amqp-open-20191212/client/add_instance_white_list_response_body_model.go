// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddInstanceWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddInstanceWhiteListResponseBody
	GetCode() *string
	SetData(v interface{}) *AddInstanceWhiteListResponseBody
	GetData() interface{}
	SetMessage(v string) *AddInstanceWhiteListResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddInstanceWhiteListResponseBody
	GetRequestId() *string
	SetStatusCode(v string) *AddInstanceWhiteListResponseBody
	GetStatusCode() *string
	SetSuccess(v bool) *AddInstanceWhiteListResponseBody
	GetSuccess() *bool
}

type AddInstanceWhiteListResponseBody struct {
	// The status code. A value of `200` indicates a successful request.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Data returned by the operation.
	//
	// example:
	//
	// true
	Data interface{} `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned with the response. This is typically an error message if the request fails.
	//
	// example:
	//
	// The specified parameter \\"[\\"vpc-xxx\\"]\\" is not valid
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8BFB1C9D-08A2-4859-A47C-403C9EFA2***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The HTTP status code for the response. A value of `200` indicates success.
	//
	// example:
	//
	// 200
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s AddInstanceWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddInstanceWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *AddInstanceWhiteListResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddInstanceWhiteListResponseBody) GetData() interface{} {
	return s.Data
}

func (s *AddInstanceWhiteListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddInstanceWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddInstanceWhiteListResponseBody) GetStatusCode() *string {
	return s.StatusCode
}

func (s *AddInstanceWhiteListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AddInstanceWhiteListResponseBody) SetCode(v string) *AddInstanceWhiteListResponseBody {
	s.Code = &v
	return s
}

func (s *AddInstanceWhiteListResponseBody) SetData(v interface{}) *AddInstanceWhiteListResponseBody {
	s.Data = v
	return s
}

func (s *AddInstanceWhiteListResponseBody) SetMessage(v string) *AddInstanceWhiteListResponseBody {
	s.Message = &v
	return s
}

func (s *AddInstanceWhiteListResponseBody) SetRequestId(v string) *AddInstanceWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddInstanceWhiteListResponseBody) SetStatusCode(v string) *AddInstanceWhiteListResponseBody {
	s.StatusCode = &v
	return s
}

func (s *AddInstanceWhiteListResponseBody) SetSuccess(v bool) *AddInstanceWhiteListResponseBody {
	s.Success = &v
	return s
}

func (s *AddInstanceWhiteListResponseBody) Validate() error {
	return dara.Validate(s)
}
