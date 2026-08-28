// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSyncMCPServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGatewayId(v string) *ListSyncMCPServerRequest
	GetGatewayId() *string
	SetNamespace(v string) *ListSyncMCPServerRequest
	GetNamespace() *string
	SetSourceId(v string) *ListSyncMCPServerRequest
	GetSourceId() *string
}

type ListSyncMCPServerRequest struct {
	// example:
	//
	// gw-xxxx
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// example:
	//
	// public
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// example:
	//
	// src-xxxx
	SourceId *string `json:"sourceId,omitempty" xml:"sourceId,omitempty"`
}

func (s ListSyncMCPServerRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSyncMCPServerRequest) GoString() string {
	return s.String()
}

func (s *ListSyncMCPServerRequest) GetGatewayId() *string {
	return s.GatewayId
}

func (s *ListSyncMCPServerRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *ListSyncMCPServerRequest) GetSourceId() *string {
	return s.SourceId
}

func (s *ListSyncMCPServerRequest) SetGatewayId(v string) *ListSyncMCPServerRequest {
	s.GatewayId = &v
	return s
}

func (s *ListSyncMCPServerRequest) SetNamespace(v string) *ListSyncMCPServerRequest {
	s.Namespace = &v
	return s
}

func (s *ListSyncMCPServerRequest) SetSourceId(v string) *ListSyncMCPServerRequest {
	s.SourceId = &v
	return s
}

func (s *ListSyncMCPServerRequest) Validate() error {
	return dara.Validate(s)
}
