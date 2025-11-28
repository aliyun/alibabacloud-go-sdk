// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStorageDomainRoutingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateStorageDomainRoutingRuleResponseBody
	GetCode() *string
	SetRequestId(v string) *UpdateStorageDomainRoutingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateStorageDomainRoutingRuleResponseBody
	GetSuccess() *bool
}

type UpdateStorageDomainRoutingRuleResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// D4978DCC-ECBD-40B0-A714-EE6959*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateStorageDomainRoutingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateStorageDomainRoutingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStorageDomainRoutingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateStorageDomainRoutingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateStorageDomainRoutingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateStorageDomainRoutingRuleResponseBody) SetCode(v string) *UpdateStorageDomainRoutingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateStorageDomainRoutingRuleResponseBody) SetRequestId(v string) *UpdateStorageDomainRoutingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateStorageDomainRoutingRuleResponseBody) SetSuccess(v bool) *UpdateStorageDomainRoutingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateStorageDomainRoutingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
