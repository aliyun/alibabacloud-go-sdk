// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteImageResponseBody
	GetCode() *string
	SetData(v bool) *DeleteImageResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteImageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteImageResponseBody
	GetSuccess() *bool
}

type DeleteImageResponseBody struct {
	// Status code:
	//
	// 	- `200`: Success.
	//
	// 	- `400`: An error occurred.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Whether the deletion is successful.
	//
	// Enumerated values:
	//
	// 	- true:Deleted successfully.
	//
	// 	- false:Deletion failure.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The error message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteImageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteImageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteImageResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteImageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteImageResponseBody) SetCode(v string) *DeleteImageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteImageResponseBody) SetData(v bool) *DeleteImageResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteImageResponseBody) SetMessage(v string) *DeleteImageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteImageResponseBody) SetRequestId(v string) *DeleteImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteImageResponseBody) SetSuccess(v bool) *DeleteImageResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteImageResponseBody) Validate() error {
	return dara.Validate(s)
}
