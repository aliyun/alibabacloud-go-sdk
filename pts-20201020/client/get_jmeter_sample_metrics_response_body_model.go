// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJMeterSampleMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetJMeterSampleMetricsResponseBody
	GetCode() *string
	SetMessage(v string) *GetJMeterSampleMetricsResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetJMeterSampleMetricsResponseBody
	GetRequestId() *string
	SetSampleMetricList(v []*string) *GetJMeterSampleMetricsResponseBody
	GetSampleMetricList() []*string
	SetSamplerMap(v map[string]interface{}) *GetJMeterSampleMetricsResponseBody
	GetSamplerMap() map[string]interface{}
	SetSuccess(v bool) *GetJMeterSampleMetricsResponseBody
	GetSuccess() *bool
}

type GetJMeterSampleMetricsResponseBody struct {
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message. If the request was successful, this parameter is left empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A8E16480-15C1-555A-922F-B736A005E52D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The metrics of the samplers.
	SampleMetricList []*string `json:"SampleMetricList,omitempty" xml:"SampleMetricList,omitempty" type:"Repeated"`
	// The sampler list. You can find the sampler to be queried based on this list.
	//
	// example:
	//
	// {0:"Http Request"}
	SamplerMap map[string]interface{} `json:"SamplerMap,omitempty" xml:"SamplerMap,omitempty"`
	// Indicates whether the request was successful. Valid values:
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

func (s GetJMeterSampleMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetJMeterSampleMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetJMeterSampleMetricsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetJMeterSampleMetricsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetJMeterSampleMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetJMeterSampleMetricsResponseBody) GetSampleMetricList() []*string {
	return s.SampleMetricList
}

func (s *GetJMeterSampleMetricsResponseBody) GetSamplerMap() map[string]interface{} {
	return s.SamplerMap
}

func (s *GetJMeterSampleMetricsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetJMeterSampleMetricsResponseBody) SetCode(v string) *GetJMeterSampleMetricsResponseBody {
	s.Code = &v
	return s
}

func (s *GetJMeterSampleMetricsResponseBody) SetMessage(v string) *GetJMeterSampleMetricsResponseBody {
	s.Message = &v
	return s
}

func (s *GetJMeterSampleMetricsResponseBody) SetRequestId(v string) *GetJMeterSampleMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetJMeterSampleMetricsResponseBody) SetSampleMetricList(v []*string) *GetJMeterSampleMetricsResponseBody {
	s.SampleMetricList = v
	return s
}

func (s *GetJMeterSampleMetricsResponseBody) SetSamplerMap(v map[string]interface{}) *GetJMeterSampleMetricsResponseBody {
	s.SamplerMap = v
	return s
}

func (s *GetJMeterSampleMetricsResponseBody) SetSuccess(v bool) *GetJMeterSampleMetricsResponseBody {
	s.Success = &v
	return s
}

func (s *GetJMeterSampleMetricsResponseBody) Validate() error {
	return dara.Validate(s)
}
