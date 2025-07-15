// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserConnectionRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConnectionRecords(v []*DescribeUserConnectionRecordsResponseBodyConnectionRecords) *DescribeUserConnectionRecordsResponseBody
	GetConnectionRecords() []*DescribeUserConnectionRecordsResponseBodyConnectionRecords
	SetNextToken(v string) *DescribeUserConnectionRecordsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeUserConnectionRecordsResponseBody
	GetRequestId() *string
}

type DescribeUserConnectionRecordsResponseBody struct {
	// The connection records.
	ConnectionRecords []*DescribeUserConnectionRecordsResponseBodyConnectionRecords `json:"ConnectionRecords,omitempty" xml:"ConnectionRecords,omitempty" type:"Repeated"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2CC66B0A-BA3B-5D87-BFBE-11AAAD7A8E03
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserConnectionRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConnectionRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsResponseBody) GetConnectionRecords() []*DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	return s.ConnectionRecords
}

func (s *DescribeUserConnectionRecordsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUserConnectionRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserConnectionRecordsResponseBody) SetConnectionRecords(v []*DescribeUserConnectionRecordsResponseBodyConnectionRecords) *DescribeUserConnectionRecordsResponseBody {
	s.ConnectionRecords = v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBody) SetNextToken(v string) *DescribeUserConnectionRecordsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBody) SetRequestId(v string) *DescribeUserConnectionRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeUserConnectionRecordsResponseBodyConnectionRecords struct {
	// The connection duration. Unit: milliseconds.
	//
	// example:
	//
	// 3405035000
	ConnectDuration *string `json:"ConnectDuration,omitempty" xml:"ConnectDuration,omitempty"`
	// The time when the end user disconnected from the cloud computer.
	//
	// example:
	//
	// 2022-02-27T20:03:13Z
	ConnectEndTime *string `json:"ConnectEndTime,omitempty" xml:"ConnectEndTime,omitempty"`
	// The time when the end user connected to the cloud computer.
	//
	// example:
	//
	// 2022-01-19T10:12:38Z
	ConnectStartTime *string `json:"ConnectStartTime,omitempty" xml:"ConnectStartTime,omitempty"`
	// The ID of the connection record.
	//
	// example:
	//
	// 528
	ConnectionRecordId *string `json:"ConnectionRecordId,omitempty" xml:"ConnectionRecordId,omitempty"`
	// The ID of the cloud computer to which the end user connected.
	//
	// example:
	//
	// ud-2hawufy3uedi1****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The name of the cloud computer to which the end user connected.
	//
	// example:
	//
	// testName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
}

func (s DescribeUserConnectionRecordsResponseBodyConnectionRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserConnectionRecordsResponseBodyConnectionRecords) GoString() string {
	return s.String()
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) GetConnectDuration() *string {
	return s.ConnectDuration
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) GetConnectEndTime() *string {
	return s.ConnectEndTime
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) GetConnectStartTime() *string {
	return s.ConnectStartTime
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) GetConnectionRecordId() *string {
	return s.ConnectionRecordId
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetConnectDuration(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.ConnectDuration = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetConnectEndTime(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.ConnectEndTime = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetConnectStartTime(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.ConnectStartTime = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetConnectionRecordId(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.ConnectionRecordId = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetDesktopId(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.DesktopId = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) SetDesktopName(v string) *DescribeUserConnectionRecordsResponseBodyConnectionRecords {
	s.DesktopName = &v
	return s
}

func (s *DescribeUserConnectionRecordsResponseBodyConnectionRecords) Validate() error {
	return dara.Validate(s)
}
