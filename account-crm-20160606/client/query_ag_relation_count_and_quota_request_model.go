// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAgRelationCountAndQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *QueryAgRelationCountAndQuotaRequest
	GetAppName() *string
	SetCallerBid(v int64) *QueryAgRelationCountAndQuotaRequest
	GetCallerBid() *int64
	SetCallerParentId(v int64) *QueryAgRelationCountAndQuotaRequest
	GetCallerParentId() *int64
	SetCallerType(v string) *QueryAgRelationCountAndQuotaRequest
	GetCallerType() *string
	SetCallerUid(v int64) *QueryAgRelationCountAndQuotaRequest
	GetCallerUid() *int64
	SetMpk(v string) *QueryAgRelationCountAndQuotaRequest
	GetMpk() *string
	SetNullObject(v bool) *QueryAgRelationCountAndQuotaRequest
	GetNullObject() *bool
	SetRequestId(v string) *QueryAgRelationCountAndQuotaRequest
	GetRequestId() *string
	SetSecurityToken(v string) *QueryAgRelationCountAndQuotaRequest
	GetSecurityToken() *string
	SetSourceIp(v string) *QueryAgRelationCountAndQuotaRequest
	GetSourceIp() *string
	SetStsTokenCallerBid(v int64) *QueryAgRelationCountAndQuotaRequest
	GetStsTokenCallerBid() *int64
	SetStsTokenCallerUid(v int64) *QueryAgRelationCountAndQuotaRequest
	GetStsTokenCallerUid() *int64
	SetStsTokenRoleId(v int64) *QueryAgRelationCountAndQuotaRequest
	GetStsTokenRoleId() *int64
	SetVersion(v string) *QueryAgRelationCountAndQuotaRequest
	GetVersion() *string
}

type QueryAgRelationCountAndQuotaRequest struct {
	AppName           *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	CallerBid         *int64  `json:"CallerBid,omitempty" xml:"CallerBid,omitempty"`
	CallerParentId    *int64  `json:"CallerParentId,omitempty" xml:"CallerParentId,omitempty"`
	CallerType        *string `json:"CallerType,omitempty" xml:"CallerType,omitempty"`
	CallerUid         *int64  `json:"CallerUid,omitempty" xml:"CallerUid,omitempty"`
	Mpk               *string `json:"Mpk,omitempty" xml:"Mpk,omitempty"`
	NullObject        *bool   `json:"NullObject,omitempty" xml:"NullObject,omitempty"`
	RequestId         *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityToken     *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	SourceIp          *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	StsTokenCallerBid *int64  `json:"StsTokenCallerBid,omitempty" xml:"StsTokenCallerBid,omitempty"`
	StsTokenCallerUid *int64  `json:"StsTokenCallerUid,omitempty" xml:"StsTokenCallerUid,omitempty"`
	StsTokenRoleId    *int64  `json:"StsTokenRoleId,omitempty" xml:"StsTokenRoleId,omitempty"`
	Version           *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s QueryAgRelationCountAndQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryAgRelationCountAndQuotaRequest) GoString() string {
	return s.String()
}

func (s *QueryAgRelationCountAndQuotaRequest) GetAppName() *string {
	return s.AppName
}

func (s *QueryAgRelationCountAndQuotaRequest) GetCallerBid() *int64 {
	return s.CallerBid
}

func (s *QueryAgRelationCountAndQuotaRequest) GetCallerParentId() *int64 {
	return s.CallerParentId
}

func (s *QueryAgRelationCountAndQuotaRequest) GetCallerType() *string {
	return s.CallerType
}

func (s *QueryAgRelationCountAndQuotaRequest) GetCallerUid() *int64 {
	return s.CallerUid
}

func (s *QueryAgRelationCountAndQuotaRequest) GetMpk() *string {
	return s.Mpk
}

func (s *QueryAgRelationCountAndQuotaRequest) GetNullObject() *bool {
	return s.NullObject
}

func (s *QueryAgRelationCountAndQuotaRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAgRelationCountAndQuotaRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *QueryAgRelationCountAndQuotaRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *QueryAgRelationCountAndQuotaRequest) GetStsTokenCallerBid() *int64 {
	return s.StsTokenCallerBid
}

func (s *QueryAgRelationCountAndQuotaRequest) GetStsTokenCallerUid() *int64 {
	return s.StsTokenCallerUid
}

func (s *QueryAgRelationCountAndQuotaRequest) GetStsTokenRoleId() *int64 {
	return s.StsTokenRoleId
}

func (s *QueryAgRelationCountAndQuotaRequest) GetVersion() *string {
	return s.Version
}

func (s *QueryAgRelationCountAndQuotaRequest) SetAppName(v string) *QueryAgRelationCountAndQuotaRequest {
	s.AppName = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetCallerBid(v int64) *QueryAgRelationCountAndQuotaRequest {
	s.CallerBid = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetCallerParentId(v int64) *QueryAgRelationCountAndQuotaRequest {
	s.CallerParentId = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetCallerType(v string) *QueryAgRelationCountAndQuotaRequest {
	s.CallerType = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetCallerUid(v int64) *QueryAgRelationCountAndQuotaRequest {
	s.CallerUid = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetMpk(v string) *QueryAgRelationCountAndQuotaRequest {
	s.Mpk = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetNullObject(v bool) *QueryAgRelationCountAndQuotaRequest {
	s.NullObject = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetRequestId(v string) *QueryAgRelationCountAndQuotaRequest {
	s.RequestId = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetSecurityToken(v string) *QueryAgRelationCountAndQuotaRequest {
	s.SecurityToken = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetSourceIp(v string) *QueryAgRelationCountAndQuotaRequest {
	s.SourceIp = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetStsTokenCallerBid(v int64) *QueryAgRelationCountAndQuotaRequest {
	s.StsTokenCallerBid = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetStsTokenCallerUid(v int64) *QueryAgRelationCountAndQuotaRequest {
	s.StsTokenCallerUid = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetStsTokenRoleId(v int64) *QueryAgRelationCountAndQuotaRequest {
	s.StsTokenRoleId = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) SetVersion(v string) *QueryAgRelationCountAndQuotaRequest {
	s.Version = &v
	return s
}

func (s *QueryAgRelationCountAndQuotaRequest) Validate() error {
	return dara.Validate(s)
}
