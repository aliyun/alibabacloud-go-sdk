// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociatePrincipalWithPortfolioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPortfolioId(v string) *AssociatePrincipalWithPortfolioRequest
	GetPortfolioId() *string
	SetPrincipalId(v string) *AssociatePrincipalWithPortfolioRequest
	GetPrincipalId() *string
	SetPrincipalPattern(v string) *AssociatePrincipalWithPortfolioRequest
	GetPrincipalPattern() *string
	SetPrincipalType(v string) *AssociatePrincipalWithPortfolioRequest
	GetPrincipalType() *string
}

type AssociatePrincipalWithPortfolioRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 24477111603637****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// example:
	//
	// user/userName
	PrincipalPattern *string `json:"PrincipalPattern,omitempty" xml:"PrincipalPattern,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// RamUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s AssociatePrincipalWithPortfolioRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociatePrincipalWithPortfolioRequest) GoString() string {
	return s.String()
}

func (s *AssociatePrincipalWithPortfolioRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *AssociatePrincipalWithPortfolioRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *AssociatePrincipalWithPortfolioRequest) GetPrincipalPattern() *string {
	return s.PrincipalPattern
}

func (s *AssociatePrincipalWithPortfolioRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *AssociatePrincipalWithPortfolioRequest) SetPortfolioId(v string) *AssociatePrincipalWithPortfolioRequest {
	s.PortfolioId = &v
	return s
}

func (s *AssociatePrincipalWithPortfolioRequest) SetPrincipalId(v string) *AssociatePrincipalWithPortfolioRequest {
	s.PrincipalId = &v
	return s
}

func (s *AssociatePrincipalWithPortfolioRequest) SetPrincipalPattern(v string) *AssociatePrincipalWithPortfolioRequest {
	s.PrincipalPattern = &v
	return s
}

func (s *AssociatePrincipalWithPortfolioRequest) SetPrincipalType(v string) *AssociatePrincipalWithPortfolioRequest {
	s.PrincipalType = &v
	return s
}

func (s *AssociatePrincipalWithPortfolioRequest) Validate() error {
	return dara.Validate(s)
}
