// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataServiceAppGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupName(v string) *CreateDataServiceAppGroupRequest
	GetGroupName() *string
	SetOpTenantId(v int64) *CreateDataServiceAppGroupRequest
	GetOpTenantId() *int64
}

type CreateDataServiceAppGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// default_app_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s CreateDataServiceAppGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataServiceAppGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateDataServiceAppGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *CreateDataServiceAppGroupRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *CreateDataServiceAppGroupRequest) SetGroupName(v string) *CreateDataServiceAppGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateDataServiceAppGroupRequest) SetOpTenantId(v int64) *CreateDataServiceAppGroupRequest {
	s.OpTenantId = &v
	return s
}

func (s *CreateDataServiceAppGroupRequest) Validate() error {
	return dara.Validate(s)
}
