// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventBridgeIntegrationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageBean(v *ListEventBridgeIntegrationsResponseBodyPageBean) *ListEventBridgeIntegrationsResponseBody
	GetPageBean() *ListEventBridgeIntegrationsResponseBodyPageBean
	SetRequestId(v string) *ListEventBridgeIntegrationsResponseBody
	GetRequestId() *string
}

type ListEventBridgeIntegrationsResponseBody struct {
	// The information about EventBridge integrations that is returned on each page.
	PageBean *ListEventBridgeIntegrationsResponseBodyPageBean `json:"PageBean,omitempty" xml:"PageBean,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 2B289756-E791-5842-BCBD-AD414C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEventBridgeIntegrationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEventBridgeIntegrationsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEventBridgeIntegrationsResponseBody) GetPageBean() *ListEventBridgeIntegrationsResponseBodyPageBean {
	return s.PageBean
}

func (s *ListEventBridgeIntegrationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEventBridgeIntegrationsResponseBody) SetPageBean(v *ListEventBridgeIntegrationsResponseBodyPageBean) *ListEventBridgeIntegrationsResponseBody {
	s.PageBean = v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBody) SetRequestId(v string) *ListEventBridgeIntegrationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListEventBridgeIntegrationsResponseBodyPageBean struct {
	// The EventBridge integrations.
	EventBridgeIntegrations []*ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations `json:"EventBridgeIntegrations,omitempty" xml:"EventBridgeIntegrations,omitempty" type:"Repeated"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 15
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The total number of EventBridge integrations that are returned.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListEventBridgeIntegrationsResponseBodyPageBean) String() string {
	return dara.Prettify(s)
}

func (s ListEventBridgeIntegrationsResponseBodyPageBean) GoString() string {
	return s.String()
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) GetEventBridgeIntegrations() []*ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations {
	return s.EventBridgeIntegrations
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) GetPage() *int64 {
	return s.Page
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) GetSize() *int64 {
	return s.Size
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) GetTotal() *int64 {
	return s.Total
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) SetEventBridgeIntegrations(v []*ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) *ListEventBridgeIntegrationsResponseBodyPageBean {
	s.EventBridgeIntegrations = v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) SetPage(v int64) *ListEventBridgeIntegrationsResponseBodyPageBean {
	s.Page = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) SetSize(v int64) *ListEventBridgeIntegrationsResponseBodyPageBean {
	s.Size = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) SetTotal(v int64) *ListEventBridgeIntegrationsResponseBodyPageBean {
	s.Total = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBean) Validate() error {
	return dara.Validate(s)
}

type ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations struct {
	// The description of the EventBridge integration.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the EventBridge integration.
	//
	// example:
	//
	// 123
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the EventBridge integration.
	//
	// example:
	//
	// EventBridge_Test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) String() string {
	return dara.Prettify(s)
}

func (s ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) GoString() string {
	return s.String()
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) GetDescription() *string {
	return s.Description
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) GetId() *int64 {
	return s.Id
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) GetName() *string {
	return s.Name
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) SetDescription(v string) *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations {
	s.Description = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) SetId(v int64) *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations {
	s.Id = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) SetName(v string) *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations {
	s.Name = &v
	return s
}

func (s *ListEventBridgeIntegrationsResponseBodyPageBeanEventBridgeIntegrations) Validate() error {
	return dara.Validate(s)
}
