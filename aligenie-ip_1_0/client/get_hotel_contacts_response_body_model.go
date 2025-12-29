// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetHotelContactsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelContactsResponseBody
	GetRequestId() *string
	SetResult(v []*GetHotelContactsResponseBodyResult) *GetHotelContactsResponseBody
	GetResult() []*GetHotelContactsResponseBodyResult
	SetStatusCode(v int32) *GetHotelContactsResponseBody
	GetStatusCode() *int32
}

type GetHotelContactsResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 0EC7*726E
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetHotelContactsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetHotelContactsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactsResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelContactsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelContactsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelContactsResponseBody) GetResult() []*GetHotelContactsResponseBodyResult {
	return s.Result
}

func (s *GetHotelContactsResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelContactsResponseBody) SetMessage(v string) *GetHotelContactsResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelContactsResponseBody) SetRequestId(v string) *GetHotelContactsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelContactsResponseBody) SetResult(v []*GetHotelContactsResponseBodyResult) *GetHotelContactsResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelContactsResponseBody) SetStatusCode(v int32) *GetHotelContactsResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetHotelContactsResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHotelContactsResponseBodyResult struct {
	// example:
	//
	// 1649472283046
	ExpireAt *string `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	// example:
	//
	// cf24***96e7
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
	// 0862***A809
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetHotelContactsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelContactsResponseBodyResult) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *GetHotelContactsResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *GetHotelContactsResponseBodyResult) GetIcon() *string {
	return s.Icon
}

func (s *GetHotelContactsResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetHotelContactsResponseBodyResult) GetNumber() *string {
	return s.Number
}

func (s *GetHotelContactsResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *GetHotelContactsResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetHotelContactsResponseBodyResult) GetUuid() *string {
	return s.Uuid
}

func (s *GetHotelContactsResponseBodyResult) SetExpireAt(v string) *GetHotelContactsResponseBodyResult {
	s.ExpireAt = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetHotelId(v string) *GetHotelContactsResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetIcon(v string) *GetHotelContactsResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetName(v string) *GetHotelContactsResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetNumber(v string) *GetHotelContactsResponseBodyResult {
	s.Number = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetStatus(v int32) *GetHotelContactsResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetType(v string) *GetHotelContactsResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) SetUuid(v string) *GetHotelContactsResponseBodyResult {
	s.Uuid = &v
	return s
}

func (s *GetHotelContactsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
