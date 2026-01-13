// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrivateRAGServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePrivateRAGServiceResponseBody
	GetRequestId() *string
}

type DeletePrivateRAGServiceResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePrivateRAGServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePrivateRAGServiceResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePrivateRAGServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePrivateRAGServiceResponseBody) SetRequestId(v string) *DeletePrivateRAGServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePrivateRAGServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
