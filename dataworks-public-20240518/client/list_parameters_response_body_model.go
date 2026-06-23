// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListParametersResponseBodyPagingInfo) *ListParametersResponseBody
	GetPagingInfo() *ListParametersResponseBodyPagingInfo
	SetRequestId(v string) *ListParametersResponseBody
	GetRequestId() *string
}

type ListParametersResponseBody struct {
	// The pagination information.
	PagingInfo *ListParametersResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListParametersResponseBody) GoString() string {
	return s.String()
}

func (s *ListParametersResponseBody) GetPagingInfo() *ListParametersResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListParametersResponseBody) SetPagingInfo(v *ListParametersResponseBodyPagingInfo) *ListParametersResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListParametersResponseBody) SetRequestId(v string) *ListParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParametersResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListParametersResponseBodyPagingInfo struct {
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// A list of parameters.
	Parameters []*ListParametersResponseBodyPagingInfoParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// The total number of parameters.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListParametersResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListParametersResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListParametersResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListParametersResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListParametersResponseBodyPagingInfo) GetParameters() []*ListParametersResponseBodyPagingInfoParameters {
	return s.Parameters
}

func (s *ListParametersResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListParametersResponseBodyPagingInfo) SetPageNumber(v int32) *ListParametersResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfo) SetPageSize(v int32) *ListParametersResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfo) SetParameters(v []*ListParametersResponseBodyPagingInfoParameters) *ListParametersResponseBodyPagingInfo {
	s.Parameters = v
	return s
}

func (s *ListParametersResponseBodyPagingInfo) SetTotalCount(v int32) *ListParametersResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfo) Validate() error {
	if s.Parameters != nil {
		for _, item := range s.Parameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListParametersResponseBodyPagingInfoParameters struct {
	// The time when the parameter was created. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1640000000000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The Alibaba Cloud account ID of the creator.
	//
	// example:
	//
	// 123456789
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// The parameter description.
	//
	// example:
	//
	// This is a test parameter.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The parameter ID.
	//
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The time when the parameter was last modified. This value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1640000000000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The Alibaba Cloud account ID of the user who last modified the parameter.
	//
	// example:
	//
	// 123456789
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// The parameter name.
	//
	// example:
	//
	// workspace.para
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The Alibaba Cloud account ID of the owner.
	//
	// example:
	//
	// 123456789
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The workspace ID.
	//
	// example:
	//
	// 1000
	ProjectId *int64 `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	// The environment-specific values for the parameter.
	Properties []*ListParametersResponseBodyPagingInfoParametersProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// The parameter scope.
	//
	// example:
	//
	// Project
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The parameter type. Valid values:
	//
	// - `PlainConstant`: a plaintext constant
	//
	// - `SecretConstant`: an encrypted constant
	//
	// - `Variable`: a variable
	//
	// example:
	//
	// PlainConstant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The version number.
	//
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListParametersResponseBodyPagingInfoParameters) String() string {
	return dara.Prettify(s)
}

func (s ListParametersResponseBodyPagingInfoParameters) GoString() string {
	return s.String()
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetDescription() *string {
	return s.Description
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetId() *int64 {
	return s.Id
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetName() *string {
	return s.Name
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetOwner() *string {
	return s.Owner
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetProperties() []*ListParametersResponseBodyPagingInfoParametersProperties {
	return s.Properties
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetScope() *string {
	return s.Scope
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetType() *string {
	return s.Type
}

func (s *ListParametersResponseBodyPagingInfoParameters) GetVersion() *int32 {
	return s.Version
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetCreateTime(v int64) *ListParametersResponseBodyPagingInfoParameters {
	s.CreateTime = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetCreateUser(v string) *ListParametersResponseBodyPagingInfoParameters {
	s.CreateUser = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetDescription(v string) *ListParametersResponseBodyPagingInfoParameters {
	s.Description = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetId(v int64) *ListParametersResponseBodyPagingInfoParameters {
	s.Id = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetModifyTime(v int64) *ListParametersResponseBodyPagingInfoParameters {
	s.ModifyTime = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetModifyUser(v string) *ListParametersResponseBodyPagingInfoParameters {
	s.ModifyUser = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetName(v string) *ListParametersResponseBodyPagingInfoParameters {
	s.Name = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetOwner(v string) *ListParametersResponseBodyPagingInfoParameters {
	s.Owner = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetProjectId(v int64) *ListParametersResponseBodyPagingInfoParameters {
	s.ProjectId = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetProperties(v []*ListParametersResponseBodyPagingInfoParametersProperties) *ListParametersResponseBodyPagingInfoParameters {
	s.Properties = v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetScope(v string) *ListParametersResponseBodyPagingInfoParameters {
	s.Scope = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetType(v string) *ListParametersResponseBodyPagingInfoParameters {
	s.Type = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) SetVersion(v int32) *ListParametersResponseBodyPagingInfoParameters {
	s.Version = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParameters) Validate() error {
	if s.Properties != nil {
		for _, item := range s.Properties {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListParametersResponseBodyPagingInfoParametersProperties struct {
	// The environment type. Valid values:
	//
	// - `Prod`: production environment
	//
	// - `Dev`: development environment
	//
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// The parameter value.
	//
	// example:
	//
	// value123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListParametersResponseBodyPagingInfoParametersProperties) String() string {
	return dara.Prettify(s)
}

func (s ListParametersResponseBodyPagingInfoParametersProperties) GoString() string {
	return s.String()
}

func (s *ListParametersResponseBodyPagingInfoParametersProperties) GetEnvType() *string {
	return s.EnvType
}

func (s *ListParametersResponseBodyPagingInfoParametersProperties) GetValue() *string {
	return s.Value
}

func (s *ListParametersResponseBodyPagingInfoParametersProperties) SetEnvType(v string) *ListParametersResponseBodyPagingInfoParametersProperties {
	s.EnvType = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParametersProperties) SetValue(v string) *ListParametersResponseBodyPagingInfoParametersProperties {
	s.Value = &v
	return s
}

func (s *ListParametersResponseBodyPagingInfoParametersProperties) Validate() error {
	return dara.Validate(s)
}
