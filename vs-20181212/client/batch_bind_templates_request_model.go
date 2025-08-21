// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyAll(v bool) *BatchBindTemplatesRequest
	GetApplyAll() *bool
	SetInstanceId(v string) *BatchBindTemplatesRequest
	GetInstanceId() *string
	SetInstanceType(v string) *BatchBindTemplatesRequest
	GetInstanceType() *string
	SetOwnerId(v int64) *BatchBindTemplatesRequest
	GetOwnerId() *int64
	SetReplace(v bool) *BatchBindTemplatesRequest
	GetReplace() *bool
	SetTemplateId(v string) *BatchBindTemplatesRequest
	GetTemplateId() *string
	SetTemplateType(v string) *BatchBindTemplatesRequest
	GetTemplateType() *string
}

type BatchBindTemplatesRequest struct {
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
	// stream
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// example:
	//
	// true
	Replace *bool `json:"Replace,omitempty" xml:"Replace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 323*****998-cn-qingdao
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// timeshift
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s BatchBindTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchBindTemplatesRequest) GoString() string {
	return s.String()
}

func (s *BatchBindTemplatesRequest) GetApplyAll() *bool {
	return s.ApplyAll
}

func (s *BatchBindTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchBindTemplatesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *BatchBindTemplatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchBindTemplatesRequest) GetReplace() *bool {
	return s.Replace
}

func (s *BatchBindTemplatesRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *BatchBindTemplatesRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *BatchBindTemplatesRequest) SetApplyAll(v bool) *BatchBindTemplatesRequest {
	s.ApplyAll = &v
	return s
}

func (s *BatchBindTemplatesRequest) SetInstanceId(v string) *BatchBindTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchBindTemplatesRequest) SetInstanceType(v string) *BatchBindTemplatesRequest {
	s.InstanceType = &v
	return s
}

func (s *BatchBindTemplatesRequest) SetOwnerId(v int64) *BatchBindTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchBindTemplatesRequest) SetReplace(v bool) *BatchBindTemplatesRequest {
	s.Replace = &v
	return s
}

func (s *BatchBindTemplatesRequest) SetTemplateId(v string) *BatchBindTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *BatchBindTemplatesRequest) SetTemplateType(v string) *BatchBindTemplatesRequest {
	s.TemplateType = &v
	return s
}

func (s *BatchBindTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
