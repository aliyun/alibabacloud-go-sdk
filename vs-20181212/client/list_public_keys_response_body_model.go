// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int64) *ListPublicKeysResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListPublicKeysResponseBody
	GetPageSize() *int64
	SetPublicKeys(v []*ListPublicKeysResponseBodyPublicKeys) *ListPublicKeysResponseBody
	GetPublicKeys() []*ListPublicKeysResponseBodyPublicKeys
	SetRequestId(v string) *ListPublicKeysResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListPublicKeysResponseBody
	GetTotalCount() *int64
}

type ListPublicKeysResponseBody struct {
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize   *int64                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	PublicKeys []*ListPublicKeysResponseBodyPublicKeys `json:"PublicKeys,omitempty" xml:"PublicKeys,omitempty" type:"Repeated"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListPublicKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPublicKeysResponseBody) GoString() string {
	return s.String()
}

func (s *ListPublicKeysResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListPublicKeysResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPublicKeysResponseBody) GetPublicKeys() []*ListPublicKeysResponseBodyPublicKeys {
	return s.PublicKeys
}

func (s *ListPublicKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPublicKeysResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListPublicKeysResponseBody) SetPageNumber(v int64) *ListPublicKeysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListPublicKeysResponseBody) SetPageSize(v int64) *ListPublicKeysResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListPublicKeysResponseBody) SetPublicKeys(v []*ListPublicKeysResponseBodyPublicKeys) *ListPublicKeysResponseBody {
	s.PublicKeys = v
	return s
}

func (s *ListPublicKeysResponseBody) SetRequestId(v string) *ListPublicKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPublicKeysResponseBody) SetTotalCount(v int64) *ListPublicKeysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPublicKeysResponseBody) Validate() error {
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

type ListPublicKeysResponseBodyPublicKeys struct {
	// example:
	//
	// verify_30d89ccb0905c8c7882c1d14a991954b
	Content     *string `json:"Content,omitempty" xml:"Content,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// mygroup
	KeyGroup *string `json:"KeyGroup,omitempty" xml:"KeyGroup,omitempty"`
	// example:
	//
	// mykey
	KeyName *string `json:"KeyName,omitempty" xml:"KeyName,omitempty"`
	KeyType *string `json:"KeyType,omitempty" xml:"KeyType,omitempty"`
	// example:
	//
	// 2024-06-11T14:26:48+08:00
	UploadTime *string `json:"UploadTime,omitempty" xml:"UploadTime,omitempty"`
}

func (s ListPublicKeysResponseBodyPublicKeys) String() string {
	return dara.Prettify(s)
}

func (s ListPublicKeysResponseBodyPublicKeys) GoString() string {
	return s.String()
}

func (s *ListPublicKeysResponseBodyPublicKeys) GetContent() *string {
	return s.Content
}

func (s *ListPublicKeysResponseBodyPublicKeys) GetDescription() *string {
	return s.Description
}

func (s *ListPublicKeysResponseBodyPublicKeys) GetKeyGroup() *string {
	return s.KeyGroup
}

func (s *ListPublicKeysResponseBodyPublicKeys) GetKeyName() *string {
	return s.KeyName
}

func (s *ListPublicKeysResponseBodyPublicKeys) GetKeyType() *string {
	return s.KeyType
}

func (s *ListPublicKeysResponseBodyPublicKeys) GetUploadTime() *string {
	return s.UploadTime
}

func (s *ListPublicKeysResponseBodyPublicKeys) SetContent(v string) *ListPublicKeysResponseBodyPublicKeys {
	s.Content = &v
	return s
}

func (s *ListPublicKeysResponseBodyPublicKeys) SetDescription(v string) *ListPublicKeysResponseBodyPublicKeys {
	s.Description = &v
	return s
}

func (s *ListPublicKeysResponseBodyPublicKeys) SetKeyGroup(v string) *ListPublicKeysResponseBodyPublicKeys {
	s.KeyGroup = &v
	return s
}

func (s *ListPublicKeysResponseBodyPublicKeys) SetKeyName(v string) *ListPublicKeysResponseBodyPublicKeys {
	s.KeyName = &v
	return s
}

func (s *ListPublicKeysResponseBodyPublicKeys) SetKeyType(v string) *ListPublicKeysResponseBodyPublicKeys {
	s.KeyType = &v
	return s
}

func (s *ListPublicKeysResponseBodyPublicKeys) SetUploadTime(v string) *ListPublicKeysResponseBodyPublicKeys {
	s.UploadTime = &v
	return s
}

func (s *ListPublicKeysResponseBodyPublicKeys) Validate() error {
	return dara.Validate(s)
}
