// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelSampleUtterancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetHotelSampleUtterancesResponseBody
	GetCode() *int32
	SetMessage(v string) *GetHotelSampleUtterancesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelSampleUtterancesResponseBody
	GetRequestId() *string
	SetResult(v []*string) *GetHotelSampleUtterancesResponseBody
	GetResult() []*string
}

type GetHotelSampleUtterancesResponseBody struct {
	Code      *int32    `json:"Code,omitempty" xml:"Code,omitempty"`
	Message   *string   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*string `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetHotelSampleUtterancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelSampleUtterancesResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelSampleUtterancesResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetHotelSampleUtterancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelSampleUtterancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelSampleUtterancesResponseBody) GetResult() []*string {
	return s.Result
}

func (s *GetHotelSampleUtterancesResponseBody) SetCode(v int32) *GetHotelSampleUtterancesResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelSampleUtterancesResponseBody) SetMessage(v string) *GetHotelSampleUtterancesResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelSampleUtterancesResponseBody) SetRequestId(v string) *GetHotelSampleUtterancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelSampleUtterancesResponseBody) SetResult(v []*string) *GetHotelSampleUtterancesResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelSampleUtterancesResponseBody) Validate() error {
	return dara.Validate(s)
}
