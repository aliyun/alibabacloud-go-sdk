// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnAssociateEnsEipAddressResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnAssociateEnsEipAddressResponseBody
	GetRequestId() *string
}

type UnAssociateEnsEipAddressResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 4A431388-2D4B-46F4-A96B-D4E6BD0688C1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnAssociateEnsEipAddressResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnAssociateEnsEipAddressResponseBody) GoString() string {
	return s.String()
}

func (s *UnAssociateEnsEipAddressResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnAssociateEnsEipAddressResponseBody) SetRequestId(v string) *UnAssociateEnsEipAddressResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnAssociateEnsEipAddressResponseBody) Validate() error {
	return dara.Validate(s)
}
