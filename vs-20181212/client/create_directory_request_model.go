// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDirectoryRequest
	GetDescription() *string
	SetGroupId(v string) *CreateDirectoryRequest
	GetGroupId() *string
	SetName(v string) *CreateDirectoryRequest
	GetName() *string
	SetOwnerId(v int64) *CreateDirectoryRequest
	GetOwnerId() *int64
	SetParentId(v string) *CreateDirectoryRequest
	GetParentId() *string
}

type CreateDirectoryRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 348*****174-cn-qingdao
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// This parameter is required.
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 399*****774-cn-qingdao
	ParentId *string `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s CreateDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDirectoryRequest) GoString() string {
	return s.String()
}

func (s *CreateDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDirectoryRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *CreateDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *CreateDirectoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateDirectoryRequest) GetParentId() *string {
	return s.ParentId
}

func (s *CreateDirectoryRequest) SetDescription(v string) *CreateDirectoryRequest {
	s.Description = &v
	return s
}

func (s *CreateDirectoryRequest) SetGroupId(v string) *CreateDirectoryRequest {
	s.GroupId = &v
	return s
}

func (s *CreateDirectoryRequest) SetName(v string) *CreateDirectoryRequest {
	s.Name = &v
	return s
}

func (s *CreateDirectoryRequest) SetOwnerId(v int64) *CreateDirectoryRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateDirectoryRequest) SetParentId(v string) *CreateDirectoryRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
