// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyCompliancePacksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCopyRulesResult(v bool) *CopyCompliancePacksResponseBody
	GetCopyRulesResult() *bool
	SetRequestId(v string) *CopyCompliancePacksResponseBody
	GetRequestId() *string
}

type CopyCompliancePacksResponseBody struct {
	// Indicates whether the compliance packages are replicated. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	CopyRulesResult *bool `json:"CopyRulesResult,omitempty" xml:"CopyRulesResult,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9E1E69DE-BDED-581E-B559-0C15690901D0
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyCompliancePacksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyCompliancePacksResponseBody) GoString() string {
	return s.String()
}

func (s *CopyCompliancePacksResponseBody) GetCopyRulesResult() *bool {
	return s.CopyRulesResult
}

func (s *CopyCompliancePacksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyCompliancePacksResponseBody) SetCopyRulesResult(v bool) *CopyCompliancePacksResponseBody {
	s.CopyRulesResult = &v
	return s
}

func (s *CopyCompliancePacksResponseBody) SetRequestId(v string) *CopyCompliancePacksResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyCompliancePacksResponseBody) Validate() error {
	return dara.Validate(s)
}
