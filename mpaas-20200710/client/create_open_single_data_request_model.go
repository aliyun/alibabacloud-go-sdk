// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOpenSingleDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *CreateOpenSingleDataRequest
	GetAppId() *string
	SetAppMaxVersion(v string) *CreateOpenSingleDataRequest
	GetAppMaxVersion() *string
	SetAppMinVersion(v string) *CreateOpenSingleDataRequest
	GetAppMinVersion() *string
	SetBizType(v string) *CreateOpenSingleDataRequest
	GetBizType() *string
	SetCheckOnline(v bool) *CreateOpenSingleDataRequest
	GetCheckOnline() *bool
	SetExtAttrStr(v string) *CreateOpenSingleDataRequest
	GetExtAttrStr() *string
	SetLinkToken(v string) *CreateOpenSingleDataRequest
	GetLinkToken() *string
	SetOsType(v string) *CreateOpenSingleDataRequest
	GetOsType() *string
	SetPayload(v string) *CreateOpenSingleDataRequest
	GetPayload() *string
	SetTenantId(v string) *CreateOpenSingleDataRequest
	GetTenantId() *string
	SetThirdMsgId(v string) *CreateOpenSingleDataRequest
	GetThirdMsgId() *string
	SetValidTimeEnd(v int64) *CreateOpenSingleDataRequest
	GetValidTimeEnd() *int64
	SetValidTimeStart(v int64) *CreateOpenSingleDataRequest
	GetValidTimeStart() *int64
	SetWorkspaceId(v string) *CreateOpenSingleDataRequest
	GetWorkspaceId() *string
}

type CreateOpenSingleDataRequest struct {
	// This parameter is required.
	AppId         *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppMaxVersion *string `json:"AppMaxVersion,omitempty" xml:"AppMaxVersion,omitempty"`
	AppMinVersion *string `json:"AppMinVersion,omitempty" xml:"AppMinVersion,omitempty"`
	// This parameter is required.
	BizType     *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	CheckOnline *bool   `json:"CheckOnline,omitempty" xml:"CheckOnline,omitempty"`
	ExtAttrStr  *string `json:"ExtAttrStr,omitempty" xml:"ExtAttrStr,omitempty"`
	// This parameter is required.
	LinkToken *string `json:"LinkToken,omitempty" xml:"LinkToken,omitempty"`
	OsType    *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// This parameter is required.
	Payload  *string `json:"Payload,omitempty" xml:"Payload,omitempty"`
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
	// This parameter is required.
	ThirdMsgId     *string `json:"ThirdMsgId,omitempty" xml:"ThirdMsgId,omitempty"`
	ValidTimeEnd   *int64  `json:"ValidTimeEnd,omitempty" xml:"ValidTimeEnd,omitempty"`
	ValidTimeStart *int64  `json:"ValidTimeStart,omitempty" xml:"ValidTimeStart,omitempty"`
	// This parameter is required.
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s CreateOpenSingleDataRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateOpenSingleDataRequest) GoString() string {
	return s.String()
}

func (s *CreateOpenSingleDataRequest) GetAppId() *string {
	return s.AppId
}

func (s *CreateOpenSingleDataRequest) GetAppMaxVersion() *string {
	return s.AppMaxVersion
}

func (s *CreateOpenSingleDataRequest) GetAppMinVersion() *string {
	return s.AppMinVersion
}

func (s *CreateOpenSingleDataRequest) GetBizType() *string {
	return s.BizType
}

func (s *CreateOpenSingleDataRequest) GetCheckOnline() *bool {
	return s.CheckOnline
}

func (s *CreateOpenSingleDataRequest) GetExtAttrStr() *string {
	return s.ExtAttrStr
}

func (s *CreateOpenSingleDataRequest) GetLinkToken() *string {
	return s.LinkToken
}

func (s *CreateOpenSingleDataRequest) GetOsType() *string {
	return s.OsType
}

func (s *CreateOpenSingleDataRequest) GetPayload() *string {
	return s.Payload
}

func (s *CreateOpenSingleDataRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateOpenSingleDataRequest) GetThirdMsgId() *string {
	return s.ThirdMsgId
}

func (s *CreateOpenSingleDataRequest) GetValidTimeEnd() *int64 {
	return s.ValidTimeEnd
}

func (s *CreateOpenSingleDataRequest) GetValidTimeStart() *int64 {
	return s.ValidTimeStart
}

func (s *CreateOpenSingleDataRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *CreateOpenSingleDataRequest) SetAppId(v string) *CreateOpenSingleDataRequest {
	s.AppId = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetAppMaxVersion(v string) *CreateOpenSingleDataRequest {
	s.AppMaxVersion = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetAppMinVersion(v string) *CreateOpenSingleDataRequest {
	s.AppMinVersion = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetBizType(v string) *CreateOpenSingleDataRequest {
	s.BizType = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetCheckOnline(v bool) *CreateOpenSingleDataRequest {
	s.CheckOnline = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetExtAttrStr(v string) *CreateOpenSingleDataRequest {
	s.ExtAttrStr = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetLinkToken(v string) *CreateOpenSingleDataRequest {
	s.LinkToken = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetOsType(v string) *CreateOpenSingleDataRequest {
	s.OsType = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetPayload(v string) *CreateOpenSingleDataRequest {
	s.Payload = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetTenantId(v string) *CreateOpenSingleDataRequest {
	s.TenantId = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetThirdMsgId(v string) *CreateOpenSingleDataRequest {
	s.ThirdMsgId = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetValidTimeEnd(v int64) *CreateOpenSingleDataRequest {
	s.ValidTimeEnd = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetValidTimeStart(v int64) *CreateOpenSingleDataRequest {
	s.ValidTimeStart = &v
	return s
}

func (s *CreateOpenSingleDataRequest) SetWorkspaceId(v string) *CreateOpenSingleDataRequest {
	s.WorkspaceId = &v
	return s
}

func (s *CreateOpenSingleDataRequest) Validate() error {
	return dara.Validate(s)
}
