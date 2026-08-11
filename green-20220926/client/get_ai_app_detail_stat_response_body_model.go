// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAiAppDetailStatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetAiAppDetailStatResponseBody
	GetAppId() *string
	SetAvgModelDuration(v float32) *GetAiAppDetailStatResponseBody
	GetAvgModelDuration() *float32
	SetAvgModelDurationDau(v float32) *GetAiAppDetailStatResponseBody
	GetAvgModelDurationDau() *float32
	SetModelCount(v int64) *GetAiAppDetailStatResponseBody
	GetModelCount() *int64
	SetModelCountDau(v float32) *GetAiAppDetailStatResponseBody
	GetModelCountDau() *float32
	SetRequestId(v string) *GetAiAppDetailStatResponseBody
	GetRequestId() *string
	SetRiskEventCount(v int64) *GetAiAppDetailStatResponseBody
	GetRiskEventCount() *int64
	SetTokenCount(v int64) *GetAiAppDetailStatResponseBody
	GetTokenCount() *int64
	SetTokenCountDau(v float32) *GetAiAppDetailStatResponseBody
	GetTokenCountDau() *float32
}

type GetAiAppDetailStatResponseBody struct {
	// The application ID.
	//
	// example:
	//
	// id-xxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The average duration of model calls.
	//
	// example:
	//
	// 2.55
	AvgModelDuration *float32 `json:"AvgModelDuration,omitempty" xml:"AvgModelDuration,omitempty"`
	// The day-over-day change ratio of average model call duration.
	//
	// example:
	//
	// 0.05
	AvgModelDurationDau *float32 `json:"AvgModelDurationDau,omitempty" xml:"AvgModelDurationDau,omitempty"`
	// The number of model calls.
	//
	// example:
	//
	// 15
	ModelCount *int64 `json:"ModelCount,omitempty" xml:"ModelCount,omitempty"`
	// The day-over-day change ratio of model call count.
	//
	// example:
	//
	// 0.15
	ModelCountDau *float32 `json:"ModelCountDau,omitempty" xml:"ModelCountDau,omitempty"`
	// The ID assigned by the backend to uniquely identify a request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of risk events.
	//
	// example:
	//
	// 1
	RiskEventCount *int64 `json:"RiskEventCount,omitempty" xml:"RiskEventCount,omitempty"`
	// The number of tokens consumed.
	//
	// example:
	//
	// 11000
	TokenCount *int64 `json:"TokenCount,omitempty" xml:"TokenCount,omitempty"`
	// The day-over-day change ratio of token consumption count.
	//
	// example:
	//
	// -0.15
	TokenCountDau *float32 `json:"TokenCountDau,omitempty" xml:"TokenCountDau,omitempty"`
}

func (s GetAiAppDetailStatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAiAppDetailStatResponseBody) GoString() string {
	return s.String()
}

func (s *GetAiAppDetailStatResponseBody) GetAppId() *string {
	return s.AppId
}

func (s *GetAiAppDetailStatResponseBody) GetAvgModelDuration() *float32 {
	return s.AvgModelDuration
}

func (s *GetAiAppDetailStatResponseBody) GetAvgModelDurationDau() *float32 {
	return s.AvgModelDurationDau
}

func (s *GetAiAppDetailStatResponseBody) GetModelCount() *int64 {
	return s.ModelCount
}

func (s *GetAiAppDetailStatResponseBody) GetModelCountDau() *float32 {
	return s.ModelCountDau
}

func (s *GetAiAppDetailStatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAiAppDetailStatResponseBody) GetRiskEventCount() *int64 {
	return s.RiskEventCount
}

func (s *GetAiAppDetailStatResponseBody) GetTokenCount() *int64 {
	return s.TokenCount
}

func (s *GetAiAppDetailStatResponseBody) GetTokenCountDau() *float32 {
	return s.TokenCountDau
}

func (s *GetAiAppDetailStatResponseBody) SetAppId(v string) *GetAiAppDetailStatResponseBody {
	s.AppId = &v
	return s
}

func (s *GetAiAppDetailStatResponseBody) SetAvgModelDuration(v float32) *GetAiAppDetailStatResponseBody {
	s.AvgModelDuration = &v
	return s
}

func (s *GetAiAppDetailStatResponseBody) SetAvgModelDurationDau(v float32) *GetAiAppDetailStatResponseBody {
	s.AvgModelDurationDau = &v
	return s
}

func (s *GetAiAppDetailStatResponseBody) SetModelCount(v int64) *GetAiAppDetailStatResponseBody {
	s.ModelCount = &v
	return s
}

func (s *GetAiAppDetailStatResponseBody) SetModelCountDau(v float32) *GetAiAppDetailStatResponseBody {
	s.ModelCountDau = &v
	return s
}

func (s *GetAiAppDetailStatResponseBody) SetRequestId(v string) *GetAiAppDetailStatResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAiAppDetailStatResponseBody) SetRiskEventCount(v int64) *GetAiAppDetailStatResponseBody {
	s.RiskEventCount = &v
	return s
}

func (s *GetAiAppDetailStatResponseBody) SetTokenCount(v int64) *GetAiAppDetailStatResponseBody {
	s.TokenCount = &v
	return s
}

func (s *GetAiAppDetailStatResponseBody) SetTokenCountDau(v float32) *GetAiAppDetailStatResponseBody {
	s.TokenCountDau = &v
	return s
}

func (s *GetAiAppDetailStatResponseBody) Validate() error {
	return dara.Validate(s)
}
