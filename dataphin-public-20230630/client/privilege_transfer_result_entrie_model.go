// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrivilegeTransferResultEntrie interface {
	dara.Model
	String() string
	GoString() string
	SetChildren(v []*PrivilegeTransferResultEntrie) *PrivilegeTransferResultEntrie
	GetChildren() []*PrivilegeTransferResultEntrie
	SetErrMsg(v string) *PrivilegeTransferResultEntrie
	GetErrMsg() *string
	SetIsLeaf(v bool) *PrivilegeTransferResultEntrie
	GetIsLeaf() *bool
	SetPrivilege(v string) *PrivilegeTransferResultEntrie
	GetPrivilege() *string
	SetPrivilegeDisplayName(v string) *PrivilegeTransferResultEntrie
	GetPrivilegeDisplayName() *string
	SetStatus(v string) *PrivilegeTransferResultEntrie
	GetStatus() *string
	SetTest(v string) *PrivilegeTransferResultEntrie
	GetTest() *string
}

type PrivilegeTransferResultEntrie struct {
	Children             []*PrivilegeTransferResultEntrie `json:"Children,omitempty" xml:"Children,omitempty" type:"Repeated"`
	ErrMsg               *string                          `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	IsLeaf               *bool                            `json:"IsLeaf,omitempty" xml:"IsLeaf,omitempty"`
	Privilege            *string                          `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
	PrivilegeDisplayName *string                          `json:"PrivilegeDisplayName,omitempty" xml:"PrivilegeDisplayName,omitempty"`
	Status               *string                          `json:"Status,omitempty" xml:"Status,omitempty"`
	Test                 *string                          `json:"Test,omitempty" xml:"Test,omitempty"`
}

func (s PrivilegeTransferResultEntrie) String() string {
	return dara.Prettify(s)
}

func (s PrivilegeTransferResultEntrie) GoString() string {
	return s.String()
}

func (s *PrivilegeTransferResultEntrie) GetChildren() []*PrivilegeTransferResultEntrie {
	return s.Children
}

func (s *PrivilegeTransferResultEntrie) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *PrivilegeTransferResultEntrie) GetIsLeaf() *bool {
	return s.IsLeaf
}

func (s *PrivilegeTransferResultEntrie) GetPrivilege() *string {
	return s.Privilege
}

func (s *PrivilegeTransferResultEntrie) GetPrivilegeDisplayName() *string {
	return s.PrivilegeDisplayName
}

func (s *PrivilegeTransferResultEntrie) GetStatus() *string {
	return s.Status
}

func (s *PrivilegeTransferResultEntrie) GetTest() *string {
	return s.Test
}

func (s *PrivilegeTransferResultEntrie) SetChildren(v []*PrivilegeTransferResultEntrie) *PrivilegeTransferResultEntrie {
	s.Children = v
	return s
}

func (s *PrivilegeTransferResultEntrie) SetErrMsg(v string) *PrivilegeTransferResultEntrie {
	s.ErrMsg = &v
	return s
}

func (s *PrivilegeTransferResultEntrie) SetIsLeaf(v bool) *PrivilegeTransferResultEntrie {
	s.IsLeaf = &v
	return s
}

func (s *PrivilegeTransferResultEntrie) SetPrivilege(v string) *PrivilegeTransferResultEntrie {
	s.Privilege = &v
	return s
}

func (s *PrivilegeTransferResultEntrie) SetPrivilegeDisplayName(v string) *PrivilegeTransferResultEntrie {
	s.PrivilegeDisplayName = &v
	return s
}

func (s *PrivilegeTransferResultEntrie) SetStatus(v string) *PrivilegeTransferResultEntrie {
	s.Status = &v
	return s
}

func (s *PrivilegeTransferResultEntrie) SetTest(v string) *PrivilegeTransferResultEntrie {
	s.Test = &v
	return s
}

func (s *PrivilegeTransferResultEntrie) Validate() error {
	return dara.Validate(s)
}
