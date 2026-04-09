// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataPermissionsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataId(v string) *CreateDataPermissionsShrinkRequest
	GetDataId() *string
	SetDataType(v string) *CreateDataPermissionsShrinkRequest
	GetDataType() *string
	SetPermissionUserInfosShrink(v string) *CreateDataPermissionsShrinkRequest
	GetPermissionUserInfosShrink() *string
	SetWorkspaceId(v string) *CreateDataPermissionsShrinkRequest
	GetWorkspaceId() *string
}

type CreateDataPermissionsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// SystemSearch.QuarkCommonNews
	DataId *string `json:"DataId,omitempty" xml:"DataId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// dataset
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// This parameter is required.
	PermissionUserInfosShrink *string `json:"PermissionUserInfos,omitempty" xml:"PermissionUserInfos,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataPermissionsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPermissionsShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataPermissionsShrinkRequest) GetDataId() *string {
	return s.DataId
}

func (s *CreateDataPermissionsShrinkRequest) GetDataType() *string {
	return s.DataType
}

func (s *CreateDataPermissionsShrinkRequest) GetPermissionUserInfosShrink() *string {
	return s.PermissionUserInfosShrink
}

func (s *CreateDataPermissionsShrinkRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDataPermissionsShrinkRequest) SetDataId(v string) *CreateDataPermissionsShrinkRequest {
	s.DataId = &v
	return s
}

func (s *CreateDataPermissionsShrinkRequest) SetDataType(v string) *CreateDataPermissionsShrinkRequest {
	s.DataType = &v
	return s
}

func (s *CreateDataPermissionsShrinkRequest) SetPermissionUserInfosShrink(v string) *CreateDataPermissionsShrinkRequest {
	s.PermissionUserInfosShrink = &v
	return s
}

func (s *CreateDataPermissionsShrinkRequest) SetWorkspaceId(v string) *CreateDataPermissionsShrinkRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataPermissionsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
