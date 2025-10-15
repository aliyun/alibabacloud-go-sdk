// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDomainsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBrandId(v string) *ListDomainsRequest
	GetBrandId() *string
	SetInstanceId(v string) *ListDomainsRequest
	GetInstanceId() *string
}

type ListDomainsRequest struct {
	BrandId *string `json:"BrandId,omitempty" xml:"BrandId,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_ue2jvisn35ea5lmthk267xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) GetBrandId() *string {
	return s.BrandId
}

func (s *ListDomainsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDomainsRequest) SetBrandId(v string) *ListDomainsRequest {
	s.BrandId = &v
	return s
}

func (s *ListDomainsRequest) SetInstanceId(v string) *ListDomainsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListDomainsRequest) Validate() error {
	return dara.Validate(s)
}
