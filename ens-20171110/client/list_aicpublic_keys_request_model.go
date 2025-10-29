// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICPublicKeysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyGroup(v string) *ListAICPublicKeysRequest
	GetKeyGroup() *string
	SetKeyName(v string) *ListAICPublicKeysRequest
	GetKeyName() *string
	SetKeyType(v string) *ListAICPublicKeysRequest
	GetKeyType() *string
	SetPageNumber(v int32) *ListAICPublicKeysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *ListAICPublicKeysRequest
	GetPageSize() *int32
}

type ListAICPublicKeysRequest struct {
	// example:
	//
	// g-test
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// example:
	//
	// mykey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// example:
	//
	// adb
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s ListAICPublicKeysRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAICPublicKeysRequest) GoString() string {
	return s.String()
}

func (s *ListAICPublicKeysRequest) GetKeyGroup() *string {
	return s.KeyGroup
}

func (s *ListAICPublicKeysRequest) GetKeyName() *string {
	return s.KeyName
}

func (s *ListAICPublicKeysRequest) GetKeyType() *string {
	return s.KeyType
}

func (s *ListAICPublicKeysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListAICPublicKeysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListAICPublicKeysRequest) SetKeyGroup(v string) *ListAICPublicKeysRequest {
	s.KeyGroup = &v
	return s
}

func (s *ListAICPublicKeysRequest) SetKeyName(v string) *ListAICPublicKeysRequest {
	s.KeyName = &v
	return s
}

func (s *ListAICPublicKeysRequest) SetKeyType(v string) *ListAICPublicKeysRequest {
	s.KeyType = &v
	return s
}

func (s *ListAICPublicKeysRequest) SetPageNumber(v int32) *ListAICPublicKeysRequest {
	s.PageNumber = &v
	return s
}

func (s *ListAICPublicKeysRequest) SetPageSize(v int32) *ListAICPublicKeysRequest {
	s.PageSize = &v
	return s
}

func (s *ListAICPublicKeysRequest) Validate() error {
	return dara.Validate(s)
}
