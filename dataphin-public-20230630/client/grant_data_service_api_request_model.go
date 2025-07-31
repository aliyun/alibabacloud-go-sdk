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
	// This parameter is required.
	GrantCommand *GrantDataServiceApiRequestGrantCommand `json:"GrantCommand,omitempty" xml:"GrantCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
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
	return dara.Validate(s)
}

type GrantDataServiceApiRequestGrantCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 1021
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// AppID
	//
	// This parameter is required.
	//
	// example:
	//
	// 1201
	AppId        *int32                                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DevFieldList []*GrantDataServiceApiRequestGrantCommandDevFieldList `json:"DevFieldList,omitempty" xml:"DevFieldList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30
	ExpireDate    *string                                                `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	ProdFieldList []*GrantDataServiceApiRequestGrantCommandProdFieldList `json:"ProdFieldList,omitempty" xml:"ProdFieldList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
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

func (s *GrantDataServiceApiRequestGrantCommand) GetDevFieldList() []*GrantDataServiceApiRequestGrantCommandDevFieldList {
	return s.DevFieldList
}

func (s *GrantDataServiceApiRequestGrantCommand) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *GrantDataServiceApiRequestGrantCommand) GetProdFieldList() []*GrantDataServiceApiRequestGrantCommandProdFieldList {
	return s.ProdFieldList
}

func (s *GrantDataServiceApiRequestGrantCommand) GetReason() *string {
	return s.Reason
}

func (s *GrantDataServiceApiRequestGrantCommand) SetApiId(v int64) *GrantDataServiceApiRequestGrantCommand {
	s.ApiId = &v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) SetAppId(v int32) *GrantDataServiceApiRequestGrantCommand {
	s.AppId = &v
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

func (s *GrantDataServiceApiRequestGrantCommand) SetProdFieldList(v []*GrantDataServiceApiRequestGrantCommandProdFieldList) *GrantDataServiceApiRequestGrantCommand {
	s.ProdFieldList = v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) SetReason(v string) *GrantDataServiceApiRequestGrantCommand {
	s.Reason = &v
	return s
}

func (s *GrantDataServiceApiRequestGrantCommand) Validate() error {
	return dara.Validate(s)
}

type GrantDataServiceApiRequestGrantCommandDevFieldList struct {
	// This parameter is required.
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
	// This parameter is required.
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
