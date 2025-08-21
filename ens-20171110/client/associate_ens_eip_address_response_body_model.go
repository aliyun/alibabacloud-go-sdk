// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEnsEipAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateEnsEipAddressResponseBody
	GetRequestId() *string
}

type AssociateEnsEipAddressResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateEnsEipAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateEnsEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateEnsEipAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateEnsEipAddressResponseBody) SetRequestId(v string) *AssociateEnsEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateEnsEipAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
