// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainId(v string) *GetDomainRequest
	GetDomainId() *string
	SetFields(v string) *GetDomainRequest
	GetFields() *string
	SetGetQuotaUsed(v bool) *GetDomainRequest
	GetGetQuotaUsed() *bool
}

type GetDomainRequest struct {
	// The ID of the domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	Fields   *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// Specifies whether to return the used quota of the domain. Default value: false. If the quota of the domain is greater than 0 and you set this parameter to true, the used quota of the domain is returned.
	//
	// example:
	//
	// true
	GetQuotaUsed *bool `json:"get_quota_used,omitempty" xml:"get_quota_used,omitempty"`
}

func (s GetDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDomainRequest) GoString() string {
	return s.String()
}

func (s *GetDomainRequest) GetDomainId() *string {
	return s.DomainId
}

func (s *GetDomainRequest) GetFields() *string {
	return s.Fields
}

func (s *GetDomainRequest) GetGetQuotaUsed() *bool {
	return s.GetQuotaUsed
}

func (s *GetDomainRequest) SetDomainId(v string) *GetDomainRequest {
	s.DomainId = &v
	return s
}

func (s *GetDomainRequest) SetFields(v string) *GetDomainRequest {
	s.Fields = &v
	return s
}

func (s *GetDomainRequest) SetGetQuotaUsed(v bool) *GetDomainRequest {
	s.GetQuotaUsed = &v
	return s
}

func (s *GetDomainRequest) Validate() error {
	return dara.Validate(s)
}
