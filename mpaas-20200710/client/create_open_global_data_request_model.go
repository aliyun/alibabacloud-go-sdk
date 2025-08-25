// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenGlobalDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateOpenGlobalDataRequest
	GetAppId() *string
	SetAppMaxVersion(v string) *CreateOpenGlobalDataRequest
	GetAppMaxVersion() *string
	SetAppMinVersion(v string) *CreateOpenGlobalDataRequest
	GetAppMinVersion() *string
	SetBizType(v string) *CreateOpenGlobalDataRequest
	GetBizType() *string
	SetExtAttrStr(v string) *CreateOpenGlobalDataRequest
	GetExtAttrStr() *string
	SetMaxUid(v int64) *CreateOpenGlobalDataRequest
	GetMaxUid() *int64
	SetMinUid(v int64) *CreateOpenGlobalDataRequest
	GetMinUid() *int64
	SetOsType(v string) *CreateOpenGlobalDataRequest
	GetOsType() *string
	SetPayload(v string) *CreateOpenGlobalDataRequest
	GetPayload() *string
	SetTenantId(v string) *CreateOpenGlobalDataRequest
	GetTenantId() *string
	SetThirdMsgId(v string) *CreateOpenGlobalDataRequest
	GetThirdMsgId() *string
	SetUids(v string) *CreateOpenGlobalDataRequest
	GetUids() *string
	SetValidTimeEnd(v int64) *CreateOpenGlobalDataRequest
	GetValidTimeEnd() *int64
	SetValidTimeStart(v int64) *CreateOpenGlobalDataRequest
	GetValidTimeStart() *int64
	SetWorkspaceId(v string) *CreateOpenGlobalDataRequest
	GetWorkspaceId() *string
}

type CreateOpenGlobalDataRequest struct {
	// This parameter is required.
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppMaxVersion *string `json:"AppMaxVersion,omitempty" xml:"AppMaxVersion,omitempty"`
	AppMinVersion *string `json:"AppMinVersion,omitempty" xml:"AppMinVersion,omitempty"`
	// This parameter is required.
	BizType    *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	ExtAttrStr *string `json:"ExtAttrStr,omitempty" xml:"ExtAttrStr,omitempty"`
	MaxUid     *int64  `json:"MaxUid,omitempty" xml:"MaxUid,omitempty"`
	MinUid     *int64  `json:"MinUid,omitempty" xml:"MinUid,omitempty"`
	OsType     *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// This parameter is required.
	Payload  *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	ThirdMsgId     *string `json:"ThirdMsgId,omitempty" xml:"ThirdMsgId,omitempty"`
	Uids           *string `json:"Uids,omitempty" xml:"Uids,omitempty"`
	ValidTimeEnd   *int64  `json:"ValidTimeEnd,omitempty" xml:"ValidTimeEnd,omitempty"`
	ValidTimeStart *int64  `json:"ValidTimeStart,omitempty" xml:"ValidTimeStart,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateOpenGlobalDataRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenGlobalDataRequest) GoString() string {
	return s.String()
}

func (s *CreateOpenGlobalDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateOpenGlobalDataRequest) GetAppMaxVersion() *string {
	return s.AppMaxVersion
}

func (s *CreateOpenGlobalDataRequest) GetAppMinVersion() *string {
	return s.AppMinVersion
}

func (s *CreateOpenGlobalDataRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateOpenGlobalDataRequest) GetExtAttrStr() *string {
	return s.ExtAttrStr
}

func (s *CreateOpenGlobalDataRequest) GetMaxUid() *int64 {
	return s.MaxUid
}

func (s *CreateOpenGlobalDataRequest) GetMinUid() *int64 {
	return s.MinUid
}

func (s *CreateOpenGlobalDataRequest) GetOsType() *string {
	return s.OsType
}

func (s *CreateOpenGlobalDataRequest) GetPayload() *string {
	return s.Payload
}

func (s *CreateOpenGlobalDataRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateOpenGlobalDataRequest) GetThirdMsgId() *string {
	return s.ThirdMsgId
}

func (s *CreateOpenGlobalDataRequest) GetUids() *string {
	return s.Uids
}

func (s *CreateOpenGlobalDataRequest) GetValidTimeEnd() *int64 {
	return s.ValidTimeEnd
}

func (s *CreateOpenGlobalDataRequest) GetValidTimeStart() *int64 {
	return s.ValidTimeStart
}

func (s *CreateOpenGlobalDataRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateOpenGlobalDataRequest) SetAppId(v string) *CreateOpenGlobalDataRequest {
	s.AppId = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetAppMaxVersion(v string) *CreateOpenGlobalDataRequest {
	s.AppMaxVersion = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetAppMinVersion(v string) *CreateOpenGlobalDataRequest {
	s.AppMinVersion = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetBizType(v string) *CreateOpenGlobalDataRequest {
	s.BizType = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetExtAttrStr(v string) *CreateOpenGlobalDataRequest {
	s.ExtAttrStr = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetMaxUid(v int64) *CreateOpenGlobalDataRequest {
	s.MaxUid = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetMinUid(v int64) *CreateOpenGlobalDataRequest {
	s.MinUid = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetOsType(v string) *CreateOpenGlobalDataRequest {
	s.OsType = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetPayload(v string) *CreateOpenGlobalDataRequest {
	s.Payload = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetTenantId(v string) *CreateOpenGlobalDataRequest {
	s.TenantId = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetThirdMsgId(v string) *CreateOpenGlobalDataRequest {
	s.ThirdMsgId = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetUids(v string) *CreateOpenGlobalDataRequest {
	s.Uids = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetValidTimeEnd(v int64) *CreateOpenGlobalDataRequest {
	s.ValidTimeEnd = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetValidTimeStart(v int64) *CreateOpenGlobalDataRequest {
	s.ValidTimeStart = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) SetWorkspaceId(v string) *CreateOpenGlobalDataRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateOpenGlobalDataRequest) Validate() error {
	return dara.Validate(s)
}
