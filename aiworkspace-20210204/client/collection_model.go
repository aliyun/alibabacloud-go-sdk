// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCollection interface {
	dara.Model
	String() string
	GoString() string
	SetCollectionName(v string) *Collection
	GetCollectionName() *string
	SetGmtCreateTime(v string) *Collection
	GetGmtCreateTime() *string
	SetGmtModifiedTime(v string) *Collection
	GetGmtModifiedTime() *string
	SetOwnerId(v string) *Collection
	GetOwnerId() *string
	SetUserId(v string) *Collection
	GetUserId() *string
}

type Collection struct {
	// example:
	//
	// AI4D
	CollectionName *string `json:"CollectionName,omitempty" xml:"CollectionName,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtCreateTime *string `json:"GmtCreateTime,omitempty" xml:"GmtCreateTime,omitempty"`
	// example:
	//
	// 2021-01-21T17:12:35Z
	GmtModifiedTime *string `json:"GmtModifiedTime,omitempty" xml:"GmtModifiedTime,omitempty"`
	// example:
	//
	// 155770209******
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// 155770209******
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s Collection) String() string {
	return dara.Prettify(s)
}

func (s Collection) GoString() string {
	return s.String()
}

func (s *Collection) GetCollectionName() *string {
	return s.CollectionName
}

func (s *Collection) GetGmtCreateTime() *string {
	return s.GmtCreateTime
}

func (s *Collection) GetGmtModifiedTime() *string {
	return s.GmtModifiedTime
}

func (s *Collection) GetOwnerId() *string {
	return s.OwnerId
}

func (s *Collection) GetUserId() *string {
	return s.UserId
}

func (s *Collection) SetCollectionName(v string) *Collection {
	s.CollectionName = &v
	return s
}

func (s *Collection) SetGmtCreateTime(v string) *Collection {
	s.GmtCreateTime = &v
	return s
}

func (s *Collection) SetGmtModifiedTime(v string) *Collection {
	s.GmtModifiedTime = &v
	return s
}

func (s *Collection) SetOwnerId(v string) *Collection {
	s.OwnerId = &v
	return s
}

func (s *Collection) SetUserId(v string) *Collection {
	s.UserId = &v
	return s
}

func (s *Collection) Validate() error {
	return dara.Validate(s)
}
