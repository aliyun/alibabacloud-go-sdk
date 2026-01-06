// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDingtalkProjectionInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTenantContext(v *GetDingtalkProjectionInfoRequestTenantContext) *GetDingtalkProjectionInfoRequest
	GetTenantContext() *GetDingtalkProjectionInfoRequestTenantContext
	SetClient(v string) *GetDingtalkProjectionInfoRequest
	GetClient() *string
	SetEndTs(v int64) *GetDingtalkProjectionInfoRequest
	GetEndTs() *int64
	SetOrgId(v string) *GetDingtalkProjectionInfoRequest
	GetOrgId() *string
	SetPubWorkNo(v string) *GetDingtalkProjectionInfoRequest
	GetPubWorkNo() *string
	SetRoomId(v string) *GetDingtalkProjectionInfoRequest
	GetRoomId() *string
	SetStartTs(v int64) *GetDingtalkProjectionInfoRequest
	GetStartTs() *int64
	SetSubUid(v string) *GetDingtalkProjectionInfoRequest
	GetSubUid() *string
}

type GetDingtalkProjectionInfoRequest struct {
	TenantContext *GetDingtalkProjectionInfoRequestTenantContext `json:"TenantContext,omitempty" xml:"TenantContext,omitempty" type:"Struct"`
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

func (s GetDingtalkProjectionInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionInfoRequest) GetTenantContext() *GetDingtalkProjectionInfoRequestTenantContext {
	return s.TenantContext
}

func (s *GetDingtalkProjectionInfoRequest) GetClient() *string {
	return s.Client
}

func (s *GetDingtalkProjectionInfoRequest) GetEndTs() *int64 {
	return s.EndTs
}

func (s *GetDingtalkProjectionInfoRequest) GetOrgId() *string {
	return s.OrgId
}

func (s *GetDingtalkProjectionInfoRequest) GetPubWorkNo() *string {
	return s.PubWorkNo
}

func (s *GetDingtalkProjectionInfoRequest) GetRoomId() *string {
	return s.RoomId
}

func (s *GetDingtalkProjectionInfoRequest) GetStartTs() *int64 {
	return s.StartTs
}

func (s *GetDingtalkProjectionInfoRequest) GetSubUid() *string {
	return s.SubUid
}

func (s *GetDingtalkProjectionInfoRequest) SetTenantContext(v *GetDingtalkProjectionInfoRequestTenantContext) *GetDingtalkProjectionInfoRequest {
	s.TenantContext = v
	return s
}

func (s *GetDingtalkProjectionInfoRequest) SetClient(v string) *GetDingtalkProjectionInfoRequest {
	s.Client = &v
	return s
}

func (s *GetDingtalkProjectionInfoRequest) SetEndTs(v int64) *GetDingtalkProjectionInfoRequest {
	s.EndTs = &v
	return s
}

func (s *GetDingtalkProjectionInfoRequest) SetOrgId(v string) *GetDingtalkProjectionInfoRequest {
	s.OrgId = &v
	return s
}

func (s *GetDingtalkProjectionInfoRequest) SetPubWorkNo(v string) *GetDingtalkProjectionInfoRequest {
	s.PubWorkNo = &v
	return s
}

func (s *GetDingtalkProjectionInfoRequest) SetRoomId(v string) *GetDingtalkProjectionInfoRequest {
	s.RoomId = &v
	return s
}

func (s *GetDingtalkProjectionInfoRequest) SetStartTs(v int64) *GetDingtalkProjectionInfoRequest {
	s.StartTs = &v
	return s
}

func (s *GetDingtalkProjectionInfoRequest) SetSubUid(v string) *GetDingtalkProjectionInfoRequest {
	s.SubUid = &v
	return s
}

func (s *GetDingtalkProjectionInfoRequest) Validate() error {
	if s.TenantContext != nil {
		if err := s.TenantContext.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDingtalkProjectionInfoRequestTenantContext struct {
	// example:
	//
	// xxxxxx
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s GetDingtalkProjectionInfoRequestTenantContext) String() string {
	return dara.Prettify(s)
}

func (s GetDingtalkProjectionInfoRequestTenantContext) GoString() string {
	return s.String()
}

func (s *GetDingtalkProjectionInfoRequestTenantContext) GetTenantId() *string {
	return s.TenantId
}

func (s *GetDingtalkProjectionInfoRequestTenantContext) SetTenantId(v string) *GetDingtalkProjectionInfoRequestTenantContext {
	s.TenantId = &v
	return s
}

func (s *GetDingtalkProjectionInfoRequestTenantContext) Validate() error {
	return dara.Validate(s)
}
