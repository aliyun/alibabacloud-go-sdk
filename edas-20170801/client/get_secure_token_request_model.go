// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSecureTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNamespaceId(v string) *GetSecureTokenRequest
	GetNamespaceId() *string
}

type GetSecureTokenRequest struct {
	// The ID of the namespace, such as cn-beijing or cn-beijing:prod````.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen:x*****
	NamespaceId *string `json:"NamespaceId,omitempty" xml:"NamespaceId,omitempty"`
}

func (s GetSecureTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSecureTokenRequest) GoString() string {
	return s.String()
}

func (s *GetSecureTokenRequest) GetNamespaceId() *string {
	return s.NamespaceId
}

func (s *GetSecureTokenRequest) SetNamespaceId(v string) *GetSecureTokenRequest {
	s.NamespaceId = &v
	return s
}

func (s *GetSecureTokenRequest) Validate() error {
	return dara.Validate(s)
}
