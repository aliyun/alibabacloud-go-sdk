// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotelControlDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListHotelControlDeviceResponseBody
	GetCode() *int32
	SetMessage(v string) *ListHotelControlDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListHotelControlDeviceResponseBody
	GetRequestId() *string
	SetResult(v []map[string]*string) *ListHotelControlDeviceResponseBody
	GetResult() []map[string]*string
}

type ListHotelControlDeviceResponseBody struct {
	Code      *int32               `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []map[string]*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListHotelControlDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListHotelControlDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ListHotelControlDeviceResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListHotelControlDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListHotelControlDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListHotelControlDeviceResponseBody) GetResult() []map[string]*string {
	return s.Result
}

func (s *ListHotelControlDeviceResponseBody) SetCode(v int32) *ListHotelControlDeviceResponseBody {
	s.Code = &v
	return s
}

func (s *ListHotelControlDeviceResponseBody) SetMessage(v string) *ListHotelControlDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *ListHotelControlDeviceResponseBody) SetRequestId(v string) *ListHotelControlDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListHotelControlDeviceResponseBody) SetResult(v []map[string]*string) *ListHotelControlDeviceResponseBody {
	s.Result = v
	return s
}

func (s *ListHotelControlDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
