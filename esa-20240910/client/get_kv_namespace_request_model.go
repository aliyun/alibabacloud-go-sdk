// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKvNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespace(v string) *GetKvNamespaceRequest
	GetNamespace() *string
}

type GetKvNamespaceRequest struct {
	// The name of the namespace that you specify when you call the [CreateKvNamespace](https://help.aliyun.com/document_detail/2850317.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_namespace
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s GetKvNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetKvNamespaceRequest) GoString() string {
	return s.String()
}

func (s *GetKvNamespaceRequest) GetNamespace() *string {
	return s.Namespace
}

func (s *GetKvNamespaceRequest) SetNamespace(v string) *GetKvNamespaceRequest {
	s.Namespace = &v
	return s
}

func (s *GetKvNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
