// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlgorithmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAlgorithmResponseBody
	GetRequestId() *string
}

type DeleteAlgorithmResponseBody struct {
	// example:
	//
	// FFB1D4B4-B253-540A-9B3B-AA711C48A1B7
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeleteAlgorithmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlgorithmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAlgorithmResponseBody) SetRequestId(v string) *DeleteAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlgorithmResponseBody) Validate() error {
	return dara.Validate(s)
}
