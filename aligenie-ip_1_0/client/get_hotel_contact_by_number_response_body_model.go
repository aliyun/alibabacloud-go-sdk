// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactByNumberResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetHotelContactByNumberResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelContactByNumberResponseBody
	GetRequestId() *string
	SetResult(v *GetHotelContactByNumberResponseBodyResult) *GetHotelContactByNumberResponseBody
	GetResult() *GetHotelContactByNumberResponseBodyResult
	SetStatusCode(v int32) *GetHotelContactByNumberResponseBody
	GetStatusCode() *int32
}

type GetHotelContactByNumberResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelContactByNumberResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetHotelContactByNumberResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByNumberResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelContactByNumberResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelContactByNumberResponseBody) GetResult() *GetHotelContactByNumberResponseBodyResult {
	return s.Result
}

func (s *GetHotelContactByNumberResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelContactByNumberResponseBody) SetMessage(v string) *GetHotelContactByNumberResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelContactByNumberResponseBody) SetRequestId(v string) *GetHotelContactByNumberResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelContactByNumberResponseBody) SetResult(v *GetHotelContactByNumberResponseBodyResult) *GetHotelContactByNumberResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelContactByNumberResponseBody) SetStatusCode(v int32) *GetHotelContactByNumberResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetHotelContactByNumberResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelContactByNumberResponseBodyResult struct {
	// example:
	//
	// 1649316479098
	ExpireAt *string `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// xxx.icon
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 101
	Number *string `json:"Number,omitempty" xml:"Number,omitempty"`
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// group
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 2E57***D45F9
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetHotelContactByNumberResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByNumberResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelContactByNumberResponseBodyResult) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *GetHotelContactByNumberResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *GetHotelContactByNumberResponseBodyResult) GetIcon() *string {
	return s.Icon
}

func (s *GetHotelContactByNumberResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetHotelContactByNumberResponseBodyResult) GetNumber() *string {
	return s.Number
}

func (s *GetHotelContactByNumberResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *GetHotelContactByNumberResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetHotelContactByNumberResponseBodyResult) GetUuid() *string {
	return s.Uuid
}

func (s *GetHotelContactByNumberResponseBodyResult) SetExpireAt(v string) *GetHotelContactByNumberResponseBodyResult {
	s.ExpireAt = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetHotelId(v string) *GetHotelContactByNumberResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetIcon(v string) *GetHotelContactByNumberResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetName(v string) *GetHotelContactByNumberResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetNumber(v string) *GetHotelContactByNumberResponseBodyResult {
	s.Number = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetStatus(v int32) *GetHotelContactByNumberResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetType(v string) *GetHotelContactByNumberResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) SetUuid(v string) *GetHotelContactByNumberResponseBodyResult {
	s.Uuid = &v
	return s
}

func (s *GetHotelContactByNumberResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
