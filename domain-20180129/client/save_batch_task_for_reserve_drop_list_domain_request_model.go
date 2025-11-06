// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForReserveDropListDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactTemplateId(v string) *SaveBatchTaskForReserveDropListDomainRequest
	GetContactTemplateId() *string
	SetDomains(v []*SaveBatchTaskForReserveDropListDomainRequestDomains) *SaveBatchTaskForReserveDropListDomainRequest
	GetDomains() []*SaveBatchTaskForReserveDropListDomainRequestDomains
}

type SaveBatchTaskForReserveDropListDomainRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123123
	ContactTemplateId *string `json:"ContactTemplateId,omitempty" xml:"ContactTemplateId,omitempty"`
	// This parameter is required.
	Domains []*SaveBatchTaskForReserveDropListDomainRequestDomains `json:"Domains,omitempty" xml:"Domains,omitempty" type:"Repeated"`
}

func (s SaveBatchTaskForReserveDropListDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForReserveDropListDomainRequest) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForReserveDropListDomainRequest) GetContactTemplateId() *string {
	return s.ContactTemplateId
}

func (s *SaveBatchTaskForReserveDropListDomainRequest) GetDomains() []*SaveBatchTaskForReserveDropListDomainRequestDomains {
	return s.Domains
}

func (s *SaveBatchTaskForReserveDropListDomainRequest) SetContactTemplateId(v string) *SaveBatchTaskForReserveDropListDomainRequest {
	s.ContactTemplateId = &v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainRequest) SetDomains(v []*SaveBatchTaskForReserveDropListDomainRequestDomains) *SaveBatchTaskForReserveDropListDomainRequest {
	s.Domains = v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainRequest) Validate() error {
	if s.Domains != nil {
		for _, item := range s.Domains {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SaveBatchTaskForReserveDropListDomainRequestDomains struct {
	Dns1 *string `json:"Dns1,omitempty" xml:"Dns1,omitempty"`
	Dns2 *string `json:"Dns2,omitempty" xml:"Dns2,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s SaveBatchTaskForReserveDropListDomainRequestDomains) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForReserveDropListDomainRequestDomains) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForReserveDropListDomainRequestDomains) GetDns1() *string {
	return s.Dns1
}

func (s *SaveBatchTaskForReserveDropListDomainRequestDomains) GetDns2() *string {
	return s.Dns2
}

func (s *SaveBatchTaskForReserveDropListDomainRequestDomains) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveBatchTaskForReserveDropListDomainRequestDomains) SetDns1(v string) *SaveBatchTaskForReserveDropListDomainRequestDomains {
	s.Dns1 = &v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainRequestDomains) SetDns2(v string) *SaveBatchTaskForReserveDropListDomainRequestDomains {
	s.Dns2 = &v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainRequestDomains) SetDomainName(v string) *SaveBatchTaskForReserveDropListDomainRequestDomains {
	s.DomainName = &v
	return s
}

func (s *SaveBatchTaskForReserveDropListDomainRequestDomains) Validate() error {
	return dara.Validate(s)
}
