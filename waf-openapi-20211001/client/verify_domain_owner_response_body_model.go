// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyDomainOwnerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *VerifyDomainOwnerResponseBody
	GetRequestId() *string
	SetVerifyResult(v *VerifyDomainOwnerResponseBodyVerifyResult) *VerifyDomainOwnerResponseBody
	GetVerifyResult() *VerifyDomainOwnerResponseBodyVerifyResult
}

type VerifyDomainOwnerResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// F35F45B0-5D6B-4238-BE02-A62D****E840
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The verification result.
	VerifyResult *VerifyDomainOwnerResponseBodyVerifyResult `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty" type:"Struct"`
}

func (s VerifyDomainOwnerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyDomainOwnerResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyDomainOwnerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyDomainOwnerResponseBody) GetVerifyResult() *VerifyDomainOwnerResponseBodyVerifyResult {
	return s.VerifyResult
}

func (s *VerifyDomainOwnerResponseBody) SetRequestId(v string) *VerifyDomainOwnerResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyDomainOwnerResponseBody) SetVerifyResult(v *VerifyDomainOwnerResponseBodyVerifyResult) *VerifyDomainOwnerResponseBody {
	s.VerifyResult = v
	return s
}

func (s *VerifyDomainOwnerResponseBody) Validate() error {
	if s.VerifyResult != nil {
		if err := s.VerifyResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type VerifyDomainOwnerResponseBodyVerifyResult struct {
	// The reasons why the verification fails. Valid values:
	//
	// 	- DnsTxtVerifyFailed: The DNS TXT record and the domain name do not match.
	//
	// 	- DnsServerError: The DNS server is abnormal.
	//
	// 	- VerifyFileNotExist: The verification file does not exist.
	//
	// 	- VerifyDomainNotAccess: The access to the domain name failed.
	//
	// 	- FileContentVerifyFailed: The content of the verification file and the domain name do not match.
	//
	// example:
	//
	// DnsTxtVerifyFailed
	FailCode *string `json:"FailCode,omitempty" xml:"FailCode,omitempty"`
	// The verification result. Valid values:
	//
	// 	- **true**: The verification succeeds.
	//
	// 	- **false**: The verification fails.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s VerifyDomainOwnerResponseBodyVerifyResult) String() string {
	return dara.Prettify(s)
}

func (s VerifyDomainOwnerResponseBodyVerifyResult) GoString() string {
	return s.String()
}

func (s *VerifyDomainOwnerResponseBodyVerifyResult) GetFailCode() *string {
	return s.FailCode
}

func (s *VerifyDomainOwnerResponseBodyVerifyResult) GetResult() *bool {
	return s.Result
}

func (s *VerifyDomainOwnerResponseBodyVerifyResult) SetFailCode(v string) *VerifyDomainOwnerResponseBodyVerifyResult {
	s.FailCode = &v
	return s
}

func (s *VerifyDomainOwnerResponseBodyVerifyResult) SetResult(v bool) *VerifyDomainOwnerResponseBodyVerifyResult {
	s.Result = &v
	return s
}

func (s *VerifyDomainOwnerResponseBodyVerifyResult) Validate() error {
	return dara.Validate(s)
}
