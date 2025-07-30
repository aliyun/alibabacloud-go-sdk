// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInitDomainAutoRedirectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableInitDomainAutoRedirectResponseBody
	GetRequestId() *string
}

type DisableInitDomainAutoRedirectResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableInitDomainAutoRedirectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableInitDomainAutoRedirectResponseBody) GoString() string {
	return s.String()
}

func (s *DisableInitDomainAutoRedirectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableInitDomainAutoRedirectResponseBody) SetRequestId(v string) *DisableInitDomainAutoRedirectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableInitDomainAutoRedirectResponseBody) Validate() error {
	return dara.Validate(s)
}
