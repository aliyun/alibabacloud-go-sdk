// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttackedAssetDealRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDateType(v string) *GetAttackedAssetDealRequest
	GetDateType() *string
	SetEndDate(v int64) *GetAttackedAssetDealRequest
	GetEndDate() *int64
	SetStartDate(v int64) *GetAttackedAssetDealRequest
	GetStartDate() *int64
	SetSuspEventSource(v string) *GetAttackedAssetDealRequest
	GetSuspEventSource() *string
}

type GetAttackedAssetDealRequest struct {
	// Time filter type, supporting filtering by the last 7 days, the last 30 days, the last half year, or custom time periods.
	//
	// This parameter is required.
	//
	// example:
	//
	// month
	DateType *string `json:"DateType,omitempty" xml:"DateType,omitempty"`
	// End time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732268720000
	EndDate *int64 `json:"EndDate,omitempty" xml:"EndDate,omitempty"`
	// Start time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1732268720000
	StartDate *int64 `json:"StartDate,omitempty" xml:"StartDate,omitempty"`
	// Source of the alert event.
	//
	// example:
	//
	// 暂时无需传参，有问题请联系管理员
	SuspEventSource *string `json:"SuspEventSource,omitempty" xml:"SuspEventSource,omitempty"`
}

func (s GetAttackedAssetDealRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAttackedAssetDealRequest) GoString() string {
	return s.String()
}

func (s *GetAttackedAssetDealRequest) GetDateType() *string {
	return s.DateType
}

func (s *GetAttackedAssetDealRequest) GetEndDate() *int64 {
	return s.EndDate
}

func (s *GetAttackedAssetDealRequest) GetStartDate() *int64 {
	return s.StartDate
}

func (s *GetAttackedAssetDealRequest) GetSuspEventSource() *string {
	return s.SuspEventSource
}

func (s *GetAttackedAssetDealRequest) SetDateType(v string) *GetAttackedAssetDealRequest {
	s.DateType = &v
	return s
}

func (s *GetAttackedAssetDealRequest) SetEndDate(v int64) *GetAttackedAssetDealRequest {
	s.EndDate = &v
	return s
}

func (s *GetAttackedAssetDealRequest) SetStartDate(v int64) *GetAttackedAssetDealRequest {
	s.StartDate = &v
	return s
}

func (s *GetAttackedAssetDealRequest) SetSuspEventSource(v string) *GetAttackedAssetDealRequest {
	s.SuspEventSource = &v
	return s
}

func (s *GetAttackedAssetDealRequest) Validate() error {
	return dara.Validate(s)
}
