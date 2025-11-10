// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFolderName(v string) *CreateFolderRequest
	GetFolderName() *string
	SetParentFolderId(v string) *CreateFolderRequest
	GetParentFolderId() *string
	SetTag(v []*CreateFolderRequestTag) *CreateFolderRequest
	GetTag() []*CreateFolderRequestTag
}

type CreateFolderRequest struct {
	// The name of the folder.
	//
	// The name must be 1 to 24 characters in length and can contain letters, digits, underscores (_), periods (.),and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// rdFolder
	FolderName *string `json:"FolderName,omitempty" xml:"FolderName,omitempty"`
	// The ID of the parent folder.
	//
	// example:
	//
	// r-b1****
	ParentFolderId *string `json:"ParentFolderId,omitempty" xml:"ParentFolderId,omitempty"`
	// The tag to add to the folder.
	Tag []*CreateFolderRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateFolderRequest) GetFolderName() *string {
	return s.FolderName
}

func (s *CreateFolderRequest) GetParentFolderId() *string {
	return s.ParentFolderId
}

func (s *CreateFolderRequest) GetTag() []*CreateFolderRequestTag {
	return s.Tag
}

func (s *CreateFolderRequest) SetFolderName(v string) *CreateFolderRequest {
	s.FolderName = &v
	return s
}

func (s *CreateFolderRequest) SetParentFolderId(v string) *CreateFolderRequest {
	s.ParentFolderId = &v
	return s
}

func (s *CreateFolderRequest) SetTag(v []*CreateFolderRequestTag) *CreateFolderRequest {
	s.Tag = v
	return s
}

func (s *CreateFolderRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateFolderRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// k1
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// v1
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateFolderRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderRequestTag) GoString() string {
	return s.String()
}

func (s *CreateFolderRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateFolderRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateFolderRequestTag) SetKey(v string) *CreateFolderRequestTag {
	s.Key = &v
	return s
}

func (s *CreateFolderRequestTag) SetValue(v string) *CreateFolderRequestTag {
	s.Value = &v
	return s
}

func (s *CreateFolderRequestTag) Validate() error {
	return dara.Validate(s)
}
