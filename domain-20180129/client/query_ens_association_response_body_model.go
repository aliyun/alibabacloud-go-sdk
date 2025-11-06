// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEnsAssociationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddress(v string) *QueryEnsAssociationResponseBody
	GetAddress() *string
	SetRequestId(v string) *QueryEnsAssociationResponseBody
	GetRequestId() *string
}

type QueryEnsAssociationResponseBody struct {
	// example:
	//
	// 0x123456789012345678901234567890123456****
	Address *string `json:"Address,omitempty" xml:"Address,omitempty"`
	// example:
	//
	// 3ECD5439-39A2-477D-9A19-64FCA1F77EEB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryEnsAssociationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryEnsAssociationResponseBody) GoString() string {
	return s.String()
}

func (s *QueryEnsAssociationResponseBody) GetAddress() *string {
	return s.Address
}

func (s *QueryEnsAssociationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryEnsAssociationResponseBody) SetAddress(v string) *QueryEnsAssociationResponseBody {
	s.Address = &v
	return s
}

func (s *QueryEnsAssociationResponseBody) SetRequestId(v string) *QueryEnsAssociationResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryEnsAssociationResponseBody) Validate() error {
	return dara.Validate(s)
}
