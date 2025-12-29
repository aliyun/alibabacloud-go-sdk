// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelSceneBookItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateHotelSceneBookItemResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateHotelSceneBookItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHotelSceneBookItemResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateHotelSceneBookItemResponseBody
	GetResult() *bool
}

type UpdateHotelSceneBookItemResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateHotelSceneBookItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneBookItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneBookItemResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateHotelSceneBookItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHotelSceneBookItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHotelSceneBookItemResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateHotelSceneBookItemResponseBody) SetCode(v int32) *UpdateHotelSceneBookItemResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHotelSceneBookItemResponseBody) SetMessage(v string) *UpdateHotelSceneBookItemResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHotelSceneBookItemResponseBody) SetRequestId(v string) *UpdateHotelSceneBookItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHotelSceneBookItemResponseBody) SetResult(v bool) *UpdateHotelSceneBookItemResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateHotelSceneBookItemResponseBody) Validate() error {
	return dara.Validate(s)
}
