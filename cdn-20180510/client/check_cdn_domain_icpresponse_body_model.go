// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCdnDomainICPResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckCdnDomainICPResponseBody
	GetRequestId() *string
	SetStatus(v string) *CheckCdnDomainICPResponseBody
	GetStatus() *string
}

type CheckCdnDomainICPResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the resource plan. Valid value:
	//
	// 	- **DomainIsRegistration**: An ICP filing is obtained for the domain name.
	//
	// 	- **DomainNotRegistration**: No ICP filing is obtained for the domain name.
	//
	// example:
	//
	// DomainIsRegistration
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CheckCdnDomainICPResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCdnDomainICPResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCdnDomainICPResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCdnDomainICPResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CheckCdnDomainICPResponseBody) SetRequestId(v string) *CheckCdnDomainICPResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCdnDomainICPResponseBody) SetStatus(v string) *CheckCdnDomainICPResponseBody {
	s.Status = &v
	return s
}

func (s *CheckCdnDomainICPResponseBody) Validate() error {
	return dara.Validate(s)
}
