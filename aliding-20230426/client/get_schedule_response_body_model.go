// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetScheduleResponseBody
	GetRequestId() *string
	SetScheduleInformation(v []*GetScheduleResponseBodyScheduleInformation) *GetScheduleResponseBody
	GetScheduleInformation() []*GetScheduleResponseBodyScheduleInformation
	SetVendorRequestId(v string) *GetScheduleResponseBody
	GetVendorRequestId() *string
	SetVendorType(v string) *GetScheduleResponseBody
	GetVendorType() *string
}

type GetScheduleResponseBody struct {
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	RequestId           *string                                       `json:"requestId,omitempty" xml:"requestId,omitempty"`
	ScheduleInformation []*GetScheduleResponseBodyScheduleInformation `json:"scheduleInformation,omitempty" xml:"scheduleInformation,omitempty" type:"Repeated"`
	// example:
	//
	// 0FAAEC9C-C6C8-5C87-AF8E-1195889BBXXX
	VendorRequestId *string `json:"vendorRequestId,omitempty" xml:"vendorRequestId,omitempty"`
	// example:
	//
	// dingtalk
	VendorType *string `json:"vendorType,omitempty" xml:"vendorType,omitempty"`
}

func (s GetScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *GetScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetScheduleResponseBody) GetScheduleInformation() []*GetScheduleResponseBodyScheduleInformation {
	return s.ScheduleInformation
}

func (s *GetScheduleResponseBody) GetVendorRequestId() *string {
	return s.VendorRequestId
}

func (s *GetScheduleResponseBody) GetVendorType() *string {
	return s.VendorType
}

func (s *GetScheduleResponseBody) SetRequestId(v string) *GetScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetScheduleResponseBody) SetScheduleInformation(v []*GetScheduleResponseBodyScheduleInformation) *GetScheduleResponseBody {
	s.ScheduleInformation = v
	return s
}

func (s *GetScheduleResponseBody) SetVendorRequestId(v string) *GetScheduleResponseBody {
	s.VendorRequestId = &v
	return s
}

func (s *GetScheduleResponseBody) SetVendorType(v string) *GetScheduleResponseBody {
	s.VendorType = &v
	return s
}

func (s *GetScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetScheduleResponseBodyScheduleInformation struct {
	// example:
	//
	// 无权限
	Error         *string                                                    `json:"Error,omitempty" xml:"Error,omitempty"`
	ScheduleItems []*GetScheduleResponseBodyScheduleInformationScheduleItems `json:"ScheduleItems,omitempty" xml:"ScheduleItems,omitempty" type:"Repeated"`
	// example:
	//
	// 012345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GetScheduleResponseBodyScheduleInformation) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleResponseBodyScheduleInformation) GoString() string {
	return s.String()
}

func (s *GetScheduleResponseBodyScheduleInformation) GetError() *string {
	return s.Error
}

func (s *GetScheduleResponseBodyScheduleInformation) GetScheduleItems() []*GetScheduleResponseBodyScheduleInformationScheduleItems {
	return s.ScheduleItems
}

func (s *GetScheduleResponseBodyScheduleInformation) GetUserId() *string {
	return s.UserId
}

func (s *GetScheduleResponseBodyScheduleInformation) SetError(v string) *GetScheduleResponseBodyScheduleInformation {
	s.Error = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformation) SetScheduleItems(v []*GetScheduleResponseBodyScheduleInformationScheduleItems) *GetScheduleResponseBodyScheduleInformation {
	s.ScheduleItems = v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformation) SetUserId(v string) *GetScheduleResponseBodyScheduleInformation {
	s.UserId = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformation) Validate() error {
	return dara.Validate(s)
}

type GetScheduleResponseBodyScheduleInformationScheduleItems struct {
	End   *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd   `json:"End,omitempty" xml:"End,omitempty" type:"Struct"`
	Start *GetScheduleResponseBodyScheduleInformationScheduleItemsStart `json:"Start,omitempty" xml:"Start,omitempty" type:"Struct"`
	// example:
	//
	// BUSY
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s GetScheduleResponseBodyScheduleInformationScheduleItems) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleResponseBodyScheduleInformationScheduleItems) GoString() string {
	return s.String()
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItems) GetEnd() *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd {
	return s.End
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItems) GetStart() *GetScheduleResponseBodyScheduleInformationScheduleItemsStart {
	return s.Start
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItems) GetStatus() *string {
	return s.Status
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItems) SetEnd(v *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) *GetScheduleResponseBodyScheduleInformationScheduleItems {
	s.End = v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItems) SetStart(v *GetScheduleResponseBodyScheduleInformationScheduleItemsStart) *GetScheduleResponseBodyScheduleInformationScheduleItems {
	s.Start = v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItems) SetStatus(v string) *GetScheduleResponseBodyScheduleInformationScheduleItems {
	s.Status = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItems) Validate() error {
	return dara.Validate(s)
}

type GetScheduleResponseBodyScheduleInformationScheduleItemsEnd struct {
	// example:
	//
	// 2020-01-01
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) GoString() string {
	return s.String()
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) GetDate() *string {
	return s.Date
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) GetDateTime() *string {
	return s.DateTime
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) GetTimeZone() *string {
	return s.TimeZone
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) SetDate(v string) *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd {
	s.Date = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) SetDateTime(v string) *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd {
	s.DateTime = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) SetTimeZone(v string) *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd {
	s.TimeZone = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsEnd) Validate() error {
	return dara.Validate(s)
}

type GetScheduleResponseBodyScheduleInformationScheduleItemsStart struct {
	// example:
	//
	// 2020-01-01
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// example:
	//
	// 2020-01-01T10:15:30+08:00
	DateTime *string `json:"DateTime,omitempty" xml:"DateTime,omitempty"`
	// example:
	//
	// Asia/Shanghai
	TimeZone *string `json:"TimeZone,omitempty" xml:"TimeZone,omitempty"`
}

func (s GetScheduleResponseBodyScheduleInformationScheduleItemsStart) String() string {
	return dara.Prettify(s)
}

func (s GetScheduleResponseBodyScheduleInformationScheduleItemsStart) GoString() string {
	return s.String()
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsStart) GetDate() *string {
	return s.Date
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsStart) GetDateTime() *string {
	return s.DateTime
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsStart) GetTimeZone() *string {
	return s.TimeZone
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsStart) SetDate(v string) *GetScheduleResponseBodyScheduleInformationScheduleItemsStart {
	s.Date = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsStart) SetDateTime(v string) *GetScheduleResponseBodyScheduleInformationScheduleItemsStart {
	s.DateTime = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsStart) SetTimeZone(v string) *GetScheduleResponseBodyScheduleInformationScheduleItemsStart {
	s.TimeZone = &v
	return s
}

func (s *GetScheduleResponseBodyScheduleInformationScheduleItemsStart) Validate() error {
	return dara.Validate(s)
}
