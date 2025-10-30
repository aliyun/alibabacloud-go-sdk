// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataReDistributeInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataReDistributeInfo(v *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) *DescribeDataReDistributeInfoResponseBody
	GetDataReDistributeInfo() *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo
	SetRequestId(v string) *DescribeDataReDistributeInfoResponseBody
	GetRequestId() *string
}

type DescribeDataReDistributeInfoResponseBody struct {
	// The data redistribution information.
	DataReDistributeInfo *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo `json:"DataReDistributeInfo,omitempty" xml:"DataReDistributeInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04836A02-ADC9-1AA7-AC36-DE5E048BF505
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDataReDistributeInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataReDistributeInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDataReDistributeInfoResponseBody) GetDataReDistributeInfo() *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	return s.DataReDistributeInfo
}

func (s *DescribeDataReDistributeInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDataReDistributeInfoResponseBody) SetDataReDistributeInfo(v *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) *DescribeDataReDistributeInfoResponseBody {
	s.DataReDistributeInfo = v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBody) SetRequestId(v string) *DescribeDataReDistributeInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBody) Validate() error {
	if s.DataReDistributeInfo != nil {
		if err := s.DataReDistributeInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo struct {
	// The execution information. If an error occurs, the error message is returned.
	//
	// example:
	//
	// redistributing
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The progress of data redistribution. Unit: %.
	//
	// example:
	//
	// 33
	Progress *int64 `json:"Progress,omitempty" xml:"Progress,omitempty"`
	// The estimated remaining time for data redistribution.
	//
	// example:
	//
	// 00:01:28
	RemainTime *string `json:"RemainTime,omitempty" xml:"RemainTime,omitempty"`
	// This parameter is not supported.
	//
	// example:
	//
	// null
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The status of data redistribution.
	//
	// example:
	//
	// RUNNING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The execution type. The value **immediate*	- is returned, indicating immediate execution.
	//
	// example:
	//
	// immediate
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) GoString() string {
	return s.String()
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) GetMessage() *string {
	return s.Message
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) GetProgress() *int64 {
	return s.Progress
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) GetRemainTime() *string {
	return s.RemainTime
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) GetStatus() *string {
	return s.Status
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) GetType() *string {
	return s.Type
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) SetMessage(v string) *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	s.Message = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) SetProgress(v int64) *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	s.Progress = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) SetRemainTime(v string) *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	s.RemainTime = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) SetStartTime(v string) *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	s.StartTime = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) SetStatus(v string) *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	s.Status = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) SetType(v string) *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo {
	s.Type = &v
	return s
}

func (s *DescribeDataReDistributeInfoResponseBodyDataReDistributeInfo) Validate() error {
	return dara.Validate(s)
}
