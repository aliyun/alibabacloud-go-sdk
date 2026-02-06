// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomResponseCodeRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateCustomResponseCodeRuleResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateCustomResponseCodeRuleResponseBody
	GetRequestId() *string
}

type CreateCustomResponseCodeRuleResponseBody struct {
	// The configuration ID.
	//
	// example:
	//
	// 3528160969****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// C2B2F8CF-3074-5BBC-891A-AAD292E2624F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCustomResponseCodeRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomResponseCodeRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomResponseCodeRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateCustomResponseCodeRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCustomResponseCodeRuleResponseBody) SetConfigId(v int64) *CreateCustomResponseCodeRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateCustomResponseCodeRuleResponseBody) SetRequestId(v string) *CreateCustomResponseCodeRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCustomResponseCodeRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
