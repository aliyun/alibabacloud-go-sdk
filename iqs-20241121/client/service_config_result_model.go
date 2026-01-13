// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceConfigResult interface {
	dara.Model
	String() string
	GoString() string
	SetPoiSearchTotalQuota(v string) *ServiceConfigResult
	GetPoiSearchTotalQuota() *string
	SetPoiSearchUsedQuota(v string) *ServiceConfigResult
	GetPoiSearchUsedQuota() *string
	SetSearchTotalQuota(v string) *ServiceConfigResult
	GetSearchTotalQuota() *string
	SetSearchUsedQuota(v string) *ServiceConfigResult
	GetSearchUsedQuota() *string
	SetStatus(v string) *ServiceConfigResult
	GetStatus() *string
}

type ServiceConfigResult struct {
	PoiSearchTotalQuota *string `json:"poiSearchTotalQuota,omitempty" xml:"poiSearchTotalQuota,omitempty"`
	PoiSearchUsedQuota  *string `json:"poiSearchUsedQuota,omitempty" xml:"poiSearchUsedQuota,omitempty"`
	SearchTotalQuota    *string `json:"searchTotalQuota,omitempty" xml:"searchTotalQuota,omitempty"`
	SearchUsedQuota     *string `json:"searchUsedQuota,omitempty" xml:"searchUsedQuota,omitempty"`
	Status              *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ServiceConfigResult) String() string {
	return dara.Prettify(s)
}

func (s ServiceConfigResult) GoString() string {
	return s.String()
}

func (s *ServiceConfigResult) GetPoiSearchTotalQuota() *string {
	return s.PoiSearchTotalQuota
}

func (s *ServiceConfigResult) GetPoiSearchUsedQuota() *string {
	return s.PoiSearchUsedQuota
}

func (s *ServiceConfigResult) GetSearchTotalQuota() *string {
	return s.SearchTotalQuota
}

func (s *ServiceConfigResult) GetSearchUsedQuota() *string {
	return s.SearchUsedQuota
}

func (s *ServiceConfigResult) GetStatus() *string {
	return s.Status
}

func (s *ServiceConfigResult) SetPoiSearchTotalQuota(v string) *ServiceConfigResult {
	s.PoiSearchTotalQuota = &v
	return s
}

func (s *ServiceConfigResult) SetPoiSearchUsedQuota(v string) *ServiceConfigResult {
	s.PoiSearchUsedQuota = &v
	return s
}

func (s *ServiceConfigResult) SetSearchTotalQuota(v string) *ServiceConfigResult {
	s.SearchTotalQuota = &v
	return s
}

func (s *ServiceConfigResult) SetSearchUsedQuota(v string) *ServiceConfigResult {
	s.SearchUsedQuota = &v
	return s
}

func (s *ServiceConfigResult) SetStatus(v string) *ServiceConfigResult {
	s.Status = &v
	return s
}

func (s *ServiceConfigResult) Validate() error {
	return dara.Validate(s)
}
