// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotNodeMetricListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *GetHoneypotNodeMetricListRequest
	GetEndTime() *string
	SetExpress(v string) *GetHoneypotNodeMetricListRequest
	GetExpress() *string
	SetLength(v string) *GetHoneypotNodeMetricListRequest
	GetLength() *string
	SetMetricName(v string) *GetHoneypotNodeMetricListRequest
	GetMetricName() *string
	SetNamespace(v string) *GetHoneypotNodeMetricListRequest
	GetNamespace() *string
	SetNodeId(v string) *GetHoneypotNodeMetricListRequest
	GetNodeId() *string
	SetPeriod(v string) *GetHoneypotNodeMetricListRequest
	GetPeriod() *string
	SetStartTime(v string) *GetHoneypotNodeMetricListRequest
	GetStartTime() *string
}

type GetHoneypotNodeMetricListRequest struct {
	// The end of the time range to query. Valid values:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC
	//
	// 	- Date format: YYYY-MM-DDThh:mm:ssZ
	//
	// example:
	//
	// 2019-01-30 00:10:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The expression that is used to compute the query results in real time.
	//
	// >  Only the groupby expression is supported. This expression is similar to the GROUP BY statement that applies to databases.
	//
	// example:
	//
	// {"groupby":["userId","instanceId"]}
	Express *string `json:"Express,omitempty" xml:"Express,omitempty"`
	// The number of entries per page.
	//
	// >  The maximum value of the Length parameter in a request is 1440.
	//
	// example:
	//
	// 100
	Length *string `json:"Length,omitempty" xml:"Length,omitempty"`
	// The metric that is used to monitor the cloud service.
	//
	// example:
	//
	// cpu_total
	MetricName *string `json:"MetricName,omitempty" xml:"MetricName,omitempty"`
	// The namespace of the cloud service. Format: acs_cloud service name.
	//
	// example:
	//
	// acs_ecs_dashboard
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// The management node ID.
	//
	// example:
	//
	// cc427e14-f257-4670-9d2b-d83bbbe7****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The time interval. Unit: seconds. Valid values: 60, 300, and 900.
	//
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// The beginning of the time range to query. The following formats are supported:
	//
	// 	- UNIX timestamp: the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC
	//
	// 	- Date format: YYYY-MM-DDThh:mm:ssZ
	//
	// 	- The interval between the start time and the end time is less than or equal to 31 days.
	//
	// example:
	//
	// 2019-01-30 00:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s GetHoneypotNodeMetricListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotNodeMetricListRequest) GoString() string {
	return s.String()
}

func (s *GetHoneypotNodeMetricListRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *GetHoneypotNodeMetricListRequest) GetExpress() *string {
	return s.Express
}

func (s *GetHoneypotNodeMetricListRequest) GetLength() *string {
	return s.Length
}

func (s *GetHoneypotNodeMetricListRequest) GetMetricName() *string {
	return s.MetricName
}

func (s *GetHoneypotNodeMetricListRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetHoneypotNodeMetricListRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *GetHoneypotNodeMetricListRequest) GetPeriod() *string {
	return s.Period
}

func (s *GetHoneypotNodeMetricListRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *GetHoneypotNodeMetricListRequest) SetEndTime(v string) *GetHoneypotNodeMetricListRequest {
	s.EndTime = &v
	return s
}

func (s *GetHoneypotNodeMetricListRequest) SetExpress(v string) *GetHoneypotNodeMetricListRequest {
	s.Express = &v
	return s
}

func (s *GetHoneypotNodeMetricListRequest) SetLength(v string) *GetHoneypotNodeMetricListRequest {
	s.Length = &v
	return s
}

func (s *GetHoneypotNodeMetricListRequest) SetMetricName(v string) *GetHoneypotNodeMetricListRequest {
	s.MetricName = &v
	return s
}

func (s *GetHoneypotNodeMetricListRequest) SetNamespace(v string) *GetHoneypotNodeMetricListRequest {
	s.Namespace = &v
	return s
}

func (s *GetHoneypotNodeMetricListRequest) SetNodeId(v string) *GetHoneypotNodeMetricListRequest {
	s.NodeId = &v
	return s
}

func (s *GetHoneypotNodeMetricListRequest) SetPeriod(v string) *GetHoneypotNodeMetricListRequest {
	s.Period = &v
	return s
}

func (s *GetHoneypotNodeMetricListRequest) SetStartTime(v string) *GetHoneypotNodeMetricListRequest {
	s.StartTime = &v
	return s
}

func (s *GetHoneypotNodeMetricListRequest) Validate() error {
	return dara.Validate(s)
}
