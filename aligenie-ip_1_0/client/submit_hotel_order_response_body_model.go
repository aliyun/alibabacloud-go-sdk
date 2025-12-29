// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotelOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *SubmitHotelOrderResponseBody
	GetCode() *int32
	SetMessage(v string) *SubmitHotelOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitHotelOrderResponseBody
	GetRequestId() *string
	SetResult(v string) *SubmitHotelOrderResponseBody
	GetResult() *string
	SetStatusCode(v int32) *SubmitHotelOrderResponseBody
	GetStatusCode() *int32
}

type SubmitHotelOrderResponseBody struct {
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
	// CCCF2E86-D9B5-12A6-AD25-8A06933D2B0F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20220809104752000114671478353329
	Result     *string `json:"Result,omitempty" xml:"Result,omitempty"`
	StatusCode *int32  `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s SubmitHotelOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitHotelOrderResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *SubmitHotelOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitHotelOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitHotelOrderResponseBody) GetResult() *string {
	return s.Result
}

func (s *SubmitHotelOrderResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitHotelOrderResponseBody) SetCode(v int32) *SubmitHotelOrderResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitHotelOrderResponseBody) SetMessage(v string) *SubmitHotelOrderResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitHotelOrderResponseBody) SetRequestId(v string) *SubmitHotelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitHotelOrderResponseBody) SetResult(v string) *SubmitHotelOrderResponseBody {
	s.Result = &v
	return s
}

func (s *SubmitHotelOrderResponseBody) SetStatusCode(v int32) *SubmitHotelOrderResponseBody {
	s.StatusCode = &v
	return s
}

func (s *SubmitHotelOrderResponseBody) Validate() error {
	return dara.Validate(s)
}
