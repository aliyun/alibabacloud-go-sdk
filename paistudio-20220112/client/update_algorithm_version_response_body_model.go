// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlgorithmVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmId(v string) *UpdateAlgorithmVersionResponseBody
	GetAlgorithmId() *string
	SetAlgorithmVersion(v string) *UpdateAlgorithmVersionResponseBody
	GetAlgorithmVersion() *string
}

type UpdateAlgorithmVersionResponseBody struct {
	// example:
	//
	// algo-sidjc8134hv
	AlgorithmId *string `json:"AlgorithmId,omitempty" xml:"AlgorithmId,omitempty"`
	// example:
	//
	// v0.1.0
	AlgorithmVersion *string `json:"AlgorithmVersion,omitempty" xml:"AlgorithmVersion,omitempty"`
}

func (s UpdateAlgorithmVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlgorithmVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmVersionResponseBody) GetAlgorithmId() *string {
	return s.AlgorithmId
}

func (s *UpdateAlgorithmVersionResponseBody) GetAlgorithmVersion() *string {
	return s.AlgorithmVersion
}

func (s *UpdateAlgorithmVersionResponseBody) SetAlgorithmId(v string) *UpdateAlgorithmVersionResponseBody {
	s.AlgorithmId = &v
	return s
}

func (s *UpdateAlgorithmVersionResponseBody) SetAlgorithmVersion(v string) *UpdateAlgorithmVersionResponseBody {
	s.AlgorithmVersion = &v
	return s
}

func (s *UpdateAlgorithmVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
