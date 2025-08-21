// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebPreciseAccessRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebPreciseAccessRuleResponseBody
	GetRequestId() *string
}

type ModifyWebPreciseAccessRuleResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F908E959-ADA8-4D7B-8A05-FF2F67F50964
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebPreciseAccessRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebPreciseAccessRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebPreciseAccessRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebPreciseAccessRuleResponseBody) SetRequestId(v string) *ModifyWebPreciseAccessRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebPreciseAccessRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
