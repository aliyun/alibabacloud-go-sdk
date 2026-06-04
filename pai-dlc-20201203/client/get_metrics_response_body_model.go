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
	SetDataPoints(v string) *GetMetricsResponseBody
	GetDataPoints() *string
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
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// [{\\"Content\\": \\"\\", \\"OperationName\\": \\"purchase\\", \\"Success\\": 1, \\"Id\\": \\"217\\", \\"LogDatetime\\": 1687679582923}]
	DataPoints *string `json:"DataPoints,omitempty" xml:"DataPoints,omitempty"`
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6lESTRpd5hnHNnmKOP/+w9F
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 5
	Period *string `json:"Period,omitempty" xml:"Period,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
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

func (s *GetMetricsResponseBody) GetDataPoints() *string {
	return s.DataPoints
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

func (s *GetMetricsResponseBody) SetDataPoints(v string) *GetMetricsResponseBody {
	s.DataPoints = &v
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
