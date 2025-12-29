// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHotelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *CreateHotelResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *CreateHotelResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateHotelResponseBody
	GetRequestId() *string
	SetResult(v string) *CreateHotelResponseBody
	GetResult() *string
	SetStatusCode(v int32) *CreateHotelResponseBody
	GetStatusCode() *int32
}

type CreateHotelResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 5abfd9***2c38661
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s CreateHotelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHotelResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHotelResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *CreateHotelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateHotelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHotelResponseBody) GetResult() *string {
	return s.Result
}

func (s *CreateHotelResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHotelResponseBody) SetExtentions(v map[string]interface{}) *CreateHotelResponseBody {
	s.Extentions = v
	return s
}

func (s *CreateHotelResponseBody) SetMessage(v string) *CreateHotelResponseBody {
	s.Message = &v
	return s
}

func (s *CreateHotelResponseBody) SetRequestId(v string) *CreateHotelResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHotelResponseBody) SetResult(v string) *CreateHotelResponseBody {
	s.Result = &v
	return s
}

func (s *CreateHotelResponseBody) SetStatusCode(v int32) *CreateHotelResponseBody {
	s.StatusCode = &v
	return s
}

func (s *CreateHotelResponseBody) Validate() error {
	return dara.Validate(s)
}
