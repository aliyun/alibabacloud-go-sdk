// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPrometheusAlertTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListPrometheusAlertTemplatesRequest
	GetClusterId() *string
	SetRegionId(v string) *ListPrometheusAlertTemplatesRequest
	GetRegionId() *string
}

type ListPrometheusAlertTemplatesRequest struct {
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListPrometheusAlertTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPrometheusAlertTemplatesRequest) GoString() string {
	return s.String()
}

func (s *ListPrometheusAlertTemplatesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListPrometheusAlertTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPrometheusAlertTemplatesRequest) SetClusterId(v string) *ListPrometheusAlertTemplatesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListPrometheusAlertTemplatesRequest) SetRegionId(v string) *ListPrometheusAlertTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListPrometheusAlertTemplatesRequest) Validate() error {
	return dara.Validate(s)
}
