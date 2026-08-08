// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrossAccountsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCreateTargets(v []*UpdateCrossAccountsRequestCreateTargets) *UpdateCrossAccountsRequest
	GetCreateTargets() []*UpdateCrossAccountsRequestCreateTargets
	SetDeleteTargets(v []*UpdateCrossAccountsRequestDeleteTargets) *UpdateCrossAccountsRequest
	GetDeleteTargets() []*UpdateCrossAccountsRequestDeleteTargets
}

type UpdateCrossAccountsRequest struct {
	CreateTargets []*UpdateCrossAccountsRequestCreateTargets `json:"CreateTargets,omitempty" xml:"CreateTargets,omitempty" type:"Repeated"`
	DeleteTargets []*UpdateCrossAccountsRequestDeleteTargets `json:"DeleteTargets,omitempty" xml:"DeleteTargets,omitempty" type:"Repeated"`
}

func (s UpdateCrossAccountsRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrossAccountsRequest) GoString() string {
	return s.String()
}

func (s *UpdateCrossAccountsRequest) GetCreateTargets() []*UpdateCrossAccountsRequestCreateTargets {
	return s.CreateTargets
}

func (s *UpdateCrossAccountsRequest) GetDeleteTargets() []*UpdateCrossAccountsRequestDeleteTargets {
	return s.DeleteTargets
}

func (s *UpdateCrossAccountsRequest) SetCreateTargets(v []*UpdateCrossAccountsRequestCreateTargets) *UpdateCrossAccountsRequest {
	s.CreateTargets = v
	return s
}

func (s *UpdateCrossAccountsRequest) SetDeleteTargets(v []*UpdateCrossAccountsRequestDeleteTargets) *UpdateCrossAccountsRequest {
	s.DeleteTargets = v
	return s
}

func (s *UpdateCrossAccountsRequest) Validate() error {
	if s.CreateTargets != nil {
		for _, item := range s.CreateTargets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DeleteTargets != nil {
		for _, item := range s.DeleteTargets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type UpdateCrossAccountsRequestCreateTargets struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***7890
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ACCOUNT
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateCrossAccountsRequestCreateTargets) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrossAccountsRequestCreateTargets) GoString() string {
	return s.String()
}

func (s *UpdateCrossAccountsRequestCreateTargets) GetTargetId() *string {
	return s.TargetId
}

func (s *UpdateCrossAccountsRequestCreateTargets) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateCrossAccountsRequestCreateTargets) SetTargetId(v string) *UpdateCrossAccountsRequestCreateTargets {
	s.TargetId = &v
	return s
}

func (s *UpdateCrossAccountsRequestCreateTargets) SetTargetType(v string) *UpdateCrossAccountsRequestCreateTargets {
	s.TargetType = &v
	return s
}

func (s *UpdateCrossAccountsRequestCreateTargets) Validate() error {
	return dara.Validate(s)
}

type UpdateCrossAccountsRequestDeleteTargets struct {
	// This parameter is required.
	//
	// example:
	//
	// 123***7890
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ACCOUNT
	TargetType *string `json:"TargetType,omitempty" xml:"TargetType,omitempty"`
}

func (s UpdateCrossAccountsRequestDeleteTargets) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrossAccountsRequestDeleteTargets) GoString() string {
	return s.String()
}

func (s *UpdateCrossAccountsRequestDeleteTargets) GetTargetId() *string {
	return s.TargetId
}

func (s *UpdateCrossAccountsRequestDeleteTargets) GetTargetType() *string {
	return s.TargetType
}

func (s *UpdateCrossAccountsRequestDeleteTargets) SetTargetId(v string) *UpdateCrossAccountsRequestDeleteTargets {
	s.TargetId = &v
	return s
}

func (s *UpdateCrossAccountsRequestDeleteTargets) SetTargetType(v string) *UpdateCrossAccountsRequestDeleteTargets {
	s.TargetType = &v
	return s
}

func (s *UpdateCrossAccountsRequestDeleteTargets) Validate() error {
	return dara.Validate(s)
}
