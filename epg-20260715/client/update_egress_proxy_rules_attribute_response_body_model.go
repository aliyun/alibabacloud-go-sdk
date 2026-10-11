// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyRulesAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateEgressProxyRulesAttributeResponseBody
	GetRequestId() *string
}

type UpdateEgressProxyRulesAttributeResponseBody struct {
	// example:
	//
	// 01A05AB5-5BED-56ED-8EE9-304FBC2F8537
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateEgressProxyRulesAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyRulesAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyRulesAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateEgressProxyRulesAttributeResponseBody) SetRequestId(v string) *UpdateEgressProxyRulesAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateEgressProxyRulesAttributeResponseBody) Validate() error {
	return dara.Validate(s)
}
