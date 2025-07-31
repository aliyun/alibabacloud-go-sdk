// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyDataServiceApiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyCommand(v *ApplyDataServiceApiRequestApplyCommand) *ApplyDataServiceApiRequest
	GetApplyCommand() *ApplyDataServiceApiRequestApplyCommand
	SetOpTenantId(v int64) *ApplyDataServiceApiRequest
	GetOpTenantId() *int64
	SetProjectId(v int32) *ApplyDataServiceApiRequest
	GetProjectId() *int32
}

type ApplyDataServiceApiRequest struct {
	// This parameter is required.
	ApplyCommand *ApplyDataServiceApiRequestApplyCommand `json:"ApplyCommand,omitempty" xml:"ApplyCommand,omitempty" type:"Struct"`
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

func (s ApplyDataServiceApiRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataServiceApiRequest) GoString() string {
	return s.String()
}

func (s *ApplyDataServiceApiRequest) GetApplyCommand() *ApplyDataServiceApiRequestApplyCommand {
	return s.ApplyCommand
}

func (s *ApplyDataServiceApiRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ApplyDataServiceApiRequest) GetProjectId() *int32 {
	return s.ProjectId
}

func (s *ApplyDataServiceApiRequest) SetApplyCommand(v *ApplyDataServiceApiRequestApplyCommand) *ApplyDataServiceApiRequest {
	s.ApplyCommand = v
	return s
}

func (s *ApplyDataServiceApiRequest) SetOpTenantId(v int64) *ApplyDataServiceApiRequest {
	s.OpTenantId = &v
	return s
}

func (s *ApplyDataServiceApiRequest) SetProjectId(v int32) *ApplyDataServiceApiRequest {
	s.ProjectId = &v
	return s
}

func (s *ApplyDataServiceApiRequest) Validate() error {
	return dara.Validate(s)
}

type ApplyDataServiceApiRequestApplyCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 1021
	ApiId *int64 `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// AppId
	//
	// This parameter is required.
	//
	// example:
	//
	// 1203
	AppId        *int32                                                `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DevFieldList []*ApplyDataServiceApiRequestApplyCommandDevFieldList `json:"DevFieldList,omitempty" xml:"DevFieldList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 2025-06-30
	ExpireDate    *string                                                `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	ProdFieldList []*ApplyDataServiceApiRequestApplyCommandProdFieldList `json:"ProdFieldList,omitempty" xml:"ProdFieldList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
}

func (s ApplyDataServiceApiRequestApplyCommand) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataServiceApiRequestApplyCommand) GoString() string {
	return s.String()
}

func (s *ApplyDataServiceApiRequestApplyCommand) GetApiId() *int64 {
	return s.ApiId
}

func (s *ApplyDataServiceApiRequestApplyCommand) GetAppId() *int32 {
	return s.AppId
}

func (s *ApplyDataServiceApiRequestApplyCommand) GetDevFieldList() []*ApplyDataServiceApiRequestApplyCommandDevFieldList {
	return s.DevFieldList
}

func (s *ApplyDataServiceApiRequestApplyCommand) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *ApplyDataServiceApiRequestApplyCommand) GetProdFieldList() []*ApplyDataServiceApiRequestApplyCommandProdFieldList {
	return s.ProdFieldList
}

func (s *ApplyDataServiceApiRequestApplyCommand) GetReason() *string {
	return s.Reason
}

func (s *ApplyDataServiceApiRequestApplyCommand) SetApiId(v int64) *ApplyDataServiceApiRequestApplyCommand {
	s.ApiId = &v
	return s
}

func (s *ApplyDataServiceApiRequestApplyCommand) SetAppId(v int32) *ApplyDataServiceApiRequestApplyCommand {
	s.AppId = &v
	return s
}

func (s *ApplyDataServiceApiRequestApplyCommand) SetDevFieldList(v []*ApplyDataServiceApiRequestApplyCommandDevFieldList) *ApplyDataServiceApiRequestApplyCommand {
	s.DevFieldList = v
	return s
}

func (s *ApplyDataServiceApiRequestApplyCommand) SetExpireDate(v string) *ApplyDataServiceApiRequestApplyCommand {
	s.ExpireDate = &v
	return s
}

func (s *ApplyDataServiceApiRequestApplyCommand) SetProdFieldList(v []*ApplyDataServiceApiRequestApplyCommandProdFieldList) *ApplyDataServiceApiRequestApplyCommand {
	s.ProdFieldList = v
	return s
}

func (s *ApplyDataServiceApiRequestApplyCommand) SetReason(v string) *ApplyDataServiceApiRequestApplyCommand {
	s.Reason = &v
	return s
}

func (s *ApplyDataServiceApiRequestApplyCommand) Validate() error {
	return dara.Validate(s)
}

type ApplyDataServiceApiRequestApplyCommandDevFieldList struct {
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ApplyDataServiceApiRequestApplyCommandDevFieldList) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataServiceApiRequestApplyCommandDevFieldList) GoString() string {
	return s.String()
}

func (s *ApplyDataServiceApiRequestApplyCommandDevFieldList) GetId() *int32 {
	return s.Id
}

func (s *ApplyDataServiceApiRequestApplyCommandDevFieldList) SetId(v int32) *ApplyDataServiceApiRequestApplyCommandDevFieldList {
	s.Id = &v
	return s
}

func (s *ApplyDataServiceApiRequestApplyCommandDevFieldList) Validate() error {
	return dara.Validate(s)
}

type ApplyDataServiceApiRequestApplyCommandProdFieldList struct {
	// This parameter is required.
	//
	// example:
	//
	// 22
	Id *int32 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ApplyDataServiceApiRequestApplyCommandProdFieldList) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataServiceApiRequestApplyCommandProdFieldList) GoString() string {
	return s.String()
}

func (s *ApplyDataServiceApiRequestApplyCommandProdFieldList) GetId() *int32 {
	return s.Id
}

func (s *ApplyDataServiceApiRequestApplyCommandProdFieldList) SetId(v int32) *ApplyDataServiceApiRequestApplyCommandProdFieldList {
	s.Id = &v
	return s
}

func (s *ApplyDataServiceApiRequestApplyCommandProdFieldList) Validate() error {
	return dara.Validate(s)
}
