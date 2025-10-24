// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *GetQuotaUsageResponseBodyData) *GetQuotaUsageResponseBody
	GetData() *GetQuotaUsageResponseBodyData
	SetErrorCode(v string) *GetQuotaUsageResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetQuotaUsageResponseBody
	GetErrorMsg() *string
	SetHttpCode(v int32) *GetQuotaUsageResponseBody
	GetHttpCode() *int32
	SetRequestId(v string) *GetQuotaUsageResponseBody
	GetRequestId() *string
}

type GetQuotaUsageResponseBody struct {
	// The data returned.
	Data *GetQuotaUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// OBJECT_NOT_EXIST
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// This object does not exist.
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The HTTP status code.
	//
	// 	- 1xx: informational response. The request is received and is being processed.
	//
	// 	- 2xx: success. The request is successfully received, understood, and accepted by the server.
	//
	// 	- 3xx: redirection. The request is redirected, and further actions are required to complete the request.
	//
	// 	- 4xx: client error. The request contains invalid request parameters and syntaxes, or specific request conditions cannot be met.
	//
	// 	- 5xx: server error. The server cannot meet requirements due to other reasons.
	//
	// example:
	//
	// 200
	HttpCode *int32 `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0b87b7a416652014358483492eea0b
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetQuotaUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaUsageResponseBody) GetData() *GetQuotaUsageResponseBodyData {
	return s.Data
}

func (s *GetQuotaUsageResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetQuotaUsageResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetQuotaUsageResponseBody) GetHttpCode() *int32 {
	return s.HttpCode
}

func (s *GetQuotaUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQuotaUsageResponseBody) SetData(v *GetQuotaUsageResponseBodyData) *GetQuotaUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetQuotaUsageResponseBody) SetErrorCode(v string) *GetQuotaUsageResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetQuotaUsageResponseBody) SetErrorMsg(v string) *GetQuotaUsageResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetQuotaUsageResponseBody) SetHttpCode(v int32) *GetQuotaUsageResponseBody {
	s.HttpCode = &v
	return s
}

func (s *GetQuotaUsageResponseBody) SetRequestId(v string) *GetQuotaUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaUsageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQuotaUsageResponseBodyData struct {
	// The metric results.
	Metrics map[string]interface{} `json:"metrics,omitempty" xml:"metrics,omitempty"`
	// The information about the chart.
	Plot []*GetQuotaUsageResponseBodyDataPlot `json:"plot,omitempty" xml:"plot,omitempty" type:"Repeated"`
}

func (s GetQuotaUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetQuotaUsageResponseBodyData) GetMetrics() map[string]interface{} {
	return s.Metrics
}

func (s *GetQuotaUsageResponseBodyData) GetPlot() []*GetQuotaUsageResponseBodyDataPlot {
	return s.Plot
}

func (s *GetQuotaUsageResponseBodyData) SetMetrics(v map[string]interface{}) *GetQuotaUsageResponseBodyData {
	s.Metrics = v
	return s
}

func (s *GetQuotaUsageResponseBodyData) SetPlot(v []*GetQuotaUsageResponseBodyDataPlot) *GetQuotaUsageResponseBodyData {
	s.Plot = v
	return s
}

func (s *GetQuotaUsageResponseBodyData) Validate() error {
	if s.Plot != nil {
		for _, item := range s.Plot {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetQuotaUsageResponseBodyDataPlot struct {
	// The title of the chart.
	//
	// example:
	//
	// request
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
	// The type of the chart.
	//
	// example:
	//
	// request
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The data metric field.
	YAxis []*string `json:"yAxis,omitempty" xml:"yAxis,omitempty" type:"Repeated"`
}

func (s GetQuotaUsageResponseBodyDataPlot) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUsageResponseBodyDataPlot) GoString() string {
	return s.String()
}

func (s *GetQuotaUsageResponseBodyDataPlot) GetTitle() *string {
	return s.Title
}

func (s *GetQuotaUsageResponseBodyDataPlot) GetType() *string {
	return s.Type
}

func (s *GetQuotaUsageResponseBodyDataPlot) GetYAxis() []*string {
	return s.YAxis
}

func (s *GetQuotaUsageResponseBodyDataPlot) SetTitle(v string) *GetQuotaUsageResponseBodyDataPlot {
	s.Title = &v
	return s
}

func (s *GetQuotaUsageResponseBodyDataPlot) SetType(v string) *GetQuotaUsageResponseBodyDataPlot {
	s.Type = &v
	return s
}

func (s *GetQuotaUsageResponseBodyDataPlot) SetYAxis(v []*string) *GetQuotaUsageResponseBodyDataPlot {
	s.YAxis = v
	return s
}

func (s *GetQuotaUsageResponseBodyDataPlot) Validate() error {
	return dara.Validate(s)
}
