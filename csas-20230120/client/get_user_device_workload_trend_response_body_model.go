// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserDeviceWorkloadTrendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetUserDeviceWorkloadTrendResponseBody
	GetRequestId() *string
	SetTitleEn(v string) *GetUserDeviceWorkloadTrendResponseBody
	GetTitleEn() *string
	SetTitleZh(v string) *GetUserDeviceWorkloadTrendResponseBody
	GetTitleZh() *string
	SetWorkloadList(v []*GetUserDeviceWorkloadTrendResponseBodyWorkloadList) *GetUserDeviceWorkloadTrendResponseBody
	GetWorkloadList() []*GetUserDeviceWorkloadTrendResponseBodyWorkloadList
}

type GetUserDeviceWorkloadTrendResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 102350E7-1A20-58F5-9D63-ABEA820AE6E1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The English name of the workload type. Valid values:
	//
	// - **CPU Usage**: returned when WorkloadType is set to cpu.
	//
	// - **Memory Usage**: returned when WorkloadType is set to mem.
	//
	// example:
	//
	// CPU Usage
	TitleEn *string `json:"TitleEn,omitempty" xml:"TitleEn,omitempty"`
	// The Chinese name of the workload type. Valid values:
	//
	// - **CPU使用率**: returned when WorkloadType is set to cpu.
	//
	// - **内存使用率**: returned when WorkloadType is set to mem.
	//
	// example:
	//
	// CPU 使用率
	TitleZh *string `json:"TitleZh,omitempty" xml:"TitleZh,omitempty"`
	// The list of workload trend data points, sorted by time in ascending order.
	WorkloadList []*GetUserDeviceWorkloadTrendResponseBodyWorkloadList `json:"WorkloadList,omitempty" xml:"WorkloadList,omitempty" type:"Repeated"`
}

func (s GetUserDeviceWorkloadTrendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeviceWorkloadTrendResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserDeviceWorkloadTrendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserDeviceWorkloadTrendResponseBody) GetTitleEn() *string {
	return s.TitleEn
}

func (s *GetUserDeviceWorkloadTrendResponseBody) GetTitleZh() *string {
	return s.TitleZh
}

func (s *GetUserDeviceWorkloadTrendResponseBody) GetWorkloadList() []*GetUserDeviceWorkloadTrendResponseBodyWorkloadList {
	return s.WorkloadList
}

func (s *GetUserDeviceWorkloadTrendResponseBody) SetRequestId(v string) *GetUserDeviceWorkloadTrendResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserDeviceWorkloadTrendResponseBody) SetTitleEn(v string) *GetUserDeviceWorkloadTrendResponseBody {
	s.TitleEn = &v
	return s
}

func (s *GetUserDeviceWorkloadTrendResponseBody) SetTitleZh(v string) *GetUserDeviceWorkloadTrendResponseBody {
	s.TitleZh = &v
	return s
}

func (s *GetUserDeviceWorkloadTrendResponseBody) SetWorkloadList(v []*GetUserDeviceWorkloadTrendResponseBodyWorkloadList) *GetUserDeviceWorkloadTrendResponseBody {
	s.WorkloadList = v
	return s
}

func (s *GetUserDeviceWorkloadTrendResponseBody) Validate() error {
	if s.WorkloadList != nil {
		for _, item := range s.WorkloadList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUserDeviceWorkloadTrendResponseBodyWorkloadList struct {
	// The collection time of the data point. This value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1755360600
	Timestamp *int64 `json:"Timestamp,omitempty" xml:"Timestamp,omitempty"`
	// The workload usage percentage. Valid values: 0 to 100, with two decimal places.
	//
	// example:
	//
	// 35.27
	Workload *float64 `json:"Workload,omitempty" xml:"Workload,omitempty"`
}

func (s GetUserDeviceWorkloadTrendResponseBodyWorkloadList) String() string {
	return dara.Prettify(s)
}

func (s GetUserDeviceWorkloadTrendResponseBodyWorkloadList) GoString() string {
	return s.String()
}

func (s *GetUserDeviceWorkloadTrendResponseBodyWorkloadList) GetTimestamp() *int64 {
	return s.Timestamp
}

func (s *GetUserDeviceWorkloadTrendResponseBodyWorkloadList) GetWorkload() *float64 {
	return s.Workload
}

func (s *GetUserDeviceWorkloadTrendResponseBodyWorkloadList) SetTimestamp(v int64) *GetUserDeviceWorkloadTrendResponseBodyWorkloadList {
	s.Timestamp = &v
	return s
}

func (s *GetUserDeviceWorkloadTrendResponseBodyWorkloadList) SetWorkload(v float64) *GetUserDeviceWorkloadTrendResponseBodyWorkloadList {
	s.Workload = &v
	return s
}

func (s *GetUserDeviceWorkloadTrendResponseBodyWorkloadList) Validate() error {
	return dara.Validate(s)
}
