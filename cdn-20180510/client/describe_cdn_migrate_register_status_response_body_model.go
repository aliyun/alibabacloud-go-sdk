// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCdnMigrateRegisterStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainName(v string) *DescribeCdnMigrateRegisterStatusResponseBody
	GetDomainName() *string
	SetRequestId(v string) *DescribeCdnMigrateRegisterStatusResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeCdnMigrateRegisterStatusResponseBody
	GetStatus() *string
}

type DescribeCdnMigrateRegisterStatusResponseBody struct {
	// The accelerated domain name.
	//
	// example:
	//
	// example.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// 64D28B53-5902-409B-94F6-FD46680144FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The registration status. Valid values:
	//
	// 	- **not exist**
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

func (s DescribeCdnMigrateRegisterStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCdnMigrateRegisterStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCdnMigrateRegisterStatusResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeCdnMigrateRegisterStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCdnMigrateRegisterStatusResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeCdnMigrateRegisterStatusResponseBody) SetDomainName(v string) *DescribeCdnMigrateRegisterStatusResponseBody {
	s.DomainName = &v
	return s
}

func (s *DescribeCdnMigrateRegisterStatusResponseBody) SetRequestId(v string) *DescribeCdnMigrateRegisterStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCdnMigrateRegisterStatusResponseBody) SetStatus(v string) *DescribeCdnMigrateRegisterStatusResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeCdnMigrateRegisterStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
