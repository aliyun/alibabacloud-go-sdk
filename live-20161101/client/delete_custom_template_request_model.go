// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *DeleteCustomTemplateRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCustomTemplateRequest
	GetRegionId() *string
	SetTemplate(v string) *DeleteCustomTemplateRequest
	GetTemplate() *string
}

type DeleteCustomTemplateRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the template that you want to delete. The value is a string.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestTemplate
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s DeleteCustomTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteCustomTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCustomTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCustomTemplateRequest) GetTemplate() *string {
	return s.Template
}

func (s *DeleteCustomTemplateRequest) SetOwnerId(v int64) *DeleteCustomTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCustomTemplateRequest) SetRegionId(v string) *DeleteCustomTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCustomTemplateRequest) SetTemplate(v string) *DeleteCustomTemplateRequest {
	s.Template = &v
	return s
}

func (s *DeleteCustomTemplateRequest) Validate() error {
	return dara.Validate(s)
}
