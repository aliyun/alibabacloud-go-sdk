// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebRuleResponseBody
	GetRequestId() *string
}

type ModifyWebRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CB3261D2-7D1B-4ADA-9E98-A200B2CDA2DC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebRuleResponseBody) SetRequestId(v string) *ModifyWebRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
