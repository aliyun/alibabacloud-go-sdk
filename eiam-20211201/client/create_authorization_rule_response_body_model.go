// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizationRuleId(v string) *CreateAuthorizationRuleResponseBody
	GetAuthorizationRuleId() *string
	SetRequestId(v string) *CreateAuthorizationRuleResponseBody
	GetRequestId() *string
}

type CreateAuthorizationRuleResponseBody struct {
	// example:
	//
	// arrule_01kf143ug06fg7m9f43u7vahxxxx
	AuthorizationRuleId *string `json:"AuthorizationRuleId,omitempty" xml:"AuthorizationRuleId,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAuthorizationRuleResponseBody) GetAuthorizationRuleId() *string {
	return s.AuthorizationRuleId
}

func (s *CreateAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAuthorizationRuleResponseBody) SetAuthorizationRuleId(v string) *CreateAuthorizationRuleResponseBody {
	s.AuthorizationRuleId = &v
	return s
}

func (s *CreateAuthorizationRuleResponseBody) SetRequestId(v string) *CreateAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAuthorizationRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
