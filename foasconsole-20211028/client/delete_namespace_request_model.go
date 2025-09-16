// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteNamespaceRequest
	GetInstanceId() *string
	SetNamespace(v string) *DeleteNamespaceRequest
	GetNamespace() *string
	SetRegion(v string) *DeleteNamespaceRequest
	GetRegion() *string
}

type DeleteNamespaceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// f-cn-wwo36qj4g06
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// di-593439443804417
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
}

func (s DeleteNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteNamespaceRequest) GetRegion() *string {
	return s.Region
}

func (s *DeleteNamespaceRequest) SetInstanceId(v string) *DeleteNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteNamespaceRequest) SetNamespace(v string) *DeleteNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteNamespaceRequest) SetRegion(v string) *DeleteNamespaceRequest {
	s.Region = &v
	return s
}

func (s *DeleteNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
