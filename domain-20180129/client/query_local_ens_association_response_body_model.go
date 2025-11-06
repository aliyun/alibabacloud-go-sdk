// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryLocalEnsAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *QueryLocalEnsAssociationResponseBody
	GetAddress() *string
	SetRequestId(v string) *QueryLocalEnsAssociationResponseBody
	GetRequestId() *string
}

type QueryLocalEnsAssociationResponseBody struct {
	// example:
	//
	// 3ECD5439-39A2-477D-9A19-64FCA1F77EEB
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 0x1234567890123456789012345678901234567890
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryLocalEnsAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryLocalEnsAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryLocalEnsAssociationResponseBody) GetAddress() *string {
	return s.Address
}

func (s *QueryLocalEnsAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryLocalEnsAssociationResponseBody) SetAddress(v string) *QueryLocalEnsAssociationResponseBody {
	s.Address = &v
	return s
}

func (s *QueryLocalEnsAssociationResponseBody) SetRequestId(v string) *QueryLocalEnsAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryLocalEnsAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}
