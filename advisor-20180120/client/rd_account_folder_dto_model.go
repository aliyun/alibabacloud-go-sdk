// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRdAccountFolderDTO interface {
	dara.Model
	String() string
	GoString() string
	SetAccountCount(v int32) *RdAccountFolderDTO
	GetAccountCount() *int32
	SetAccountList(v []*RdAccountDTO) *RdAccountFolderDTO
	GetAccountList() []*RdAccountDTO
	SetFolderId(v string) *RdAccountFolderDTO
	GetFolderId() *string
	SetFolderList(v []*RdAccountFolderDTO) *RdAccountFolderDTO
	GetFolderList() []*RdAccountFolderDTO
	SetFolderName(v string) *RdAccountFolderDTO
	GetFolderName() *string
	SetResourceDirectoryId(v string) *RdAccountFolderDTO
	GetResourceDirectoryId() *string
	SetResourceDirectoryPath(v string) *RdAccountFolderDTO
	GetResourceDirectoryPath() *string
	SetResourceDirectoryPathName(v string) *RdAccountFolderDTO
	GetResourceDirectoryPathName() *string
	SetSelectedCount(v int32) *RdAccountFolderDTO
	GetSelectedCount() *int32
}

type RdAccountFolderDTO struct {
	AccountCount              *int32                `json:"AccountCount,omitempty" xml:"AccountCount,omitempty"`
	AccountList               []*RdAccountDTO       `json:"AccountList,omitempty" xml:"AccountList,omitempty" type:"Repeated"`
	FolderId                  *string               `json:"FolderId,omitempty" xml:"FolderId,omitempty"`
	FolderList                []*RdAccountFolderDTO `json:"FolderList,omitempty" xml:"FolderList,omitempty" type:"Repeated"`
	FolderName                *string               `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	ResourceDirectoryId       *string               `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	ResourceDirectoryPath     *string               `json:"ResourceDirectoryPath,omitempty" xml:"ResourceDirectoryPath,omitempty"`
	ResourceDirectoryPathName *string               `json:"ResourceDirectoryPathName,omitempty" xml:"ResourceDirectoryPathName,omitempty"`
	SelectedCount             *int32                `json:"SelectedCount,omitempty" xml:"SelectedCount,omitempty"`
}

func (s RdAccountFolderDTO) String() string {
	return dara.Prettify(s)
}

func (s RdAccountFolderDTO) GoString() string {
	return s.String()
}

func (s *RdAccountFolderDTO) GetAccountCount() *int32 {
	return s.AccountCount
}

func (s *RdAccountFolderDTO) GetAccountList() []*RdAccountDTO {
	return s.AccountList
}

func (s *RdAccountFolderDTO) GetFolderId() *string {
	return s.FolderId
}

func (s *RdAccountFolderDTO) GetFolderList() []*RdAccountFolderDTO {
	return s.FolderList
}

func (s *RdAccountFolderDTO) GetFolderName() *string {
	return s.FolderName
}

func (s *RdAccountFolderDTO) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *RdAccountFolderDTO) GetResourceDirectoryPath() *string {
	return s.ResourceDirectoryPath
}

func (s *RdAccountFolderDTO) GetResourceDirectoryPathName() *string {
	return s.ResourceDirectoryPathName
}

func (s *RdAccountFolderDTO) GetSelectedCount() *int32 {
	return s.SelectedCount
}

func (s *RdAccountFolderDTO) SetAccountCount(v int32) *RdAccountFolderDTO {
	s.AccountCount = &v
	return s
}

func (s *RdAccountFolderDTO) SetAccountList(v []*RdAccountDTO) *RdAccountFolderDTO {
	s.AccountList = v
	return s
}

func (s *RdAccountFolderDTO) SetFolderId(v string) *RdAccountFolderDTO {
	s.FolderId = &v
	return s
}

func (s *RdAccountFolderDTO) SetFolderList(v []*RdAccountFolderDTO) *RdAccountFolderDTO {
	s.FolderList = v
	return s
}

func (s *RdAccountFolderDTO) SetFolderName(v string) *RdAccountFolderDTO {
	s.FolderName = &v
	return s
}

func (s *RdAccountFolderDTO) SetResourceDirectoryId(v string) *RdAccountFolderDTO {
	s.ResourceDirectoryId = &v
	return s
}

func (s *RdAccountFolderDTO) SetResourceDirectoryPath(v string) *RdAccountFolderDTO {
	s.ResourceDirectoryPath = &v
	return s
}

func (s *RdAccountFolderDTO) SetResourceDirectoryPathName(v string) *RdAccountFolderDTO {
	s.ResourceDirectoryPathName = &v
	return s
}

func (s *RdAccountFolderDTO) SetSelectedCount(v int32) *RdAccountFolderDTO {
	s.SelectedCount = &v
	return s
}

func (s *RdAccountFolderDTO) Validate() error {
	if s.AccountList != nil {
		for _, item := range s.AccountList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.FolderList != nil {
		for _, item := range s.FolderList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
