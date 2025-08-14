// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTransitRouterPrefixListAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateTransitRouterPrefixListAssociationResponseBody
	GetRequestId() *string
}

type CreateTransitRouterPrefixListAssociationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0C2EE7A8-74D4-4081-8236-CEBDE3BBCF50
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateTransitRouterPrefixListAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateTransitRouterPrefixListAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransitRouterPrefixListAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateTransitRouterPrefixListAssociationResponseBody) SetRequestId(v string) *CreateTransitRouterPrefixListAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransitRouterPrefixListAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}
