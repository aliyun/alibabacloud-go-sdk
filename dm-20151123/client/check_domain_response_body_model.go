// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDomainResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainStatus(v int32) *CheckDomainResponseBody
	GetDomainStatus() *int32
	SetRequestId(v string) *CheckDomainResponseBody
	GetRequestId() *string
}

type CheckDomainResponseBody struct {
	// Domain status. Indicates whether the verification was successful, with values as follows:
	//
	// - **0**: Available, verified successfully
	//
	// - **1**: Unavailable, verification failed
	//
	// example:
	//
	// 1
	DomainStatus *int32 `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
	// Request ID
	//
	// example:
	//
	// F0B82E83-A1D9-4FE6-97D2-F4B231F80B02
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDomainResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainResponseBody) GetDomainStatus() *int32 {
	return s.DomainStatus
}

func (s *CheckDomainResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDomainResponseBody) SetDomainStatus(v int32) *CheckDomainResponseBody {
	s.DomainStatus = &v
	return s
}

func (s *CheckDomainResponseBody) SetRequestId(v string) *CheckDomainResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDomainResponseBody) Validate() error {
	return dara.Validate(s)
}
