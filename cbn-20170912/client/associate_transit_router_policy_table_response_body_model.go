// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateTransitRouterPolicyTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAssociationId(v string) *AssociateTransitRouterPolicyTableResponseBody
	GetAssociationId() *string
	SetRequestId(v string) *AssociateTransitRouterPolicyTableResponseBody
	GetRequestId() *string
}

type AssociateTransitRouterPolicyTableResponseBody struct {
	// The ID of the association between the policy route table and the network instance connection.
	//
	// example:
	//
	// pta-123asfzc
	AssociationId *string `json:"AssociationId,omitempty" xml:"AssociationId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ADD98358-D265-4060-87CB-A2427F5A8944
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateTransitRouterPolicyTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateTransitRouterPolicyTableResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateTransitRouterPolicyTableResponseBody) GetAssociationId() *string {
	return s.AssociationId
}

func (s *AssociateTransitRouterPolicyTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateTransitRouterPolicyTableResponseBody) SetAssociationId(v string) *AssociateTransitRouterPolicyTableResponseBody {
	s.AssociationId = &v
	return s
}

func (s *AssociateTransitRouterPolicyTableResponseBody) SetRequestId(v string) *AssociateTransitRouterPolicyTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateTransitRouterPolicyTableResponseBody) Validate() error {
	return dara.Validate(s)
}
