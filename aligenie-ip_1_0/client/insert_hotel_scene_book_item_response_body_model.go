// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertHotelSceneBookItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *InsertHotelSceneBookItemResponseBody
	GetCode() *int32
	SetMessage(v string) *InsertHotelSceneBookItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *InsertHotelSceneBookItemResponseBody
	GetRequestId() *string
	SetResult(v bool) *InsertHotelSceneBookItemResponseBody
	GetResult() *bool
}

type InsertHotelSceneBookItemResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// RequestId
	//
	// example:
	//
	// 36FB***80C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s InsertHotelSceneBookItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InsertHotelSceneBookItemResponseBody) GoString() string {
	return s.String()
}

func (s *InsertHotelSceneBookItemResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *InsertHotelSceneBookItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *InsertHotelSceneBookItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InsertHotelSceneBookItemResponseBody) GetResult() *bool {
	return s.Result
}

func (s *InsertHotelSceneBookItemResponseBody) SetCode(v int32) *InsertHotelSceneBookItemResponseBody {
	s.Code = &v
	return s
}

func (s *InsertHotelSceneBookItemResponseBody) SetMessage(v string) *InsertHotelSceneBookItemResponseBody {
	s.Message = &v
	return s
}

func (s *InsertHotelSceneBookItemResponseBody) SetRequestId(v string) *InsertHotelSceneBookItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *InsertHotelSceneBookItemResponseBody) SetResult(v bool) *InsertHotelSceneBookItemResponseBody {
	s.Result = &v
	return s
}

func (s *InsertHotelSceneBookItemResponseBody) Validate() error {
	return dara.Validate(s)
}
