// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDirectory(v *CreateDirectoryResponseBodyDirectory) *CreateDirectoryResponseBody
	GetDirectory() *CreateDirectoryResponseBodyDirectory
	SetRequestId(v string) *CreateDirectoryResponseBody
	GetRequestId() *string
}

type CreateDirectoryResponseBody struct {
	// The information about the directory.
	Directory *CreateDirectoryResponseBodyDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ADADC31D-90EE-5459-99B0-D83DF07769A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDirectoryResponseBody) GetDirectory() *CreateDirectoryResponseBodyDirectory {
	return s.Directory
}

func (s *CreateDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDirectoryResponseBody) SetDirectory(v *CreateDirectoryResponseBodyDirectory) *CreateDirectoryResponseBody {
	s.Directory = v
	return s
}

func (s *CreateDirectoryResponseBody) SetRequestId(v string) *CreateDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDirectoryResponseBody) Validate() error {
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDirectoryResponseBodyDirectory struct {
	// The time when the directory was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-10-10T04:04:04Z
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
	// example
	DirectoryName *string `json:"DirectoryName,omitempty" xml:"DirectoryName,omitempty"`
	// The region ID of the directory.
	//
	// example:
	//
	// cn-shanghai
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The time when the directory was modified. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-10-10T04:04:04Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s CreateDirectoryResponseBodyDirectory) String() string {
	return dara.Prettify(s)
}

func (s CreateDirectoryResponseBodyDirectory) GoString() string {
	return s.String()
}

func (s *CreateDirectoryResponseBodyDirectory) GetCreateTime() *string {
	return s.CreateTime
}

func (s *CreateDirectoryResponseBodyDirectory) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *CreateDirectoryResponseBodyDirectory) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *CreateDirectoryResponseBodyDirectory) GetRegion() *string {
	return s.Region
}

func (s *CreateDirectoryResponseBodyDirectory) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *CreateDirectoryResponseBodyDirectory) SetCreateTime(v string) *CreateDirectoryResponseBodyDirectory {
	s.CreateTime = &v
	return s
}

func (s *CreateDirectoryResponseBodyDirectory) SetDirectoryId(v string) *CreateDirectoryResponseBodyDirectory {
	s.DirectoryId = &v
	return s
}

func (s *CreateDirectoryResponseBodyDirectory) SetDirectoryName(v string) *CreateDirectoryResponseBodyDirectory {
	s.DirectoryName = &v
	return s
}

func (s *CreateDirectoryResponseBodyDirectory) SetRegion(v string) *CreateDirectoryResponseBodyDirectory {
	s.Region = &v
	return s
}

func (s *CreateDirectoryResponseBodyDirectory) SetUpdateTime(v string) *CreateDirectoryResponseBodyDirectory {
	s.UpdateTime = &v
	return s
}

func (s *CreateDirectoryResponseBodyDirectory) Validate() error {
	return dara.Validate(s)
}
