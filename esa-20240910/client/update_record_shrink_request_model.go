// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRecordShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthConfShrink(v string) *UpdateRecordShrinkRequest
	GetAuthConfShrink() *string
	SetBizName(v string) *UpdateRecordShrinkRequest
	GetBizName() *string
	SetComment(v string) *UpdateRecordShrinkRequest
	GetComment() *string
	SetDataShrink(v string) *UpdateRecordShrinkRequest
	GetDataShrink() *string
	SetHostPolicy(v string) *UpdateRecordShrinkRequest
	GetHostPolicy() *string
	SetProxied(v bool) *UpdateRecordShrinkRequest
	GetProxied() *bool
	SetRecordId(v int64) *UpdateRecordShrinkRequest
	GetRecordId() *int64
	SetSourceType(v string) *UpdateRecordShrinkRequest
	GetSourceType() *string
	SetTtl(v int32) *UpdateRecordShrinkRequest
	GetTtl() *int32
	SetType(v string) *UpdateRecordShrinkRequest
	GetType() *string
}

type UpdateRecordShrinkRequest struct {
	AuthConfShrink *string `json:"AuthConf,omitempty" xml:"AuthConf,omitempty"`
	BizName        *string `json:"BizName,omitempty" xml:"BizName,omitempty"`
	Comment        *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	DataShrink *string `json:"Data,omitempty" xml:"Data,omitempty"`
	HostPolicy *string `json:"HostPolicy,omitempty" xml:"HostPolicy,omitempty"`
	// 是否代理加速
	Proxied *bool `json:"Proxied,omitempty" xml:"Proxied,omitempty"`
	// This parameter is required.
	RecordId   *int64  `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	SourceType *string `json:"SourceType,omitempty" xml:"SourceType,omitempty"`
	Ttl        *int32  `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
	Type       *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s UpdateRecordShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateRecordShrinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateRecordShrinkRequest) GetAuthConfShrink() *string {
	return s.AuthConfShrink
}

func (s *UpdateRecordShrinkRequest) GetBizName() *string {
	return s.BizName
}

func (s *UpdateRecordShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *UpdateRecordShrinkRequest) GetDataShrink() *string {
	return s.DataShrink
}

func (s *UpdateRecordShrinkRequest) GetHostPolicy() *string {
	return s.HostPolicy
}

func (s *UpdateRecordShrinkRequest) GetProxied() *bool {
	return s.Proxied
}

func (s *UpdateRecordShrinkRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *UpdateRecordShrinkRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateRecordShrinkRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *UpdateRecordShrinkRequest) GetType() *string {
	return s.Type
}

func (s *UpdateRecordShrinkRequest) SetAuthConfShrink(v string) *UpdateRecordShrinkRequest {
	s.AuthConfShrink = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetBizName(v string) *UpdateRecordShrinkRequest {
	s.BizName = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetComment(v string) *UpdateRecordShrinkRequest {
	s.Comment = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetDataShrink(v string) *UpdateRecordShrinkRequest {
	s.DataShrink = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetHostPolicy(v string) *UpdateRecordShrinkRequest {
	s.HostPolicy = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetProxied(v bool) *UpdateRecordShrinkRequest {
	s.Proxied = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetRecordId(v int64) *UpdateRecordShrinkRequest {
	s.RecordId = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetSourceType(v string) *UpdateRecordShrinkRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetTtl(v int32) *UpdateRecordShrinkRequest {
	s.Ttl = &v
	return s
}

func (s *UpdateRecordShrinkRequest) SetType(v string) *UpdateRecordShrinkRequest {
	s.Type = &v
	return s
}

func (s *UpdateRecordShrinkRequest) Validate() error {
	return dara.Validate(s)
}
