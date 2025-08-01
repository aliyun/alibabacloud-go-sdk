// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *QueryMonitorRequest
	GetAcceptLanguage() *string
	SetEndTime(v int64) *QueryMonitorRequest
	GetEndTime() *int64
	SetInstanceId(v string) *QueryMonitorRequest
	GetInstanceId() *string
	SetMonitorType(v string) *QueryMonitorRequest
	GetMonitorType() *string
	SetRequestPars(v string) *QueryMonitorRequest
	GetRequestPars() *string
	SetStartTime(v int64) *QueryMonitorRequest
	GetStartTime() *int64
	SetStep(v int64) *QueryMonitorRequest
	GetStep() *int64
}

type QueryMonitorRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The timestamp when the monitoring ends.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1666678376
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// mse-cn-****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The metric type. The following metric types are supported:
	//
	// [Basic system metrics]
	//
	// 	- cpuUsage
	//
	// 	- memoryUsage
	//
	// 	- diskUsage
	//
	// 	- gcCount
	//
	// 	- gcTime
	//
	// [Nacos registry]
	//
	// 	- serviceCount
	//
	// 	- writeCostTime
	//
	// 	- readCostTime
	//
	// 	- TPS regCenterTps
	//
	// 	- QPS regCenterQps
	//
	// [Nacos configuration center]
	//
	// 	- publish
	//
	// 	- getConfig
	//
	// [zookeeper]
	//
	// 	- TPS zk_TpsCount
	//
	// 	- QPS zk_QpsCount
	//
	// 	- zookeeper_AvgRequestLatency
	//
	// This parameter is required.
	//
	// example:
	//
	// regCenterQps
	MonitorType *string `json:"MonitorType,omitempty" xml:"MonitorType,omitempty"`
	// The extended request parameters in the JSON format.
	//
	// example:
	//
	// {}
	RequestPars *string `json:"RequestPars,omitempty" xml:"RequestPars,omitempty"`
	// The timestamp when the monitoring starts.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1666678376
	StartTime *int64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The interval between data points. Unit: seconds.
	//
	// example:
	//
	// 7
	Step *int64 `json:"Step,omitempty" xml:"Step,omitempty"`
}

func (s QueryMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryMonitorRequest) GoString() string {
	return s.String()
}

func (s *QueryMonitorRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *QueryMonitorRequest) GetEndTime() *int64 {
	return s.EndTime
}

func (s *QueryMonitorRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryMonitorRequest) GetMonitorType() *string {
	return s.MonitorType
}

func (s *QueryMonitorRequest) GetRequestPars() *string {
	return s.RequestPars
}

func (s *QueryMonitorRequest) GetStartTime() *int64 {
	return s.StartTime
}

func (s *QueryMonitorRequest) GetStep() *int64 {
	return s.Step
}

func (s *QueryMonitorRequest) SetAcceptLanguage(v string) *QueryMonitorRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *QueryMonitorRequest) SetEndTime(v int64) *QueryMonitorRequest {
	s.EndTime = &v
	return s
}

func (s *QueryMonitorRequest) SetInstanceId(v string) *QueryMonitorRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryMonitorRequest) SetMonitorType(v string) *QueryMonitorRequest {
	s.MonitorType = &v
	return s
}

func (s *QueryMonitorRequest) SetRequestPars(v string) *QueryMonitorRequest {
	s.RequestPars = &v
	return s
}

func (s *QueryMonitorRequest) SetStartTime(v int64) *QueryMonitorRequest {
	s.StartTime = &v
	return s
}

func (s *QueryMonitorRequest) SetStep(v int64) *QueryMonitorRequest {
	s.Step = &v
	return s
}

func (s *QueryMonitorRequest) Validate() error {
	return dara.Validate(s)
}
