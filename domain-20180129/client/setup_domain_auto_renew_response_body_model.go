// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupDomainAutoRenewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetupDomainAutoRenewResponseBody
	GetRequestId() *string
	SetResult(v bool) *SetupDomainAutoRenewResponseBody
	GetResult() *bool
}

type SetupDomainAutoRenewResponseBody struct {
	// example:
	//
	// 8fc97e44-837a-447d-ac61-ea28d2fe8a38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s SetupDomainAutoRenewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetupDomainAutoRenewResponseBody) GoString() string {
	return s.String()
}

func (s *SetupDomainAutoRenewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetupDomainAutoRenewResponseBody) GetResult() *bool {
	return s.Result
}

func (s *SetupDomainAutoRenewResponseBody) SetRequestId(v string) *SetupDomainAutoRenewResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetupDomainAutoRenewResponseBody) SetResult(v bool) *SetupDomainAutoRenewResponseBody {
	s.Result = &v
	return s
}

func (s *SetupDomainAutoRenewResponseBody) Validate() error {
	return dara.Validate(s)
}
