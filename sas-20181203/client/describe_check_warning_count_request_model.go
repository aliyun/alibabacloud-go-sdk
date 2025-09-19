// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCheckWarningCountRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAliUid(v int64) *DescribeCheckWarningCountRequest
	GetAliUid() *int64
	SetCheckId(v int64) *DescribeCheckWarningCountRequest
	GetCheckId() *int64
	SetRiskId(v int64) *DescribeCheckWarningCountRequest
	GetRiskId() *int64
	SetStatus(v int32) *DescribeCheckWarningCountRequest
	GetStatus() *int32
}

type DescribeCheckWarningCountRequest struct {
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 103784262032****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the check item.
	//
	// >  You can call the [ListCheckItemWarningSummary](~~ListCheckItemWarningSummary~~) operation to query the IDs of check items.
	//
	// example:
	//
	// 926
	CheckId *int64 `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	// The ID of the risk item.
	//
	// >  You can call the [DescribeCheckWarningSummary](~~DescribeCheckWarningSummary~~) operation to query the IDs of risk items.
	//
	// example:
	//
	// 43
	RiskId *int64 `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The status of the check item. Valid values:
	//
	// 	- **1**: failed
	//
	// 	- **2**: verifying
	//
	// 	- **3**: passed
	//
	// 	- **6**: ignored
	//
	// example:
	//
	// 3
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeCheckWarningCountRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCheckWarningCountRequest) GoString() string {
	return s.String()
}

func (s *DescribeCheckWarningCountRequest) GetAliUid() *int64 {
	return s.AliUid
}

func (s *DescribeCheckWarningCountRequest) GetCheckId() *int64 {
	return s.CheckId
}

func (s *DescribeCheckWarningCountRequest) GetRiskId() *int64 {
	return s.RiskId
}

func (s *DescribeCheckWarningCountRequest) GetStatus() *int32 {
	return s.Status
}

func (s *DescribeCheckWarningCountRequest) SetAliUid(v int64) *DescribeCheckWarningCountRequest {
	s.AliUid = &v
	return s
}

func (s *DescribeCheckWarningCountRequest) SetCheckId(v int64) *DescribeCheckWarningCountRequest {
	s.CheckId = &v
	return s
}

func (s *DescribeCheckWarningCountRequest) SetRiskId(v int64) *DescribeCheckWarningCountRequest {
	s.RiskId = &v
	return s
}

func (s *DescribeCheckWarningCountRequest) SetStatus(v int32) *DescribeCheckWarningCountRequest {
	s.Status = &v
	return s
}

func (s *DescribeCheckWarningCountRequest) Validate() error {
	return dara.Validate(s)
}
