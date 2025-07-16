// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCdnMigrateRegisterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *CdnMigrateRegisterResponseBody
	GetDomainName() *string
	SetRequestId(v string) *CdnMigrateRegisterResponseBody
	GetRequestId() *string
	SetStatus(v string) *CdnMigrateRegisterResponseBody
	GetStatus() *string
}

type CdnMigrateRegisterResponseBody struct {
	// The accelerated domain name. You can specify only one domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 64D28B53-5902-409B-94F6-FD46680144FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The registration status. Valid values:
	//
	// 	- **running**
	//
	// 	- **succeed**
	//
	// 	- **failed**
	//
	// example:
	//
	// succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CdnMigrateRegisterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CdnMigrateRegisterResponseBody) GoString() string {
	return s.String()
}

func (s *CdnMigrateRegisterResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *CdnMigrateRegisterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CdnMigrateRegisterResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CdnMigrateRegisterResponseBody) SetDomainName(v string) *CdnMigrateRegisterResponseBody {
	s.DomainName = &v
	return s
}

func (s *CdnMigrateRegisterResponseBody) SetRequestId(v string) *CdnMigrateRegisterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CdnMigrateRegisterResponseBody) SetStatus(v string) *CdnMigrateRegisterResponseBody {
	s.Status = &v
	return s
}

func (s *CdnMigrateRegisterResponseBody) Validate() error {
	return dara.Validate(s)
}
