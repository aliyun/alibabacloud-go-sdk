// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNamespacedConfigMapsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *ListNamespacedConfigMapsRequest
	GetNamespaceId() *string
}

type ListNamespacedConfigMapsRequest struct {
	// cn-hangzhou
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ListNamespacedConfigMapsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListNamespacedConfigMapsRequest) GoString() string {
	return s.String()
}

func (s *ListNamespacedConfigMapsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListNamespacedConfigMapsRequest) SetNamespaceId(v string) *ListNamespacedConfigMapsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListNamespacedConfigMapsRequest) Validate() error {
	return dara.Validate(s)
}
