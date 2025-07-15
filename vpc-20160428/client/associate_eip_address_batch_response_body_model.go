// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateEipAddressBatchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateEipAddressBatchResponseBody
	GetRequestId() *string
}

type AssociateEipAddressBatchResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateEipAddressBatchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateEipAddressBatchResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateEipAddressBatchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateEipAddressBatchResponseBody) SetRequestId(v string) *AssociateEipAddressBatchResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateEipAddressBatchResponseBody) Validate() error {
	return dara.Validate(s)
}
