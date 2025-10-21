// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLock interface {
	dara.Model
	String() string
	GoString() string
	SetHolderId(v string) *Lock
	GetHolderId() *string
	SetHolderName(v string) *Lock
	GetHolderName() *string
	SetId(v string) *Lock
	GetId() *string
	SetNamespace(v string) *Lock
	GetNamespace() *string
	SetWorkspace(v string) *Lock
	GetWorkspace() *string
}

type Lock struct {
	HolderId   *string `json:"holderId,omitempty" xml:"holderId,omitempty"`
	HolderName *string `json:"holderName,omitempty" xml:"holderName,omitempty"`
	Id         *string `json:"id,omitempty" xml:"id,omitempty"`
	Namespace  *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	Workspace  *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s Lock) String() string {
	return dara.Prettify(s)
}

func (s Lock) GoString() string {
	return s.String()
}

func (s *Lock) GetHolderId() *string {
	return s.HolderId
}

func (s *Lock) GetHolderName() *string {
	return s.HolderName
}

func (s *Lock) GetId() *string {
	return s.Id
}

func (s *Lock) GetNamespace() *string {
	return s.Namespace
}

func (s *Lock) GetWorkspace() *string {
	return s.Workspace
}

func (s *Lock) SetHolderId(v string) *Lock {
	s.HolderId = &v
	return s
}

func (s *Lock) SetHolderName(v string) *Lock {
	s.HolderName = &v
	return s
}

func (s *Lock) SetId(v string) *Lock {
	s.Id = &v
	return s
}

func (s *Lock) SetNamespace(v string) *Lock {
	s.Namespace = &v
	return s
}

func (s *Lock) SetWorkspace(v string) *Lock {
	s.Workspace = &v
	return s
}

func (s *Lock) Validate() error {
	return dara.Validate(s)
}
