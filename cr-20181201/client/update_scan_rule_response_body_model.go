// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScanRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *UpdateScanRuleResponseBody
	GetCode() *int64
	SetRequestId(v string) *UpdateScanRuleResponseBody
	GetRequestId() *string
	SetScanRuleId(v string) *UpdateScanRuleResponseBody
	GetScanRuleId() *string
}

type UpdateScanRuleResponseBody struct {
	// example:
	//
	// success
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Request Id
	//
	// example:
	//
	// 33EF1695-E2B7-5D57-B4E1-3D655FE1EBD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// crscnr-3qmkeiuggfpjkfrq
	ScanRuleId *string `json:"ScanRuleId,omitempty" xml:"ScanRuleId,omitempty"`
}

func (s UpdateScanRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateScanRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateScanRuleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *UpdateScanRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateScanRuleResponseBody) GetScanRuleId() *string {
	return s.ScanRuleId
}

func (s *UpdateScanRuleResponseBody) SetCode(v int64) *UpdateScanRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateScanRuleResponseBody) SetRequestId(v string) *UpdateScanRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateScanRuleResponseBody) SetScanRuleId(v string) *UpdateScanRuleResponseBody {
	s.ScanRuleId = &v
	return s
}

func (s *UpdateScanRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
