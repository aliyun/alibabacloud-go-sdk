// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelContactByGenieDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *GetHotelContactByGenieDeviceResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelContactByGenieDeviceResponseBody
	GetRequestId() *string
	SetResult(v *GetHotelContactByGenieDeviceResponseBodyResult) *GetHotelContactByGenieDeviceResponseBody
	GetResult() *GetHotelContactByGenieDeviceResponseBodyResult
	SetStatusCode(v int32) *GetHotelContactByGenieDeviceResponseBody
	GetStatusCode() *int32
}

type GetHotelContactByGenieDeviceResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73C6***E6FA
	RequestId *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetHotelContactByGenieDeviceResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s GetHotelContactByGenieDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByGenieDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelContactByGenieDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelContactByGenieDeviceResponseBody) GetResult() *GetHotelContactByGenieDeviceResponseBodyResult {
	return s.Result
}

func (s *GetHotelContactByGenieDeviceResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHotelContactByGenieDeviceResponseBody) SetMessage(v string) *GetHotelContactByGenieDeviceResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBody) SetRequestId(v string) *GetHotelContactByGenieDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBody) SetResult(v *GetHotelContactByGenieDeviceResponseBodyResult) *GetHotelContactByGenieDeviceResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBody) SetStatusCode(v int32) *GetHotelContactByGenieDeviceResponseBody {
	s.StatusCode = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHotelContactByGenieDeviceResponseBodyResult struct {
	// example:
	//
	// 1649472283046
	ExpireAt *string `json:"ExpireAt,omitempty" xml:"ExpireAt,omitempty"`
	// example:
	//
	// 2022-07-21 20:02:12
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// 2022-07-27 14:06:27
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// xxx.icon
	Icon *string `json:"Icon,omitempty" xml:"Icon,omitempty"`
	// example:
	//
	// 1
	Id   *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
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

func (s GetHotelContactByGenieDeviceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetHotelContactByGenieDeviceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) GetExpireAt() *string {
	return s.ExpireAt
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) GetIcon() *string {
	return s.Icon
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) GetNumber() *string {
	return s.Number
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) GetStatus() *int32 {
	return s.Status
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) GetType() *string {
	return s.Type
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) GetUuid() *string {
	return s.Uuid
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetExpireAt(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.ExpireAt = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetGmtCreate(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetGmtModified(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.GmtModified = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetHotelId(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetIcon(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Icon = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetId(v int64) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Id = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetName(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetNumber(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Number = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetStatus(v int32) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Status = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetType(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Type = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) SetUuid(v string) *GetHotelContactByGenieDeviceResponseBodyResult {
	s.Uuid = &v
	return s
}

func (s *GetHotelContactByGenieDeviceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
