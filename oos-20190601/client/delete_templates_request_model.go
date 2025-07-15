// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoDeleteExecutions(v bool) *DeleteTemplatesRequest
	GetAutoDeleteExecutions() *bool
	SetRegionId(v string) *DeleteTemplatesRequest
	GetRegionId() *string
	SetTemplateNames(v string) *DeleteTemplatesRequest
	GetTemplateNames() *string
}

type DeleteTemplatesRequest struct {
	// Specifies whether to delete the related executions when a template is deleted.
	//
	// example:
	//
	// false
	AutoDeleteExecutions *bool `json:"AutoDeleteExecutions,omitempty" xml:"AutoDeleteExecutions,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The names of the templates to be deleted.
	//
	// This parameter is required.
	//
	// example:
	//
	// ["t1","t2"]
	TemplateNames *string `json:"TemplateNames,omitempty" xml:"TemplateNames,omitempty"`
}

func (s DeleteTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteTemplatesRequest) GetAutoDeleteExecutions() *bool {
	return s.AutoDeleteExecutions
}

func (s *DeleteTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteTemplatesRequest) GetTemplateNames() *string {
	return s.TemplateNames
}

func (s *DeleteTemplatesRequest) SetAutoDeleteExecutions(v bool) *DeleteTemplatesRequest {
	s.AutoDeleteExecutions = &v
	return s
}

func (s *DeleteTemplatesRequest) SetRegionId(v string) *DeleteTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteTemplatesRequest) SetTemplateNames(v string) *DeleteTemplatesRequest {
	s.TemplateNames = &v
	return s
}

func (s *DeleteTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
