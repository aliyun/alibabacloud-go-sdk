// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *UnbindTemplateRequest
	GetInstanceId() *string
	SetInstanceType(v string) *UnbindTemplateRequest
	GetInstanceType() *string
	SetOwnerId(v int64) *UnbindTemplateRequest
	GetOwnerId() *int64
	SetTemplateId(v string) *UnbindTemplateRequest
	GetTemplateId() *string
	SetTemplateType(v string) *UnbindTemplateRequest
	GetTemplateType() *string
}

type UnbindTemplateRequest struct {
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

func (s UnbindTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindTemplateRequest) GoString() string {
	return s.String()
}

func (s *UnbindTemplateRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UnbindTemplateRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *UnbindTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnbindTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *UnbindTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *UnbindTemplateRequest) SetInstanceId(v string) *UnbindTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *UnbindTemplateRequest) SetInstanceType(v string) *UnbindTemplateRequest {
	s.InstanceType = &v
	return s
}

func (s *UnbindTemplateRequest) SetOwnerId(v int64) *UnbindTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *UnbindTemplateRequest) SetTemplateId(v string) *UnbindTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *UnbindTemplateRequest) SetTemplateType(v string) *UnbindTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *UnbindTemplateRequest) Validate() error {
	return dara.Validate(s)
}
