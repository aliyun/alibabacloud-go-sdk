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
	Data      *GetQuotaUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	ErrorCode *string                        `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	ErrorMsg  *string                        `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	HttpCode  *int32                         `json:"httpCode,omitempty" xml:"httpCode,omitempty"`
	RequestId *string                        `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
	Metrics map[string]interface{}               `json:"metrics,omitempty" xml:"metrics,omitempty"`
	Plot    []*GetQuotaUsageResponseBodyDataPlot `json:"plot,omitempty" xml:"plot,omitempty" type:"Repeated"`
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
	Title *string   `json:"title,omitempty" xml:"title,omitempty"`
	Type  *string   `json:"type,omitempty" xml:"type,omitempty"`
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
