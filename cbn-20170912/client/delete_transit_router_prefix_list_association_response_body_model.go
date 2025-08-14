// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPrefixListAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTransitRouterPrefixListAssociationResponseBody
	GetRequestId() *string
}

type DeleteTransitRouterPrefixListAssociationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 835E7F4B-B380-4E0F-96A5-6EA572388047
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTransitRouterPrefixListAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPrefixListAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPrefixListAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTransitRouterPrefixListAssociationResponseBody) SetRequestId(v string) *DeleteTransitRouterPrefixListAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTransitRouterPrefixListAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}
