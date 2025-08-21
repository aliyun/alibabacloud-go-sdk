// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssUsageDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetOssUsageDataResponseBody
	GetRequestId() *string
	SetUsageList(v []*GetOssUsageDataResponseBodyUsageList) *GetOssUsageDataResponseBody
	GetUsageList() []*GetOssUsageDataResponseBodyUsageList
}

type GetOssUsageDataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 2A8CCD48-14F9-0309-B957-7B1D74A8119D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of usage data.
	UsageList []*GetOssUsageDataResponseBodyUsageList `json:"UsageList,omitempty" xml:"UsageList,omitempty" type:"Repeated"`
}

func (s GetOssUsageDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssUsageDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssUsageDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssUsageDataResponseBody) GetUsageList() []*GetOssUsageDataResponseBodyUsageList {
	return s.UsageList
}

func (s *GetOssUsageDataResponseBody) SetRequestId(v string) *GetOssUsageDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssUsageDataResponseBody) SetUsageList(v []*GetOssUsageDataResponseBodyUsageList) *GetOssUsageDataResponseBody {
	s.UsageList = v
	return s
}

func (s *GetOssUsageDataResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetOssUsageDataResponseBodyUsageList struct {
	// The inbound bandwidth over the internal network. Unit: bit/s.
	//
	// example:
	//
	// 37865147
	LanRxBw *int64 `json:"LanRxBw,omitempty" xml:"LanRxBw,omitempty"`
	// The outbound bandwidth over the internal network. Unit: bit/s.
	//
	// example:
	//
	// 22012187890
	LanTxBw *int64 `json:"LanTxBw,omitempty" xml:"LanTxBw,omitempty"`
	// The number of time points within a day.
	//
	// example:
	//
	// 144
	Point *int64 `json:"Point,omitempty" xml:"Point,omitempty"`
	// The point in time, in UTC. Format: 2010-01-21T09:50:23Z.
	//
	// example:
	//
	// 2022-01-12T00:00:00Z
	PointTs *string `json:"PointTs,omitempty" xml:"PointTs,omitempty"`
	// The storage usage. Unit: bytes.
	//
	// example:
	//
	// 85462146217
	StorageUsageByte *int64 `json:"StorageUsageByte,omitempty" xml:"StorageUsageByte,omitempty"`
	// The outbound bandwidth over the Internet. Unit: bit/s.
	//
	// example:
	//
	// 542155715
	WanRxBw *int64 `json:"WanRxBw,omitempty" xml:"WanRxBw,omitempty"`
	// The outbound bandwidth over the Internet. Unit: bit/s.
	//
	// example:
	//
	// 547126175217
	WanTxBw *int64 `json:"WanTxBw,omitempty" xml:"WanTxBw,omitempty"`
}

func (s GetOssUsageDataResponseBodyUsageList) String() string {
	return dara.Prettify(s)
}

func (s GetOssUsageDataResponseBodyUsageList) GoString() string {
	return s.String()
}

func (s *GetOssUsageDataResponseBodyUsageList) GetLanRxBw() *int64 {
	return s.LanRxBw
}

func (s *GetOssUsageDataResponseBodyUsageList) GetLanTxBw() *int64 {
	return s.LanTxBw
}

func (s *GetOssUsageDataResponseBodyUsageList) GetPoint() *int64 {
	return s.Point
}

func (s *GetOssUsageDataResponseBodyUsageList) GetPointTs() *string {
	return s.PointTs
}

func (s *GetOssUsageDataResponseBodyUsageList) GetStorageUsageByte() *int64 {
	return s.StorageUsageByte
}

func (s *GetOssUsageDataResponseBodyUsageList) GetWanRxBw() *int64 {
	return s.WanRxBw
}

func (s *GetOssUsageDataResponseBodyUsageList) GetWanTxBw() *int64 {
	return s.WanTxBw
}

func (s *GetOssUsageDataResponseBodyUsageList) SetLanRxBw(v int64) *GetOssUsageDataResponseBodyUsageList {
	s.LanRxBw = &v
	return s
}

func (s *GetOssUsageDataResponseBodyUsageList) SetLanTxBw(v int64) *GetOssUsageDataResponseBodyUsageList {
	s.LanTxBw = &v
	return s
}

func (s *GetOssUsageDataResponseBodyUsageList) SetPoint(v int64) *GetOssUsageDataResponseBodyUsageList {
	s.Point = &v
	return s
}

func (s *GetOssUsageDataResponseBodyUsageList) SetPointTs(v string) *GetOssUsageDataResponseBodyUsageList {
	s.PointTs = &v
	return s
}

func (s *GetOssUsageDataResponseBodyUsageList) SetStorageUsageByte(v int64) *GetOssUsageDataResponseBodyUsageList {
	s.StorageUsageByte = &v
	return s
}

func (s *GetOssUsageDataResponseBodyUsageList) SetWanRxBw(v int64) *GetOssUsageDataResponseBodyUsageList {
	s.WanRxBw = &v
	return s
}

func (s *GetOssUsageDataResponseBodyUsageList) SetWanTxBw(v int64) *GetOssUsageDataResponseBodyUsageList {
	s.WanTxBw = &v
	return s
}

func (s *GetOssUsageDataResponseBodyUsageList) Validate() error {
	return dara.Validate(s)
}
