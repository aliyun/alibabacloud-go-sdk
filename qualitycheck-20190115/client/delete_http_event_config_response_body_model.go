// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHttpEventConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteHttpEventConfigResponseBody
	GetCode() *string
	SetData(v string) *DeleteHttpEventConfigResponseBody
	GetData() *string
	SetMessage(v string) *DeleteHttpEventConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHttpEventConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteHttpEventConfigResponseBody
	GetSuccess() *bool
}

type DeleteHttpEventConfigResponseBody struct {
	// The result code. A value of **200*	- indicates success. Any other value indicates failure. You can use this field to determine the cause of the failure.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 95
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message returned when an error occurs.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 106C6CA0-282D-4AF7-85F0-D2D24***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// You can determine whether the request is successful based on this field:
	//
	// - **true**: Successful.
	//
	// - **false/null**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteHttpEventConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHttpEventConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHttpEventConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteHttpEventConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteHttpEventConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHttpEventConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHttpEventConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteHttpEventConfigResponseBody) SetCode(v string) *DeleteHttpEventConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHttpEventConfigResponseBody) SetData(v string) *DeleteHttpEventConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteHttpEventConfigResponseBody) SetMessage(v string) *DeleteHttpEventConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHttpEventConfigResponseBody) SetRequestId(v string) *DeleteHttpEventConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHttpEventConfigResponseBody) SetSuccess(v bool) *DeleteHttpEventConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteHttpEventConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
