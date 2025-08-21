// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBindTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyAll(v bool) *BindTemplateRequest
	GetApplyAll() *bool
	SetInstanceId(v string) *BindTemplateRequest
	GetInstanceId() *string
	SetInstanceType(v string) *BindTemplateRequest
	GetInstanceType() *string
	SetOwnerId(v int64) *BindTemplateRequest
	GetOwnerId() *int64
	SetReplace(v bool) *BindTemplateRequest
	GetReplace() *bool
	SetTemplateId(v string) *BindTemplateRequest
	GetTemplateId() *string
	SetTemplateType(v string) *BindTemplateRequest
	GetTemplateType() *string
}

type BindTemplateRequest struct {
	// example:
	//
	// false
	ApplyAll *bool `json:"ApplyAll,omitempty" xml:"ApplyAll,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 323*****994-cn-qingdao
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// group
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// false
	Replace *bool `json:"Replace,omitempty" xml:"Replace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 323*****998-cn-qingdao
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// record
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s BindTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s BindTemplateRequest) GoString() string {
	return s.String()
}

func (s *BindTemplateRequest) GetApplyAll() *bool {
	return s.ApplyAll
}

func (s *BindTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BindTemplateRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *BindTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BindTemplateRequest) GetReplace() *bool {
	return s.Replace
}

func (s *BindTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *BindTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *BindTemplateRequest) SetApplyAll(v bool) *BindTemplateRequest {
	s.ApplyAll = &v
	return s
}

func (s *BindTemplateRequest) SetInstanceId(v string) *BindTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *BindTemplateRequest) SetInstanceType(v string) *BindTemplateRequest {
	s.InstanceType = &v
	return s
}

func (s *BindTemplateRequest) SetOwnerId(v int64) *BindTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *BindTemplateRequest) SetReplace(v bool) *BindTemplateRequest {
	s.Replace = &v
	return s
}

func (s *BindTemplateRequest) SetTemplateId(v string) *BindTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *BindTemplateRequest) SetTemplateType(v string) *BindTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *BindTemplateRequest) Validate() error {
	return dara.Validate(s)
}
