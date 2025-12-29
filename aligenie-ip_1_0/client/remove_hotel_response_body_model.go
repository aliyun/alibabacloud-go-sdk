// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveHotelResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *RemoveHotelResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *RemoveHotelResponseBody
	GetMessage() *string
	SetRequestId(v string) *RemoveHotelResponseBody
	GetRequestId() *string
	SetResult(v bool) *RemoveHotelResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *RemoveHotelResponseBody
	GetStatusCode() *int32
}

type RemoveHotelResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 73C67BD9-175A-1324-8202-9FAABBB3E6FA
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

func (s RemoveHotelResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveHotelResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveHotelResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *RemoveHotelResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RemoveHotelResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveHotelResponseBody) GetResult() *bool {
	return s.Result
}

func (s *RemoveHotelResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveHotelResponseBody) SetExtentions(v map[string]interface{}) *RemoveHotelResponseBody {
	s.Extentions = v
	return s
}

func (s *RemoveHotelResponseBody) SetMessage(v string) *RemoveHotelResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveHotelResponseBody) SetRequestId(v string) *RemoveHotelResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveHotelResponseBody) SetResult(v bool) *RemoveHotelResponseBody {
	s.Result = &v
	return s
}

func (s *RemoveHotelResponseBody) SetStatusCode(v int32) *RemoveHotelResponseBody {
	s.StatusCode = &v
	return s
}

func (s *RemoveHotelResponseBody) Validate() error {
	return dara.Validate(s)
}
