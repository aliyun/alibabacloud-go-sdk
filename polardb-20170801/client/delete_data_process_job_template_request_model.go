// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataProcessJobTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteDataProcessJobTemplateRequest
	GetRegionId() *string
	SetTemplateId(v string) *DeleteDataProcessJobTemplateRequest
	GetTemplateId() *string
}

type DeleteDataProcessJobTemplateRequest struct {
	// The ID of the region where the template is located.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the task template to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// tpl-123456
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s DeleteDataProcessJobTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataProcessJobTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteDataProcessJobTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDataProcessJobTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *DeleteDataProcessJobTemplateRequest) SetRegionId(v string) *DeleteDataProcessJobTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDataProcessJobTemplateRequest) SetTemplateId(v string) *DeleteDataProcessJobTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *DeleteDataProcessJobTemplateRequest) Validate() error {
	return dara.Validate(s)
}
