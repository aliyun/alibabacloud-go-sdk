// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLhSpaceByNameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetLhSpaceByNameResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetLhSpaceByNameResponseBody
	GetErrorMessage() *string
	SetLakehouseSpace(v *GetLhSpaceByNameResponseBodyLakehouseSpace) *GetLhSpaceByNameResponseBody
	GetLakehouseSpace() *GetLhSpaceByNameResponseBodyLakehouseSpace
	SetRequestId(v string) *GetLhSpaceByNameResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLhSpaceByNameResponseBody
	GetSuccess() *bool
}

type GetLhSpaceByNameResponseBody struct {
	// The error code returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The workspace for data warehouse development.
	LakehouseSpace *GetLhSpaceByNameResponseBodyLakehouseSpace `json:"LakehouseSpace,omitempty" xml:"LakehouseSpace,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// FE8CA4A8-AB2D-55B7-BD30-01A4609F40D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**: The request is successful.
	//
	// 	- **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetLhSpaceByNameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLhSpaceByNameResponseBody) GoString() string {
	return s.String()
}

func (s *GetLhSpaceByNameResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetLhSpaceByNameResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetLhSpaceByNameResponseBody) GetLakehouseSpace() *GetLhSpaceByNameResponseBodyLakehouseSpace {
	return s.LakehouseSpace
}

func (s *GetLhSpaceByNameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLhSpaceByNameResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLhSpaceByNameResponseBody) SetErrorCode(v string) *GetLhSpaceByNameResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetLhSpaceByNameResponseBody) SetErrorMessage(v string) *GetLhSpaceByNameResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetLhSpaceByNameResponseBody) SetLakehouseSpace(v *GetLhSpaceByNameResponseBodyLakehouseSpace) *GetLhSpaceByNameResponseBody {
	s.LakehouseSpace = v
	return s
}

func (s *GetLhSpaceByNameResponseBody) SetRequestId(v string) *GetLhSpaceByNameResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLhSpaceByNameResponseBody) SetSuccess(v bool) *GetLhSpaceByNameResponseBody {
	s.Success = &v
	return s
}

func (s *GetLhSpaceByNameResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetLhSpaceByNameResponseBodyLakehouseSpace struct {
	// The ID of the user who creates the workspace.
	//
	// example:
	//
	// 51***
	CreatorId *string `json:"CreatorId,omitempty" xml:"CreatorId,omitempty"`
	// The description of the workspace.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the development database.
	//
	// example:
	//
	// 2435****
	DevDbId *int32 `json:"DevDbId,omitempty" xml:"DevDbId,omitempty"`
	// The type of the database. Valid values:
	//
	// 	- **14**: AnalyticDB for MySQL
	//
	// 	- **18**: AnalyticDB for PostgreSQL
	//
	// example:
	//
	// 14
	DwDbType *string `json:"DwDbType,omitempty" xml:"DwDbType,omitempty"`
	// The ID of the workspace.
	//
	// example:
	//
	// 1
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// Indicates whether the workspace is deleted. Valid values:
	//
	// 	- **true**: The workspace is deleted.
	//
	// 	- **false**: The workspace is not deleted.
	//
	// example:
	//
	// false
	IsDeleted *bool `json:"IsDeleted,omitempty" xml:"IsDeleted,omitempty"`
	// The mode in which the workspace runs. Valid values:
	//
	// 	- **0**: basic mode
	//
	// 	- **1**: standard mode
	//
	// example:
	//
	// 1
	Mode *int32 `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The ID of the production database.
	//
	// example:
	//
	// 2442****
	ProdDbId *int32 `json:"ProdDbId,omitempty" xml:"ProdDbId,omitempty"`
	// The configuration of the workspace. Valid values:
	//
	// 	- **skipManualRunCheck**: No security rule check is required in the trial run phase.
	//
	// 	- **skipPublishApprove**: No approval is required for publishing and O\\&M.
	//
	// example:
	//
	// {\\"skipManualRunCheck\\":true,\\"skipPublishApprove\\":true}
	SpaceConfig *string `json:"SpaceConfig,omitempty" xml:"SpaceConfig,omitempty"`
	// The name of the workspace.
	//
	// example:
	//
	// test_space
	SpaceName *string `json:"SpaceName,omitempty" xml:"SpaceName,omitempty"`
	// The ID of the tenant to which the workspace belongs.
	//
	// example:
	//
	// 3***
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s GetLhSpaceByNameResponseBodyLakehouseSpace) String() string {
	return dara.Prettify(s)
}

func (s GetLhSpaceByNameResponseBodyLakehouseSpace) GoString() string {
	return s.String()
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) GetCreatorId() *string {
	return s.CreatorId
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) GetDescription() *string {
	return s.Description
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) GetDevDbId() *int32 {
	return s.DevDbId
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) GetDwDbType() *string {
	return s.DwDbType
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) GetId() *int64 {
	return s.Id
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) GetIsDeleted() *bool {
	return s.IsDeleted
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) GetMode() *int32 {
	return s.Mode
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) GetProdDbId() *int32 {
	return s.ProdDbId
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) GetSpaceConfig() *string {
	return s.SpaceConfig
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) GetSpaceName() *string {
	return s.SpaceName
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) GetTenantId() *string {
	return s.TenantId
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) SetCreatorId(v string) *GetLhSpaceByNameResponseBodyLakehouseSpace {
	s.CreatorId = &v
	return s
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) SetDescription(v string) *GetLhSpaceByNameResponseBodyLakehouseSpace {
	s.Description = &v
	return s
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) SetDevDbId(v int32) *GetLhSpaceByNameResponseBodyLakehouseSpace {
	s.DevDbId = &v
	return s
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) SetDwDbType(v string) *GetLhSpaceByNameResponseBodyLakehouseSpace {
	s.DwDbType = &v
	return s
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) SetId(v int64) *GetLhSpaceByNameResponseBodyLakehouseSpace {
	s.Id = &v
	return s
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) SetIsDeleted(v bool) *GetLhSpaceByNameResponseBodyLakehouseSpace {
	s.IsDeleted = &v
	return s
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) SetMode(v int32) *GetLhSpaceByNameResponseBodyLakehouseSpace {
	s.Mode = &v
	return s
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) SetProdDbId(v int32) *GetLhSpaceByNameResponseBodyLakehouseSpace {
	s.ProdDbId = &v
	return s
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) SetSpaceConfig(v string) *GetLhSpaceByNameResponseBodyLakehouseSpace {
	s.SpaceConfig = &v
	return s
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) SetSpaceName(v string) *GetLhSpaceByNameResponseBodyLakehouseSpace {
	s.SpaceName = &v
	return s
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) SetTenantId(v string) *GetLhSpaceByNameResponseBodyLakehouseSpace {
	s.TenantId = &v
	return s
}

func (s *GetLhSpaceByNameResponseBodyLakehouseSpace) Validate() error {
	return dara.Validate(s)
}
