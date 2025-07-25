// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMoveGtmResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *MoveGtmResourceGroupRequest
	GetLang() *string
	SetNewResourceGroupId(v string) *MoveGtmResourceGroupRequest
	GetNewResourceGroupId() *string
	SetResourceId(v string) *MoveGtmResourceGroupRequest
	GetResourceId() *string
}

type MoveGtmResourceGroupRequest struct {
	// example:
	//
	// en
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AgIDE1MA_XXX
	NewResourceGroupId *string `json:"NewResourceGroupId,omitempty" xml:"NewResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// rg-aekzzk7hx3*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
}

func (s MoveGtmResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s MoveGtmResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *MoveGtmResourceGroupRequest) GetLang() *string {
	return s.Lang
}

func (s *MoveGtmResourceGroupRequest) GetNewResourceGroupId() *string {
	return s.NewResourceGroupId
}

func (s *MoveGtmResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *MoveGtmResourceGroupRequest) SetLang(v string) *MoveGtmResourceGroupRequest {
	s.Lang = &v
	return s
}

func (s *MoveGtmResourceGroupRequest) SetNewResourceGroupId(v string) *MoveGtmResourceGroupRequest {
	s.NewResourceGroupId = &v
	return s
}

func (s *MoveGtmResourceGroupRequest) SetResourceId(v string) *MoveGtmResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *MoveGtmResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}
