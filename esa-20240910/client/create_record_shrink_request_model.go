// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfShrink(v string) *CreateRecordShrinkRequest
	GetAuthConfShrink() *string
	SetBizName(v string) *CreateRecordShrinkRequest
	GetBizName() *string
	SetComment(v string) *CreateRecordShrinkRequest
	GetComment() *string
	SetDataShrink(v string) *CreateRecordShrinkRequest
	GetDataShrink() *string
	SetHostPolicy(v string) *CreateRecordShrinkRequest
	GetHostPolicy() *string
	SetProxied(v bool) *CreateRecordShrinkRequest
	GetProxied() *bool
	SetRecordName(v string) *CreateRecordShrinkRequest
	GetRecordName() *string
	SetSiteId(v int64) *CreateRecordShrinkRequest
	GetSiteId() *int64
	SetSourceType(v string) *CreateRecordShrinkRequest
	GetSourceType() *string
	SetTtl(v int32) *CreateRecordShrinkRequest
	GetTtl() *int32
	SetType(v string) *CreateRecordShrinkRequest
	GetType() *string
}

type CreateRecordShrinkRequest struct {
	AuthConfShrink *string `json:"AuthConf,omitempty" xml:"AuthConf,omitempty"`
	// 业务场景
	BizName *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	DataShrink *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// 是否代理加速
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// 记录名称
	//
	// This parameter is required.
	RecordName *string `json:"RecordName,omitempty" xml:"RecordName,omitempty"`
	// This parameter is required.
	SiteId     *int64  `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	// This parameter is required.
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	// 记录类型
	//
	// This parameter is required.
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateRecordShrinkRequest) GetAuthConfShrink() *string {
	return s.AuthConfShrink
}

func (s *CreateRecordShrinkRequest) GetBizName() *string {
	return s.BizName
}

func (s *CreateRecordShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateRecordShrinkRequest) GetDataShrink() *string {
	return s.DataShrink
}

func (s *CreateRecordShrinkRequest) GetHostPolicy() *string {
	return s.HostPolicy
}

func (s *CreateRecordShrinkRequest) GetProxied() *bool {
	return s.Proxied
}

func (s *CreateRecordShrinkRequest) GetRecordName() *string {
	return s.RecordName
}

func (s *CreateRecordShrinkRequest) GetSiteId() *int64 {
	return s.SiteId
}

func (s *CreateRecordShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *CreateRecordShrinkRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *CreateRecordShrinkRequest) GetType() *string {
	return s.Type
}

func (s *CreateRecordShrinkRequest) SetAuthConfShrink(v string) *CreateRecordShrinkRequest {
	s.AuthConfShrink = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetBizName(v string) *CreateRecordShrinkRequest {
	s.BizName = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetComment(v string) *CreateRecordShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetDataShrink(v string) *CreateRecordShrinkRequest {
	s.DataShrink = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetHostPolicy(v string) *CreateRecordShrinkRequest {
	s.HostPolicy = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetProxied(v bool) *CreateRecordShrinkRequest {
	s.Proxied = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetRecordName(v string) *CreateRecordShrinkRequest {
	s.RecordName = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetSiteId(v int64) *CreateRecordShrinkRequest {
	s.SiteId = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetSourceType(v string) *CreateRecordShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetTtl(v int32) *CreateRecordShrinkRequest {
	s.Ttl = &v
	return s
}

func (s *CreateRecordShrinkRequest) SetType(v string) *CreateRecordShrinkRequest {
	s.Type = &v
	return s
}

func (s *CreateRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
