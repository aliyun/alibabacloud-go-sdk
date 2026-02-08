// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaterialDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *CreateMaterialDirectoryRequest
	GetBizId() *string
	SetName(v string) *CreateMaterialDirectoryRequest
	GetName() *string
	SetParentDirectoryId(v string) *CreateMaterialDirectoryRequest
	GetParentDirectoryId() *string
}

type CreateMaterialDirectoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// WS20260206134746000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 68157a0a-769a-4364-bbdc-a0e2cf3d5ad
	ParentDirectoryId *string `json:"ParentDirectoryId,omitempty" xml:"ParentDirectoryId,omitempty"`
}

func (s CreateMaterialDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateMaterialDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateMaterialDirectoryRequest) GetBizId() *string {
	return s.BizId
}

func (s *CreateMaterialDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateMaterialDirectoryRequest) GetParentDirectoryId() *string {
	return s.ParentDirectoryId
}

func (s *CreateMaterialDirectoryRequest) SetBizId(v string) *CreateMaterialDirectoryRequest {
	s.BizId = &v
	return s
}

func (s *CreateMaterialDirectoryRequest) SetName(v string) *CreateMaterialDirectoryRequest {
	s.Name = &v
	return s
}

func (s *CreateMaterialDirectoryRequest) SetParentDirectoryId(v string) *CreateMaterialDirectoryRequest {
	s.ParentDirectoryId = &v
	return s
}

func (s *CreateMaterialDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
