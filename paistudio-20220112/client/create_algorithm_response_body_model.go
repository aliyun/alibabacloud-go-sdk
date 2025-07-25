// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlgorithmResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmId(v string) *CreateAlgorithmResponseBody
	GetAlgorithmId() *string
	SetRequestId(v string) *CreateAlgorithmResponseBody
	GetRequestId() *string
}

type CreateAlgorithmResponseBody struct {
	// example:
	//
	// algo-xsldfvu1334
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// E7C42CC7-2E85-508A-84F4-923B605FD10F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAlgorithmResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlgorithmResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmResponseBody) GetAlgorithmId() *string {
	return s.AlgorithmId
}

func (s *CreateAlgorithmResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlgorithmResponseBody) SetAlgorithmId(v string) *CreateAlgorithmResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *CreateAlgorithmResponseBody) SetRequestId(v string) *CreateAlgorithmResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlgorithmResponseBody) Validate() error {
	return dara.Validate(s)
}
