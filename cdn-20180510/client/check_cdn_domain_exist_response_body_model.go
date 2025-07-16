// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCdnDomainExistResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckCdnDomainExistResponseBody
	GetRequestId() *string
	SetStatus(v string) *CheckCdnDomainExistResponseBody
	GetStatus() *string
}

type CheckCdnDomainExistResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 64D28B53-5902-409B-94F6-FD46680144FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the domain name. Valid values:
	//
	// 	- **DomainNotExist**: The domain name is not added.
	//
	// 	- **DomainExistOtherUser**: The domain name has been added by another account.
	//
	// 	- **DomainExistCdnProduct**: The domain name has been added to Alibaba Cloud CDN.
	//
	// 	- **DomainExistDcdnProduct**: The domain name has been added to Dynamic Content Delivery Network (DCDN).
	//
	// 	- **DomainExistScdnProduct**: The domain name has been added to Secure CDN (SCDN).
	//
	// 	- **DomainExistVodProduct**: The domain name has been added to ApsaraVideo VOD.
	//
	// 	- **DomainExistLiveProduct**: The domain name has been added to ApsaraVideo Live.
	//
	// 	- **DomainExistDcdnipaProduct**: The domain name has been added to DCDN IP Application Accelerator (IPA).
	//
	// example:
	//
	// DomainNotExist
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckCdnDomainExistResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCdnDomainExistResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCdnDomainExistResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCdnDomainExistResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CheckCdnDomainExistResponseBody) SetRequestId(v string) *CheckCdnDomainExistResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCdnDomainExistResponseBody) SetStatus(v string) *CheckCdnDomainExistResponseBody {
	s.Status = &v
	return s
}

func (s *CheckCdnDomainExistResponseBody) Validate() error {
	return dara.Validate(s)
}
