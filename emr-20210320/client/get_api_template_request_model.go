// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApiTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetApiTemplateRequest
	GetRegionId() *string
	SetTemplateId(v string) *GetApiTemplateRequest
	GetTemplateId() *string
}

type GetApiTemplateRequest struct {
	// Region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Cluster template ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// at-41b4c6a0fc63****
	TemplateId *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
}

func (s GetApiTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s GetApiTemplateRequest) GoString() string {
	return s.String()
}

func (s *GetApiTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetApiTemplateRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetApiTemplateRequest) SetRegionId(v string) *GetApiTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *GetApiTemplateRequest) SetTemplateId(v string) *GetApiTemplateRequest {
	s.TemplateId = &v
	return s
}

func (s *GetApiTemplateRequest) Validate() error {
	return dara.Validate(s)
}
