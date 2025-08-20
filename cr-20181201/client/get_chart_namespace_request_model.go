// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChartNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetChartNamespaceRequest
	GetInstanceId() *string
	SetNamespaceName(v string) *GetChartNamespaceRequest
	GetNamespaceName() *string
}

type GetChartNamespaceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the namespace.
	//
	// This parameter is required.
	//
	// example:
	//
	// ns1
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s GetChartNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChartNamespaceRequest) GoString() string {
	return s.String()
}

func (s *GetChartNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetChartNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *GetChartNamespaceRequest) SetInstanceId(v string) *GetChartNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetChartNamespaceRequest) SetNamespaceName(v string) *GetChartNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *GetChartNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
