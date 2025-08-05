// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventBusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteEventBusResponseBody
	GetCode() *string
	SetData(v bool) *DeleteEventBusResponseBody
	GetData() *bool
	SetMessage(v string) *DeleteEventBusResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteEventBusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteEventBusResponseBody
	GetSuccess() *bool
}

type DeleteEventBusResponseBody struct {
	// The response code. The code 200 indicates that the request was successful. Other codes indicate that the request failed. For information about error codes, see Error codes.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *bool   `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned error message.
	//
	// example:
	//
	// EventBusNotExist
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C229E140-1A5C-5D55-8904-CFC5BA4CAA98
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteEventBusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventBusResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventBusResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteEventBusResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteEventBusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteEventBusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventBusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteEventBusResponseBody) SetCode(v string) *DeleteEventBusResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventBusResponseBody) SetData(v bool) *DeleteEventBusResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteEventBusResponseBody) SetMessage(v string) *DeleteEventBusResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventBusResponseBody) SetRequestId(v string) *DeleteEventBusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventBusResponseBody) SetSuccess(v bool) *DeleteEventBusResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteEventBusResponseBody) Validate() error {
	return dara.Validate(s)
}
