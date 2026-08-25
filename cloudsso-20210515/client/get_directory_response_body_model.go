// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDirectory(v *GetDirectoryResponseBodyDirectory) *GetDirectoryResponseBody
	GetDirectory() *GetDirectoryResponseBodyDirectory
	SetRequestId(v string) *GetDirectoryResponseBody
	GetRequestId() *string
}

type GetDirectoryResponseBody struct {
	// The information about the directory.
	Directory *GetDirectoryResponseBodyDirectory `json:"Directory,omitempty" xml:"Directory,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// AA6A9E4B-8A61-59E1-AA87-F61CA18258A3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetDirectoryResponseBody) GetDirectory() *GetDirectoryResponseBodyDirectory {
	return s.Directory
}

func (s *GetDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDirectoryResponseBody) SetDirectory(v *GetDirectoryResponseBodyDirectory) *GetDirectoryResponseBody {
	s.Directory = v
	return s
}

func (s *GetDirectoryResponseBody) SetRequestId(v string) *GetDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDirectoryResponseBody) Validate() error {
	if s.Directory != nil {
		if err := s.Directory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDirectoryResponseBodyDirectory struct {
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
	// example
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
	// 2021-10-25T07:18:46Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetDirectoryResponseBodyDirectory) String() string {
	return dara.Prettify(s)
}

func (s GetDirectoryResponseBodyDirectory) GoString() string {
	return s.String()
}

func (s *GetDirectoryResponseBodyDirectory) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetDirectoryResponseBodyDirectory) GetDirectoryId() *string {
	return s.DirectoryId
}

func (s *GetDirectoryResponseBodyDirectory) GetDirectoryName() *string {
	return s.DirectoryName
}

func (s *GetDirectoryResponseBodyDirectory) GetRegion() *string {
	return s.Region
}

func (s *GetDirectoryResponseBodyDirectory) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetDirectoryResponseBodyDirectory) SetCreateTime(v string) *GetDirectoryResponseBodyDirectory {
	s.CreateTime = &v
	return s
}

func (s *GetDirectoryResponseBodyDirectory) SetDirectoryId(v string) *GetDirectoryResponseBodyDirectory {
	s.DirectoryId = &v
	return s
}

func (s *GetDirectoryResponseBodyDirectory) SetDirectoryName(v string) *GetDirectoryResponseBodyDirectory {
	s.DirectoryName = &v
	return s
}

func (s *GetDirectoryResponseBodyDirectory) SetRegion(v string) *GetDirectoryResponseBodyDirectory {
	s.Region = &v
	return s
}

func (s *GetDirectoryResponseBodyDirectory) SetUpdateTime(v string) *GetDirectoryResponseBodyDirectory {
	s.UpdateTime = &v
	return s
}

func (s *GetDirectoryResponseBodyDirectory) Validate() error {
	return dara.Validate(s)
}
