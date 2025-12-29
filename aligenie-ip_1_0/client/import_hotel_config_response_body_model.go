// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportHotelConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *ImportHotelConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *ImportHotelConfigResponseBody
	GetRequestId() *string
	SetResult(v bool) *ImportHotelConfigResponseBody
	GetResult() *bool
	SetStatusCode(v int32) *ImportHotelConfigResponseBody
	GetStatusCode() *int32
}

type ImportHotelConfigResponseBody struct {
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

func (s ImportHotelConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ImportHotelConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ImportHotelConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ImportHotelConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportHotelConfigResponseBody) GetResult() *bool {
	return s.Result
}

func (s *ImportHotelConfigResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportHotelConfigResponseBody) SetMessage(v string) *ImportHotelConfigResponseBody {
	s.Message = &v
	return s
}

func (s *ImportHotelConfigResponseBody) SetRequestId(v string) *ImportHotelConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ImportHotelConfigResponseBody) SetResult(v bool) *ImportHotelConfigResponseBody {
	s.Result = &v
	return s
}

func (s *ImportHotelConfigResponseBody) SetStatusCode(v int32) *ImportHotelConfigResponseBody {
	s.StatusCode = &v
	return s
}

func (s *ImportHotelConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
