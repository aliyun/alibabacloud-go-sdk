// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferResourcesOutofGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransferResourcesOutofGroupResponseBody
	GetRequestId() *string
}

type TransferResourcesOutofGroupResponseBody struct {
	// example:
	//
	// D2B1DF9E-576B-5B91-BAF6-2C3DD1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TransferResourcesOutofGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferResourcesOutofGroupResponseBody) GoString() string {
	return s.String()
}

func (s *TransferResourcesOutofGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferResourcesOutofGroupResponseBody) SetRequestId(v string) *TransferResourcesOutofGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferResourcesOutofGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
