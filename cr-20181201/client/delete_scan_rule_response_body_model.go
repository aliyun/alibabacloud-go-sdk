// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScanRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteScanRuleResponseBody
	GetCode() *string
	SetRequestId(v string) *DeleteScanRuleResponseBody
	GetRequestId() *string
}

type DeleteScanRuleResponseBody struct {
	// The HTTP status code
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Request Id
	//
	// example:
	//
	// 12589EF7-96E2-4554-AAD7-F7209E88CAD3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScanRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScanRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScanRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteScanRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScanRuleResponseBody) SetCode(v string) *DeleteScanRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteScanRuleResponseBody) SetRequestId(v string) *DeleteScanRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScanRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
