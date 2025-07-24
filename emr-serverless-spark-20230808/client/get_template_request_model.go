// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetTemplateRequest
	GetRegionId() *string
	SetTemplateBizId(v string) *GetTemplateRequest
	GetTemplateBizId() *string
	SetTemplateType(v string) *GetTemplateRequest
	GetTemplateType() *string
}

type GetTemplateRequest struct {
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId      *string `json:"regionId,omitempty" xml:"regionId,omitempty"`
	TemplateBizId *string `json:"templateBizId,omitempty" xml:"templateBizId,omitempty"`
	// The template type.
	//
	// Valid values:
	//
	// 	- TASK
	//
	// 	- SESSION
	//
	// example:
	//
	// TASK
	TemplateType *string `json:"templateType,omitempty" xml:"templateType,omitempty"`
}

func (s GetTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTemplateRequest) GetTemplateBizId() *string {
	return s.TemplateBizId
}

func (s *GetTemplateRequest) GetTemplateType() *string {
	return s.TemplateType
}

func (s *GetTemplateRequest) SetRegionId(v string) *GetTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateBizId(v string) *GetTemplateRequest {
	s.TemplateBizId = &v
	return s
}

func (s *GetTemplateRequest) SetTemplateType(v string) *GetTemplateRequest {
	s.TemplateType = &v
	return s
}

func (s *GetTemplateRequest) Validate() error {
	return dara.Validate(s)
}
