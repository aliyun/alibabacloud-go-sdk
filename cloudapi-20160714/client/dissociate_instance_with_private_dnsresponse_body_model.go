// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateInstanceWithPrivateDNSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DissociateInstanceWithPrivateDNSResponseBody
	GetRequestId() *string
}

type DissociateInstanceWithPrivateDNSResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6C87A26A-6A18-4B8E-8099-705278381A2C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DissociateInstanceWithPrivateDNSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DissociateInstanceWithPrivateDNSResponseBody) GoString() string {
	return s.String()
}

func (s *DissociateInstanceWithPrivateDNSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DissociateInstanceWithPrivateDNSResponseBody) SetRequestId(v string) *DissociateInstanceWithPrivateDNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *DissociateInstanceWithPrivateDNSResponseBody) Validate() error {
	return dara.Validate(s)
}
