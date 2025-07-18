// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConnectorIds(v []*string) *ListConnectorsRequest
	GetConnectorIds() []*string
	SetCurrentPage(v int32) *ListConnectorsRequest
	GetCurrentPage() *int32
	SetName(v string) *ListConnectorsRequest
	GetName() *string
	SetPageSize(v int32) *ListConnectorsRequest
	GetPageSize() *int32
	SetStatus(v string) *ListConnectorsRequest
	GetStatus() *string
	SetSwitchStatus(v string) *ListConnectorsRequest
	GetSwitchStatus() *string
}

type ListConnectorsRequest struct {
	ConnectorIds []*string `json:"ConnectorIds,omitempty" xml:"ConnectorIds,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	CurrentPage *int32 `json:"CurrentPage,omitempty" xml:"CurrentPage,omitempty"`
	// example:
	//
	// connector_name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 10
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SwitchStatus *string `json:"SwitchStatus,omitempty" xml:"SwitchStatus,omitempty"`
}

func (s ListConnectorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListConnectorsRequest) GoString() string {
	return s.String()
}

func (s *ListConnectorsRequest) GetConnectorIds() []*string {
	return s.ConnectorIds
}

func (s *ListConnectorsRequest) GetCurrentPage() *int32 {
	return s.CurrentPage
}

func (s *ListConnectorsRequest) GetName() *string {
	return s.Name
}

func (s *ListConnectorsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConnectorsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListConnectorsRequest) GetSwitchStatus() *string {
	return s.SwitchStatus
}

func (s *ListConnectorsRequest) SetConnectorIds(v []*string) *ListConnectorsRequest {
	s.ConnectorIds = v
	return s
}

func (s *ListConnectorsRequest) SetCurrentPage(v int32) *ListConnectorsRequest {
	s.CurrentPage = &v
	return s
}

func (s *ListConnectorsRequest) SetName(v string) *ListConnectorsRequest {
	s.Name = &v
	return s
}

func (s *ListConnectorsRequest) SetPageSize(v int32) *ListConnectorsRequest {
	s.PageSize = &v
	return s
}

func (s *ListConnectorsRequest) SetStatus(v string) *ListConnectorsRequest {
	s.Status = &v
	return s
}

func (s *ListConnectorsRequest) SetSwitchStatus(v string) *ListConnectorsRequest {
	s.SwitchStatus = &v
	return s
}

func (s *ListConnectorsRequest) Validate() error {
	return dara.Validate(s)
}
