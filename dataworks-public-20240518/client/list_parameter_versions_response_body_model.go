// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListParameterVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPagingInfo(v *ListParameterVersionsResponseBodyPagingInfo) *ListParameterVersionsResponseBody
	GetPagingInfo() *ListParameterVersionsResponseBodyPagingInfo
	SetRequestId(v string) *ListParameterVersionsResponseBody
	GetRequestId() *string
}

type ListParameterVersionsResponseBody struct {
	PagingInfo *ListParameterVersionsResponseBodyPagingInfo `json:"PagingInfo,omitempty" xml:"PagingInfo,omitempty" type:"Struct"`
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListParameterVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListParameterVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponseBody) GetPagingInfo() *ListParameterVersionsResponseBodyPagingInfo {
	return s.PagingInfo
}

func (s *ListParameterVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListParameterVersionsResponseBody) SetPagingInfo(v *ListParameterVersionsResponseBodyPagingInfo) *ListParameterVersionsResponseBody {
	s.PagingInfo = v
	return s
}

func (s *ListParameterVersionsResponseBody) SetRequestId(v string) *ListParameterVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListParameterVersionsResponseBody) Validate() error {
	if s.PagingInfo != nil {
		if err := s.PagingInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListParameterVersionsResponseBodyPagingInfo struct {
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize         *int32                                                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	ParameterVersion []*ListParameterVersionsResponseBodyPagingInfoParameterVersion `json:"ParameterVersion,omitempty" xml:"ParameterVersion,omitempty" type:"Repeated"`
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListParameterVersionsResponseBodyPagingInfo) String() string {
	return dara.Prettify(s)
}

func (s ListParameterVersionsResponseBodyPagingInfo) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponseBodyPagingInfo) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListParameterVersionsResponseBodyPagingInfo) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListParameterVersionsResponseBodyPagingInfo) GetParameterVersion() []*ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	return s.ParameterVersion
}

func (s *ListParameterVersionsResponseBodyPagingInfo) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListParameterVersionsResponseBodyPagingInfo) SetPageNumber(v int32) *ListParameterVersionsResponseBodyPagingInfo {
	s.PageNumber = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfo) SetPageSize(v int32) *ListParameterVersionsResponseBodyPagingInfo {
	s.PageSize = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfo) SetParameterVersion(v []*ListParameterVersionsResponseBodyPagingInfoParameterVersion) *ListParameterVersionsResponseBodyPagingInfo {
	s.ParameterVersion = v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfo) SetTotalCount(v int32) *ListParameterVersionsResponseBodyPagingInfo {
	s.TotalCount = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfo) Validate() error {
	if s.ParameterVersion != nil {
		for _, item := range s.ParameterVersion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListParameterVersionsResponseBodyPagingInfoParameterVersion struct {
	// example:
	//
	// 1640000000000
	CreateTime *int64 `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 123456789
	CreateUser *string `json:"CreateUser,omitempty" xml:"CreateUser,omitempty"`
	// example:
	//
	// 这是一个测试参数
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 12345
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// 1640000000000
	ModifyTime *int64 `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 123456789
	ModifyUser *string `json:"ModifyUser,omitempty" xml:"ModifyUser,omitempty"`
	// example:
	//
	// workspace.para
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 123456789
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// example:
	//
	// 1000
	ProjectId  *int64                                                                   `json:"ProjectId,omitempty" xml:"ProjectId,omitempty"`
	Properties []*ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties `json:"Properties,omitempty" xml:"Properties,omitempty" type:"Repeated"`
	// example:
	//
	// Project
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// example:
	//
	// PlainConstant
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// example:
	//
	// 1
	Version *int32 `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListParameterVersionsResponseBodyPagingInfoParameterVersion) String() string {
	return dara.Prettify(s)
}

func (s ListParameterVersionsResponseBodyPagingInfoParameterVersion) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetCreateUser() *string {
	return s.CreateUser
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetDescription() *string {
	return s.Description
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetId() *int64 {
	return s.Id
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetModifyTime() *int64 {
	return s.ModifyTime
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetModifyUser() *string {
	return s.ModifyUser
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetName() *string {
	return s.Name
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetOwner() *string {
	return s.Owner
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetProjectId() *int64 {
	return s.ProjectId
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetProperties() []*ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties {
	return s.Properties
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetScope() *string {
	return s.Scope
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetType() *string {
	return s.Type
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) GetVersion() *int32 {
	return s.Version
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetCreateTime(v int64) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.CreateTime = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetCreateUser(v string) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.CreateUser = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetDescription(v string) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.Description = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetId(v int64) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.Id = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetModifyTime(v int64) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.ModifyTime = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetModifyUser(v string) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.ModifyUser = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetName(v string) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.Name = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetOwner(v string) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.Owner = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetProjectId(v int64) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.ProjectId = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetProperties(v []*ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.Properties = v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetScope(v string) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.Scope = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetType(v string) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.Type = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) SetVersion(v int32) *ListParameterVersionsResponseBodyPagingInfoParameterVersion {
	s.Version = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersion) Validate() error {
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

type ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties struct {
	// example:
	//
	// Prod
	EnvType *string `json:"EnvType,omitempty" xml:"EnvType,omitempty"`
	// example:
	//
	// value123
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties) String() string {
	return dara.Prettify(s)
}

func (s ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties) GoString() string {
	return s.String()
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties) GetEnvType() *string {
	return s.EnvType
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties) GetValue() *string {
	return s.Value
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties) SetEnvType(v string) *ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties {
	s.EnvType = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties) SetValue(v string) *ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties {
	s.Value = &v
	return s
}

func (s *ListParameterVersionsResponseBodyPagingInfoParameterVersionProperties) Validate() error {
	return dara.Validate(s)
}
