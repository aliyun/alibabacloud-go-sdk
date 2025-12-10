// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEnableResourceDirectoryResponseBody interface {
  dara.Model
  String() string
  GoString() string
  SetRequestId(v string) *EnableResourceDirectoryResponseBody
  GetRequestId() *string 
  SetResourceDirectory(v *EnableResourceDirectoryResponseBodyResourceDirectory) *EnableResourceDirectoryResponseBody
  GetResourceDirectory() *EnableResourceDirectoryResponseBodyResourceDirectory 
}

type EnableResourceDirectoryResponseBody struct {
  // The ID of the request.
  // 
  // example:
  // 
  // EC2FE94D-A4A2-51A1-A493-5C273A36C46A
  RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
  // The information of the resource directory.
  ResourceDirectory *EnableResourceDirectoryResponseBodyResourceDirectory `json:"ResourceDirectory,omitempty" xml:"ResourceDirectory,omitempty" type:"Struct"`
}

func (s EnableResourceDirectoryResponseBody) String() string {
  return dara.Prettify(s)
}

func (s EnableResourceDirectoryResponseBody) GoString() string {
  return s.String()
}

func (s *EnableResourceDirectoryResponseBody) GetRequestId() *string  {
  return s.RequestId
}

func (s *EnableResourceDirectoryResponseBody) GetResourceDirectory() *EnableResourceDirectoryResponseBodyResourceDirectory  {
  return s.ResourceDirectory
}

func (s *EnableResourceDirectoryResponseBody) SetRequestId(v string) *EnableResourceDirectoryResponseBody {
  s.RequestId = &v
  return s
}

func (s *EnableResourceDirectoryResponseBody) SetResourceDirectory(v *EnableResourceDirectoryResponseBodyResourceDirectory) *EnableResourceDirectoryResponseBody {
  s.ResourceDirectory = v
  return s
}

func (s *EnableResourceDirectoryResponseBody) Validate() error {
  if s.ResourceDirectory != nil {
    if err := s.ResourceDirectory.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EnableResourceDirectoryResponseBodyResourceDirectory struct {
  // The time when the resource directory was enabled.
  // 
  // example:
  // 
  // 2021-12-08T02:15:31.744Z
  CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
  // The ID of the management account.
  // 
  // example:
  // 
  // 507408460615****
  MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
  // The name of the management account.
  // 
  // example:
  // 
  // alice@example.com
  MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
  // The ID of the resource directory.
  // 
  // example:
  // 
  // rd-54****
  ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
  // The ID of the Root folder.
  // 
  // example:
  // 
  // r-G9****
  RootFolderId *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty"`
}

func (s EnableResourceDirectoryResponseBodyResourceDirectory) String() string {
  return dara.Prettify(s)
}

func (s EnableResourceDirectoryResponseBodyResourceDirectory) GoString() string {
  return s.String()
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) GetCreateTime() *string  {
  return s.CreateTime
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) GetMasterAccountId() *string  {
  return s.MasterAccountId
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) GetMasterAccountName() *string  {
  return s.MasterAccountName
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) GetResourceDirectoryId() *string  {
  return s.ResourceDirectoryId
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) GetRootFolderId() *string  {
  return s.RootFolderId
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetCreateTime(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
  s.CreateTime = &v
  return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountId(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
  s.MasterAccountId = &v
  return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountName(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
  s.MasterAccountName = &v
  return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetResourceDirectoryId(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
  s.ResourceDirectoryId = &v
  return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) SetRootFolderId(v string) *EnableResourceDirectoryResponseBodyResourceDirectory {
  s.RootFolderId = &v
  return s
}

func (s *EnableResourceDirectoryResponseBodyResourceDirectory) Validate() error {
  return dara.Validate(s)
}

