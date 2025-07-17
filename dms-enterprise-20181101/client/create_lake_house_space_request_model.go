// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakeHouseSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateLakeHouseSpaceRequest
	GetDescription() *string
	SetDevDbId(v string) *CreateLakeHouseSpaceRequest
	GetDevDbId() *string
	SetDwDbType(v string) *CreateLakeHouseSpaceRequest
	GetDwDbType() *string
	SetMode(v string) *CreateLakeHouseSpaceRequest
	GetMode() *string
	SetProdDbId(v string) *CreateLakeHouseSpaceRequest
	GetProdDbId() *string
	SetSpaceConfig(v string) *CreateLakeHouseSpaceRequest
	GetSpaceConfig() *string
	SetSpaceName(v string) *CreateLakeHouseSpaceRequest
	GetSpaceName() *string
	SetTid(v int64) *CreateLakeHouseSpaceRequest
	GetTid() *int64
}

type CreateLakeHouseSpaceRequest struct {
	// The description of the workspace.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the development database. You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2435****
	DevDbId *string `json:"DevDbId,omitempty" xml:"DevDbId,omitempty"`
	// The type of the database. Valid values:
	//
	// 	- **14**: AnalyticDB for MySQL
	//
	// 	- **18**: AnalyticDB for PostgreSQL
	//
	// This parameter is required.
	//
	// example:
	//
	// 14
	DwDbType *string `json:"DwDbType,omitempty" xml:"DwDbType,omitempty"`
	// The mode in which the workspace runs. Valid values:
	//
	// 	- **0**: basic mode. This mode is unavailable.
	//
	// 	- **1**: standard mode.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The ID of the production database. You can call the [ListDatabases](https://help.aliyun.com/document_detail/141873.html) or [SearchDatabase](https://help.aliyun.com/document_detail/141876.html) operation to obtain the ID.
	//
	// example:
	//
	// 2442****
	ProdDbId *string `json:"ProdDbId,omitempty" xml:"ProdDbId,omitempty"`
	// The configuration of the workspace. Valid values:
	//
	// 	- **skipManualRunCheck**: No security rule check is required in the trial run phase.
	//
	// 	- **skipPublishApprove**: No approval is required for publishing and O\\&M.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"skipManualRunCheck\\":true,\\"skipPublishApprove\\":true}
	SpaceConfig *string `json:"SpaceConfig,omitempty" xml:"SpaceConfig,omitempty"`
	// The name of the workspace.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_space
	SpaceName *string `json:"SpaceName,omitempty" xml:"SpaceName,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateLakeHouseSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLakeHouseSpaceRequest) GoString() string {
	return s.String()
}

func (s *CreateLakeHouseSpaceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLakeHouseSpaceRequest) GetDevDbId() *string {
	return s.DevDbId
}

func (s *CreateLakeHouseSpaceRequest) GetDwDbType() *string {
	return s.DwDbType
}

func (s *CreateLakeHouseSpaceRequest) GetMode() *string {
	return s.Mode
}

func (s *CreateLakeHouseSpaceRequest) GetProdDbId() *string {
	return s.ProdDbId
}

func (s *CreateLakeHouseSpaceRequest) GetSpaceConfig() *string {
	return s.SpaceConfig
}

func (s *CreateLakeHouseSpaceRequest) GetSpaceName() *string {
	return s.SpaceName
}

func (s *CreateLakeHouseSpaceRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateLakeHouseSpaceRequest) SetDescription(v string) *CreateLakeHouseSpaceRequest {
	s.Description = &v
	return s
}

func (s *CreateLakeHouseSpaceRequest) SetDevDbId(v string) *CreateLakeHouseSpaceRequest {
	s.DevDbId = &v
	return s
}

func (s *CreateLakeHouseSpaceRequest) SetDwDbType(v string) *CreateLakeHouseSpaceRequest {
	s.DwDbType = &v
	return s
}

func (s *CreateLakeHouseSpaceRequest) SetMode(v string) *CreateLakeHouseSpaceRequest {
	s.Mode = &v
	return s
}

func (s *CreateLakeHouseSpaceRequest) SetProdDbId(v string) *CreateLakeHouseSpaceRequest {
	s.ProdDbId = &v
	return s
}

func (s *CreateLakeHouseSpaceRequest) SetSpaceConfig(v string) *CreateLakeHouseSpaceRequest {
	s.SpaceConfig = &v
	return s
}

func (s *CreateLakeHouseSpaceRequest) SetSpaceName(v string) *CreateLakeHouseSpaceRequest {
	s.SpaceName = &v
	return s
}

func (s *CreateLakeHouseSpaceRequest) SetTid(v int64) *CreateLakeHouseSpaceRequest {
	s.Tid = &v
	return s
}

func (s *CreateLakeHouseSpaceRequest) Validate() error {
	return dara.Validate(s)
}
