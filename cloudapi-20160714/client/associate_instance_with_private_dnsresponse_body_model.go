// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateInstanceWithPrivateDNSResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateInstanceWithPrivateDNSResponseBody
	GetRequestId() *string
}

type AssociateInstanceWithPrivateDNSResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 03442A3D-3B7D-434C-8A95-A5FEB999B529
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateInstanceWithPrivateDNSResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateInstanceWithPrivateDNSResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateInstanceWithPrivateDNSResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateInstanceWithPrivateDNSResponseBody) SetRequestId(v string) *AssociateInstanceWithPrivateDNSResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateInstanceWithPrivateDNSResponseBody) Validate() error {
	return dara.Validate(s)
}
