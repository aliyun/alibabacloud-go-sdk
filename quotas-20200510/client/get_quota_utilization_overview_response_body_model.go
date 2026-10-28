// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQuotaUtilizationOverviewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetQuotaUtilizationOverviewResponseBody
	GetErrorCode() *string
	SetErrorMsg(v string) *GetQuotaUtilizationOverviewResponseBody
	GetErrorMsg() *string
	SetGeneratedAt(v string) *GetQuotaUtilizationOverviewResponseBody
	GetGeneratedAt() *string
	SetMessage(v string) *GetQuotaUtilizationOverviewResponseBody
	GetMessage() *string
	SetQuotaOverview(v *GetQuotaUtilizationOverviewResponseBodyQuotaOverview) *GetQuotaUtilizationOverviewResponseBody
	GetQuotaOverview() *GetQuotaUtilizationOverviewResponseBodyQuotaOverview
	SetRequestId(v string) *GetQuotaUtilizationOverviewResponseBody
	GetRequestId() *string
	SetStatus(v string) *GetQuotaUtilizationOverviewResponseBody
	GetStatus() *string
	SetSuccess(v bool) *GetQuotaUtilizationOverviewResponseBody
	GetSuccess() *bool
}

type GetQuotaUtilizationOverviewResponseBody struct {
	// example:
	//
	// RAM.PERMISSION.DENIED
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// You are not authorized to do this action or the API input parameter is invalid.
	ErrorMsg *string `json:"ErrorMsg,omitempty" xml:"ErrorMsg,omitempty"`
	// example:
	//
	// 2026-10-21T07:30:10Z
	GeneratedAt *string `json:"GeneratedAt,omitempty" xml:"GeneratedAt,omitempty"`
	// example:
	//
	// Quota utilization data is available.
	Message       *string                                               `json:"Message,omitempty" xml:"Message,omitempty"`
	QuotaOverview *GetQuotaUtilizationOverviewResponseBodyQuotaOverview `json:"QuotaOverview,omitempty" xml:"QuotaOverview,omitempty" type:"Struct"`
	// example:
	//
	// 920D8A47-26BB-49FA-A09F-F98D7DAA55F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetQuotaUtilizationOverviewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUtilizationOverviewResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaUtilizationOverviewResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetQuotaUtilizationOverviewResponseBody) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *GetQuotaUtilizationOverviewResponseBody) GetGeneratedAt() *string {
	return s.GeneratedAt
}

func (s *GetQuotaUtilizationOverviewResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetQuotaUtilizationOverviewResponseBody) GetQuotaOverview() *GetQuotaUtilizationOverviewResponseBodyQuotaOverview {
	return s.QuotaOverview
}

func (s *GetQuotaUtilizationOverviewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetQuotaUtilizationOverviewResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetQuotaUtilizationOverviewResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetQuotaUtilizationOverviewResponseBody) SetErrorCode(v string) *GetQuotaUtilizationOverviewResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBody) SetErrorMsg(v string) *GetQuotaUtilizationOverviewResponseBody {
	s.ErrorMsg = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBody) SetGeneratedAt(v string) *GetQuotaUtilizationOverviewResponseBody {
	s.GeneratedAt = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBody) SetMessage(v string) *GetQuotaUtilizationOverviewResponseBody {
	s.Message = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBody) SetQuotaOverview(v *GetQuotaUtilizationOverviewResponseBodyQuotaOverview) *GetQuotaUtilizationOverviewResponseBody {
	s.QuotaOverview = v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBody) SetRequestId(v string) *GetQuotaUtilizationOverviewResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBody) SetStatus(v string) *GetQuotaUtilizationOverviewResponseBody {
	s.Status = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBody) SetSuccess(v bool) *GetQuotaUtilizationOverviewResponseBody {
	s.Success = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBody) Validate() error {
	if s.QuotaOverview != nil {
		if err := s.QuotaOverview.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQuotaUtilizationOverviewResponseBodyQuotaOverview struct {
	InUse         *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse         `json:"InUse,omitempty" xml:"InUse,omitempty" type:"Struct"`
	Over80Percent *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent `json:"Over80Percent,omitempty" xml:"Over80Percent,omitempty" type:"Struct"`
	Over95Percent *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent `json:"Over95Percent,omitempty" xml:"Over95Percent,omitempty" type:"Struct"`
}

func (s GetQuotaUtilizationOverviewResponseBodyQuotaOverview) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUtilizationOverviewResponseBodyQuotaOverview) GoString() string {
	return s.String()
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverview) GetInUse() *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse {
	return s.InUse
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverview) GetOver80Percent() *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent {
	return s.Over80Percent
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverview) GetOver95Percent() *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent {
	return s.Over95Percent
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverview) SetInUse(v *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse) *GetQuotaUtilizationOverviewResponseBodyQuotaOverview {
	s.InUse = v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverview) SetOver80Percent(v *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent) *GetQuotaUtilizationOverviewResponseBodyQuotaOverview {
	s.Over80Percent = v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverview) SetOver95Percent(v *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent) *GetQuotaUtilizationOverviewResponseBodyQuotaOverview {
	s.Over95Percent = v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverview) Validate() error {
	if s.InUse != nil {
		if err := s.InUse.Validate(); err != nil {
			return err
		}
	}
	if s.Over80Percent != nil {
		if err := s.Over80Percent.Validate(); err != nil {
			return err
		}
	}
	if s.Over95Percent != nil {
		if err := s.Over95Percent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse struct {
	// example:
	//
	// 50
	Adjustable *int64 `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	// example:
	//
	// 50
	AdjustableRatio *float64 `json:"AdjustableRatio,omitempty" xml:"AdjustableRatio,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse) GoString() string {
	return s.String()
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse) GetAdjustable() *int64 {
	return s.Adjustable
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse) GetAdjustableRatio() *float64 {
	return s.AdjustableRatio
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse) GetTotal() *int64 {
	return s.Total
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse) SetAdjustable(v int64) *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse {
	s.Adjustable = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse) SetAdjustableRatio(v float64) *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse {
	s.AdjustableRatio = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse) SetTotal(v int64) *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse {
	s.Total = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewInUse) Validate() error {
	return dara.Validate(s)
}

type GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent struct {
	// example:
	//
	// 50
	Adjustable *int64 `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	// example:
	//
	// 50
	AdjustableRatio *float64 `json:"AdjustableRatio,omitempty" xml:"AdjustableRatio,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent) GoString() string {
	return s.String()
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent) GetAdjustable() *int64 {
	return s.Adjustable
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent) GetAdjustableRatio() *float64 {
	return s.AdjustableRatio
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent) GetTotal() *int64 {
	return s.Total
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent) SetAdjustable(v int64) *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent {
	s.Adjustable = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent) SetAdjustableRatio(v float64) *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent {
	s.AdjustableRatio = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent) SetTotal(v int64) *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent {
	s.Total = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver80Percent) Validate() error {
	return dara.Validate(s)
}

type GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent struct {
	// example:
	//
	// 50
	Adjustable *int64 `json:"Adjustable,omitempty" xml:"Adjustable,omitempty"`
	// example:
	//
	// 50
	AdjustableRatio *float64 `json:"AdjustableRatio,omitempty" xml:"AdjustableRatio,omitempty"`
	// example:
	//
	// 100
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent) String() string {
	return dara.Prettify(s)
}

func (s GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent) GoString() string {
	return s.String()
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent) GetAdjustable() *int64 {
	return s.Adjustable
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent) GetAdjustableRatio() *float64 {
	return s.AdjustableRatio
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent) GetTotal() *int64 {
	return s.Total
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent) SetAdjustable(v int64) *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent {
	s.Adjustable = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent) SetAdjustableRatio(v float64) *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent {
	s.AdjustableRatio = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent) SetTotal(v int64) *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent {
	s.Total = &v
	return s
}

func (s *GetQuotaUtilizationOverviewResponseBodyQuotaOverviewOver95Percent) Validate() error {
	return dara.Validate(s)
}
