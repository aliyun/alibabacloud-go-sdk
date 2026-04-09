// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataPermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDataId(v string) *CreateDataPermissionsRequest
	GetDataId() *string
	SetDataType(v string) *CreateDataPermissionsRequest
	GetDataType() *string
	SetPermissionUserInfos(v []*CreateDataPermissionsRequestPermissionUserInfos) *CreateDataPermissionsRequest
	GetPermissionUserInfos() []*CreateDataPermissionsRequestPermissionUserInfos
	SetWorkspaceId(v string) *CreateDataPermissionsRequest
	GetWorkspaceId() *string
}

type CreateDataPermissionsRequest struct {
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
	PermissionUserInfos []*CreateDataPermissionsRequestPermissionUserInfos `json:"PermissionUserInfos,omitempty" xml:"PermissionUserInfos,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-xxx
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateDataPermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPermissionsRequest) GoString() string {
	return s.String()
}

func (s *CreateDataPermissionsRequest) GetDataId() *string {
	return s.DataId
}

func (s *CreateDataPermissionsRequest) GetDataType() *string {
	return s.DataType
}

func (s *CreateDataPermissionsRequest) GetPermissionUserInfos() []*CreateDataPermissionsRequestPermissionUserInfos {
	return s.PermissionUserInfos
}

func (s *CreateDataPermissionsRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateDataPermissionsRequest) SetDataId(v string) *CreateDataPermissionsRequest {
	s.DataId = &v
	return s
}

func (s *CreateDataPermissionsRequest) SetDataType(v string) *CreateDataPermissionsRequest {
	s.DataType = &v
	return s
}

func (s *CreateDataPermissionsRequest) SetPermissionUserInfos(v []*CreateDataPermissionsRequestPermissionUserInfos) *CreateDataPermissionsRequest {
	s.PermissionUserInfos = v
	return s
}

func (s *CreateDataPermissionsRequest) SetWorkspaceId(v string) *CreateDataPermissionsRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateDataPermissionsRequest) Validate() error {
	if s.PermissionUserInfos != nil {
		for _, item := range s.PermissionUserInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateDataPermissionsRequestPermissionUserInfos struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	PermissionUserId *string `json:"PermissionUserId,omitempty" xml:"PermissionUserId,omitempty"`
	// example:
	//
	// xxx
	PermissionUsername *string `json:"PermissionUsername,omitempty" xml:"PermissionUsername,omitempty"`
}

func (s CreateDataPermissionsRequestPermissionUserInfos) String() string {
	return dara.Prettify(s)
}

func (s CreateDataPermissionsRequestPermissionUserInfos) GoString() string {
	return s.String()
}

func (s *CreateDataPermissionsRequestPermissionUserInfos) GetPermissionUserId() *string {
	return s.PermissionUserId
}

func (s *CreateDataPermissionsRequestPermissionUserInfos) GetPermissionUsername() *string {
	return s.PermissionUsername
}

func (s *CreateDataPermissionsRequestPermissionUserInfos) SetPermissionUserId(v string) *CreateDataPermissionsRequestPermissionUserInfos {
	s.PermissionUserId = &v
	return s
}

func (s *CreateDataPermissionsRequestPermissionUserInfos) SetPermissionUsername(v string) *CreateDataPermissionsRequestPermissionUserInfos {
	s.PermissionUsername = &v
	return s
}

func (s *CreateDataPermissionsRequestPermissionUserInfos) Validate() error {
	return dara.Validate(s)
}
