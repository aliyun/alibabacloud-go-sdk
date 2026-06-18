// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteKvNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *DeleteKvNamespaceRequest
	GetNamespace() *string
}

type DeleteKvNamespaceRequest struct {
	// The name that you specified when you called [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s DeleteKvNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteKvNamespaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteKvNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *DeleteKvNamespaceRequest) SetNamespace(v string) *DeleteKvNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *DeleteKvNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
