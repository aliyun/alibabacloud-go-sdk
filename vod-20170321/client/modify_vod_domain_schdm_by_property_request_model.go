// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyVodDomainSchdmByPropertyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *ModifyVodDomainSchdmByPropertyRequest
	GetDomainName() *string
	SetOwnerId(v int64) *ModifyVodDomainSchdmByPropertyRequest
	GetOwnerId() *int64
	SetProperty(v string) *ModifyVodDomainSchdmByPropertyRequest
	GetProperty() *string
}

type ModifyVodDomainSchdmByPropertyRequest struct {
	// This parameter is required.
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	Property *string `json:"Property,omitempty" xml:"Property,omitempty"`
}

func (s ModifyVodDomainSchdmByPropertyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyVodDomainSchdmByPropertyRequest) GoString() string {
	return s.String()
}

func (s *ModifyVodDomainSchdmByPropertyRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *ModifyVodDomainSchdmByPropertyRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyVodDomainSchdmByPropertyRequest) GetProperty() *string {
	return s.Property
}

func (s *ModifyVodDomainSchdmByPropertyRequest) SetDomainName(v string) *ModifyVodDomainSchdmByPropertyRequest {
	s.DomainName = &v
	return s
}

func (s *ModifyVodDomainSchdmByPropertyRequest) SetOwnerId(v int64) *ModifyVodDomainSchdmByPropertyRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyVodDomainSchdmByPropertyRequest) SetProperty(v string) *ModifyVodDomainSchdmByPropertyRequest {
	s.Property = &v
	return s
}

func (s *ModifyVodDomainSchdmByPropertyRequest) Validate() error {
	return dara.Validate(s)
}
