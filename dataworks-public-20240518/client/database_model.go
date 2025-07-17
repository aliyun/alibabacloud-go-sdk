// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDatabase interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *Database
	GetComment() *string
	SetCreateTime(v int64) *Database
	GetCreateTime() *int64
	SetId(v string) *Database
	GetId() *string
	SetLocationUri(v string) *Database
	GetLocationUri() *string
	SetModifyTime(v int64) *Database
	GetModifyTime() *int64
	SetName(v string) *Database
	GetName() *string
	SetParentMetaEntityId(v string) *Database
	GetParentMetaEntityId() *string
}

type Database struct {
	// example:
	//
	// test comment
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// example:
	//
	// 1736852168000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// holo-database:h-xxxx::test_db
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// oss://test-bucket/test_db
	LocationUri *string `json:"LocationUri,omitempty" xml:"LocationUri,omitempty"`
	// example:
	//
	// 1736852168000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// test_db
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// holo:h-xxxx
	ParentMetaEntityId *string `json:"ParentMetaEntityId,omitempty" xml:"ParentMetaEntityId,omitempty"`
}

func (s Database) String() string {
	return dara.Prettify(s)
}

func (s Database) GoString() string {
	return s.String()
}

func (s *Database) GetComment() *string {
	return s.Comment
}

func (s *Database) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *Database) GetId() *string {
	return s.Id
}

func (s *Database) GetLocationUri() *string {
	return s.LocationUri
}

func (s *Database) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *Database) GetName() *string {
	return s.Name
}

func (s *Database) GetParentMetaEntityId() *string {
	return s.ParentMetaEntityId
}

func (s *Database) SetComment(v string) *Database {
	s.Comment = &v
	return s
}

func (s *Database) SetCreateTime(v int64) *Database {
	s.CreateTime = &v
	return s
}

func (s *Database) SetId(v string) *Database {
	s.Id = &v
	return s
}

func (s *Database) SetLocationUri(v string) *Database {
	s.LocationUri = &v
	return s
}

func (s *Database) SetModifyTime(v int64) *Database {
	s.ModifyTime = &v
	return s
}

func (s *Database) SetName(v string) *Database {
	s.Name = &v
	return s
}

func (s *Database) SetParentMetaEntityId(v string) *Database {
	s.ParentMetaEntityId = &v
	return s
}

func (s *Database) Validate() error {
	return dara.Validate(s)
}
