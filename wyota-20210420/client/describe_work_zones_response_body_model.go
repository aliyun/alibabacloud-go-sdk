// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWorkZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeWorkZonesResponseBody
	GetCode() *string
	SetData(v *DescribeWorkZonesResponseBodyData) *DescribeWorkZonesResponseBody
	GetData() *DescribeWorkZonesResponseBodyData
	SetMessage(v string) *DescribeWorkZonesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeWorkZonesResponseBody
	GetRequestId() *string
}

type DescribeWorkZonesResponseBody struct {
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *DescribeWorkZonesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWorkZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWorkZonesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeWorkZonesResponseBody) GetData() *DescribeWorkZonesResponseBodyData {
	return s.Data
}

func (s *DescribeWorkZonesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeWorkZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWorkZonesResponseBody) SetCode(v string) *DescribeWorkZonesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeWorkZonesResponseBody) SetData(v *DescribeWorkZonesResponseBodyData) *DescribeWorkZonesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWorkZonesResponseBody) SetMessage(v string) *DescribeWorkZonesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWorkZonesResponseBody) SetRequestId(v string) *DescribeWorkZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWorkZonesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkZonesResponseBodyData struct {
	TotalCount      *int64                                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	WorkZoneDTOList []*DescribeWorkZonesResponseBodyDataWorkZoneDTOList `json:"WorkZoneDTOList,omitempty" xml:"WorkZoneDTOList,omitempty" type:"Repeated"`
}

func (s DescribeWorkZonesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkZonesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWorkZonesResponseBodyData) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeWorkZonesResponseBodyData) GetWorkZoneDTOList() []*DescribeWorkZonesResponseBodyDataWorkZoneDTOList {
	return s.WorkZoneDTOList
}

func (s *DescribeWorkZonesResponseBodyData) SetTotalCount(v int64) *DescribeWorkZonesResponseBodyData {
	s.TotalCount = &v
	return s
}

func (s *DescribeWorkZonesResponseBodyData) SetWorkZoneDTOList(v []*DescribeWorkZonesResponseBodyDataWorkZoneDTOList) *DescribeWorkZonesResponseBodyData {
	s.WorkZoneDTOList = v
	return s
}

func (s *DescribeWorkZonesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeWorkZonesResponseBodyDataWorkZoneDTOList struct {
	SeatCol  *int32  `json:"SeatCol,omitempty" xml:"SeatCol,omitempty"`
	SeatRow  *int32  `json:"SeatRow,omitempty" xml:"SeatRow,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	ZoneId   *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
	ZoneName *string `json:"ZoneName,omitempty" xml:"ZoneName,omitempty"`
}

func (s DescribeWorkZonesResponseBodyDataWorkZoneDTOList) String() string {
	return dara.Prettify(s)
}

func (s DescribeWorkZonesResponseBodyDataWorkZoneDTOList) GoString() string {
	return s.String()
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) GetSeatCol() *int32 {
	return s.SeatCol
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) GetSeatRow() *int32 {
	return s.SeatRow
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) GetTenantId() *string {
	return s.TenantId
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) GetZoneName() *string {
	return s.ZoneName
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) SetSeatCol(v int32) *DescribeWorkZonesResponseBodyDataWorkZoneDTOList {
	s.SeatCol = &v
	return s
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) SetSeatRow(v int32) *DescribeWorkZonesResponseBodyDataWorkZoneDTOList {
	s.SeatRow = &v
	return s
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) SetTenantId(v string) *DescribeWorkZonesResponseBodyDataWorkZoneDTOList {
	s.TenantId = &v
	return s
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) SetZoneId(v string) *DescribeWorkZonesResponseBodyDataWorkZoneDTOList {
	s.ZoneId = &v
	return s
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) SetZoneName(v string) *DescribeWorkZonesResponseBodyDataWorkZoneDTOList {
	s.ZoneName = &v
	return s
}

func (s *DescribeWorkZonesResponseBodyDataWorkZoneDTOList) Validate() error {
	return dara.Validate(s)
}
