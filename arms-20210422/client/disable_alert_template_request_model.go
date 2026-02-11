// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAlertTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DisableAlertTemplateRequest
	GetId() *int64
	SetRegionId(v string) *DisableAlertTemplateRequest
	GetRegionId() *string
}

type DisableAlertTemplateRequest struct {
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DisableAlertTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableAlertTemplateRequest) GoString() string {
	return s.String()
}

func (s *DisableAlertTemplateRequest) GetId() *int64 {
	return s.Id
}

func (s *DisableAlertTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DisableAlertTemplateRequest) SetId(v int64) *DisableAlertTemplateRequest {
	s.Id = &v
	return s
}

func (s *DisableAlertTemplateRequest) SetRegionId(v string) *DisableAlertTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DisableAlertTemplateRequest) Validate() error {
	return dara.Validate(s)
}
