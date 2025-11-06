// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDomainSunriseClaimResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClaimKey(v string) *CheckDomainSunriseClaimResponseBody
	GetClaimKey() *string
	SetRequestId(v string) *CheckDomainSunriseClaimResponseBody
	GetRequestId() *string
	SetResult(v int32) *CheckDomainSunriseClaimResponseBody
	GetResult() *int32
}

type CheckDomainSunriseClaimResponseBody struct {
	// example:
	//
	// 2017092100/8/2/1/kDfu9htHGEx_y-LJ3XSlKMZ70000020001
	ClaimKey *string `json:"ClaimKey,omitempty" xml:"ClaimKey,omitempty"`
	// example:
	//
	// BA7A4FD4-EB9A-4A20-BB0C-9AEB15634DC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	Result *int32 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CheckDomainSunriseClaimResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDomainSunriseClaimResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDomainSunriseClaimResponseBody) GetClaimKey() *string {
	return s.ClaimKey
}

func (s *CheckDomainSunriseClaimResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDomainSunriseClaimResponseBody) GetResult() *int32 {
	return s.Result
}

func (s *CheckDomainSunriseClaimResponseBody) SetClaimKey(v string) *CheckDomainSunriseClaimResponseBody {
	s.ClaimKey = &v
	return s
}

func (s *CheckDomainSunriseClaimResponseBody) SetRequestId(v string) *CheckDomainSunriseClaimResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDomainSunriseClaimResponseBody) SetResult(v int32) *CheckDomainSunriseClaimResponseBody {
	s.Result = &v
	return s
}

func (s *CheckDomainSunriseClaimResponseBody) Validate() error {
	return dara.Validate(s)
}
