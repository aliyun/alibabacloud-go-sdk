// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeyVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyVersions(v *ListKeyVersionsResponseBodyKeyVersions) *ListKeyVersionsResponseBody
	GetKeyVersions() *ListKeyVersionsResponseBodyKeyVersions
	SetPageNumber(v int32) *ListKeyVersionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKeyVersionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListKeyVersionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListKeyVersionsResponseBody
	GetTotalCount() *int32
}

type ListKeyVersionsResponseBody struct {
	// An array that consists of key versions.
	KeyVersions *ListKeyVersionsResponseBodyKeyVersions `json:"KeyVersions,omitempty" xml:"KeyVersions,omitempty" type:"Struct"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// f71204c4-53cd-4eea-b405-653ba2db7e86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of returned key versions.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKeyVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKeyVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsResponseBody) GetKeyVersions() *ListKeyVersionsResponseBodyKeyVersions {
	return s.KeyVersions
}

func (s *ListKeyVersionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKeyVersionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKeyVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKeyVersionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKeyVersionsResponseBody) SetKeyVersions(v *ListKeyVersionsResponseBodyKeyVersions) *ListKeyVersionsResponseBody {
	s.KeyVersions = v
	return s
}

func (s *ListKeyVersionsResponseBody) SetPageNumber(v int32) *ListKeyVersionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKeyVersionsResponseBody) SetPageSize(v int32) *ListKeyVersionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKeyVersionsResponseBody) SetRequestId(v string) *ListKeyVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKeyVersionsResponseBody) SetTotalCount(v int32) *ListKeyVersionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKeyVersionsResponseBody) Validate() error {
	if s.KeyVersions != nil {
		if err := s.KeyVersions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListKeyVersionsResponseBodyKeyVersions struct {
	KeyVersion []*ListKeyVersionsResponseBodyKeyVersionsKeyVersion `json:"KeyVersion,omitempty" xml:"KeyVersion,omitempty" type:"Repeated"`
}

func (s ListKeyVersionsResponseBodyKeyVersions) String() string {
	return dara.Prettify(s)
}

func (s ListKeyVersionsResponseBodyKeyVersions) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsResponseBodyKeyVersions) GetKeyVersion() []*ListKeyVersionsResponseBodyKeyVersionsKeyVersion {
	return s.KeyVersion
}

func (s *ListKeyVersionsResponseBodyKeyVersions) SetKeyVersion(v []*ListKeyVersionsResponseBodyKeyVersionsKeyVersion) *ListKeyVersionsResponseBodyKeyVersions {
	s.KeyVersion = v
	return s
}

func (s *ListKeyVersionsResponseBodyKeyVersions) Validate() error {
	if s.KeyVersion != nil {
		for _, item := range s.KeyVersion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKeyVersionsResponseBodyKeyVersionsKeyVersion struct {
	// The date and time when the CMK version was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2016-03-25T10:42:40Z
	CreationDate *string `json:"CreationDate,omitempty" xml:"CreationDate,omitempty"`
	// The globally unique ID of the CMK.
	//
	// >  If you set the KeyId parameter to the alias of the CMK, the ID of the CMK to which the alias is bound is returned.
	//
	// example:
	//
	// 0b30658a-ed1a-4922-b8f7-a673ca9c****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The globally unique ID of the CMK version.
	//
	// example:
	//
	// 1e3304fd-68ac-4d5b-8886-ae5f01a1****
	KeyVersionId *string `json:"KeyVersionId,omitempty" xml:"KeyVersionId,omitempty"`
}

func (s ListKeyVersionsResponseBodyKeyVersionsKeyVersion) String() string {
	return dara.Prettify(s)
}

func (s ListKeyVersionsResponseBodyKeyVersionsKeyVersion) GoString() string {
	return s.String()
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) GetCreationDate() *string {
	return s.CreationDate
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) GetKeyId() *string {
	return s.KeyId
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) GetKeyVersionId() *string {
	return s.KeyVersionId
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) SetCreationDate(v string) *ListKeyVersionsResponseBodyKeyVersionsKeyVersion {
	s.CreationDate = &v
	return s
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) SetKeyId(v string) *ListKeyVersionsResponseBodyKeyVersionsKeyVersion {
	s.KeyId = &v
	return s
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) SetKeyVersionId(v string) *ListKeyVersionsResponseBodyKeyVersionsKeyVersion {
	s.KeyVersionId = &v
	return s
}

func (s *ListKeyVersionsResponseBodyKeyVersionsKeyVersion) Validate() error {
	return dara.Validate(s)
}
