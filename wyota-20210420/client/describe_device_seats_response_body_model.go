// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeviceSeatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDeviceSeatsResponseBody
	GetCode() *string
	SetData(v []*DescribeDeviceSeatsResponseBodyData) *DescribeDeviceSeatsResponseBody
	GetData() []*DescribeDeviceSeatsResponseBodyData
	SetMessage(v string) *DescribeDeviceSeatsResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *DescribeDeviceSeatsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDeviceSeatsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDeviceSeatsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeDeviceSeatsResponseBody
	GetTotalCount() *int64
}

type DescribeDeviceSeatsResponseBody struct {
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       []*DescribeDeviceSeatsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	PageNumber *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int64                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDeviceSeatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceSeatsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDeviceSeatsResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDeviceSeatsResponseBody) GetData() []*DescribeDeviceSeatsResponseBodyData {
	return s.Data
}

func (s *DescribeDeviceSeatsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDeviceSeatsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDeviceSeatsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeviceSeatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDeviceSeatsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeDeviceSeatsResponseBody) SetCode(v string) *DescribeDeviceSeatsResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) SetData(v []*DescribeDeviceSeatsResponseBodyData) *DescribeDeviceSeatsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) SetMessage(v string) *DescribeDeviceSeatsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) SetPageNumber(v int32) *DescribeDeviceSeatsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) SetPageSize(v int32) *DescribeDeviceSeatsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) SetRequestId(v string) *DescribeDeviceSeatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) SetTotalCount(v int64) *DescribeDeviceSeatsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDeviceSeatsResponseBodyData struct {
	SeatCol  *int32  `json:"SeatCol,omitempty" xml:"SeatCol,omitempty"`
	SeatName *string `json:"SeatName,omitempty" xml:"SeatName,omitempty"`
	SeatNo   *string `json:"SeatNo,omitempty" xml:"SeatNo,omitempty"`
	SeatRow  *int32  `json:"SeatRow,omitempty" xml:"SeatRow,omitempty"`
	SerialNo *string `json:"SerialNo,omitempty" xml:"SerialNo,omitempty"`
	SiteId   *string `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SiteName *string `json:"SiteName,omitempty" xml:"SiteName,omitempty"`
}

func (s DescribeDeviceSeatsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeviceSeatsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDeviceSeatsResponseBodyData) GetSeatCol() *int32 {
	return s.SeatCol
}

func (s *DescribeDeviceSeatsResponseBodyData) GetSeatName() *string {
	return s.SeatName
}

func (s *DescribeDeviceSeatsResponseBodyData) GetSeatNo() *string {
	return s.SeatNo
}

func (s *DescribeDeviceSeatsResponseBodyData) GetSeatRow() *int32 {
	return s.SeatRow
}

func (s *DescribeDeviceSeatsResponseBodyData) GetSerialNo() *string {
	return s.SerialNo
}

func (s *DescribeDeviceSeatsResponseBodyData) GetSiteId() *string {
	return s.SiteId
}

func (s *DescribeDeviceSeatsResponseBodyData) GetSiteName() *string {
	return s.SiteName
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSeatCol(v int32) *DescribeDeviceSeatsResponseBodyData {
	s.SeatCol = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSeatName(v string) *DescribeDeviceSeatsResponseBodyData {
	s.SeatName = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSeatNo(v string) *DescribeDeviceSeatsResponseBodyData {
	s.SeatNo = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSeatRow(v int32) *DescribeDeviceSeatsResponseBodyData {
	s.SeatRow = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSerialNo(v string) *DescribeDeviceSeatsResponseBodyData {
	s.SerialNo = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSiteId(v string) *DescribeDeviceSeatsResponseBodyData {
	s.SiteId = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) SetSiteName(v string) *DescribeDeviceSeatsResponseBodyData {
	s.SiteName = &v
	return s
}

func (s *DescribeDeviceSeatsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
