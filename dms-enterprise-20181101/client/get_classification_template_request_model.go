// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClassificationTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v int64) *GetClassificationTemplateRequest
	GetInstanceId() *int64
	SetTid(v int64) *GetClassificationTemplateRequest
	GetTid() *int64
}

type GetClassificationTemplateRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 169****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 23***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetClassificationTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetClassificationTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetClassificationTemplateRequest) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetClassificationTemplateRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetClassificationTemplateRequest) SetInstanceId(v int64) *GetClassificationTemplateRequest {
	s.InstanceId = &v
	return s
}

func (s *GetClassificationTemplateRequest) SetTid(v int64) *GetClassificationTemplateRequest {
	s.Tid = &v
	return s
}

func (s *GetClassificationTemplateRequest) Validate() error {
	return dara.Validate(s)
}
