// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomOrgRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCorpId(v string) *CreateCustomOrgRequest
	GetCorpId() *string
	SetCorpName(v string) *CreateCustomOrgRequest
	GetCorpName() *string
	SetTenantId(v string) *CreateCustomOrgRequest
	GetTenantId() *string
}

type CreateCustomOrgRequest struct {
	// The corpId of the activated enterprise.
	//
	// This parameter is required.
	//
	// example:
	//
	// exampleCorpId
	CorpId *string `json:"corpId,omitempty" xml:"corpId,omitempty"`
	// The organization name.
	//
	// example:
	//
	// string_value
	CorpName *string `json:"corpName,omitempty" xml:"corpName,omitempty"`
	// The tenant ID.
	//
	// example:
	//
	// 692318833855074
	TenantId *string `json:"tenantId,omitempty" xml:"tenantId,omitempty"`
}

func (s CreateCustomOrgRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomOrgRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomOrgRequest) GetCorpId() *string {
	return s.CorpId
}

func (s *CreateCustomOrgRequest) GetCorpName() *string {
	return s.CorpName
}

func (s *CreateCustomOrgRequest) GetTenantId() *string {
	return s.TenantId
}

func (s *CreateCustomOrgRequest) SetCorpId(v string) *CreateCustomOrgRequest {
	s.CorpId = &v
	return s
}

func (s *CreateCustomOrgRequest) SetCorpName(v string) *CreateCustomOrgRequest {
	s.CorpName = &v
	return s
}

func (s *CreateCustomOrgRequest) SetTenantId(v string) *CreateCustomOrgRequest {
	s.TenantId = &v
	return s
}

func (s *CreateCustomOrgRequest) Validate() error {
	return dara.Validate(s)
}
