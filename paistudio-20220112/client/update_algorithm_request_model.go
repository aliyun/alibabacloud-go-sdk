// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAlgorithmRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgorithmDescription(v string) *UpdateAlgorithmRequest
	GetAlgorithmDescription() *string
	SetDisplayName(v string) *UpdateAlgorithmRequest
	GetDisplayName() *string
}

type UpdateAlgorithmRequest struct {
	AlgorithmDescription *string `json:"AlgorithmDescription,omitempty" xml:"AlgorithmDescription,omitempty"`
	// example:
	//
	// LLM Train
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
}

func (s UpdateAlgorithmRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAlgorithmRequest) GoString() string {
	return s.String()
}

func (s *UpdateAlgorithmRequest) GetAlgorithmDescription() *string {
	return s.AlgorithmDescription
}

func (s *UpdateAlgorithmRequest) GetDisplayName() *string {
	return s.DisplayName
}

func (s *UpdateAlgorithmRequest) SetAlgorithmDescription(v string) *UpdateAlgorithmRequest {
	s.AlgorithmDescription = &v
	return s
}

func (s *UpdateAlgorithmRequest) SetDisplayName(v string) *UpdateAlgorithmRequest {
	s.DisplayName = &v
	return s
}

func (s *UpdateAlgorithmRequest) Validate() error {
	return dara.Validate(s)
}
