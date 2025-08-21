// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchUnbindTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *BatchUnbindTemplatesRequest
	GetInstanceId() *string
	SetInstanceType(v string) *BatchUnbindTemplatesRequest
	GetInstanceType() *string
	SetOwnerId(v int64) *BatchUnbindTemplatesRequest
	GetOwnerId() *int64
	SetTemplateId(v string) *BatchUnbindTemplatesRequest
	GetTemplateId() *string
	SetTemplateType(v string) *BatchUnbindTemplatesRequest
	GetTemplateType() *string
}

type BatchUnbindTemplatesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 400941290881239938-cn-beijing
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
	// 323*****998-cn-qingdao
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	// example:
	//
	// record
	TemplateType *string `json:"TemplateType,omitempty" xml:"TemplateType,omitempty"`
}

func (s BatchUnbindTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchUnbindTemplatesRequest) GoString() string {
	return s.String()
}

func (s *BatchUnbindTemplatesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *BatchUnbindTemplatesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *BatchUnbindTemplatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *BatchUnbindTemplatesRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *BatchUnbindTemplatesRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *BatchUnbindTemplatesRequest) SetInstanceId(v string) *BatchUnbindTemplatesRequest {
	s.InstanceId = &v
	return s
}

func (s *BatchUnbindTemplatesRequest) SetInstanceType(v string) *BatchUnbindTemplatesRequest {
	s.InstanceType = &v
	return s
}

func (s *BatchUnbindTemplatesRequest) SetOwnerId(v int64) *BatchUnbindTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *BatchUnbindTemplatesRequest) SetTemplateId(v string) *BatchUnbindTemplatesRequest {
	s.TemplateId = &v
	return s
}

func (s *BatchUnbindTemplatesRequest) SetTemplateType(v string) *BatchUnbindTemplatesRequest {
	s.TemplateType = &v
	return s
}

func (s *BatchUnbindTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
