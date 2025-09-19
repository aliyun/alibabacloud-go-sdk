// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateAutoTagRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateOrUpdateAutoTagRuleResponseBody
	GetRequestId() *string
}

type CreateOrUpdateAutoTagRuleResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 33DCC98C-824D-55D6-8DC5-47F3A71AD867
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateOrUpdateAutoTagRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateAutoTagRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateAutoTagRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrUpdateAutoTagRuleResponseBody) SetRequestId(v string) *CreateOrUpdateAutoTagRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrUpdateAutoTagRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
