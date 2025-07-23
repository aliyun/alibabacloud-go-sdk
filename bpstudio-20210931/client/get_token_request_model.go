// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *GetTokenRequest
	GetResourceGroupId() *string
}

type GetTokenRequest struct {
	// The ID of the resource group.
	//
	// example:
	//
	// rg-aekzhfgmw4e6fwq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s GetTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTokenRequest) GoString() string {
	return s.String()
}

func (s *GetTokenRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetTokenRequest) SetResourceGroupId(v string) *GetTokenRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *GetTokenRequest) Validate() error {
	return dara.Validate(s)
}
