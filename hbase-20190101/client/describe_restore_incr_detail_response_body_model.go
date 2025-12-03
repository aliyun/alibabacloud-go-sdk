// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRestoreIncrDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRestoreIncrDetailResponseBody
	GetRequestId() *string
	SetRestoreIncrDetail(v *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) *DescribeRestoreIncrDetailResponseBody
	GetRestoreIncrDetail() *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail
}

type DescribeRestoreIncrDetailResponseBody struct {
	// example:
	//
	// D0FE2717-E194-465A-B27B-7373F96E580B
	RequestId         *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RestoreIncrDetail *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail `json:"RestoreIncrDetail,omitempty" xml:"RestoreIncrDetail,omitempty" type:"Struct"`
}

func (s DescribeRestoreIncrDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreIncrDetailResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRestoreIncrDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRestoreIncrDetailResponseBody) GetRestoreIncrDetail() *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	return s.RestoreIncrDetail
}

func (s *DescribeRestoreIncrDetailResponseBody) SetRequestId(v string) *DescribeRestoreIncrDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBody) SetRestoreIncrDetail(v *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) *DescribeRestoreIncrDetailResponseBody {
	s.RestoreIncrDetail = v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBody) Validate() error {
	if s.RestoreIncrDetail != nil {
		if err := s.RestoreIncrDetail.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail struct {
	// example:
	//
	// 2020-11-05T06:45:44Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 0/0
	Process *string `json:"Process,omitempty" xml:"Process,omitempty"`
	// example:
	//
	// 0 ms
	RestoreDelay *string `json:"RestoreDelay,omitempty" xml:"RestoreDelay,omitempty"`
	// example:
	//
	// 2020-11-02T18:00:00Z
	RestoreStartTs *string `json:"RestoreStartTs,omitempty" xml:"RestoreStartTs,omitempty"`
	// example:
	//
	// \\"\\"
	RestoredTs *string `json:"RestoredTs,omitempty" xml:"RestoredTs,omitempty"`
	// example:
	//
	// 2020-11-05T06:45:44Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// example:
	//
	// SUCCEEDED
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) GoString() string {
	return s.String()
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) GetProcess() *string {
	return s.Process
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) GetRestoreDelay() *string {
	return s.RestoreDelay
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) GetRestoreStartTs() *string {
	return s.RestoreStartTs
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) GetRestoredTs() *string {
	return s.RestoredTs
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) GetState() *string {
	return s.State
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetEndTime(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.EndTime = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetProcess(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.Process = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetRestoreDelay(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.RestoreDelay = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetRestoreStartTs(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.RestoreStartTs = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetRestoredTs(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.RestoredTs = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetStartTime(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.StartTime = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) SetState(v string) *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail {
	s.State = &v
	return s
}

func (s *DescribeRestoreIncrDetailResponseBodyRestoreIncrDetail) Validate() error {
	return dara.Validate(s)
}
