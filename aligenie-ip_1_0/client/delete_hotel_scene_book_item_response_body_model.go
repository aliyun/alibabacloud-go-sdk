// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHotelSceneBookItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DeleteHotelSceneBookItemResponseBody
	GetCode() *int32
	SetMessage(v string) *DeleteHotelSceneBookItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteHotelSceneBookItemResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteHotelSceneBookItemResponseBody
	GetResult() *bool
}

type DeleteHotelSceneBookItemResponseBody struct {
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

func (s DeleteHotelSceneBookItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHotelSceneBookItemResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHotelSceneBookItemResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DeleteHotelSceneBookItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteHotelSceneBookItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHotelSceneBookItemResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteHotelSceneBookItemResponseBody) SetCode(v int32) *DeleteHotelSceneBookItemResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteHotelSceneBookItemResponseBody) SetMessage(v string) *DeleteHotelSceneBookItemResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteHotelSceneBookItemResponseBody) SetRequestId(v string) *DeleteHotelSceneBookItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHotelSceneBookItemResponseBody) SetResult(v bool) *DeleteHotelSceneBookItemResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteHotelSceneBookItemResponseBody) Validate() error {
	return dara.Validate(s)
}
