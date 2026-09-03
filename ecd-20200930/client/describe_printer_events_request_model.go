// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrinterEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDesktopId(v string) *DescribePrinterEventsRequest
	GetDesktopId() *string
	SetDesktopName(v string) *DescribePrinterEventsRequest
	GetDesktopName() *string
	SetEndTime(v string) *DescribePrinterEventsRequest
	GetEndTime() *string
	SetEndUserId(v string) *DescribePrinterEventsRequest
	GetEndUserId() *string
	SetEndUserIds(v []*string) *DescribePrinterEventsRequest
	GetEndUserIds() []*string
	SetMaxResults(v int32) *DescribePrinterEventsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePrinterEventsRequest
	GetNextToken() *string
	SetPrinterDriver(v string) *DescribePrinterEventsRequest
	GetPrinterDriver() *string
	SetPrinterName(v string) *DescribePrinterEventsRequest
	GetPrinterName() *string
	SetPrinterRedirType(v int32) *DescribePrinterEventsRequest
	GetPrinterRedirType() *int32
	SetRegionId(v string) *DescribePrinterEventsRequest
	GetRegionId() *string
	SetStartTime(v string) *DescribePrinterEventsRequest
	GetStartTime() *string
}

type DescribePrinterEventsRequest struct {
	// The cloud computer ID. If you do not specify this parameter, all cloud computers in the region are queried.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The cloud computer name.
	//
	// example:
	//
	// desktop-001
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The end time. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC+0. If you do not specify this parameter, the current time is used.
	//
	// example:
	//
	// 2020-11-31T06:32:31Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The logon user information, which is a RAM user ID or an Active Directory (AD) username. If you do not specify this parameter, events of all users in the region are queried.
	//
	// example:
	//
	// user001
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The list of end user IDs.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The number of entries per page in a paged query. Default value: 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token. Set this parameter to the value of NextToken that was returned in the previous API call.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nmB7qrRFJ8vmttjxPL****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The printer driver name.
	//
	// example:
	//
	// HP LaserJet PCL 6
	PrinterDriver *string `json:"PrinterDriver,omitempty" xml:"PrinterDriver,omitempty"`
	// The printer name.
	//
	// example:
	//
	// HP LaserJet Pro
	PrinterName *string `json:"PrinterName,omitempty" xml:"PrinterName,omitempty"`
	// The printer redirection type.
	//
	// example:
	//
	// 1
	PrinterRedirType *int32 `json:"PrinterRedirType,omitempty" xml:"PrinterRedirType,omitempty"`
	// The region ID. You can call DescribeRegions to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The start time. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the YYYY-MM-DDThh:mm:ssZ format. The time must be in UTC+0. If you do not specify this parameter, the query starts from the time that is calculated backward from the time specified by `EndTime`.
	//
	// example:
	//
	// 2022-03-23T04:10:21Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribePrinterEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrinterEventsRequest) GoString() string {
	return s.String()
}

func (s *DescribePrinterEventsRequest) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribePrinterEventsRequest) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribePrinterEventsRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribePrinterEventsRequest) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribePrinterEventsRequest) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribePrinterEventsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePrinterEventsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePrinterEventsRequest) GetPrinterDriver() *string {
	return s.PrinterDriver
}

func (s *DescribePrinterEventsRequest) GetPrinterName() *string {
	return s.PrinterName
}

func (s *DescribePrinterEventsRequest) GetPrinterRedirType() *int32 {
	return s.PrinterRedirType
}

func (s *DescribePrinterEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePrinterEventsRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribePrinterEventsRequest) SetDesktopId(v string) *DescribePrinterEventsRequest {
	s.DesktopId = &v
	return s
}

func (s *DescribePrinterEventsRequest) SetDesktopName(v string) *DescribePrinterEventsRequest {
	s.DesktopName = &v
	return s
}

func (s *DescribePrinterEventsRequest) SetEndTime(v string) *DescribePrinterEventsRequest {
	s.EndTime = &v
	return s
}

func (s *DescribePrinterEventsRequest) SetEndUserId(v string) *DescribePrinterEventsRequest {
	s.EndUserId = &v
	return s
}

func (s *DescribePrinterEventsRequest) SetEndUserIds(v []*string) *DescribePrinterEventsRequest {
	s.EndUserIds = v
	return s
}

func (s *DescribePrinterEventsRequest) SetMaxResults(v int32) *DescribePrinterEventsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePrinterEventsRequest) SetNextToken(v string) *DescribePrinterEventsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePrinterEventsRequest) SetPrinterDriver(v string) *DescribePrinterEventsRequest {
	s.PrinterDriver = &v
	return s
}

func (s *DescribePrinterEventsRequest) SetPrinterName(v string) *DescribePrinterEventsRequest {
	s.PrinterName = &v
	return s
}

func (s *DescribePrinterEventsRequest) SetPrinterRedirType(v int32) *DescribePrinterEventsRequest {
	s.PrinterRedirType = &v
	return s
}

func (s *DescribePrinterEventsRequest) SetRegionId(v string) *DescribePrinterEventsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePrinterEventsRequest) SetStartTime(v string) *DescribePrinterEventsRequest {
	s.StartTime = &v
	return s
}

func (s *DescribePrinterEventsRequest) Validate() error {
	return dara.Validate(s)
}
