// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetaDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatabaseInfo(v *ListMetaDBResponseBodyDatabaseInfo) *ListMetaDBResponseBody
	GetDatabaseInfo() *ListMetaDBResponseBodyDatabaseInfo
	SetRequestId(v string) *ListMetaDBResponseBody
	GetRequestId() *string
}

type ListMetaDBResponseBody struct {
	// The database information.
	DatabaseInfo *ListMetaDBResponseBodyDatabaseInfo `json:"DatabaseInfo,omitempty" xml:"DatabaseInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// abc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMetaDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListMetaDBResponseBody) GoString() string {
	return s.String()
}

func (s *ListMetaDBResponseBody) GetDatabaseInfo() *ListMetaDBResponseBodyDatabaseInfo {
	return s.DatabaseInfo
}

func (s *ListMetaDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMetaDBResponseBody) SetDatabaseInfo(v *ListMetaDBResponseBodyDatabaseInfo) *ListMetaDBResponseBody {
	s.DatabaseInfo = v
	return s
}

func (s *ListMetaDBResponseBody) SetRequestId(v string) *ListMetaDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMetaDBResponseBody) Validate() error {
	if s.DatabaseInfo != nil {
		if err := s.DatabaseInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListMetaDBResponseBodyDatabaseInfo struct {
	// The list of databases.
	DbList []*ListMetaDBResponseBodyDatabaseInfoDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
	// The total number of databases.
	//
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListMetaDBResponseBodyDatabaseInfo) String() string {
	return dara.Prettify(s)
}

func (s ListMetaDBResponseBodyDatabaseInfo) GoString() string {
	return s.String()
}

func (s *ListMetaDBResponseBodyDatabaseInfo) GetDbList() []*ListMetaDBResponseBodyDatabaseInfoDbList {
	return s.DbList
}

func (s *ListMetaDBResponseBodyDatabaseInfo) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListMetaDBResponseBodyDatabaseInfo) SetDbList(v []*ListMetaDBResponseBodyDatabaseInfoDbList) *ListMetaDBResponseBodyDatabaseInfo {
	s.DbList = v
	return s
}

func (s *ListMetaDBResponseBodyDatabaseInfo) SetTotalCount(v int64) *ListMetaDBResponseBodyDatabaseInfo {
	s.TotalCount = &v
	return s
}

func (s *ListMetaDBResponseBodyDatabaseInfo) Validate() error {
	if s.DbList != nil {
		for _, item := range s.DbList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListMetaDBResponseBodyDatabaseInfoDbList struct {
	// The time when the database was created. The value is a timestamp. You can convert the timestamp to a date based on your time zone.
	//
	// example:
	//
	// 1388776825
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// The storage address of the database.
	//
	// example:
	//
	// hdfs://localhost:777/user/hadoop/test.txt
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The time when the database was last updated.
	//
	// example:
	//
	// 1388776837
	ModifiedTimeStamp *int64 `json:"ModifiedTimeStamp,omitempty" xml:"ModifiedTimeStamp,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the owner.
	//
	// example:
	//
	// 1232
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of the database.
	//
	// example:
	//
	// HIVE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The unique identifier of the database.
	//
	// example:
	//
	// 32342
	UUID *string `json:"UUID,omitempty" xml:"UUID,omitempty"`
}

func (s ListMetaDBResponseBodyDatabaseInfoDbList) String() string {
	return dara.Prettify(s)
}

func (s ListMetaDBResponseBodyDatabaseInfoDbList) GoString() string {
	return s.String()
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) GetCreateTimeStamp() *int64 {
	return s.CreateTimeStamp
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) GetLocation() *string {
	return s.Location
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) GetModifiedTimeStamp() *int64 {
	return s.ModifiedTimeStamp
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) GetName() *string {
	return s.Name
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) GetOwnerId() *string {
	return s.OwnerId
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) GetType() *string {
	return s.Type
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) GetUUID() *string {
	return s.UUID
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) SetCreateTimeStamp(v int64) *ListMetaDBResponseBodyDatabaseInfoDbList {
	s.CreateTimeStamp = &v
	return s
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) SetLocation(v string) *ListMetaDBResponseBodyDatabaseInfoDbList {
	s.Location = &v
	return s
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) SetModifiedTimeStamp(v int64) *ListMetaDBResponseBodyDatabaseInfoDbList {
	s.ModifiedTimeStamp = &v
	return s
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) SetName(v string) *ListMetaDBResponseBodyDatabaseInfoDbList {
	s.Name = &v
	return s
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) SetOwnerId(v string) *ListMetaDBResponseBodyDatabaseInfoDbList {
	s.OwnerId = &v
	return s
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) SetType(v string) *ListMetaDBResponseBodyDatabaseInfoDbList {
	s.Type = &v
	return s
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) SetUUID(v string) *ListMetaDBResponseBodyDatabaseInfoDbList {
	s.UUID = &v
	return s
}

func (s *ListMetaDBResponseBodyDatabaseInfoDbList) Validate() error {
	return dara.Validate(s)
}
