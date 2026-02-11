// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportPrometheusRulesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ImportPrometheusRulesRequest
	GetClusterId() *string
	SetContent(v string) *ImportPrometheusRulesRequest
	GetContent() *string
	SetName(v string) *ImportPrometheusRulesRequest
	GetName() *string
	SetNameSpace(v string) *ImportPrometheusRulesRequest
	GetNameSpace() *string
	SetRegionId(v string) *ImportPrometheusRulesRequest
	GetRegionId() *string
}

type ImportPrometheusRulesRequest struct {
	// This parameter is required.
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	NameSpace *string `json:"NameSpace,omitempty" xml:"NameSpace,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ImportPrometheusRulesRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportPrometheusRulesRequest) GoString() string {
	return s.String()
}

func (s *ImportPrometheusRulesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ImportPrometheusRulesRequest) GetContent() *string {
	return s.Content
}

func (s *ImportPrometheusRulesRequest) GetName() *string {
	return s.Name
}

func (s *ImportPrometheusRulesRequest) GetNameSpace() *string {
	return s.NameSpace
}

func (s *ImportPrometheusRulesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportPrometheusRulesRequest) SetClusterId(v string) *ImportPrometheusRulesRequest {
	s.ClusterId = &v
	return s
}

func (s *ImportPrometheusRulesRequest) SetContent(v string) *ImportPrometheusRulesRequest {
	s.Content = &v
	return s
}

func (s *ImportPrometheusRulesRequest) SetName(v string) *ImportPrometheusRulesRequest {
	s.Name = &v
	return s
}

func (s *ImportPrometheusRulesRequest) SetNameSpace(v string) *ImportPrometheusRulesRequest {
	s.NameSpace = &v
	return s
}

func (s *ImportPrometheusRulesRequest) SetRegionId(v string) *ImportPrometheusRulesRequest {
	s.RegionId = &v
	return s
}

func (s *ImportPrometheusRulesRequest) Validate() error {
	return dara.Validate(s)
}
