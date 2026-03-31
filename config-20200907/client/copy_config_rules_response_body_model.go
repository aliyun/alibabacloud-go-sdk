// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyConfigRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCopyRulesResult(v bool) *CopyConfigRulesResponseBody
	GetCopyRulesResult() *bool
	SetRequestId(v string) *CopyConfigRulesResponseBody
	GetRequestId() *string
}

type CopyConfigRulesResponseBody struct {
	// Indicates whether the rules are replicated. Valid values:
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
	// 7BD81ECF-3194-5A6B-8719-9FC283167AD4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyConfigRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyConfigRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CopyConfigRulesResponseBody) GetCopyRulesResult() *bool {
	return s.CopyRulesResult
}

func (s *CopyConfigRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyConfigRulesResponseBody) SetCopyRulesResult(v bool) *CopyConfigRulesResponseBody {
	s.CopyRulesResult = &v
	return s
}

func (s *CopyConfigRulesResponseBody) SetRequestId(v string) *CopyConfigRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyConfigRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
