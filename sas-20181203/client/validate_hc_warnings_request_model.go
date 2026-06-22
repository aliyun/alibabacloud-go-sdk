// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iValidateHcWarningsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCheckIds(v string) *ValidateHcWarningsRequest
	GetCheckIds() *string
	SetRiskIds(v string) *ValidateHcWarningsRequest
	GetRiskIds() *string
	SetStatus(v int32) *ValidateHcWarningsRequest
	GetStatus() *int32
	SetUuids(v string) *ValidateHcWarningsRequest
	GetUuids() *string
}

type ValidateHcWarningsRequest struct {
	// The IDs of the check items. Separate multiple check item IDs with commas (,).
	//
	// > You can call the [DescribeCheckWarningSummary](https://help.aliyun.com/document_detail/116179.html) operation to obtain check item IDs.
	//
	// example:
	//
	// 695,234
	CheckIds *string `json:"CheckIds,omitempty" xml:"CheckIds,omitempty"`
	// The IDs of the risk items to verify. Separate multiple IDs with commas (,).
	//
	// >  You can call the [DescribeCheckWarnings](~~DescribeCheckWarnings~~) operation to obtain this parameter.
	//
	// example:
	//
	// 43
	RiskIds *string `json:"RiskIds,omitempty" xml:"RiskIds,omitempty"`
	// The status of the check item to verify. Valid values:
	//
	// - 1: not passed
	//
	// - 3: passed
	//
	// - 5: expired.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUIDs of the servers to verify. Separate multiple UUIDs with commas (,).
	//
	// >  You can call the [DescribeCloudCenterInstances](~~DescribeCloudCenterInstances~~) operation to obtain server UUIDs.
	//
	// example:
	//
	// 78645c8e-2e89-441b-8eb***,a9622a6b-adb5-4dd3-929e,0136460a-1cb5-44e8-****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s ValidateHcWarningsRequest) String() string {
	return dara.Prettify(s)
}

func (s ValidateHcWarningsRequest) GoString() string {
	return s.String()
}

func (s *ValidateHcWarningsRequest) GetCheckIds() *string {
	return s.CheckIds
}

func (s *ValidateHcWarningsRequest) GetRiskIds() *string {
	return s.RiskIds
}

func (s *ValidateHcWarningsRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ValidateHcWarningsRequest) GetUuids() *string {
	return s.Uuids
}

func (s *ValidateHcWarningsRequest) SetCheckIds(v string) *ValidateHcWarningsRequest {
	s.CheckIds = &v
	return s
}

func (s *ValidateHcWarningsRequest) SetRiskIds(v string) *ValidateHcWarningsRequest {
	s.RiskIds = &v
	return s
}

func (s *ValidateHcWarningsRequest) SetStatus(v int32) *ValidateHcWarningsRequest {
	s.Status = &v
	return s
}

func (s *ValidateHcWarningsRequest) SetUuids(v string) *ValidateHcWarningsRequest {
	s.Uuids = &v
	return s
}

func (s *ValidateHcWarningsRequest) Validate() error {
	return dara.Validate(s)
}
