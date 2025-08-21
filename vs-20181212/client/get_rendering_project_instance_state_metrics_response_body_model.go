// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRenderingProjectInstanceStateMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetRenderingProjectInstanceStateMetricsResponseBody
	GetRequestId() *string
	SetStateMetrics(v []*GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics) *GetRenderingProjectInstanceStateMetricsResponseBody
	GetStateMetrics() []*GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics
}

type GetRenderingProjectInstanceStateMetricsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId    *string                                                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StateMetrics []*GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics `json:"StateMetrics,omitempty" xml:"StateMetrics,omitempty" type:"Repeated"`
}

func (s GetRenderingProjectInstanceStateMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRenderingProjectInstanceStateMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetRenderingProjectInstanceStateMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRenderingProjectInstanceStateMetricsResponseBody) GetStateMetrics() []*GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics {
	return s.StateMetrics
}

func (s *GetRenderingProjectInstanceStateMetricsResponseBody) SetRequestId(v string) *GetRenderingProjectInstanceStateMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRenderingProjectInstanceStateMetricsResponseBody) SetStateMetrics(v []*GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics) *GetRenderingProjectInstanceStateMetricsResponseBody {
	s.StateMetrics = v
	return s
}

func (s *GetRenderingProjectInstanceStateMetricsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics struct {
	// example:
	//
	// 10
	Count *string `json:"Count,omitempty" xml:"Count,omitempty"`
	// example:
	//
	// Idle
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics) String() string {
	return dara.Prettify(s)
}

func (s GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics) GoString() string {
	return s.String()
}

func (s *GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics) GetCount() *string {
	return s.Count
}

func (s *GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics) GetState() *string {
	return s.State
}

func (s *GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics) SetCount(v string) *GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics {
	s.Count = &v
	return s
}

func (s *GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics) SetState(v string) *GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics {
	s.State = &v
	return s
}

func (s *GetRenderingProjectInstanceStateMetricsResponseBodyStateMetrics) Validate() error {
	return dara.Validate(s)
}
