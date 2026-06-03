// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForReserveDropListDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContactTemplateId(v string) *SaveSingleTaskForReserveDropListDomainRequest
	GetContactTemplateId() *string
	SetDns1(v string) *SaveSingleTaskForReserveDropListDomainRequest
	GetDns1() *string
	SetDns2(v string) *SaveSingleTaskForReserveDropListDomainRequest
	GetDns2() *string
	SetDomainName(v string) *SaveSingleTaskForReserveDropListDomainRequest
	GetDomainName() *string
}

type SaveSingleTaskForReserveDropListDomainRequest struct {
	// This parameter is required.
	ContactTemplateId *string `json:"ContactTemplateId,omitempty" xml:"ContactTemplateId,omitempty"`
	Dns1              *string `json:"Dns1,omitempty" xml:"Dns1,omitempty"`
	Dns2              *string `json:"Dns2,omitempty" xml:"Dns2,omitempty"`
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s SaveSingleTaskForReserveDropListDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForReserveDropListDomainRequest) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForReserveDropListDomainRequest) GetContactTemplateId() *string {
	return s.ContactTemplateId
}

func (s *SaveSingleTaskForReserveDropListDomainRequest) GetDns1() *string {
	return s.Dns1
}

func (s *SaveSingleTaskForReserveDropListDomainRequest) GetDns2() *string {
	return s.Dns2
}

func (s *SaveSingleTaskForReserveDropListDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *SaveSingleTaskForReserveDropListDomainRequest) SetContactTemplateId(v string) *SaveSingleTaskForReserveDropListDomainRequest {
	s.ContactTemplateId = &v
	return s
}

func (s *SaveSingleTaskForReserveDropListDomainRequest) SetDns1(v string) *SaveSingleTaskForReserveDropListDomainRequest {
	s.Dns1 = &v
	return s
}

func (s *SaveSingleTaskForReserveDropListDomainRequest) SetDns2(v string) *SaveSingleTaskForReserveDropListDomainRequest {
	s.Dns2 = &v
	return s
}

func (s *SaveSingleTaskForReserveDropListDomainRequest) SetDomainName(v string) *SaveSingleTaskForReserveDropListDomainRequest {
	s.DomainName = &v
	return s
}

func (s *SaveSingleTaskForReserveDropListDomainRequest) Validate() error {
	return dara.Validate(s)
}
