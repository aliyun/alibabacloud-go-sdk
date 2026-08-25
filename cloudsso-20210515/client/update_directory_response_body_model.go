// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDirectory(v *UpdateDirectoryResponseBodyDirectory) *UpdateDirectoryResponseBody
	GetDirectory() *UpdateDirectoryResponseBodyDirectory
	SetRequestId(v string) *UpdateDirectoryResponseBody
	GetRequestId() *string
}

type UpdateDirectoryResponseBody struct {
	// The information about the directory.
	Directory *UpdateDirectoryResponseBodyDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B182C041-8C64-5F2F-A07B-FC67FAF89CF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDirectoryResponseBody) GetDirectory() *UpdateDirectoryResponseBodyDirectory {
	return s.Directory
}

func (s *UpdateDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDirectoryResponseBody) SetDirectory(v *UpdateDirectoryResponseBodyDirectory) *UpdateDirectoryResponseBody {
	s.Directory = v
	return s
}

func (s *UpdateDirectoryResponseBody) SetRequestId(v string) *UpdateDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDirectoryResponseBody) Validate() error {
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDirectoryResponseBodyDirectory struct {
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

func (s UpdateDirectoryResponseBodyDirectory) String() string {
	return dara.Prettify(s)
}

func (s UpdateDirectoryResponseBodyDirectory) GoString() string {
	return s.String()
}

func (s *UpdateDirectoryResponseBodyDirectory) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateDirectoryResponseBodyDirectory) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *UpdateDirectoryResponseBodyDirectory) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *UpdateDirectoryResponseBodyDirectory) GetRegion() *string {
	return s.Region
}

func (s *UpdateDirectoryResponseBodyDirectory) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateDirectoryResponseBodyDirectory) SetCreateTime(v string) *UpdateDirectoryResponseBodyDirectory {
	s.CreateTime = &v
	return s
}

func (s *UpdateDirectoryResponseBodyDirectory) SetDirectoryId(v string) *UpdateDirectoryResponseBodyDirectory {
	s.DirectoryId = &v
	return s
}

func (s *UpdateDirectoryResponseBodyDirectory) SetDirectoryName(v string) *UpdateDirectoryResponseBodyDirectory {
	s.DirectoryName = &v
	return s
}

func (s *UpdateDirectoryResponseBodyDirectory) SetRegion(v string) *UpdateDirectoryResponseBodyDirectory {
	s.Region = &v
	return s
}

func (s *UpdateDirectoryResponseBodyDirectory) SetUpdateTime(v string) *UpdateDirectoryResponseBodyDirectory {
	s.UpdateTime = &v
	return s
}

func (s *UpdateDirectoryResponseBodyDirectory) Validate() error {
	return dara.Validate(s)
}
