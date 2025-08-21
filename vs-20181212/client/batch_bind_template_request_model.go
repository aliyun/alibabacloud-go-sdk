// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchBindTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplyAll(v bool) *BatchBindTemplateRequest
	GetApplyAll() *bool
	SetInstanceId(v string) *BatchBindTemplateRequest
	GetInstanceId() *string
	SetInstanceType(v string) *BatchBindTemplateRequest
	GetInstanceType() *string
	SetOwnerId(v int64) *BatchBindTemplateRequest
	GetOwnerId() *int64
	SetReplace(v bool) *BatchBindTemplateRequest
	GetReplace() *bool
	SetTemplateId(v string) *BatchBindTemplateRequest
	GetTemplateId() *string
}

type BatchBindTemplateRequest struct {
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
}

func (s BatchBindTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchBindTemplateRequest) GoString() string {
	return s.String()
}

func (s *BatchBindTemplateRequest) GetApplyAll() *bool {
	return s.ApplyAll
}

func (s *BatchBindTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchBindTemplateRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *BatchBindTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchBindTemplateRequest) GetReplace() *bool {
	return s.Replace
}

func (s *BatchBindTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *BatchBindTemplateRequest) SetApplyAll(v bool) *BatchBindTemplateRequest {
	s.ApplyAll = &v
	return s
}

func (s *BatchBindTemplateRequest) SetInstanceId(v string) *BatchBindTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchBindTemplateRequest) SetInstanceType(v string) *BatchBindTemplateRequest {
	s.InstanceType = &v
	return s
}

func (s *BatchBindTemplateRequest) SetOwnerId(v int64) *BatchBindTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchBindTemplateRequest) SetReplace(v bool) *BatchBindTemplateRequest {
	s.Replace = &v
	return s
}

func (s *BatchBindTemplateRequest) SetTemplateId(v string) *BatchBindTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *BatchBindTemplateRequest) Validate() error {
	return dara.Validate(s)
}
