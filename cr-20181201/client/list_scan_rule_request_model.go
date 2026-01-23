// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListScanRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListScanRuleRequest
	GetInstanceId() *string
	SetPageNo(v int32) *ListScanRuleRequest
	GetPageNo() *int32
	SetPageSize(v int32) *ListScanRuleRequest
	GetPageSize() *int32
	SetScanType(v string) *ListScanRuleRequest
	GetScanType() *string
}

type ListScanRuleRequest struct {
	// The instance ID.
	//
	// example:
	//
	// cri-upoulewerx*****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- `VUL`: Products Cloud Security Scanner.
	//
	// 	- `SBOM`: Product Content Analysis.
	//
	// Default value: `VUL`
	//
	// example:
	//
	// SBOM
	ScanType *string `json:"ScanType,omitempty" xml:"ScanType,omitempty"`
}

func (s ListScanRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ListScanRuleRequest) GoString() string {
	return s.String()
}

func (s *ListScanRuleRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListScanRuleRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *ListScanRuleRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListScanRuleRequest) GetScanType() *string {
	return s.ScanType
}

func (s *ListScanRuleRequest) SetInstanceId(v string) *ListScanRuleRequest {
	s.InstanceId = &v
	return s
}

func (s *ListScanRuleRequest) SetPageNo(v int32) *ListScanRuleRequest {
	s.PageNo = &v
	return s
}

func (s *ListScanRuleRequest) SetPageSize(v int32) *ListScanRuleRequest {
	s.PageSize = &v
	return s
}

func (s *ListScanRuleRequest) SetScanType(v string) *ListScanRuleRequest {
	s.ScanType = &v
	return s
}

func (s *ListScanRuleRequest) Validate() error {
	return dara.Validate(s)
}
