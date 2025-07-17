// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDIJobMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListDIJobMetricsResponseBodyPagingInfo) *ListDIJobMetricsResponseBody
	GetPagingInfo() *ListDIJobMetricsResponseBodyPagingInfo
	SetRequestId(v string) *ListDIJobMetricsResponseBody
	GetRequestId() *string
}

type ListDIJobMetricsResponseBody struct {
	// The pagination information.
	PagingInfo *ListDIJobMetricsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 691CA452-D37A-4ED0-9441
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDIJobMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponseBody) GetPagingInfo() *ListDIJobMetricsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListDIJobMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDIJobMetricsResponseBody) SetPagingInfo(v *ListDIJobMetricsResponseBodyPagingInfo) *ListDIJobMetricsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListDIJobMetricsResponseBody) SetRequestId(v string) *ListDIJobMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDIJobMetricsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListDIJobMetricsResponseBodyPagingInfo struct {
	// The metrics returned.
	JobMetrics []*ListDIJobMetricsResponseBodyPagingInfoJobMetrics `json:"JobMetrics,omitempty" xml:"JobMetrics,omitempty" type:"Repeated"`
}

func (s ListDIJobMetricsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobMetricsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponseBodyPagingInfo) GetJobMetrics() []*ListDIJobMetricsResponseBodyPagingInfoJobMetrics {
	return s.JobMetrics
}

func (s *ListDIJobMetricsResponseBodyPagingInfo) SetJobMetrics(v []*ListDIJobMetricsResponseBodyPagingInfoJobMetrics) *ListDIJobMetricsResponseBodyPagingInfo {
	s.JobMetrics = v
	return s
}

func (s *ListDIJobMetricsResponseBodyPagingInfo) Validate() error {
	return dara.Validate(s)
}

type ListDIJobMetricsResponseBodyPagingInfoJobMetrics struct {
	// The name of the metric.
	//
	// example:
	//
	// JobDelay
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The metric data.
	SeriesList []*ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList `json:"SeriesList,omitempty" xml:"SeriesList,omitempty" type:"Repeated"`
}

func (s ListDIJobMetricsResponseBodyPagingInfoJobMetrics) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobMetricsResponseBodyPagingInfoJobMetrics) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetrics) GetName() *string {
	return s.Name
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetrics) GetSeriesList() []*ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList {
	return s.SeriesList
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetrics) SetName(v string) *ListDIJobMetricsResponseBodyPagingInfoJobMetrics {
	s.Name = &v
	return s
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetrics) SetSeriesList(v []*ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) *ListDIJobMetricsResponseBodyPagingInfoJobMetrics {
	s.SeriesList = v
	return s
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetrics) Validate() error {
	return dara.Validate(s)
}

type ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList struct {
	// The point in time at which data is sampled based on the metric.
	//
	// example:
	//
	// 1716881141
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The sample value.
	//
	// example:
	//
	// 10
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) String() string {
	return dara.Prettify(s)
}

func (s ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) GoString() string {
	return s.String()
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) GetTime() *int64 {
	return s.Time
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) GetValue() *float64 {
	return s.Value
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) SetTime(v int64) *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList {
	s.Time = &v
	return s
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) SetValue(v float64) *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList {
	s.Value = &v
	return s
}

func (s *ListDIJobMetricsResponseBodyPagingInfoJobMetricsSeriesList) Validate() error {
	return dara.Validate(s)
}
