// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTransitRouterPolicyTableRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteTransitRouterPolicyTableRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteTransitRouterPolicyTableRequest
	GetDryRun() *bool
	SetPolicyTableId(v string) *DeleteTransitRouterPolicyTableRequest
	GetPolicyTableId() *string
}

type DeleteTransitRouterPolicyTableRequest struct {
	// example:
	//
	// 02fb3da4****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ptbl-abc123xyz
	PolicyTableId *string `json:"PolicyTableId,omitempty" xml:"PolicyTableId,omitempty"`
}

func (s DeleteTransitRouterPolicyTableRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTransitRouterPolicyTableRequest) GoString() string {
	return s.String()
}

func (s *DeleteTransitRouterPolicyTableRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteTransitRouterPolicyTableRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteTransitRouterPolicyTableRequest) GetPolicyTableId() *string {
	return s.PolicyTableId
}

func (s *DeleteTransitRouterPolicyTableRequest) SetClientToken(v string) *DeleteTransitRouterPolicyTableRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteTransitRouterPolicyTableRequest) SetDryRun(v bool) *DeleteTransitRouterPolicyTableRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteTransitRouterPolicyTableRequest) SetPolicyTableId(v string) *DeleteTransitRouterPolicyTableRequest {
	s.PolicyTableId = &v
	return s
}

func (s *DeleteTransitRouterPolicyTableRequest) Validate() error {
	return dara.Validate(s)
}
