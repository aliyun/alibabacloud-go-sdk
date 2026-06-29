// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGrantDataServiceApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGrantCommand(v *GrantDataServiceApiRequestGrantCommand) *GrantDataServiceApiRequest
	GetGrantCommand() *GrantDataServiceApiRequestGrantCommand
	SetOpTenantId(v int64) *GrantDataServiceApiRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *GrantDataServiceApiRequest
	GetProjectId() *int32
}

type GrantDataServiceApiRequest struct {
	// The grant command.
	//
	// This parameter is required.
	GrantCommand *GrantDataServiceApiRequestGrantCommand `json:"GrantCommand,omitempty" xml:"GrantCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The data service project ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 102102
	ProjectId *int32 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
}

func (s GrantDataServiceApiRequest) String() string {
	return dara.Prettify(s)
}

func (s GrantDataServiceApiRequest) GoString() string {
	return s.String()
}

func (s *GrantDataServiceApiRequest) GetGrantCommand() *GrantDataServiceApiRequestGrantCommand {
	return s.GrantCommand
}

func (s *GrantDataServiceApiRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GrantDataServiceApiRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *GrantDataServiceApiRequest) SetGrantCommand(v *GrantDataServiceApiRequestGrantCommand) *GrantDataServiceApiRequest {
	s.GrantCommand = v
	return s
}

func (s *GrantDataServiceApiRequest) SetOpTenantId(v int64) *GrantDataServiceApiRequest {
	s.OpTenantId = &v
	return s
}

func (s *GrantDataServiceApiRequest) SetProjectId(v int32) *GrantDataServiceApiRequest {
	s.ProjectId = &v
	return s
}

func (s *GrantDataServiceApiRequest) Validate() error {
	if s.GrantCommand != nil {
		if err := s.GrantCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GrantDataServiceApiRequestGrantCommand struct {
	// The API ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1021
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The application ID.
	//
	// example:
	//
	// 1201
	AppId *int32 `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Specifies whether to request development environment permissions for operation-type APIs.
	//
	// example:
	//
	// true
	ApplyDev *bool `json:"ApplyDev,omitempty" xml:"ApplyDev,omitempty"`
	// Specifies whether to request production environment permissions for operation-type APIs.
	//
	// example:
	//
	// true
	ApplyProd *bool `json:"ApplyProd,omitempty" xml:"ApplyProd,omitempty"`
	// The list of authorization permission types. Valid values:
	//
	// - When the grantee is an application, the following permission types are supported. To grant delegation permissions, you must also grant usage permissions.
	//
	//   - USE: usage permission.
	//
	//   - DELEGATION: delegation permission.
	//
	// - When the grantee is an individual, only USE (usage) permission is supported.
	//
	// - If this parameter is not specified, the default value is USE (usage) permission.
	AuthTypes []*string `json:"AuthTypes,omitempty" xml:"AuthTypes,omitempty" type:"Repeated"`
	// The list of development environment permission fields for query-type APIs. This parameter is required in dev-prod mode. DevFieldList and ProdFieldList cannot both be empty. This parameter is not required for operation-type APIs.
	DevFieldList []*GrantDataServiceApiRequestGrantCommandDevFieldList `json:"DevFieldList,omitempty" xml:"DevFieldList,omitempty" type:"Repeated"`
	// The expiration date in the format of yyyy-MM-dd.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// The authorization object type. Valid values:
	//
	// - APP: application.
	//
	// - USER: user.
	//
	// example:
	//
	// APP
	GranteeType *string `json:"GranteeType,omitempty" xml:"GranteeType,omitempty"`
	// The list of production environment permission fields for query-type APIs. This parameter is required in basic mode. This parameter is not required for operation-type APIs.
	ProdFieldList []*GrantDataServiceApiRequestGrantCommandProdFieldList `json:"ProdFieldList,omitempty" xml:"ProdFieldList,omitempty" type:"Repeated"`
	// The reason for the authorization request.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 12345
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s GrantDataServiceApiRequestGrantCommand) String() string {
	return dara.Prettify(s)
}

func (s GrantDataServiceApiRequestGrantCommand) GoString() string {
	return s.String()
}

func (s *GrantDataServiceApiRequestGrantCommand) GetApiId() *int64 {
	return s.ApiId
}

func (s *GrantDataServiceApiRequestGrantCommand) GetAppId() *int32 {
	return s.AppId
}

func (s *GrantDataServiceApiRequestGrantCommand) GetApplyDev() *bool {
	return s.ApplyDev
}

func (s *GrantDataServiceApiRequestGrantCommand) GetApplyProd() *bool {
	return s.ApplyProd
}

func (s *GrantDataServiceApiRequestGrantCommand) GetAuthTypes() []*string {
	return s.AuthTypes
}

func (s *GrantDataServiceApiRequestGrantCommand) GetDevFieldList() []*GrantDataServiceApiRequestGrantCommandDevFieldList {
	return s.DevFieldList
}

func (s *GrantDataServiceApiRequestGrantCommand) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *GrantDataServiceApiRequestGrantCommand) GetGranteeType() *string {
	return s.GranteeType
}

func (s *GrantDataServiceApiRequestGrantCommand) GetProdFieldList() []*GrantDataServiceApiRequestGrantCommandProdFieldList {
	return s.ProdFieldList
}

func (s *GrantDataServiceApiRequestGrantCommand) GetReason() *string {
	return s.Reason
}

func (s *GrantDataServiceApiRequestGrantCommand) GetUserId() *string {
	return s.UserId
}

func (s *GrantDataServiceApiRequestGrantCommand) SetApiId(v int64) *GrantDataServiceApiRequestGrantCommand {
	s.ApiId = &v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) SetAppId(v int32) *GrantDataServiceApiRequestGrantCommand {
	s.AppId = &v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) SetApplyDev(v bool) *GrantDataServiceApiRequestGrantCommand {
	s.ApplyDev = &v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) SetApplyProd(v bool) *GrantDataServiceApiRequestGrantCommand {
	s.ApplyProd = &v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) SetAuthTypes(v []*string) *GrantDataServiceApiRequestGrantCommand {
	s.AuthTypes = v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) SetDevFieldList(v []*GrantDataServiceApiRequestGrantCommandDevFieldList) *GrantDataServiceApiRequestGrantCommand {
	s.DevFieldList = v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) SetExpireDate(v string) *GrantDataServiceApiRequestGrantCommand {
	s.ExpireDate = &v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) SetGranteeType(v string) *GrantDataServiceApiRequestGrantCommand {
	s.GranteeType = &v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) SetProdFieldList(v []*GrantDataServiceApiRequestGrantCommandProdFieldList) *GrantDataServiceApiRequestGrantCommand {
	s.ProdFieldList = v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) SetReason(v string) *GrantDataServiceApiRequestGrantCommand {
	s.Reason = &v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) SetUserId(v string) *GrantDataServiceApiRequestGrantCommand {
	s.UserId = &v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) Validate() error {
	if s.DevFieldList != nil {
		for _, item := range s.DevFieldList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ProdFieldList != nil {
		for _, item := range s.ProdFieldList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GrantDataServiceApiRequestGrantCommandDevFieldList struct {
	// The API permission field ID.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GrantDataServiceApiRequestGrantCommandDevFieldList) String() string {
	return dara.Prettify(s)
}

func (s GrantDataServiceApiRequestGrantCommandDevFieldList) GoString() string {
	return s.String()
}

func (s *GrantDataServiceApiRequestGrantCommandDevFieldList) GetId() *int32 {
	return s.Id
}

func (s *GrantDataServiceApiRequestGrantCommandDevFieldList) SetId(v int32) *GrantDataServiceApiRequestGrantCommandDevFieldList {
	s.Id = &v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommandDevFieldList) Validate() error {
	return dara.Validate(s)
}

type GrantDataServiceApiRequestGrantCommandProdFieldList struct {
	// The API permission field ID.
	//
	// example:
	//
	// 1
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GrantDataServiceApiRequestGrantCommandProdFieldList) String() string {
	return dara.Prettify(s)
}

func (s GrantDataServiceApiRequestGrantCommandProdFieldList) GoString() string {
	return s.String()
}

func (s *GrantDataServiceApiRequestGrantCommandProdFieldList) GetId() *int32 {
	return s.Id
}

func (s *GrantDataServiceApiRequestGrantCommandProdFieldList) SetId(v int32) *GrantDataServiceApiRequestGrantCommandProdFieldList {
	s.Id = &v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommandProdFieldList) Validate() error {
	return dara.Validate(s)
}
