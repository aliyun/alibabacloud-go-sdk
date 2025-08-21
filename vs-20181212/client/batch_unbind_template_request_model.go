// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *BatchUnbindTemplateRequest
	GetInstanceId() *string
	SetInstanceType(v string) *BatchUnbindTemplateRequest
	GetInstanceType() *string
	SetOwnerId(v int64) *BatchUnbindTemplateRequest
	GetOwnerId() *int64
	SetTemplateId(v string) *BatchUnbindTemplateRequest
	GetTemplateId() *string
	SetTemplateType(v string) *BatchUnbindTemplateRequest
	GetTemplateType() *string
}

type BatchUnbindTemplateRequest struct {
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
	// 323*****998-cn-qingdao
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// record
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s BatchUnbindTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindTemplateRequest) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchUnbindTemplateRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *BatchUnbindTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchUnbindTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *BatchUnbindTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *BatchUnbindTemplateRequest) SetInstanceId(v string) *BatchUnbindTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchUnbindTemplateRequest) SetInstanceType(v string) *BatchUnbindTemplateRequest {
	s.InstanceType = &v
	return s
}

func (s *BatchUnbindTemplateRequest) SetOwnerId(v int64) *BatchUnbindTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchUnbindTemplateRequest) SetTemplateId(v string) *BatchUnbindTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *BatchUnbindTemplateRequest) SetTemplateType(v string) *BatchUnbindTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *BatchUnbindTemplateRequest) Validate() error {
	return dara.Validate(s)
}
