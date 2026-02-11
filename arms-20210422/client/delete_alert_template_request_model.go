// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAlertTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteAlertTemplateRequest
	GetId() *int64
	SetRegionId(v string) *DeleteAlertTemplateRequest
	GetRegionId() *string
}

type DeleteAlertTemplateRequest struct {
	// This parameter is required.
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteAlertTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAlertTemplateRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlertTemplateRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteAlertTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteAlertTemplateRequest) SetId(v int64) *DeleteAlertTemplateRequest {
	s.Id = &v
	return s
}

func (s *DeleteAlertTemplateRequest) SetRegionId(v string) *DeleteAlertTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteAlertTemplateRequest) Validate() error {
	return dara.Validate(s)
}
