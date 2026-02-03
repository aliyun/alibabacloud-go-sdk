// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisassociatePrincipalFromPortfolioRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPortfolioId(v string) *DisassociatePrincipalFromPortfolioRequest
	GetPortfolioId() *string
	SetPrincipalId(v string) *DisassociatePrincipalFromPortfolioRequest
	GetPrincipalId() *string
	SetPrincipalPattern(v string) *DisassociatePrincipalFromPortfolioRequest
	GetPrincipalPattern() *string
	SetPrincipalType(v string) *DisassociatePrincipalFromPortfolioRequest
	GetPrincipalType() *string
}

type DisassociatePrincipalFromPortfolioRequest struct {
	// The ID of the product portfolio.
	//
	// This parameter is required.
	//
	// example:
	//
	// port-bp1yt7582g****
	PortfolioId *string `json:"PortfolioId,omitempty" xml:"PortfolioId,omitempty"`
	// The ID of the RAM entity.
	//
	// For more information about how to obtain the ID of a RAM user, see [GetUser](https://help.aliyun.com/document_detail/28681.html).
	//
	// For more information about how to obtain the ID of a RAM role, see [GetRole](https://help.aliyun.com/document_detail/28711.html).
	//
	// example:
	//
	// 24477111603637****
	PrincipalId *string `json:"PrincipalId,omitempty" xml:"PrincipalId,omitempty"`
	// example:
	//
	// user/userName
	PrincipalPattern *string `json:"PrincipalPattern,omitempty" xml:"PrincipalPattern,omitempty"`
	// The type of the Resource Access Management (RAM) entity. Valid values:
	//
	// 	- RamUser: a RAM user
	//
	// 	- RamRole: a RAM role
	//
	// This parameter is required.
	//
	// example:
	//
	// RamUser
	PrincipalType *string `json:"PrincipalType,omitempty" xml:"PrincipalType,omitempty"`
}

func (s DisassociatePrincipalFromPortfolioRequest) String() string {
	return dara.Prettify(s)
}

func (s DisassociatePrincipalFromPortfolioRequest) GoString() string {
	return s.String()
}

func (s *DisassociatePrincipalFromPortfolioRequest) GetPortfolioId() *string {
	return s.PortfolioId
}

func (s *DisassociatePrincipalFromPortfolioRequest) GetPrincipalId() *string {
	return s.PrincipalId
}

func (s *DisassociatePrincipalFromPortfolioRequest) GetPrincipalPattern() *string {
	return s.PrincipalPattern
}

func (s *DisassociatePrincipalFromPortfolioRequest) GetPrincipalType() *string {
	return s.PrincipalType
}

func (s *DisassociatePrincipalFromPortfolioRequest) SetPortfolioId(v string) *DisassociatePrincipalFromPortfolioRequest {
	s.PortfolioId = &v
	return s
}

func (s *DisassociatePrincipalFromPortfolioRequest) SetPrincipalId(v string) *DisassociatePrincipalFromPortfolioRequest {
	s.PrincipalId = &v
	return s
}

func (s *DisassociatePrincipalFromPortfolioRequest) SetPrincipalPattern(v string) *DisassociatePrincipalFromPortfolioRequest {
	s.PrincipalPattern = &v
	return s
}

func (s *DisassociatePrincipalFromPortfolioRequest) SetPrincipalType(v string) *DisassociatePrincipalFromPortfolioRequest {
	s.PrincipalType = &v
	return s
}

func (s *DisassociatePrincipalFromPortfolioRequest) Validate() error {
	return dara.Validate(s)
}
