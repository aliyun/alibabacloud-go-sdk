// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenClawInstanceVO interface {
	dara.Model
	String() string
	GoString() string
	SetAliyunAccountUid(v string) *OpenClawInstanceVO
	GetAliyunAccountUid() *string
	SetAuthType(v string) *OpenClawInstanceVO
	GetAuthType() *string
	SetBasicAuthPassword(v string) *OpenClawInstanceVO
	GetBasicAuthPassword() *string
	SetBasicAuthUsername(v string) *OpenClawInstanceVO
	GetBasicAuthUsername() *string
	SetCpu(v float64) *OpenClawInstanceVO
	GetCpu() *float64
	SetGmtCreate(v string) *OpenClawInstanceVO
	GetGmtCreate() *string
	SetGmtModified(v string) *OpenClawInstanceVO
	GetGmtModified() *string
	SetImageInfo(v *OpenClawInstanceVOImageInfo) *OpenClawInstanceVO
	GetImageInfo() *OpenClawInstanceVOImageInfo
	SetInstanceDesc(v string) *OpenClawInstanceVO
	GetInstanceDesc() *string
	SetInstanceId(v string) *OpenClawInstanceVO
	GetInstanceId() *string
	SetInstanceName(v string) *OpenClawInstanceVO
	GetInstanceName() *string
	SetInstanceRegion(v string) *OpenClawInstanceVO
	GetInstanceRegion() *string
	SetLastActiveTime(v string) *OpenClawInstanceVO
	GetLastActiveTime() *string
	SetMemorySize(v int32) *OpenClawInstanceVO
	GetMemorySize() *int32
	SetOpenclawToken(v string) *OpenClawInstanceVO
	GetOpenclawToken() *string
	SetOwnerUid(v string) *OpenClawInstanceVO
	GetOwnerUid() *string
	SetPublicDomain(v string) *OpenClawInstanceVO
	GetPublicDomain() *string
	SetStatus(v int32) *OpenClawInstanceVO
	GetStatus() *int32
	SetStatusDesc(v string) *OpenClawInstanceVO
	GetStatusDesc() *string
	SetStatusMessage(v string) *OpenClawInstanceVO
	GetStatusMessage() *string
	SetVariables(v string) *OpenClawInstanceVO
	GetVariables() *string
}

type OpenClawInstanceVO struct {
	AliyunAccountUid  *string                      `json:"AliyunAccountUid,omitempty" xml:"AliyunAccountUid,omitempty"`
	AuthType          *string                      `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	BasicAuthPassword *string                      `json:"BasicAuthPassword,omitempty" xml:"BasicAuthPassword,omitempty"`
	BasicAuthUsername *string                      `json:"BasicAuthUsername,omitempty" xml:"BasicAuthUsername,omitempty"`
	Cpu               *float64                     `json:"Cpu,omitempty" xml:"Cpu,omitempty"`
	GmtCreate         *string                      `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified       *string                      `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	ImageInfo         *OpenClawInstanceVOImageInfo `json:"ImageInfo,omitempty" xml:"ImageInfo,omitempty" type:"Struct"`
	InstanceDesc      *string                      `json:"InstanceDesc,omitempty" xml:"InstanceDesc,omitempty"`
	InstanceId        *string                      `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName      *string                      `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InstanceRegion    *string                      `json:"InstanceRegion,omitempty" xml:"InstanceRegion,omitempty"`
	LastActiveTime    *string                      `json:"LastActiveTime,omitempty" xml:"LastActiveTime,omitempty"`
	MemorySize        *int32                       `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
	OpenclawToken     *string                      `json:"OpenclawToken,omitempty" xml:"OpenclawToken,omitempty"`
	OwnerUid          *string                      `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	PublicDomain      *string                      `json:"PublicDomain,omitempty" xml:"PublicDomain,omitempty"`
	Status            *int32                       `json:"Status,omitempty" xml:"Status,omitempty"`
	StatusDesc        *string                      `json:"StatusDesc,omitempty" xml:"StatusDesc,omitempty"`
	StatusMessage     *string                      `json:"StatusMessage,omitempty" xml:"StatusMessage,omitempty"`
	Variables         *string                      `json:"Variables,omitempty" xml:"Variables,omitempty"`
}

func (s OpenClawInstanceVO) String() string {
	return dara.Prettify(s)
}

func (s OpenClawInstanceVO) GoString() string {
	return s.String()
}

func (s *OpenClawInstanceVO) GetAliyunAccountUid() *string {
	return s.AliyunAccountUid
}

func (s *OpenClawInstanceVO) GetAuthType() *string {
	return s.AuthType
}

func (s *OpenClawInstanceVO) GetBasicAuthPassword() *string {
	return s.BasicAuthPassword
}

func (s *OpenClawInstanceVO) GetBasicAuthUsername() *string {
	return s.BasicAuthUsername
}

func (s *OpenClawInstanceVO) GetCpu() *float64 {
	return s.Cpu
}

func (s *OpenClawInstanceVO) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *OpenClawInstanceVO) GetGmtModified() *string {
	return s.GmtModified
}

func (s *OpenClawInstanceVO) GetImageInfo() *OpenClawInstanceVOImageInfo {
	return s.ImageInfo
}

func (s *OpenClawInstanceVO) GetInstanceDesc() *string {
	return s.InstanceDesc
}

func (s *OpenClawInstanceVO) GetInstanceId() *string {
	return s.InstanceId
}

func (s *OpenClawInstanceVO) GetInstanceName() *string {
	return s.InstanceName
}

func (s *OpenClawInstanceVO) GetInstanceRegion() *string {
	return s.InstanceRegion
}

func (s *OpenClawInstanceVO) GetLastActiveTime() *string {
	return s.LastActiveTime
}

func (s *OpenClawInstanceVO) GetMemorySize() *int32 {
	return s.MemorySize
}

func (s *OpenClawInstanceVO) GetOpenclawToken() *string {
	return s.OpenclawToken
}

func (s *OpenClawInstanceVO) GetOwnerUid() *string {
	return s.OwnerUid
}

func (s *OpenClawInstanceVO) GetPublicDomain() *string {
	return s.PublicDomain
}

func (s *OpenClawInstanceVO) GetStatus() *int32 {
	return s.Status
}

func (s *OpenClawInstanceVO) GetStatusDesc() *string {
	return s.StatusDesc
}

func (s *OpenClawInstanceVO) GetStatusMessage() *string {
	return s.StatusMessage
}

func (s *OpenClawInstanceVO) GetVariables() *string {
	return s.Variables
}

func (s *OpenClawInstanceVO) SetAliyunAccountUid(v string) *OpenClawInstanceVO {
	s.AliyunAccountUid = &v
	return s
}

func (s *OpenClawInstanceVO) SetAuthType(v string) *OpenClawInstanceVO {
	s.AuthType = &v
	return s
}

func (s *OpenClawInstanceVO) SetBasicAuthPassword(v string) *OpenClawInstanceVO {
	s.BasicAuthPassword = &v
	return s
}

func (s *OpenClawInstanceVO) SetBasicAuthUsername(v string) *OpenClawInstanceVO {
	s.BasicAuthUsername = &v
	return s
}

func (s *OpenClawInstanceVO) SetCpu(v float64) *OpenClawInstanceVO {
	s.Cpu = &v
	return s
}

func (s *OpenClawInstanceVO) SetGmtCreate(v string) *OpenClawInstanceVO {
	s.GmtCreate = &v
	return s
}

func (s *OpenClawInstanceVO) SetGmtModified(v string) *OpenClawInstanceVO {
	s.GmtModified = &v
	return s
}

func (s *OpenClawInstanceVO) SetImageInfo(v *OpenClawInstanceVOImageInfo) *OpenClawInstanceVO {
	s.ImageInfo = v
	return s
}

func (s *OpenClawInstanceVO) SetInstanceDesc(v string) *OpenClawInstanceVO {
	s.InstanceDesc = &v
	return s
}

func (s *OpenClawInstanceVO) SetInstanceId(v string) *OpenClawInstanceVO {
	s.InstanceId = &v
	return s
}

func (s *OpenClawInstanceVO) SetInstanceName(v string) *OpenClawInstanceVO {
	s.InstanceName = &v
	return s
}

func (s *OpenClawInstanceVO) SetInstanceRegion(v string) *OpenClawInstanceVO {
	s.InstanceRegion = &v
	return s
}

func (s *OpenClawInstanceVO) SetLastActiveTime(v string) *OpenClawInstanceVO {
	s.LastActiveTime = &v
	return s
}

func (s *OpenClawInstanceVO) SetMemorySize(v int32) *OpenClawInstanceVO {
	s.MemorySize = &v
	return s
}

func (s *OpenClawInstanceVO) SetOpenclawToken(v string) *OpenClawInstanceVO {
	s.OpenclawToken = &v
	return s
}

func (s *OpenClawInstanceVO) SetOwnerUid(v string) *OpenClawInstanceVO {
	s.OwnerUid = &v
	return s
}

func (s *OpenClawInstanceVO) SetPublicDomain(v string) *OpenClawInstanceVO {
	s.PublicDomain = &v
	return s
}

func (s *OpenClawInstanceVO) SetStatus(v int32) *OpenClawInstanceVO {
	s.Status = &v
	return s
}

func (s *OpenClawInstanceVO) SetStatusDesc(v string) *OpenClawInstanceVO {
	s.StatusDesc = &v
	return s
}

func (s *OpenClawInstanceVO) SetStatusMessage(v string) *OpenClawInstanceVO {
	s.StatusMessage = &v
	return s
}

func (s *OpenClawInstanceVO) SetVariables(v string) *OpenClawInstanceVO {
	s.Variables = &v
	return s
}

func (s *OpenClawInstanceVO) Validate() error {
	if s.ImageInfo != nil {
		if err := s.ImageInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OpenClawInstanceVOImageInfo struct {
	ImageId     *int64  `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Namespace   *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	Tag         *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
	VersionDesc *string `json:"VersionDesc,omitempty" xml:"VersionDesc,omitempty"`
}

func (s OpenClawInstanceVOImageInfo) String() string {
	return dara.Prettify(s)
}

func (s OpenClawInstanceVOImageInfo) GoString() string {
	return s.String()
}

func (s *OpenClawInstanceVOImageInfo) GetImageId() *int64 {
	return s.ImageId
}

func (s *OpenClawInstanceVOImageInfo) GetName() *string {
	return s.Name
}

func (s *OpenClawInstanceVOImageInfo) GetNamespace() *string {
	return s.Namespace
}

func (s *OpenClawInstanceVOImageInfo) GetTag() *string {
	return s.Tag
}

func (s *OpenClawInstanceVOImageInfo) GetVersionDesc() *string {
	return s.VersionDesc
}

func (s *OpenClawInstanceVOImageInfo) SetImageId(v int64) *OpenClawInstanceVOImageInfo {
	s.ImageId = &v
	return s
}

func (s *OpenClawInstanceVOImageInfo) SetName(v string) *OpenClawInstanceVOImageInfo {
	s.Name = &v
	return s
}

func (s *OpenClawInstanceVOImageInfo) SetNamespace(v string) *OpenClawInstanceVOImageInfo {
	s.Namespace = &v
	return s
}

func (s *OpenClawInstanceVOImageInfo) SetTag(v string) *OpenClawInstanceVOImageInfo {
	s.Tag = &v
	return s
}

func (s *OpenClawInstanceVOImageInfo) SetVersionDesc(v string) *OpenClawInstanceVOImageInfo {
	s.VersionDesc = &v
	return s
}

func (s *OpenClawInstanceVOImageInfo) Validate() error {
	return dara.Validate(s)
}
