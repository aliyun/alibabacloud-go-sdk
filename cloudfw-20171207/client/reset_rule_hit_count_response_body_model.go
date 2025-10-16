// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetRuleHitCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v string) *ResetRuleHitCountResponseBody
	GetModule() *string
	SetRequestId(v string) *ResetRuleHitCountResponseBody
	GetRequestId() *string
}

type ResetRuleHitCountResponseBody struct {
	// example:
	//
	// sg_server
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// 706B2093-CBA0-51B2-BEBF-58903FC6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResetRuleHitCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResetRuleHitCountResponseBody) GoString() string {
	return s.String()
}

func (s *ResetRuleHitCountResponseBody) GetModule() *string {
	return s.Module
}

func (s *ResetRuleHitCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResetRuleHitCountResponseBody) SetModule(v string) *ResetRuleHitCountResponseBody {
	s.Module = &v
	return s
}

func (s *ResetRuleHitCountResponseBody) SetRequestId(v string) *ResetRuleHitCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResetRuleHitCountResponseBody) Validate() error {
	return dara.Validate(s)
}
