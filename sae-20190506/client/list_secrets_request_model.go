// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecretsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *ListSecretsRequest
	GetNamespaceId() *string
}

type ListSecretsRequest struct {
	// The ID of the namespace in which the Secrets reside. By default, the namespace ID is the same as the region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s ListSecretsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSecretsRequest) GoString() string {
	return s.String()
}

func (s *ListSecretsRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *ListSecretsRequest) SetNamespaceId(v string) *ListSecretsRequest {
	s.NamespaceId = &v
	return s
}

func (s *ListSecretsRequest) Validate() error {
	return dara.Validate(s)
}
