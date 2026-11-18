// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIqsUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCallerId(v string) *GetIqsUsageRequest
	GetCallerId() *string
	SetEndDate(v string) *GetIqsUsageRequest
	GetEndDate() *string
	SetStartDate(v string) *GetIqsUsageRequest
	GetStartDate() *string
}

type GetIqsUsageRequest struct {
	// The ID of the Alibaba Cloud RAM user or the API key caller identifier to query.
	//
	// example:
	//
	// 138608930460****
	CallerId *string `json:"callerId,omitempty" xml:"callerId,omitempty"`
	// The end date for the usage query, inclusive. The format is YYYYMMDD. The date must be within the last 60 days.
	//
	// example:
	//
	// 20241017
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// The start date for the usage query, inclusive. The format is YYYYMMDD. The date must be within the last 60 days.
	//
	// example:
	//
	// 20241011
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s GetIqsUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIqsUsageRequest) GoString() string {
	return s.String()
}

func (s *GetIqsUsageRequest) GetCallerId() *string {
	return s.CallerId
}

func (s *GetIqsUsageRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *GetIqsUsageRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *GetIqsUsageRequest) SetCallerId(v string) *GetIqsUsageRequest {
	s.CallerId = &v
	return s
}

func (s *GetIqsUsageRequest) SetEndDate(v string) *GetIqsUsageRequest {
	s.EndDate = &v
	return s
}

func (s *GetIqsUsageRequest) SetStartDate(v string) *GetIqsUsageRequest {
	s.StartDate = &v
	return s
}

func (s *GetIqsUsageRequest) Validate() error {
	return dara.Validate(s)
}
