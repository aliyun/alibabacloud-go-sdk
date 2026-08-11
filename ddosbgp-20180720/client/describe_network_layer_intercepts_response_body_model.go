// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkLayerInterceptsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInterceptionRecordCount(v int64) *DescribeNetworkLayerInterceptsResponseBody
	GetInterceptionRecordCount() *int64
	SetInterceptionRecords(v []*DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) *DescribeNetworkLayerInterceptsResponseBody
	GetInterceptionRecords() []*DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords
	SetRequestId(v string) *DescribeNetworkLayerInterceptsResponseBody
	GetRequestId() *string
	SetTotalCnt(v string) *DescribeNetworkLayerInterceptsResponseBody
	GetTotalCnt() *string
}

type DescribeNetworkLayerInterceptsResponseBody struct {
	// The number of interception log records.
	//
	// example:
	//
	// 10
	InterceptionRecordCount *int64 `json:"InterceptionRecordCount,omitempty" xml:"InterceptionRecordCount,omitempty"`
	// The interception record details.
	InterceptionRecords []*DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords `json:"InterceptionRecords,omitempty" xml:"InterceptionRecords,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// B4B379C2-9319-4C6B-B579-FE36831B09F4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of interception logs that match the current filter conditions.
	//
	// example:
	//
	// 17
	TotalCnt *string `json:"TotalCnt,omitempty" xml:"TotalCnt,omitempty"`
}

func (s DescribeNetworkLayerInterceptsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkLayerInterceptsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetworkLayerInterceptsResponseBody) GetInterceptionRecordCount() *int64 {
	return s.InterceptionRecordCount
}

func (s *DescribeNetworkLayerInterceptsResponseBody) GetInterceptionRecords() []*DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords {
	return s.InterceptionRecords
}

func (s *DescribeNetworkLayerInterceptsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetworkLayerInterceptsResponseBody) GetTotalCnt() *string {
	return s.TotalCnt
}

func (s *DescribeNetworkLayerInterceptsResponseBody) SetInterceptionRecordCount(v int64) *DescribeNetworkLayerInterceptsResponseBody {
	s.InterceptionRecordCount = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBody) SetInterceptionRecords(v []*DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) *DescribeNetworkLayerInterceptsResponseBody {
	s.InterceptionRecords = v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBody) SetRequestId(v string) *DescribeNetworkLayerInterceptsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBody) SetTotalCnt(v string) *DescribeNetworkLayerInterceptsResponseBody {
	s.TotalCnt = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBody) Validate() error {
	if s.InterceptionRecords != nil {
		for _, item := range s.InterceptionRecords {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords struct {
	// The destination IP address.
	//
	// example:
	//
	// 47.254.56.252
	DestinationIp *string `json:"DestinationIp,omitempty" xml:"DestinationIp,omitempty"`
	// The destination port in the interception log.
	//
	// example:
	//
	// 22
	DestinationPort *string `json:"DestinationPort,omitempty" xml:"DestinationPort,omitempty"`
	// The interception action.
	//
	// example:
	//
	// drop
	InterceptAction *string `json:"InterceptAction,omitempty" xml:"InterceptAction,omitempty"`
	// The number of interceptions within the specified time range.
	//
	// example:
	//
	// 1
	InterceptCount *int64 `json:"InterceptCount,omitempty" xml:"InterceptCount,omitempty"`
	// The end time of the interception.
	//
	//  > The value is a Unix/POSIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1778830500
	InterceptEndTime *int64 `json:"InterceptEndTime,omitempty" xml:"InterceptEndTime,omitempty"`
	// The packet interception module.
	//
	// example:
	//
	// dip_blacklist
	InterceptModule *string `json:"InterceptModule,omitempty" xml:"InterceptModule,omitempty"`
	// The start time of the interception.
	//
	// > The value is a Unix/POSIX timestamp. Unit: seconds.
	//
	// example:
	//
	// 1778830500
	InterceptStartTime *int64 `json:"InterceptStartTime,omitempty" xml:"InterceptStartTime,omitempty"`
	// The network protocol.
	//
	// example:
	//
	// tcp
	NetworkProtocol *string `json:"NetworkProtocol,omitempty" xml:"NetworkProtocol,omitempty"`
	// The network protocol number. This is a standard network protocol number.
	//
	// example:
	//
	// 6
	ProtocolNumber *string `json:"ProtocolNumber,omitempty" xml:"ProtocolNumber,omitempty"`
	// The source IP address.
	//
	// example:
	//
	// 183.224.38.37
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The source port in the interception log.
	//
	// example:
	//
	// 9998
	SourcePort *string `json:"SourcePort,omitempty" xml:"SourcePort,omitempty"`
}

func (s DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) GoString() string {
	return s.String()
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) GetDestinationIp() *string {
	return s.DestinationIp
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) GetDestinationPort() *string {
	return s.DestinationPort
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) GetInterceptAction() *string {
	return s.InterceptAction
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) GetInterceptCount() *int64 {
	return s.InterceptCount
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) GetInterceptEndTime() *int64 {
	return s.InterceptEndTime
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) GetInterceptModule() *string {
	return s.InterceptModule
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) GetInterceptStartTime() *int64 {
	return s.InterceptStartTime
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) GetNetworkProtocol() *string {
	return s.NetworkProtocol
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) GetProtocolNumber() *string {
	return s.ProtocolNumber
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) GetSourcePort() *string {
	return s.SourcePort
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) SetDestinationIp(v string) *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords {
	s.DestinationIp = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) SetDestinationPort(v string) *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords {
	s.DestinationPort = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) SetInterceptAction(v string) *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords {
	s.InterceptAction = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) SetInterceptCount(v int64) *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords {
	s.InterceptCount = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) SetInterceptEndTime(v int64) *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords {
	s.InterceptEndTime = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) SetInterceptModule(v string) *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords {
	s.InterceptModule = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) SetInterceptStartTime(v int64) *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords {
	s.InterceptStartTime = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) SetNetworkProtocol(v string) *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords {
	s.NetworkProtocol = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) SetProtocolNumber(v string) *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords {
	s.ProtocolNumber = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) SetSourceIp(v string) *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords {
	s.SourceIp = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) SetSourcePort(v string) *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords {
	s.SourcePort = &v
	return s
}

func (s *DescribeNetworkLayerInterceptsResponseBodyInterceptionRecords) Validate() error {
	return dara.Validate(s)
}
