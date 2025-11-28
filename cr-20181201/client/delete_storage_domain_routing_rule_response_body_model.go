// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageDomainRoutingRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteStorageDomainRoutingRuleResponseBody
	GetCode() *string
	SetRequestId(v string) *DeleteStorageDomainRoutingRuleResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteStorageDomainRoutingRuleResponseBody
	GetSuccess() *bool
}

type DeleteStorageDomainRoutingRuleResponseBody struct {
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// D4978DCC-ECBD-40B0-A714-EE6959B*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteStorageDomainRoutingRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageDomainRoutingRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStorageDomainRoutingRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteStorageDomainRoutingRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStorageDomainRoutingRuleResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteStorageDomainRoutingRuleResponseBody) SetCode(v string) *DeleteStorageDomainRoutingRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteStorageDomainRoutingRuleResponseBody) SetRequestId(v string) *DeleteStorageDomainRoutingRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStorageDomainRoutingRuleResponseBody) SetSuccess(v bool) *DeleteStorageDomainRoutingRuleResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteStorageDomainRoutingRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
