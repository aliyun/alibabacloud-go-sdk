// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHotelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *UpdateHotelResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *UpdateHotelResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateHotelResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateHotelResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *UpdateHotelResponseBody
	GetStatusCode() *int32
}

type UpdateHotelResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 46C53AEB-B19C-5C42-B32E-A726979C126F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s UpdateHotelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHotelResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHotelResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *UpdateHotelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateHotelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHotelResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateHotelResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHotelResponseBody) SetExtentions(v map[string]interface{}) *UpdateHotelResponseBody {
	s.Extentions = v
	return s
}

func (s *UpdateHotelResponseBody) SetMessage(v string) *UpdateHotelResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateHotelResponseBody) SetRequestId(v string) *UpdateHotelResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHotelResponseBody) SetResult(v bool) *UpdateHotelResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateHotelResponseBody) SetStatusCode(v int32) *UpdateHotelResponseBody {
	s.StatusCode = &v
	return s
}

func (s *UpdateHotelResponseBody) Validate() error {
	return dara.Validate(s)
}
