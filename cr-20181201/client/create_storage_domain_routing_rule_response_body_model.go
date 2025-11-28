// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStorageDomainRoutingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateStorageDomainRoutingRuleResponseBody
	GetCode() *string
	SetRequestId(v string) *CreateStorageDomainRoutingRuleResponseBody
	GetRequestId() *string
	SetRuleId(v string) *CreateStorageDomainRoutingRuleResponseBody
	GetRuleId() *string
	SetSuccess(v bool) *CreateStorageDomainRoutingRuleResponseBody
	GetSuccess() *bool
}

type CreateStorageDomainRoutingRuleResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// C8E90AB5-0A96-5D12-9E59-11EE463*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// crsdr-n6pbhgjx*****
	RuleId *string `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStorageDomainRoutingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStorageDomainRoutingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStorageDomainRoutingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateStorageDomainRoutingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStorageDomainRoutingRuleResponseBody) GetRuleId() *string {
	return s.RuleId
}

func (s *CreateStorageDomainRoutingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStorageDomainRoutingRuleResponseBody) SetCode(v string) *CreateStorageDomainRoutingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStorageDomainRoutingRuleResponseBody) SetRequestId(v string) *CreateStorageDomainRoutingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStorageDomainRoutingRuleResponseBody) SetRuleId(v string) *CreateStorageDomainRoutingRuleResponseBody {
	s.RuleId = &v
	return s
}

func (s *CreateStorageDomainRoutingRuleResponseBody) SetSuccess(v bool) *CreateStorageDomainRoutingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStorageDomainRoutingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
