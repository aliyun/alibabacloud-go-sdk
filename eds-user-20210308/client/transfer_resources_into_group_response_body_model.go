// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferResourcesIntoGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransferResourcesIntoGroupResponseBody
	GetRequestId() *string
}

type TransferResourcesIntoGroupResponseBody struct {
	// example:
	//
	// AA8D67CB-345D-5CDA-986E-FFAC7D0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferResourcesIntoGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferResourcesIntoGroupResponseBody) GoString() string {
	return s.String()
}

func (s *TransferResourcesIntoGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferResourcesIntoGroupResponseBody) SetRequestId(v string) *TransferResourcesIntoGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferResourcesIntoGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
