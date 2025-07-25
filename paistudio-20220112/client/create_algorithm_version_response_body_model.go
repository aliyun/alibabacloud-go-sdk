// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlgorithmVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmId(v string) *CreateAlgorithmVersionResponseBody
	GetAlgorithmId() *string
	SetAlgorithmVersion(v string) *CreateAlgorithmVersionResponseBody
	GetAlgorithmVersion() *string
}

type CreateAlgorithmVersionResponseBody struct {
	// example:
	//
	// algo-xsldfvu1334
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// v0.0.1
	AlgorithmVersion *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
}

func (s CreateAlgorithmVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlgorithmVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlgorithmVersionResponseBody) GetAlgorithmId() *string {
	return s.AlgorithmId
}

func (s *CreateAlgorithmVersionResponseBody) GetAlgorithmVersion() *string {
	return s.AlgorithmVersion
}

func (s *CreateAlgorithmVersionResponseBody) SetAlgorithmId(v string) *CreateAlgorithmVersionResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *CreateAlgorithmVersionResponseBody) SetAlgorithmVersion(v string) *CreateAlgorithmVersionResponseBody {
	s.AlgorithmVersion = &v
	return s
}

func (s *CreateAlgorithmVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
