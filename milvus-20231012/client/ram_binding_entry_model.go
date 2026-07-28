// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRamBindingEntry interface {
	dara.Model
	String() string
	GoString() string
	SetBoundUid(v string) *RamBindingEntry
	GetBoundUid() *string
	SetBoundUserName(v string) *RamBindingEntry
	GetBoundUserName() *string
	SetMilvusUsername(v string) *RamBindingEntry
	GetMilvusUsername() *string
}

type RamBindingEntry struct {
	// The unique identifier (UID) of the bound RAM user.
	//
	// example:
	//
	// 1234567890123456
	BoundUid *string `json:"boundUid,omitempty" xml:"boundUid,omitempty"`
	// The bound RAM username.
	//
	// example:
	//
	// test-user
	BoundUserName *string `json:"boundUserName,omitempty" xml:"boundUserName,omitempty"`
	// The Milvus username.
	//
	// example:
	//
	// root
	MilvusUsername *string `json:"milvusUsername,omitempty" xml:"milvusUsername,omitempty"`
}

func (s RamBindingEntry) String() string {
	return dara.Prettify(s)
}

func (s RamBindingEntry) GoString() string {
	return s.String()
}

func (s *RamBindingEntry) GetBoundUid() *string {
	return s.BoundUid
}

func (s *RamBindingEntry) GetBoundUserName() *string {
	return s.BoundUserName
}

func (s *RamBindingEntry) GetMilvusUsername() *string {
	return s.MilvusUsername
}

func (s *RamBindingEntry) SetBoundUid(v string) *RamBindingEntry {
	s.BoundUid = &v
	return s
}

func (s *RamBindingEntry) SetBoundUserName(v string) *RamBindingEntry {
	s.BoundUserName = &v
	return s
}

func (s *RamBindingEntry) SetMilvusUsername(v string) *RamBindingEntry {
	s.MilvusUsername = &v
	return s
}

func (s *RamBindingEntry) Validate() error {
	return dara.Validate(s)
}
