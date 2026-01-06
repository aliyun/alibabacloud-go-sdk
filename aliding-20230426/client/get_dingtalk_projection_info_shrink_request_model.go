// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkProjectionInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContextShrink(v string) *GetDingtalkProjectionInfoShrinkRequest
	GetTenantContextShrink() *string
	SetClient(v string) *GetDingtalkProjectionInfoShrinkRequest
	GetClient() *string
	SetEndTs(v int64) *GetDingtalkProjectionInfoShrinkRequest
	GetEndTs() *int64
	SetOrgId(v string) *GetDingtalkProjectionInfoShrinkRequest
	GetOrgId() *string
	SetPubWorkNo(v string) *GetDingtalkProjectionInfoShrinkRequest
	GetPubWorkNo() *string
	SetRoomId(v string) *GetDingtalkProjectionInfoShrinkRequest
	GetRoomId() *string
	SetStartTs(v int64) *GetDingtalkProjectionInfoShrinkRequest
	GetStartTs() *int64
	SetSubUid(v string) *GetDingtalkProjectionInfoShrinkRequest
	GetSubUid() *string
}

type GetDingtalkProjectionInfoShrinkRequest struct {
	TenantContextShrink *string `json:"TenantContext,omitempty" xml:"TenantContext,omitempty"`
	// example:
	//
	// web
	Client *string `json:"client,omitempty" xml:"client,omitempty"`
	// example:
	//
	// 1640998800000
	EndTs *int64 `json:"endTs,omitempty" xml:"endTs,omitempty"`
	// example:
	//
	// 21001
	OrgId *string `json:"orgId,omitempty" xml:"orgId,omitempty"`
	// example:
	//
	// 342342
	PubWorkNo *string `json:"pubWorkNo,omitempty" xml:"pubWorkNo,omitempty"`
	// example:
	//
	// room001
	RoomId *string `json:"roomId,omitempty" xml:"roomId,omitempty"`
	// example:
	//
	// 1640995200000
	StartTs *int64 `json:"startTs,omitempty" xml:"startTs,omitempty"`
	// example:
	//
	// user456
	SubUid *string `json:"subUid,omitempty" xml:"subUid,omitempty"`
}

func (s GetDingtalkProjectionInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionInfoShrinkRequest) GetTenantContextShrink() *string {
	return s.TenantContextShrink
}

func (s *GetDingtalkProjectionInfoShrinkRequest) GetClient() *string {
	return s.Client
}

func (s *GetDingtalkProjectionInfoShrinkRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *GetDingtalkProjectionInfoShrinkRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingtalkProjectionInfoShrinkRequest) GetPubWorkNo() *string {
	return s.PubWorkNo
}

func (s *GetDingtalkProjectionInfoShrinkRequest) GetRoomId() *string {
	return s.RoomId
}

func (s *GetDingtalkProjectionInfoShrinkRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *GetDingtalkProjectionInfoShrinkRequest) GetSubUid() *string {
	return s.SubUid
}

func (s *GetDingtalkProjectionInfoShrinkRequest) SetTenantContextShrink(v string) *GetDingtalkProjectionInfoShrinkRequest {
	s.TenantContextShrink = &v
	return s
}

func (s *GetDingtalkProjectionInfoShrinkRequest) SetClient(v string) *GetDingtalkProjectionInfoShrinkRequest {
	s.Client = &v
	return s
}

func (s *GetDingtalkProjectionInfoShrinkRequest) SetEndTs(v int64) *GetDingtalkProjectionInfoShrinkRequest {
	s.EndTs = &v
	return s
}

func (s *GetDingtalkProjectionInfoShrinkRequest) SetOrgId(v string) *GetDingtalkProjectionInfoShrinkRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkProjectionInfoShrinkRequest) SetPubWorkNo(v string) *GetDingtalkProjectionInfoShrinkRequest {
	s.PubWorkNo = &v
	return s
}

func (s *GetDingtalkProjectionInfoShrinkRequest) SetRoomId(v string) *GetDingtalkProjectionInfoShrinkRequest {
	s.RoomId = &v
	return s
}

func (s *GetDingtalkProjectionInfoShrinkRequest) SetStartTs(v int64) *GetDingtalkProjectionInfoShrinkRequest {
	s.StartTs = &v
	return s
}

func (s *GetDingtalkProjectionInfoShrinkRequest) SetSubUid(v string) *GetDingtalkProjectionInfoShrinkRequest {
	s.SubUid = &v
	return s
}

func (s *GetDingtalkProjectionInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}
