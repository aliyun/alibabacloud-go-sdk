// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListK8sNamespacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListK8sNamespacesRequest
	GetClusterId() *string
}

type ListK8sNamespacesRequest struct {
	// The ID of the cluster in Enterprise Distributed Application Service (EDAS).
	//
	// example:
	//
	// 5b2b4ab4-efbc-4a81-9c45-xxxxxxxxxxxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s ListK8sNamespacesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListK8sNamespacesRequest) GoString() string {
	return s.String()
}

func (s *ListK8sNamespacesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListK8sNamespacesRequest) SetClusterId(v string) *ListK8sNamespacesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListK8sNamespacesRequest) Validate() error {
	return dara.Validate(s)
}
