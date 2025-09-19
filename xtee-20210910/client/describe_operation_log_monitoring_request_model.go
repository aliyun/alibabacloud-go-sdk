// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperationLogMonitoringRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *DescribeOperationLogMonitoringRequest
	GetLang() *string
	SetEndDate(v string) *DescribeOperationLogMonitoringRequest
	GetEndDate() *string
	SetRegId(v string) *DescribeOperationLogMonitoringRequest
	GetRegId() *string
	SetStartDate(v string) *DescribeOperationLogMonitoringRequest
	GetStartDate() *string
	SetUserNameSearch(v string) *DescribeOperationLogMonitoringRequest
	GetUserNameSearch() *string
}

type DescribeOperationLogMonitoringRequest struct {
	// Language type of the returned message. Values:
	//
	// - **zh*	- (default): Chinese
	//
	// - **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// End date (in yyyy-MM-dd format, and the interval from the start date cannot exceed 90 days)
	//
	// example:
	//
	// 2025-07-30
	EndDate *string `json:"endDate,omitempty" xml:"endDate,omitempty"`
	// Region code.
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"regId,omitempty" xml:"regId,omitempty"`
	// Start date (in yyyy-MM-dd format, and the interval from the current date cannot exceed 90 days)
	//
	// example:
	//
	// 2025-07-19
	StartDate *string `json:"startDate,omitempty" xml:"startDate,omitempty"`
	// Operator.
	//
	// example:
	//
	// root
	UserNameSearch *string `json:"userNameSearch,omitempty" xml:"userNameSearch,omitempty"`
}

func (s DescribeOperationLogMonitoringRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperationLogMonitoringRequest) GoString() string {
	return s.String()
}

func (s *DescribeOperationLogMonitoringRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeOperationLogMonitoringRequest) GetEndDate() *string {
	return s.EndDate
}

func (s *DescribeOperationLogMonitoringRequest) GetRegId() *string {
	return s.RegId
}

func (s *DescribeOperationLogMonitoringRequest) GetStartDate() *string {
	return s.StartDate
}

func (s *DescribeOperationLogMonitoringRequest) GetUserNameSearch() *string {
	return s.UserNameSearch
}

func (s *DescribeOperationLogMonitoringRequest) SetLang(v string) *DescribeOperationLogMonitoringRequest {
	s.Lang = &v
	return s
}

func (s *DescribeOperationLogMonitoringRequest) SetEndDate(v string) *DescribeOperationLogMonitoringRequest {
	s.EndDate = &v
	return s
}

func (s *DescribeOperationLogMonitoringRequest) SetRegId(v string) *DescribeOperationLogMonitoringRequest {
	s.RegId = &v
	return s
}

func (s *DescribeOperationLogMonitoringRequest) SetStartDate(v string) *DescribeOperationLogMonitoringRequest {
	s.StartDate = &v
	return s
}

func (s *DescribeOperationLogMonitoringRequest) SetUserNameSearch(v string) *DescribeOperationLogMonitoringRequest {
	s.UserNameSearch = &v
	return s
}

func (s *DescribeOperationLogMonitoringRequest) Validate() error {
	return dara.Validate(s)
}
