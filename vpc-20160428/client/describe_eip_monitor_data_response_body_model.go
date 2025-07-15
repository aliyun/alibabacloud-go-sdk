// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipMonitorDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEipMonitorDatas(v *DescribeEipMonitorDataResponseBodyEipMonitorDatas) *DescribeEipMonitorDataResponseBody
	GetEipMonitorDatas() *DescribeEipMonitorDataResponseBodyEipMonitorDatas
	SetRequestId(v string) *DescribeEipMonitorDataResponseBody
	GetRequestId() *string
}

type DescribeEipMonitorDataResponseBody struct {
	// The detailed information about the monitoring data of the EIP.
	EipMonitorDatas *DescribeEipMonitorDataResponseBodyEipMonitorDatas `json:"EipMonitorDatas,omitempty" xml:"EipMonitorDatas,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C8B26B44-0189-443E-9816-D951F59623A9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeEipMonitorDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipMonitorDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEipMonitorDataResponseBody) GetEipMonitorDatas() *DescribeEipMonitorDataResponseBodyEipMonitorDatas {
	return s.EipMonitorDatas
}

func (s *DescribeEipMonitorDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEipMonitorDataResponseBody) SetEipMonitorDatas(v *DescribeEipMonitorDataResponseBodyEipMonitorDatas) *DescribeEipMonitorDataResponseBody {
	s.EipMonitorDatas = v
	return s
}

func (s *DescribeEipMonitorDataResponseBody) SetRequestId(v string) *DescribeEipMonitorDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeEipMonitorDataResponseBodyEipMonitorDatas struct {
	EipMonitorData []*DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData `json:"EipMonitorData,omitempty" xml:"EipMonitorData,omitempty" type:"Repeated"`
}

func (s DescribeEipMonitorDataResponseBodyEipMonitorDatas) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipMonitorDataResponseBodyEipMonitorDatas) GoString() string {
	return s.String()
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatas) GetEipMonitorData() []*DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	return s.EipMonitorData
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatas) SetEipMonitorData(v []*DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) *DescribeEipMonitorDataResponseBodyEipMonitorDatas {
	s.EipMonitorData = v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatas) Validate() error {
	return dara.Validate(s)
}

type DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData struct {
	// example:
	//
	// 10
	EipBandwidth *int32 `json:"EipBandwidth,omitempty" xml:"EipBandwidth,omitempty"`
	// The sum of inbound and outbound traffic.
	//
	// example:
	//
	// 465
	EipFlow *int64 `json:"EipFlow,omitempty" xml:"EipFlow,omitempty"`
	// The number of packets.
	//
	// example:
	//
	// 3434
	EipPackets *int32 `json:"EipPackets,omitempty" xml:"EipPackets,omitempty"`
	// The inbound traffic. Unit: bytes.
	//
	// example:
	//
	// 122
	EipRX *int64 `json:"EipRX,omitempty" xml:"EipRX,omitempty"`
	// The outbound traffic. Unit: bytes.
	//
	// example:
	//
	// 343
	EipTX *int64 `json:"EipTX,omitempty" xml:"EipTX,omitempty"`
	// The timestamp of the monitoring data. Specify the time in the ISO8601 standard. Example: `2020-01-21T09:50:23Z`.
	//
	// example:
	//
	// 2020-01-21T09:50:23Z
	TimeStamp *string `json:"TimeStamp,omitempty" xml:"TimeStamp,omitempty"`
}

func (s DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GoString() string {
	return s.String()
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipBandwidth() *int32 {
	return s.EipBandwidth
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipFlow() *int64 {
	return s.EipFlow
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipPackets() *int32 {
	return s.EipPackets
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipRX() *int64 {
	return s.EipRX
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetEipTX() *int64 {
	return s.EipTX
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) GetTimeStamp() *string {
	return s.TimeStamp
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipBandwidth(v int32) *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipBandwidth = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipFlow(v int64) *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipFlow = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipPackets(v int32) *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipPackets = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipRX(v int64) *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipRX = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetEipTX(v int64) *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.EipTX = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) SetTimeStamp(v string) *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData {
	s.TimeStamp = &v
	return s
}

func (s *DescribeEipMonitorDataResponseBodyEipMonitorDatasEipMonitorData) Validate() error {
	return dara.Validate(s)
}
