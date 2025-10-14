// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceMetricsResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetInstanceMetricsResponseBody
	GetHttpStatusCode() *int32
	SetInstanceId(v string) *GetInstanceMetricsResponseBody
	GetInstanceId() *string
	SetMessage(v string) *GetInstanceMetricsResponseBody
	GetMessage() *string
	SetPodMetrics(v []*GetInstanceMetricsResponseBodyPodMetrics) *GetInstanceMetricsResponseBody
	GetPodMetrics() []*GetInstanceMetricsResponseBodyPodMetrics
	SetRequestId(v string) *GetInstanceMetricsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceMetricsResponseBody
	GetSuccess() *bool
}

type GetInstanceMetricsResponseBody struct {
	// The status code. Valid values:
	//
	// 	- InternalError: an internal error. All errors, except for parameter validation errors, are classified as internal errors.
	//
	// 	- ValidationError: a parameter validation error.
	//
	// example:
	//
	// null
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code. Valid values:
	//
	// 	- 400
	//
	// 	- 404
	//
	// example:
	//
	// null
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The response message.
	//
	// example:
	//
	// "XXX"
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The information about the metrics of the pod that corresponds to the instance.
	PodMetrics []*GetInstanceMetricsResponseBodyPodMetrics `json:"PodMetrics,omitempty" xml:"PodMetrics,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// E7D55162-4489-1619-AAF5-3F97D5FCA948
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceMetricsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceMetricsResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceMetricsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceMetricsResponseBody) GetPodMetrics() []*GetInstanceMetricsResponseBodyPodMetrics {
	return s.PodMetrics
}

func (s *GetInstanceMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceMetricsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceMetricsResponseBody) SetCode(v string) *GetInstanceMetricsResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetHttpStatusCode(v int32) *GetInstanceMetricsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetInstanceId(v string) *GetInstanceMetricsResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetMessage(v string) *GetInstanceMetricsResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetPodMetrics(v []*GetInstanceMetricsResponseBodyPodMetrics) *GetInstanceMetricsResponseBody {
	s.PodMetrics = v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetRequestId(v string) *GetInstanceMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) SetSuccess(v bool) *GetInstanceMetricsResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceMetricsResponseBody) Validate() error {
	if s.PodMetrics != nil {
		for _, item := range s.PodMetrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceMetricsResponseBodyPodMetrics struct {
	// The metrics of the pod that corresponds to the instance.
	Metrics []*GetInstanceMetricsResponseBodyPodMetricsMetrics `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The ID of the pod that corresponds to the instance.
	//
	// example:
	//
	// dsw-15870-695f44c5bc-hd6xm
	PodId *string `json:"PodId,omitempty" xml:"PodId,omitempty"`
}

func (s GetInstanceMetricsResponseBodyPodMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMetricsResponseBodyPodMetrics) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) GetMetrics() []*GetInstanceMetricsResponseBodyPodMetricsMetrics {
	return s.Metrics
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) GetPodId() *string {
	return s.PodId
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) SetMetrics(v []*GetInstanceMetricsResponseBodyPodMetricsMetrics) *GetInstanceMetricsResponseBodyPodMetrics {
	s.Metrics = v
	return s
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) SetPodId(v string) *GetInstanceMetricsResponseBodyPodMetrics {
	s.PodId = &v
	return s
}

func (s *GetInstanceMetricsResponseBodyPodMetrics) Validate() error {
	if s.Metrics != nil {
		for _, item := range s.Metrics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceMetricsResponseBodyPodMetricsMetrics struct {
	// The timestamp corresponding to the metric.
	//
	// example:
	//
	// 1670890560
	Time *int64 `json:"Time,omitempty" xml:"Time,omitempty"`
	// The metric value.
	//
	// example:
	//
	// 25.901031
	Value *float32 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s GetInstanceMetricsResponseBodyPodMetricsMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceMetricsResponseBodyPodMetricsMetrics) GoString() string {
	return s.String()
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) GetTime() *int64 {
	return s.Time
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) GetValue() *float32 {
	return s.Value
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) SetTime(v int64) *GetInstanceMetricsResponseBodyPodMetricsMetrics {
	s.Time = &v
	return s
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) SetValue(v float32) *GetInstanceMetricsResponseBodyPodMetricsMetrics {
	s.Value = &v
	return s
}

func (s *GetInstanceMetricsResponseBodyPodMetricsMetrics) Validate() error {
	return dara.Validate(s)
}
