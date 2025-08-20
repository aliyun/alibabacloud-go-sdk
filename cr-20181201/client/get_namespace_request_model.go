// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNamespaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetNamespaceRequest
	GetInstanceId() *string
	SetNamespaceId(v string) *GetNamespaceRequest
	GetNamespaceId() *string
	SetNamespaceName(v string) *GetNamespaceRequest
	GetNamespaceName() *string
}

type GetNamespaceRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-kmsiwlxxdcva****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the namespace.
	//
	// example:
	//
	// crn-tiw8t3f8i5lta****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
	// The name of the namespace.
	//
	// example:
	//
	// test
	NamespaceName *string `json:"NamespaceName,omitempty" xml:"NamespaceName,omitempty"`
}

func (s GetNamespaceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetNamespaceRequest) GoString() string {
	return s.String()
}

func (s *GetNamespaceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetNamespaceRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetNamespaceRequest) GetNamespaceName() *string {
	return s.NamespaceName
}

func (s *GetNamespaceRequest) SetInstanceId(v string) *GetNamespaceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetNamespaceRequest) SetNamespaceId(v string) *GetNamespaceRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetNamespaceRequest) SetNamespaceName(v string) *GetNamespaceRequest {
	s.NamespaceName = &v
	return s
}

func (s *GetNamespaceRequest) Validate() error {
	return dara.Validate(s)
}
