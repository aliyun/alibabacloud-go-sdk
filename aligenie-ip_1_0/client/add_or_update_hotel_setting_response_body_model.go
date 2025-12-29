// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrUpdateHotelSettingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *AddOrUpdateHotelSettingResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddOrUpdateHotelSettingResponseBody
	GetRequestId() *string
	SetResult(v bool) *AddOrUpdateHotelSettingResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *AddOrUpdateHotelSettingResponseBody
	GetStatusCode() *int32
}

type AddOrUpdateHotelSettingResponseBody struct {
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
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s AddOrUpdateHotelSettingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddOrUpdateHotelSettingResponseBody) GoString() string {
	return s.String()
}

func (s *AddOrUpdateHotelSettingResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddOrUpdateHotelSettingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddOrUpdateHotelSettingResponseBody) GetResult() *bool {
	return s.Result
}

func (s *AddOrUpdateHotelSettingResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOrUpdateHotelSettingResponseBody) SetMessage(v string) *AddOrUpdateHotelSettingResponseBody {
	s.Message = &v
	return s
}

func (s *AddOrUpdateHotelSettingResponseBody) SetRequestId(v string) *AddOrUpdateHotelSettingResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddOrUpdateHotelSettingResponseBody) SetResult(v bool) *AddOrUpdateHotelSettingResponseBody {
	s.Result = &v
	return s
}

func (s *AddOrUpdateHotelSettingResponseBody) SetStatusCode(v int32) *AddOrUpdateHotelSettingResponseBody {
	s.StatusCode = &v
	return s
}

func (s *AddOrUpdateHotelSettingResponseBody) Validate() error {
	return dara.Validate(s)
}
