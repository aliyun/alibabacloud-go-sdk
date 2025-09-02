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
	// The information about the metadatabases.
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
	return dara.Validate(s)
}

type ListMetaDBResponseBodyDatabaseInfo struct {
	// The metadatabases.
	DbList []*ListMetaDBResponseBodyDatabaseInfoDbList `json:"DbList,omitempty" xml:"DbList,omitempty" type:"Repeated"`
	// The total number of the metadatabases returned.
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
	return dara.Validate(s)
}

type ListMetaDBResponseBodyDatabaseInfoDbList struct {
	// The timestamp at which the metadatabase was created. You can convert the timestamp to the date based on the time zone that you use.
	//
	// example:
	//
	// 1388776825
	CreateTimeStamp *int64 `json:"CreateTimeStamp,omitempty" xml:"CreateTimeStamp,omitempty"`
	// The URL of the metadatabase.
	//
	// example:
	//
	// hdfs://localhost:777/user/hadoop/test.txt
	Location *string `json:"Location,omitempty" xml:"Location,omitempty"`
	// The timestamp at which the metadatabase was updated.
	//
	// example:
	//
	// 1388776837
	ModifiedTimeStamp *int64 `json:"ModifiedTimeStamp,omitempty" xml:"ModifiedTimeStamp,omitempty"`
	// The name of the metadatabase.
	//
	// example:
	//
	// name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The owner ID.
	//
	// example:
	//
	// 1232
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The type of the metadatabase.
	//
	// example:
	//
	// HIVE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The UUID of the metadatabase.
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
