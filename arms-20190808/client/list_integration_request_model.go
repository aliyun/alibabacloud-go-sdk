// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIntegrationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIntegrationName(v string) *ListIntegrationRequest
	GetIntegrationName() *string
	SetIntegrationProductType(v string) *ListIntegrationRequest
	GetIntegrationProductType() *string
	SetIsDetail(v bool) *ListIntegrationRequest
	GetIsDetail() *bool
	SetPage(v int64) *ListIntegrationRequest
	GetPage() *int64
	SetSize(v int64) *ListIntegrationRequest
	GetSize() *int64
}

type ListIntegrationRequest struct {
	// The name of the alert integration.
	//
	// example:
	//
	// CloudMonitor integration
	IntegrationName *string `json:"IntegrationName,omitempty" xml:"IntegrationName,omitempty"`
	// The type of the alert integration. Valid values:
	//
	// 	- CLOUD_MONITOR: CloudMonitor
	//
	// 	- LOG_SERVICE: Log Service
	//
	// This parameter is required.
	//
	// example:
	//
	// CLOUD_MONITOR
	IntegrationProductType *string `json:"IntegrationProductType,omitempty" xml:"IntegrationProductType,omitempty"`
	// Specifies whether to display the details of each alert integration:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	IsDetail *bool `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of alert integrations to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListIntegrationRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIntegrationRequest) GoString() string {
	return s.String()
}

func (s *ListIntegrationRequest) GetIntegrationName() *string {
	return s.IntegrationName
}

func (s *ListIntegrationRequest) GetIntegrationProductType() *string {
	return s.IntegrationProductType
}

func (s *ListIntegrationRequest) GetIsDetail() *bool {
	return s.IsDetail
}

func (s *ListIntegrationRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListIntegrationRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListIntegrationRequest) SetIntegrationName(v string) *ListIntegrationRequest {
	s.IntegrationName = &v
	return s
}

func (s *ListIntegrationRequest) SetIntegrationProductType(v string) *ListIntegrationRequest {
	s.IntegrationProductType = &v
	return s
}

func (s *ListIntegrationRequest) SetIsDetail(v bool) *ListIntegrationRequest {
	s.IsDetail = &v
	return s
}

func (s *ListIntegrationRequest) SetPage(v int64) *ListIntegrationRequest {
	s.Page = &v
	return s
}

func (s *ListIntegrationRequest) SetSize(v int64) *ListIntegrationRequest {
	s.Size = &v
	return s
}

func (s *ListIntegrationRequest) Validate() error {
	return dara.Validate(s)
}
