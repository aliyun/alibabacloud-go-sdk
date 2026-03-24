// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyBroadcastSceneFromTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *CopyBroadcastSceneFromTemplateRequest
	GetName() *string
	SetRatio(v string) *CopyBroadcastSceneFromTemplateRequest
	GetRatio() *string
	SetTemplateId(v string) *CopyBroadcastSceneFromTemplateRequest
	GetTemplateId() *string
}

type CopyBroadcastSceneFromTemplateRequest struct {
	// This parameter is required.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9:16
	Ratio *string `json:"ratio,omitempty" xml:"ratio,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BS1b2WNnRMu4ouRzT4clY9Jhg
	TemplateId *string `json:"templateId,omitempty" xml:"templateId,omitempty"`
}

func (s CopyBroadcastSceneFromTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CopyBroadcastSceneFromTemplateRequest) GoString() string {
	return s.String()
}

func (s *CopyBroadcastSceneFromTemplateRequest) GetName() *string {
	return s.Name
}

func (s *CopyBroadcastSceneFromTemplateRequest) GetRatio() *string {
	return s.Ratio
}

func (s *CopyBroadcastSceneFromTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *CopyBroadcastSceneFromTemplateRequest) SetName(v string) *CopyBroadcastSceneFromTemplateRequest {
	s.Name = &v
	return s
}

func (s *CopyBroadcastSceneFromTemplateRequest) SetRatio(v string) *CopyBroadcastSceneFromTemplateRequest {
	s.Ratio = &v
	return s
}

func (s *CopyBroadcastSceneFromTemplateRequest) SetTemplateId(v string) *CopyBroadcastSceneFromTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *CopyBroadcastSceneFromTemplateRequest) Validate() error {
	return dara.Validate(s)
}
