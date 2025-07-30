// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainProxyTokensRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *ListDomainProxyTokensRequest
	GetDomainId() *string
	SetInstanceId(v string) *ListDomainProxyTokensRequest
	GetInstanceId() *string
}

type ListDomainProxyTokensRequest struct {
	// The domain ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// dm_examplexxxxx
	DomainId *string `json:"DomainId,omitempty" xml:"DomainId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListDomainProxyTokensRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDomainProxyTokensRequest) GoString() string {
	return s.String()
}

func (s *ListDomainProxyTokensRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *ListDomainProxyTokensRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDomainProxyTokensRequest) SetDomainId(v string) *ListDomainProxyTokensRequest {
	s.DomainId = &v
	return s
}

func (s *ListDomainProxyTokensRequest) SetInstanceId(v string) *ListDomainProxyTokensRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDomainProxyTokensRequest) Validate() error {
	return dara.Validate(s)
}
