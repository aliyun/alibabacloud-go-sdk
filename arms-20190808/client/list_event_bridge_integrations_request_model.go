// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventBridgeIntegrationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListEventBridgeIntegrationsRequest
	GetName() *string
	SetPage(v int64) *ListEventBridgeIntegrationsRequest
	GetPage() *int64
	SetSize(v int64) *ListEventBridgeIntegrationsRequest
	GetSize() *int64
}

type ListEventBridgeIntegrationsRequest struct {
	// The name of the EventBridge integration.
	//
	// example:
	//
	// EventBridge_Test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListEventBridgeIntegrationsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventBridgeIntegrationsRequest) GoString() string {
	return s.String()
}

func (s *ListEventBridgeIntegrationsRequest) GetName() *string {
	return s.Name
}

func (s *ListEventBridgeIntegrationsRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListEventBridgeIntegrationsRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListEventBridgeIntegrationsRequest) SetName(v string) *ListEventBridgeIntegrationsRequest {
	s.Name = &v
	return s
}

func (s *ListEventBridgeIntegrationsRequest) SetPage(v int64) *ListEventBridgeIntegrationsRequest {
	s.Page = &v
	return s
}

func (s *ListEventBridgeIntegrationsRequest) SetSize(v int64) *ListEventBridgeIntegrationsRequest {
	s.Size = &v
	return s
}

func (s *ListEventBridgeIntegrationsRequest) Validate() error {
	return dara.Validate(s)
}
