// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *ListPublicKeysRequest
	GetEndTime() *string
	SetKeyGroup(v string) *ListPublicKeysRequest
	GetKeyGroup() *string
	SetKeyName(v string) *ListPublicKeysRequest
	GetKeyName() *string
	SetKeyType(v string) *ListPublicKeysRequest
	GetKeyType() *string
	SetPageNumber(v int64) *ListPublicKeysRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListPublicKeysRequest
	GetPageSize() *int64
	SetStartTime(v string) *ListPublicKeysRequest
	GetStartTime() *string
}

type ListPublicKeysRequest struct {
	// A parameter for filtering by time range. The time must be in UTC and follow the ISO 8601 standard. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2021-09-22T02:23:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The name of the public key group.
	//
	// example:
	//
	// g-test
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// The name of the public key.
	//
	// example:
	//
	// mykey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// The type of the public key. Valid values:
	//
	// - **adb**: ADB key
	//
	// - **ssh**: SSH key
	//
	// example:
	//
	// ssh
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// The page number of the list to query. The value starts from 1.
	//
	// Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page for a paged query. Valid values: 1 to 100.
	//
	// Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A parameter for filtering by time range. The time must be in UTC and follow the ISO 8601 standard. The format is yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2021-09-21T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ListPublicKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPublicKeysRequest) GoString() string {
	return s.String()
}

func (s *ListPublicKeysRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ListPublicKeysRequest) GetKeyGroup() *string {
	return s.KeyGroup
}

func (s *ListPublicKeysRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *ListPublicKeysRequest) GetKeyType() *string {
	return s.KeyType
}

func (s *ListPublicKeysRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListPublicKeysRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPublicKeysRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ListPublicKeysRequest) SetEndTime(v string) *ListPublicKeysRequest {
	s.EndTime = &v
	return s
}

func (s *ListPublicKeysRequest) SetKeyGroup(v string) *ListPublicKeysRequest {
	s.KeyGroup = &v
	return s
}

func (s *ListPublicKeysRequest) SetKeyName(v string) *ListPublicKeysRequest {
	s.KeyName = &v
	return s
}

func (s *ListPublicKeysRequest) SetKeyType(v string) *ListPublicKeysRequest {
	s.KeyType = &v
	return s
}

func (s *ListPublicKeysRequest) SetPageNumber(v int64) *ListPublicKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPublicKeysRequest) SetPageSize(v int64) *ListPublicKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListPublicKeysRequest) SetStartTime(v string) *ListPublicKeysRequest {
	s.StartTime = &v
	return s
}

func (s *ListPublicKeysRequest) Validate() error {
	return dara.Validate(s)
}
