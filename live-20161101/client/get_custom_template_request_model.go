// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCustomTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetCustomTemplateRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetCustomTemplateRequest
	GetRegionId() *string
	SetTemplate(v string) *GetCustomTemplateRequest
	GetTemplate() *string
}

type GetCustomTemplateRequest struct {
	OwnerId  *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The name of the custom template that you want to query. The value is a string.
	//
	// This parameter is required.
	//
	// example:
	//
	// TestTemplate
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s GetCustomTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCustomTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetCustomTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetCustomTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetCustomTemplateRequest) GetTemplate() *string {
	return s.Template
}

func (s *GetCustomTemplateRequest) SetOwnerId(v int64) *GetCustomTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *GetCustomTemplateRequest) SetRegionId(v string) *GetCustomTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetCustomTemplateRequest) SetTemplate(v string) *GetCustomTemplateRequest {
	s.Template = &v
	return s
}

func (s *GetCustomTemplateRequest) Validate() error {
	return dara.Validate(s)
}
