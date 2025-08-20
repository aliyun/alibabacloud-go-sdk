// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChartNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteChartNamespaceRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *DeleteChartNamespaceRequest
	GetNamespaceName() *string
}

type DeleteChartNamespaceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the chart namespace that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns2
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s DeleteChartNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteChartNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteChartNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteChartNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *DeleteChartNamespaceRequest) SetInstanceId(v string) *DeleteChartNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteChartNamespaceRequest) SetNamespaceName(v string) *DeleteChartNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *DeleteChartNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
