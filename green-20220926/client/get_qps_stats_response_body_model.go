// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQpsStatsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCharts(v map[string]*ChartsValue) *GetQpsStatsResponseBody
	GetCharts() map[string]*ChartsValue
	SetRequestId(v string) *GetQpsStatsResponseBody
	GetRequestId() *string
}

type GetQpsStatsResponseBody struct {
	// The chart configurations.
	//
	// example:
	//
	// {"cn-shanghai":{"X":["20260701"],"Y":[{"Name":"cn-shanghai","Data":[10]}]}}
	Charts map[string]*ChartsValue `json:"Charts,omitempty" xml:"Charts,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetQpsStatsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQpsStatsResponseBody) GoString() string {
	return s.String()
}

func (s *GetQpsStatsResponseBody) GetCharts() map[string]*ChartsValue {
	return s.Charts
}

func (s *GetQpsStatsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQpsStatsResponseBody) SetCharts(v map[string]*ChartsValue) *GetQpsStatsResponseBody {
	s.Charts = v
	return s
}

func (s *GetQpsStatsResponseBody) SetRequestId(v string) *GetQpsStatsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQpsStatsResponseBody) Validate() error {
	return dara.Validate(s)
}
