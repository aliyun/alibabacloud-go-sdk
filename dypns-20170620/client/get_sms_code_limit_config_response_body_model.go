// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsCodeLimitConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetSmsCodeLimitConfigResponseBody
	GetCode() *string
	SetData(v *GetSmsCodeLimitConfigResponseBodyData) *GetSmsCodeLimitConfigResponseBody
	GetData() *GetSmsCodeLimitConfigResponseBodyData
	SetMessage(v string) *GetSmsCodeLimitConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetSmsCodeLimitConfigResponseBody
	GetRequestId() *string
}

type GetSmsCodeLimitConfigResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetSmsCodeLimitConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSmsCodeLimitConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSmsCodeLimitConfigResponseBody) GoString() string {
	return s.String()
}

func (s *GetSmsCodeLimitConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetSmsCodeLimitConfigResponseBody) GetData() *GetSmsCodeLimitConfigResponseBodyData {
	return s.Data
}

func (s *GetSmsCodeLimitConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetSmsCodeLimitConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSmsCodeLimitConfigResponseBody) SetCode(v string) *GetSmsCodeLimitConfigResponseBody {
	s.Code = &v
	return s
}

func (s *GetSmsCodeLimitConfigResponseBody) SetData(v *GetSmsCodeLimitConfigResponseBodyData) *GetSmsCodeLimitConfigResponseBody {
	s.Data = v
	return s
}

func (s *GetSmsCodeLimitConfigResponseBody) SetMessage(v string) *GetSmsCodeLimitConfigResponseBody {
	s.Message = &v
	return s
}

func (s *GetSmsCodeLimitConfigResponseBody) SetRequestId(v string) *GetSmsCodeLimitConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSmsCodeLimitConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetSmsCodeLimitConfigResponseBodyData struct {
	LimitDay    *int32  `json:"LimitDay,omitempty" xml:"LimitDay,omitempty"`
	LimitHour   *int32  `json:"LimitHour,omitempty" xml:"LimitHour,omitempty"`
	LimitId     *int64  `json:"LimitId,omitempty" xml:"LimitId,omitempty"`
	LimitMinute *int32  `json:"LimitMinute,omitempty" xml:"LimitMinute,omitempty"`
	LimitOther  *string `json:"LimitOther,omitempty" xml:"LimitOther,omitempty"`
}

func (s GetSmsCodeLimitConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetSmsCodeLimitConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetSmsCodeLimitConfigResponseBodyData) GetLimitDay() *int32 {
	return s.LimitDay
}

func (s *GetSmsCodeLimitConfigResponseBodyData) GetLimitHour() *int32 {
	return s.LimitHour
}

func (s *GetSmsCodeLimitConfigResponseBodyData) GetLimitId() *int64 {
	return s.LimitId
}

func (s *GetSmsCodeLimitConfigResponseBodyData) GetLimitMinute() *int32 {
	return s.LimitMinute
}

func (s *GetSmsCodeLimitConfigResponseBodyData) GetLimitOther() *string {
	return s.LimitOther
}

func (s *GetSmsCodeLimitConfigResponseBodyData) SetLimitDay(v int32) *GetSmsCodeLimitConfigResponseBodyData {
	s.LimitDay = &v
	return s
}

func (s *GetSmsCodeLimitConfigResponseBodyData) SetLimitHour(v int32) *GetSmsCodeLimitConfigResponseBodyData {
	s.LimitHour = &v
	return s
}

func (s *GetSmsCodeLimitConfigResponseBodyData) SetLimitId(v int64) *GetSmsCodeLimitConfigResponseBodyData {
	s.LimitId = &v
	return s
}

func (s *GetSmsCodeLimitConfigResponseBodyData) SetLimitMinute(v int32) *GetSmsCodeLimitConfigResponseBodyData {
	s.LimitMinute = &v
	return s
}

func (s *GetSmsCodeLimitConfigResponseBodyData) SetLimitOther(v string) *GetSmsCodeLimitConfigResponseBodyData {
	s.LimitOther = &v
	return s
}

func (s *GetSmsCodeLimitConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
