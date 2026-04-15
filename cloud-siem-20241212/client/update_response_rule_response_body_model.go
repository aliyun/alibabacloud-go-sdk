// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResponseRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateResponseRuleResponseBody
	GetRequestId() *string
	SetResponseRuleId(v string) *UpdateResponseRuleResponseBody
	GetResponseRuleId() *string
}

type UpdateResponseRuleResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 440918
	ResponseRuleId *string `json:"ResponseRuleId,omitempty" xml:"ResponseRuleId,omitempty"`
}

func (s UpdateResponseRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateResponseRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateResponseRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateResponseRuleResponseBody) GetResponseRuleId() *string {
	return s.ResponseRuleId
}

func (s *UpdateResponseRuleResponseBody) SetRequestId(v string) *UpdateResponseRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateResponseRuleResponseBody) SetResponseRuleId(v string) *UpdateResponseRuleResponseBody {
	s.ResponseRuleId = &v
	return s
}

func (s *UpdateResponseRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
