// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageGetHotelRoomDevicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtentions(v map[string]interface{}) *PageGetHotelRoomDevicesResponseBody
	GetExtentions() map[string]interface{}
	SetMessage(v string) *PageGetHotelRoomDevicesResponseBody
	GetMessage() *string
	SetPage(v *PageGetHotelRoomDevicesResponseBodyPage) *PageGetHotelRoomDevicesResponseBody
	GetPage() *PageGetHotelRoomDevicesResponseBodyPage
	SetRequestId(v string) *PageGetHotelRoomDevicesResponseBody
	GetRequestId() *string
	SetResult(v []*PageGetHotelRoomDevicesResponseBodyResult) *PageGetHotelRoomDevicesResponseBody
	GetResult() []*PageGetHotelRoomDevicesResponseBodyResult
	SetStatusCode(v int32) *PageGetHotelRoomDevicesResponseBody
	GetStatusCode() *int32
}

type PageGetHotelRoomDevicesResponseBody struct {
	Extentions map[string]interface{} `json:"Extentions,omitempty" xml:"Extentions,omitempty"`
	// example:
	//
	// success
	Message *string                                  `json:"Message,omitempty" xml:"Message,omitempty"`
	Page    *PageGetHotelRoomDevicesResponseBodyPage `json:"Page,omitempty" xml:"Page,omitempty" type:"Struct"`
	// example:
	//
	// 4EFBDDF4-B19D-526C-8C3D-CD8AB51974EE
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*PageGetHotelRoomDevicesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	StatusCode *int32 `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
}

func (s PageGetHotelRoomDevicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PageGetHotelRoomDevicesResponseBody) GoString() string {
	return s.String()
}

func (s *PageGetHotelRoomDevicesResponseBody) GetExtentions() map[string]interface{} {
	return s.Extentions
}

func (s *PageGetHotelRoomDevicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PageGetHotelRoomDevicesResponseBody) GetPage() *PageGetHotelRoomDevicesResponseBodyPage {
	return s.Page
}

func (s *PageGetHotelRoomDevicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PageGetHotelRoomDevicesResponseBody) GetResult() []*PageGetHotelRoomDevicesResponseBodyResult {
	return s.Result
}

func (s *PageGetHotelRoomDevicesResponseBody) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PageGetHotelRoomDevicesResponseBody) SetExtentions(v map[string]interface{}) *PageGetHotelRoomDevicesResponseBody {
	s.Extentions = v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBody) SetMessage(v string) *PageGetHotelRoomDevicesResponseBody {
	s.Message = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBody) SetPage(v *PageGetHotelRoomDevicesResponseBodyPage) *PageGetHotelRoomDevicesResponseBody {
	s.Page = v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBody) SetRequestId(v string) *PageGetHotelRoomDevicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBody) SetResult(v []*PageGetHotelRoomDevicesResponseBodyResult) *PageGetHotelRoomDevicesResponseBody {
	s.Result = v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBody) SetStatusCode(v int32) *PageGetHotelRoomDevicesResponseBody {
	s.StatusCode = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBody) Validate() error {
	if s.Page != nil {
		if err := s.Page.Validate(); err != nil {
			return err
		}
	}
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

type PageGetHotelRoomDevicesResponseBodyPage struct {
	// example:
	//
	// 4
	End *int32 `json:"End,omitempty" xml:"End,omitempty"`
	// example:
	//
	// False
	HasNext *bool `json:"HasNext,omitempty" xml:"HasNext,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 0
	Start *int32 `json:"Start,omitempty" xml:"Start,omitempty"`
	// example:
	//
	// 5
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
	// example:
	//
	// 1
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s PageGetHotelRoomDevicesResponseBodyPage) String() string {
	return dara.Prettify(s)
}

func (s PageGetHotelRoomDevicesResponseBodyPage) GoString() string {
	return s.String()
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) GetEnd() *int32 {
	return s.End
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) GetHasNext() *bool {
	return s.HasNext
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) GetPageSize() *int32 {
	return s.PageSize
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) GetStart() *int32 {
	return s.Start
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) GetTotal() *int32 {
	return s.Total
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetEnd(v int32) *PageGetHotelRoomDevicesResponseBodyPage {
	s.End = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetHasNext(v bool) *PageGetHotelRoomDevicesResponseBodyPage {
	s.HasNext = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetPageNumber(v int32) *PageGetHotelRoomDevicesResponseBodyPage {
	s.PageNumber = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetPageSize(v int32) *PageGetHotelRoomDevicesResponseBodyPage {
	s.PageSize = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetStart(v int32) *PageGetHotelRoomDevicesResponseBodyPage {
	s.Start = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetTotal(v int32) *PageGetHotelRoomDevicesResponseBodyPage {
	s.Total = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) SetTotalPage(v int32) *PageGetHotelRoomDevicesResponseBodyPage {
	s.TotalPage = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyPage) Validate() error {
	return dara.Validate(s)
}

type PageGetHotelRoomDevicesResponseBodyResult struct {
	// example:
	//
	// V21.10.00.313
	FirmwareVersion *string `json:"FirmwareVersion,omitempty" xml:"FirmwareVersion,omitempty"`
	// example:
	//
	// a7***83
	HotelId *string `json:"HotelId,omitempty" xml:"HotelId,omitempty"`
	// example:
	//
	// b4:xx:xx:xx:65:2b
	Mac *string `json:"Mac,omitempty" xml:"Mac,omitempty"`
	// example:
	//
	// 1
	OnlineStatus *int32 `json:"OnlineStatus,omitempty" xml:"OnlineStatus,omitempty"`
	// example:
	//
	// 2001
	RoomNo *string `json:"RoomNo,omitempty" xml:"RoomNo,omitempty"`
	// example:
	//
	// 1200xxx048
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
}

func (s PageGetHotelRoomDevicesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s PageGetHotelRoomDevicesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) GetFirmwareVersion() *string {
	return s.FirmwareVersion
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) GetHotelId() *string {
	return s.HotelId
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) GetMac() *string {
	return s.Mac
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) GetOnlineStatus() *int32 {
	return s.OnlineStatus
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) GetRoomNo() *string {
	return s.RoomNo
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) GetSn() *string {
	return s.Sn
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) SetFirmwareVersion(v string) *PageGetHotelRoomDevicesResponseBodyResult {
	s.FirmwareVersion = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) SetHotelId(v string) *PageGetHotelRoomDevicesResponseBodyResult {
	s.HotelId = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) SetMac(v string) *PageGetHotelRoomDevicesResponseBodyResult {
	s.Mac = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) SetOnlineStatus(v int32) *PageGetHotelRoomDevicesResponseBodyResult {
	s.OnlineStatus = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) SetRoomNo(v string) *PageGetHotelRoomDevicesResponseBodyResult {
	s.RoomNo = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) SetSn(v string) *PageGetHotelRoomDevicesResponseBodyResult {
	s.Sn = &v
	return s
}

func (s *PageGetHotelRoomDevicesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
