// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlowLogHistogramAsyncRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v int64) *DescribeSlowLogHistogramAsyncRequest
	GetEndTime() *int64
	SetFilters(v []*DescribeSlowLogHistogramAsyncRequestFilters) *DescribeSlowLogHistogramAsyncRequest
	GetFilters() []*DescribeSlowLogHistogramAsyncRequestFilters
	SetInstanceId(v string) *DescribeSlowLogHistogramAsyncRequest
	GetInstanceId() *string
	SetNodeId(v string) *DescribeSlowLogHistogramAsyncRequest
	GetNodeId() *string
	SetStartTime(v int64) *DescribeSlowLogHistogramAsyncRequest
	GetStartTime() *int64
}

type DescribeSlowLogHistogramAsyncRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1662518540764
	EndTime *int64                                         `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	Filters []*DescribeSlowLogHistogramAsyncRequestFilters `json:"Filters,omitempty" xml:"Filters,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// rm-2ze8g2am97624****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// r-****-db-0
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1596177993000
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSlowLogHistogramAsyncRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogHistogramAsyncRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogHistogramAsyncRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *DescribeSlowLogHistogramAsyncRequest) GetFilters() []*DescribeSlowLogHistogramAsyncRequestFilters {
	return s.Filters
}

func (s *DescribeSlowLogHistogramAsyncRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSlowLogHistogramAsyncRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSlowLogHistogramAsyncRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *DescribeSlowLogHistogramAsyncRequest) SetEndTime(v int64) *DescribeSlowLogHistogramAsyncRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncRequest) SetFilters(v []*DescribeSlowLogHistogramAsyncRequestFilters) *DescribeSlowLogHistogramAsyncRequest {
	s.Filters = v
	return s
}

func (s *DescribeSlowLogHistogramAsyncRequest) SetInstanceId(v string) *DescribeSlowLogHistogramAsyncRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncRequest) SetNodeId(v string) *DescribeSlowLogHistogramAsyncRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncRequest) SetStartTime(v int64) *DescribeSlowLogHistogramAsyncRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncRequest) Validate() error {
	if s.Filters != nil {
		for _, item := range s.Filters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlowLogHistogramAsyncRequestFilters struct {
	// example:
	//
	// None
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// None
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSlowLogHistogramAsyncRequestFilters) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlowLogHistogramAsyncRequestFilters) GoString() string {
	return s.String()
}

func (s *DescribeSlowLogHistogramAsyncRequestFilters) GetKey() *string {
	return s.Key
}

func (s *DescribeSlowLogHistogramAsyncRequestFilters) GetValue() *string {
	return s.Value
}

func (s *DescribeSlowLogHistogramAsyncRequestFilters) SetKey(v string) *DescribeSlowLogHistogramAsyncRequestFilters {
	s.Key = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncRequestFilters) SetValue(v string) *DescribeSlowLogHistogramAsyncRequestFilters {
	s.Value = &v
	return s
}

func (s *DescribeSlowLogHistogramAsyncRequestFilters) Validate() error {
	return dara.Validate(s)
}
