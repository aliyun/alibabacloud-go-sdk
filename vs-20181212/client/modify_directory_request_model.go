// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDirectoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyDirectoryRequest
	GetDescription() *string
	SetId(v string) *ModifyDirectoryRequest
	GetId() *string
	SetName(v string) *ModifyDirectoryRequest
	GetName() *string
	SetOwnerId(v int64) *ModifyDirectoryRequest
	GetOwnerId() *int64
}

type ModifyDirectoryRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 399*****488-cn-qingdao
	Id      *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name    *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
}

func (s ModifyDirectoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDirectoryRequest) GoString() string {
	return s.String()
}

func (s *ModifyDirectoryRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDirectoryRequest) GetId() *string {
	return s.Id
}

func (s *ModifyDirectoryRequest) GetName() *string {
	return s.Name
}

func (s *ModifyDirectoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDirectoryRequest) SetDescription(v string) *ModifyDirectoryRequest {
	s.Description = &v
	return s
}

func (s *ModifyDirectoryRequest) SetId(v string) *ModifyDirectoryRequest {
	s.Id = &v
	return s
}

func (s *ModifyDirectoryRequest) SetName(v string) *ModifyDirectoryRequest {
	s.Name = &v
	return s
}

func (s *ModifyDirectoryRequest) SetOwnerId(v int64) *ModifyDirectoryRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDirectoryRequest) Validate() error {
	return dara.Validate(s)
}
