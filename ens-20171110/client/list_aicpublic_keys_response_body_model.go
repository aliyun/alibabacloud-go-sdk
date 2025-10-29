// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAICPublicKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListAICPublicKeysResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAICPublicKeysResponseBody
	GetPageSize() *int64
	SetPublicKeys(v []*ListAICPublicKeysResponseBodyPublicKeys) *ListAICPublicKeysResponseBody
	GetPublicKeys() []*ListAICPublicKeysResponseBodyPublicKeys
	SetRequestId(v string) *ListAICPublicKeysResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListAICPublicKeysResponseBody
	GetTotalCount() *int64
}

type ListAICPublicKeysResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize   *int64                                     `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PublicKeys []*ListAICPublicKeysResponseBodyPublicKeys `json:"PublicKeys,omitempty" xml:"PublicKeys,omitempty" type:"Repeated"`
	// example:
	//
	// xxxsxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 10
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAICPublicKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAICPublicKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListAICPublicKeysResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAICPublicKeysResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAICPublicKeysResponseBody) GetPublicKeys() []*ListAICPublicKeysResponseBodyPublicKeys {
	return s.PublicKeys
}

func (s *ListAICPublicKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAICPublicKeysResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListAICPublicKeysResponseBody) SetPageNumber(v int64) *ListAICPublicKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAICPublicKeysResponseBody) SetPageSize(v int64) *ListAICPublicKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAICPublicKeysResponseBody) SetPublicKeys(v []*ListAICPublicKeysResponseBodyPublicKeys) *ListAICPublicKeysResponseBody {
	s.PublicKeys = v
	return s
}

func (s *ListAICPublicKeysResponseBody) SetRequestId(v string) *ListAICPublicKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAICPublicKeysResponseBody) SetTotalCount(v int64) *ListAICPublicKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAICPublicKeysResponseBody) Validate() error {
	if s.PublicKeys != nil {
		for _, item := range s.PublicKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAICPublicKeysResponseBodyPublicKeys struct {
	// example:
	//
	// fvsvshbvjfksvj
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// example:
	//
	// 2025-10-09T15:13:21+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// test-group
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// example:
	//
	// mykey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	// example:
	//
	// adb
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
}

func (s ListAICPublicKeysResponseBodyPublicKeys) String() string {
	return dara.Prettify(s)
}

func (s ListAICPublicKeysResponseBodyPublicKeys) GoString() string {
	return s.String()
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) GetContent() *string {
	return s.Content
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) GetDescription() *string {
	return s.Description
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) GetKeyGroup() *string {
	return s.KeyGroup
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) GetKeyName() *string {
	return s.KeyName
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) GetKeyType() *string {
	return s.KeyType
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) SetContent(v string) *ListAICPublicKeysResponseBodyPublicKeys {
	s.Content = &v
	return s
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) SetCreationTime(v string) *ListAICPublicKeysResponseBodyPublicKeys {
	s.CreationTime = &v
	return s
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) SetDescription(v string) *ListAICPublicKeysResponseBodyPublicKeys {
	s.Description = &v
	return s
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) SetKeyGroup(v string) *ListAICPublicKeysResponseBodyPublicKeys {
	s.KeyGroup = &v
	return s
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) SetKeyName(v string) *ListAICPublicKeysResponseBodyPublicKeys {
	s.KeyName = &v
	return s
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) SetKeyType(v string) *ListAICPublicKeysResponseBodyPublicKeys {
	s.KeyType = &v
	return s
}

func (s *ListAICPublicKeysResponseBodyPublicKeys) Validate() error {
	return dara.Validate(s)
}
