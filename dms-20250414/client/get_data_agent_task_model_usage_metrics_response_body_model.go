// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAgentTaskModelUsageMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*GetDataAgentTaskModelUsageMetricsResponseBodyData) *GetDataAgentTaskModelUsageMetricsResponseBody
	GetData() []*GetDataAgentTaskModelUsageMetricsResponseBodyData
	SetErrorCode(v string) *GetDataAgentTaskModelUsageMetricsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetDataAgentTaskModelUsageMetricsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *GetDataAgentTaskModelUsageMetricsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDataAgentTaskModelUsageMetricsResponseBody
	GetSuccess() *bool
}

type GetDataAgentTaskModelUsageMetricsResponseBody struct {
	// The list of TPM time series metrics for model usage, returned in chronological order.
	Data []*GetDataAgentTaskModelUsageMetricsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The error code returned when the request is abnormal.
	//
	// example:
	//
	// DMS-DA-40411
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned when the call fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 67E910F2-***7695C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDataAgentTaskModelUsageMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentTaskModelUsageMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBody) GetData() []*GetDataAgentTaskModelUsageMetricsResponseBodyData {
	return s.Data
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBody) SetData(v []*GetDataAgentTaskModelUsageMetricsResponseBodyData) *GetDataAgentTaskModelUsageMetricsResponseBody {
	s.Data = v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBody) SetErrorCode(v string) *GetDataAgentTaskModelUsageMetricsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBody) SetErrorMessage(v string) *GetDataAgentTaskModelUsageMetricsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBody) SetRequestId(v string) *GetDataAgentTaskModelUsageMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBody) SetSuccess(v bool) *GetDataAgentTaskModelUsageMetricsResponseBody {
	s.Success = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetDataAgentTaskModelUsageMetricsResponseBodyData struct {
	// The start time of the statistical interval for this time series data point. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1735660800
	BeginTime *int64 `json:"BeginTime,omitempty" xml:"BeginTime,omitempty"`
	// The end time of the statistical interval for this time series data point. The value is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1735660860
	EndTime *int64 `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The TPM for this time series data point, which is the number of tokens consumed within the statistical interval.
	//
	// example:
	//
	// 1200
	Tpm *int64 `json:"Tpm,omitempty" xml:"Tpm,omitempty"`
}

func (s GetDataAgentTaskModelUsageMetricsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetDataAgentTaskModelUsageMetricsResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBodyData) GetBeginTime() *int64 {
	return s.BeginTime
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBodyData) GetEndTime() *int64 {
	return s.EndTime
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBodyData) GetTpm() *int64 {
	return s.Tpm
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBodyData) SetBeginTime(v int64) *GetDataAgentTaskModelUsageMetricsResponseBodyData {
	s.BeginTime = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBodyData) SetEndTime(v int64) *GetDataAgentTaskModelUsageMetricsResponseBodyData {
	s.EndTime = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBodyData) SetTpm(v int64) *GetDataAgentTaskModelUsageMetricsResponseBodyData {
	s.Tpm = &v
	return s
}

func (s *GetDataAgentTaskModelUsageMetricsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
