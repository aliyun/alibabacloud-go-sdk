// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAlgorithmDefRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAlgoVersion(v string) *GetAlgorithmDefRequest
	GetAlgoVersion() *string
	SetIdentifier(v string) *GetAlgorithmDefRequest
	GetIdentifier() *string
	SetProvider(v string) *GetAlgorithmDefRequest
	GetProvider() *string
	SetSignature(v string) *GetAlgorithmDefRequest
	GetSignature() *string
}

type GetAlgorithmDefRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// v1
	AlgoVersion *string `json:"AlgoVersion,omitempty" xml:"AlgoVersion,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// logisticregression_binary
	Identifier *string `json:"Identifier,omitempty" xml:"Identifier,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pai
	Provider *string `json:"Provider,omitempty" xml:"Provider,omitempty"`
	// example:
	//
	// 5vqe4Sgtzw8E6opyK3HkK+nzYlY=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
}

func (s GetAlgorithmDefRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAlgorithmDefRequest) GoString() string {
	return s.String()
}

func (s *GetAlgorithmDefRequest) GetAlgoVersion() *string {
	return s.AlgoVersion
}

func (s *GetAlgorithmDefRequest) GetIdentifier() *string {
	return s.Identifier
}

func (s *GetAlgorithmDefRequest) GetProvider() *string {
	return s.Provider
}

func (s *GetAlgorithmDefRequest) GetSignature() *string {
	return s.Signature
}

func (s *GetAlgorithmDefRequest) SetAlgoVersion(v string) *GetAlgorithmDefRequest {
	s.AlgoVersion = &v
	return s
}

func (s *GetAlgorithmDefRequest) SetIdentifier(v string) *GetAlgorithmDefRequest {
	s.Identifier = &v
	return s
}

func (s *GetAlgorithmDefRequest) SetProvider(v string) *GetAlgorithmDefRequest {
	s.Provider = &v
	return s
}

func (s *GetAlgorithmDefRequest) SetSignature(v string) *GetAlgorithmDefRequest {
	s.Signature = &v
	return s
}

func (s *GetAlgorithmDefRequest) Validate() error {
	return dara.Validate(s)
}
