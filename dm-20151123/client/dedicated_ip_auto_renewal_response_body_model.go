// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDedicatedIpAutoRenewalResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DedicatedIpAutoRenewalResponseBody
	GetRequestId() *string
}

type DedicatedIpAutoRenewalResponseBody struct {
	// Request ID
	//
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DedicatedIpAutoRenewalResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DedicatedIpAutoRenewalResponseBody) GoString() string {
	return s.String()
}

func (s *DedicatedIpAutoRenewalResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DedicatedIpAutoRenewalResponseBody) SetRequestId(v string) *DedicatedIpAutoRenewalResponseBody {
	s.RequestId = &v
	return s
}

func (s *DedicatedIpAutoRenewalResponseBody) Validate() error {
	return dara.Validate(s)
}
