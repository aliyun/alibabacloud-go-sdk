// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFolderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CreateFolderRequest
	GetName() *string
	SetParent(v int64) *CreateFolderRequest
	GetParent() *int64
	SetType(v string) *CreateFolderRequest
	GetType() *string
}

type CreateFolderRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// gist_qc
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	Parent *int64 `json:"parent,omitempty" xml:"parent,omitempty"`
	// table, instance, template, function
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s CreateFolderRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFolderRequest) GoString() string {
	return s.String()
}

func (s *CreateFolderRequest) GetName() *string {
	return s.Name
}

func (s *CreateFolderRequest) GetParent() *int64 {
	return s.Parent
}

func (s *CreateFolderRequest) GetType() *string {
	return s.Type
}

func (s *CreateFolderRequest) SetName(v string) *CreateFolderRequest {
	s.Name = &v
	return s
}

func (s *CreateFolderRequest) SetParent(v int64) *CreateFolderRequest {
	s.Parent = &v
	return s
}

func (s *CreateFolderRequest) SetType(v string) *CreateFolderRequest {
	s.Type = &v
	return s
}

func (s *CreateFolderRequest) Validate() error {
	return dara.Validate(s)
}
