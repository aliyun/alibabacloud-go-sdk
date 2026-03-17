// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateQosResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssociateQosResponseBody
	GetRequestId() *string
}

type AssociateQosResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 551CD836-9E46-4F2C-A167-B4363180A647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AssociateQosResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssociateQosResponseBody) GoString() string {
	return s.String()
}

func (s *AssociateQosResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssociateQosResponseBody) SetRequestId(v string) *AssociateQosResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssociateQosResponseBody) Validate() error {
	return dara.Validate(s)
}
