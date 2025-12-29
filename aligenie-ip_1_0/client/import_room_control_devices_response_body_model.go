// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportRoomControlDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *ImportRoomControlDevicesResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *ImportRoomControlDevicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportRoomControlDevicesResponseBody
	GetRequestId() *string
	SetResult(v int32) *ImportRoomControlDevicesResponseBody
	GetResult() *int32
	SetStatusCode(v int32) *ImportRoomControlDevicesResponseBody
	GetStatusCode() *int32
}

type ImportRoomControlDevicesResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// fdsfregtre
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s ImportRoomControlDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportRoomControlDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *ImportRoomControlDevicesResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *ImportRoomControlDevicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportRoomControlDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportRoomControlDevicesResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *ImportRoomControlDevicesResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportRoomControlDevicesResponseBody) SetExtentions(v map[string]interface{}) *ImportRoomControlDevicesResponseBody {
	s.Extentions = v
	return s
}

func (s *ImportRoomControlDevicesResponseBody) SetMessage(v string) *ImportRoomControlDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *ImportRoomControlDevicesResponseBody) SetRequestId(v string) *ImportRoomControlDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportRoomControlDevicesResponseBody) SetResult(v int32) *ImportRoomControlDevicesResponseBody {
	s.Result = &v
	return s
}

func (s *ImportRoomControlDevicesResponseBody) SetStatusCode(v int32) *ImportRoomControlDevicesResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ImportRoomControlDevicesResponseBody) Validate() error {
	return dara.Validate(s)
}
