// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMetricsResponseBody
	GetCode() *string
	SetDatapoints(v string) *GetMetricsResponseBody
	GetDatapoints() *string
	SetMessage(v string) *GetMetricsResponseBody
	GetMessage() *string
	SetNextToken(v string) *GetMetricsResponseBody
	GetNextToken() *string
	SetPeriod(v string) *GetMetricsResponseBody
	GetPeriod() *string
	SetRequestId(v string) *GetMetricsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetMetricsResponseBody
	GetSuccess() *bool
}

type GetMetricsResponseBody struct {
	// example:
	//
	// 200
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Datapoints *string `json:"Datapoints,omitempty" xml:"Datapoints,omitempty"`
	// example:
	//
	// Succeed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 15761485350009dd70bb64cff1f0fff750b08ffff073be5fb1e785e2b020f1a949d5ea14aea7fed82f01dd8****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 60
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetMetricsResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMetricsResponseBody) GetDatapoints() *string {
	return s.Datapoints
}

func (s *GetMetricsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMetricsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetMetricsResponseBody) GetPeriod() *string {
	return s.Period
}

func (s *GetMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMetricsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMetricsResponseBody) SetCode(v string) *GetMetricsResponseBody {
	s.Code = &v
	return s
}

func (s *GetMetricsResponseBody) SetDatapoints(v string) *GetMetricsResponseBody {
	s.Datapoints = &v
	return s
}

func (s *GetMetricsResponseBody) SetMessage(v string) *GetMetricsResponseBody {
	s.Message = &v
	return s
}

func (s *GetMetricsResponseBody) SetNextToken(v string) *GetMetricsResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetMetricsResponseBody) SetPeriod(v string) *GetMetricsResponseBody {
	s.Period = &v
	return s
}

func (s *GetMetricsResponseBody) SetRequestId(v string) *GetMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMetricsResponseBody) SetSuccess(v bool) *GetMetricsResponseBody {
	s.Success = &v
	return s
}

func (s *GetMetricsResponseBody) Validate() error {
	return dara.Validate(s)
}
