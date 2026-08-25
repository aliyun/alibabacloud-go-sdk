// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDirectoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDirectories(v []*ListDirectoriesResponseBodyDirectories) *ListDirectoriesResponseBody
	GetDirectories() []*ListDirectoriesResponseBodyDirectories
	SetRequestId(v string) *ListDirectoriesResponseBody
	GetRequestId() *string
	SetTotalCounts(v int32) *ListDirectoriesResponseBody
	GetTotalCounts() *int32
}

type ListDirectoriesResponseBody struct {
	// The directories.
	Directories []*ListDirectoriesResponseBodyDirectories `json:"Directories,omitempty" xml:"Directories,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 9A504392-F06D-5029-AB64-6654CB9F1DC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of directories.
	//
	// example:
	//
	// 1
	TotalCounts *int32 `json:"TotalCounts,omitempty" xml:"TotalCounts,omitempty"`
}

func (s ListDirectoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDirectoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDirectoriesResponseBody) GetDirectories() []*ListDirectoriesResponseBodyDirectories {
	return s.Directories
}

func (s *ListDirectoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDirectoriesResponseBody) GetTotalCounts() *int32 {
	return s.TotalCounts
}

func (s *ListDirectoriesResponseBody) SetDirectories(v []*ListDirectoriesResponseBodyDirectories) *ListDirectoriesResponseBody {
	s.Directories = v
	return s
}

func (s *ListDirectoriesResponseBody) SetRequestId(v string) *ListDirectoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDirectoriesResponseBody) SetTotalCounts(v int32) *ListDirectoriesResponseBody {
	s.TotalCounts = &v
	return s
}

func (s *ListDirectoriesResponseBody) Validate() error {
	if s.Directories != nil {
		for _, item := range s.Directories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDirectoriesResponseBodyDirectories struct {
	// The time when the directory was created.
	//
	// example:
	//
	// 2021-06-30T08:35:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the directory.
	//
	// example:
	//
	// d-00fc2p61****
	DirectoryId *string `json:"DirectoryId,omitempty" xml:"DirectoryId,omitempty"`
	// The name of the directory.
	//
	// example:
	//
	// new-example
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The region ID of the directory.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the directory was modified.
	//
	// example:
	//
	// 2021-10-25T09:13:24Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListDirectoriesResponseBodyDirectories) String() string {
	return dara.Prettify(s)
}

func (s ListDirectoriesResponseBodyDirectories) GoString() string {
	return s.String()
}

func (s *ListDirectoriesResponseBodyDirectories) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListDirectoriesResponseBodyDirectories) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *ListDirectoriesResponseBodyDirectories) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *ListDirectoriesResponseBodyDirectories) GetRegion() *string {
	return s.Region
}

func (s *ListDirectoriesResponseBodyDirectories) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListDirectoriesResponseBodyDirectories) SetCreateTime(v string) *ListDirectoriesResponseBodyDirectories {
	s.CreateTime = &v
	return s
}

func (s *ListDirectoriesResponseBodyDirectories) SetDirectoryId(v string) *ListDirectoriesResponseBodyDirectories {
	s.DirectoryId = &v
	return s
}

func (s *ListDirectoriesResponseBodyDirectories) SetDirectoryName(v string) *ListDirectoriesResponseBodyDirectories {
	s.DirectoryName = &v
	return s
}

func (s *ListDirectoriesResponseBodyDirectories) SetRegion(v string) *ListDirectoriesResponseBodyDirectories {
	s.Region = &v
	return s
}

func (s *ListDirectoriesResponseBodyDirectories) SetUpdateTime(v string) *ListDirectoriesResponseBodyDirectories {
	s.UpdateTime = &v
	return s
}

func (s *ListDirectoriesResponseBodyDirectories) Validate() error {
	return dara.Validate(s)
}
