// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFaceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyFaceGroupRequest
	GetDescription() *string
	SetId(v string) *ModifyFaceGroupRequest
	GetId() *string
	SetName(v string) *ModifyFaceGroupRequest
	GetName() *string
}

type ModifyFaceGroupRequest struct {
	// example:
	//
	// description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 65c030cd54b23283ceb27b4ade5da49d
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test008
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyFaceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFaceGroupRequest) GoString() string {
	return s.String()
}

func (s *ModifyFaceGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFaceGroupRequest) GetId() *string {
	return s.Id
}

func (s *ModifyFaceGroupRequest) GetName() *string {
	return s.Name
}

func (s *ModifyFaceGroupRequest) SetDescription(v string) *ModifyFaceGroupRequest {
	s.Description = &v
	return s
}

func (s *ModifyFaceGroupRequest) SetId(v string) *ModifyFaceGroupRequest {
	s.Id = &v
	return s
}

func (s *ModifyFaceGroupRequest) SetName(v string) *ModifyFaceGroupRequest {
	s.Name = &v
	return s
}

func (s *ModifyFaceGroupRequest) Validate() error {
	return dara.Validate(s)
}
