// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlgorithmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmId(v string) *UpdateAlgorithmResponseBody
	GetAlgorithmId() *string
	SetRequestId(v string) *UpdateAlgorithmResponseBody
	GetRequestId() *string
}

type UpdateAlgorithmResponseBody struct {
	// example:
	//
	// algo-sidjc8134hv
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// E7C42CC7-2E85-508A-84F4-923B605FD10F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAlgorithmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmResponseBody) GetAlgorithmId() *string {
	return s.AlgorithmId
}

func (s *UpdateAlgorithmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAlgorithmResponseBody) SetAlgorithmId(v string) *UpdateAlgorithmResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *UpdateAlgorithmResponseBody) SetRequestId(v string) *UpdateAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAlgorithmResponseBody) Validate() error {
	return dara.Validate(s)
}
