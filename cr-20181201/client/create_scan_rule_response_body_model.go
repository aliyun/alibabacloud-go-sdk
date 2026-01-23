// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateScanRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *CreateScanRuleResponseBody
	GetCode() *int64
	SetRequestId(v string) *CreateScanRuleResponseBody
	GetRequestId() *string
	SetScanRuleId(v string) *CreateScanRuleResponseBody
	GetScanRuleId() *string
}

type CreateScanRuleResponseBody struct {
	// The returned HTTP or HTTPS status code.
	//
	// example:
	//
	// success
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Request Id
	//
	// example:
	//
	// EAEAB520-2456-5BF2-BCB5-AC998DFA3A51
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The rule ID.
	//
	// example:
	//
	// crscnr-k3gdm5vrn2nzf5hw
	ScanRuleId *string `json:"ScanRuleId,omitempty" xml:"ScanRuleId,omitempty"`
}

func (s CreateScanRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateScanRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateScanRuleResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *CreateScanRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateScanRuleResponseBody) GetScanRuleId() *string {
	return s.ScanRuleId
}

func (s *CreateScanRuleResponseBody) SetCode(v int64) *CreateScanRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateScanRuleResponseBody) SetRequestId(v string) *CreateScanRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateScanRuleResponseBody) SetScanRuleId(v string) *CreateScanRuleResponseBody {
	s.ScanRuleId = &v
	return s
}

func (s *CreateScanRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
