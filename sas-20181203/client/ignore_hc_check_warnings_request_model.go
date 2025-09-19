// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIgnoreHcCheckWarningsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckIds(v string) *IgnoreHcCheckWarningsRequest
	GetCheckIds() *string
	SetCheckWarningIds(v string) *IgnoreHcCheckWarningsRequest
	GetCheckWarningIds() *string
	SetReason(v string) *IgnoreHcCheckWarningsRequest
	GetReason() *string
	SetRiskId(v string) *IgnoreHcCheckWarningsRequest
	GetRiskId() *string
	SetSource(v string) *IgnoreHcCheckWarningsRequest
	GetSource() *string
	SetSourceIp(v string) *IgnoreHcCheckWarningsRequest
	GetSourceIp() *string
	SetType(v int64) *IgnoreHcCheckWarningsRequest
	GetType() *int64
}

type IgnoreHcCheckWarningsRequest struct {
	// The ID of the check item.
	//
	// >  You can call the [DescribeCheckWarnings](https://help.aliyun.com/document_detail/116182.html) operation to query the IDs of check items.
	//
	// example:
	//
	// 21313
	CheckIds *string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty"`
	// The ID of the alert that is triggered by the check item. Separate multiple IDs with commas (,).
	//
	// >  You can call the [DescribeCheckWarnings](https://help.aliyun.com/document_detail/116182.html) operation to query the IDs of alerts that are triggered by check items.
	//
	// example:
	//
	// 98146905,98146907
	CheckWarningIds *string `json:"CheckWarningIds,omitempty" xml:"CheckWarningIds,omitempty"`
	// The reason for the current operation.
	//
	// example:
	//
	// ignore
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The ID of the risk item that you want to ignore or cancel ignoring.
	//
	// >  You can call the [DescribeCheckWarningSummary](https://help.aliyun.com/document_detail/116179.html) operation to query the IDs of risk items.
	//
	// example:
	//
	// 51
	RiskId *string `json:"RiskId,omitempty" xml:"RiskId,omitempty"`
	// The data source. If this parameter is left empty, the server baseline results are queried by default. Valid values:
	//
	// 	- **default**: server
	//
	// 	- **agentless**
	//
	// example:
	//
	// agentless
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// The type of the operation that you want to perform. Valid values:
	//
	// 	- **1**: ignores a risk item
	//
	// 	- **2**: cancels ignoring a risk item
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Type *int64 `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s IgnoreHcCheckWarningsRequest) String() string {
	return dara.Prettify(s)
}

func (s IgnoreHcCheckWarningsRequest) GoString() string {
	return s.String()
}

func (s *IgnoreHcCheckWarningsRequest) GetCheckIds() *string {
	return s.CheckIds
}

func (s *IgnoreHcCheckWarningsRequest) GetCheckWarningIds() *string {
	return s.CheckWarningIds
}

func (s *IgnoreHcCheckWarningsRequest) GetReason() *string {
	return s.Reason
}

func (s *IgnoreHcCheckWarningsRequest) GetRiskId() *string {
	return s.RiskId
}

func (s *IgnoreHcCheckWarningsRequest) GetSource() *string {
	return s.Source
}

func (s *IgnoreHcCheckWarningsRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *IgnoreHcCheckWarningsRequest) GetType() *int64 {
	return s.Type
}

func (s *IgnoreHcCheckWarningsRequest) SetCheckIds(v string) *IgnoreHcCheckWarningsRequest {
	s.CheckIds = &v
	return s
}

func (s *IgnoreHcCheckWarningsRequest) SetCheckWarningIds(v string) *IgnoreHcCheckWarningsRequest {
	s.CheckWarningIds = &v
	return s
}

func (s *IgnoreHcCheckWarningsRequest) SetReason(v string) *IgnoreHcCheckWarningsRequest {
	s.Reason = &v
	return s
}

func (s *IgnoreHcCheckWarningsRequest) SetRiskId(v string) *IgnoreHcCheckWarningsRequest {
	s.RiskId = &v
	return s
}

func (s *IgnoreHcCheckWarningsRequest) SetSource(v string) *IgnoreHcCheckWarningsRequest {
	s.Source = &v
	return s
}

func (s *IgnoreHcCheckWarningsRequest) SetSourceIp(v string) *IgnoreHcCheckWarningsRequest {
	s.SourceIp = &v
	return s
}

func (s *IgnoreHcCheckWarningsRequest) SetType(v int64) *IgnoreHcCheckWarningsRequest {
	s.Type = &v
	return s
}

func (s *IgnoreHcCheckWarningsRequest) Validate() error {
	return dara.Validate(s)
}
