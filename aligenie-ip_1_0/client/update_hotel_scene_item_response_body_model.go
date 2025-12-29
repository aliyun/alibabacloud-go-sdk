// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelSceneItemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *UpdateHotelSceneItemResponseBody
	GetCode() *int32
	SetMessage(v string) *UpdateHotelSceneItemResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHotelSceneItemResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateHotelSceneItemResponseBody
	GetResult() *bool
}

type UpdateHotelSceneItemResponseBody struct {
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
	// 0D0C***67DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateHotelSceneItemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelSceneItemResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHotelSceneItemResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *UpdateHotelSceneItemResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHotelSceneItemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHotelSceneItemResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateHotelSceneItemResponseBody) SetCode(v int32) *UpdateHotelSceneItemResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateHotelSceneItemResponseBody) SetMessage(v string) *UpdateHotelSceneItemResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHotelSceneItemResponseBody) SetRequestId(v string) *UpdateHotelSceneItemResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHotelSceneItemResponseBody) SetResult(v bool) *UpdateHotelSceneItemResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateHotelSceneItemResponseBody) Validate() error {
	return dara.Validate(s)
}
