// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitResourceDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InitResourceDirectoryResponseBody
	GetRequestId() *string
	SetResourceDirectory(v *InitResourceDirectoryResponseBodyResourceDirectory) *InitResourceDirectoryResponseBody
	GetResourceDirectory() *InitResourceDirectoryResponseBodyResourceDirectory
}

type InitResourceDirectoryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CD76D376-2517-4924-92C5-DBC52262F93A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information of the resource directory.
	ResourceDirectory *InitResourceDirectoryResponseBodyResourceDirectory `json:"ResourceDirectory,omitempty" xml:"ResourceDirectory,omitempty" type:"Struct"`
}

func (s InitResourceDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *InitResourceDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitResourceDirectoryResponseBody) GetResourceDirectory() *InitResourceDirectoryResponseBodyResourceDirectory {
	return s.ResourceDirectory
}

func (s *InitResourceDirectoryResponseBody) SetRequestId(v string) *InitResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitResourceDirectoryResponseBody) SetResourceDirectory(v *InitResourceDirectoryResponseBodyResourceDirectory) *InitResourceDirectoryResponseBody {
	s.ResourceDirectory = v
	return s
}

func (s *InitResourceDirectoryResponseBody) Validate() error {
	if s.ResourceDirectory != nil {
		if err := s.ResourceDirectory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type InitResourceDirectoryResponseBodyResourceDirectory struct {
	// The time when the resource directory was enabled.
	//
	// example:
	//
	// 2019-02-18T15:32:10.473Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the enterprise management account.
	//
	// example:
	//
	// 172841235500****
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the enterprise management account.
	//
	// example:
	//
	// aliyun-****
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-Ss****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The ID of the root folder.
	//
	// example:
	//
	// r-Zo****
	RootFolderId *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty"`
}

func (s InitResourceDirectoryResponseBodyResourceDirectory) String() string {
	return dara.Prettify(s)
}

func (s InitResourceDirectoryResponseBodyResourceDirectory) GoString() string {
	return s.String()
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) GetCreateTime() *string {
	return s.CreateTime
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) GetMasterAccountId() *string {
	return s.MasterAccountId
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) GetMasterAccountName() *string {
	return s.MasterAccountName
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) GetRootFolderId() *string {
	return s.RootFolderId
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) SetCreateTime(v string) *InitResourceDirectoryResponseBodyResourceDirectory {
	s.CreateTime = &v
	return s
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountId(v string) *InitResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountId = &v
	return s
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountName(v string) *InitResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountName = &v
	return s
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) SetResourceDirectoryId(v string) *InitResourceDirectoryResponseBodyResourceDirectory {
	s.ResourceDirectoryId = &v
	return s
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) SetRootFolderId(v string) *InitResourceDirectoryResponseBodyResourceDirectory {
	s.RootFolderId = &v
	return s
}

func (s *InitResourceDirectoryResponseBodyResourceDirectory) Validate() error {
	return dara.Validate(s)
}
