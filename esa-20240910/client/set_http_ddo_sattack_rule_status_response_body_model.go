// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHttpDDoSAttackRuleStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetHttpDDoSAttackRuleStatusResponseBody
	GetRequestId() *string
}

type SetHttpDDoSAttackRuleStatusResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// F7B84CF8-F8A4-53F8-9B91-2643FD72042B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetHttpDDoSAttackRuleStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetHttpDDoSAttackRuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackRuleStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetHttpDDoSAttackRuleStatusResponseBody) SetRequestId(v string) *SetHttpDDoSAttackRuleStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetHttpDDoSAttackRuleStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
