// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitDomainSpecialBizCredentialsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v int64) *SubmitDomainSpecialBizCredentialsRequest
	GetBizId() *int64
	SetCredentials(v string) *SubmitDomainSpecialBizCredentialsRequest
	GetCredentials() *string
	SetExtend(v string) *SubmitDomainSpecialBizCredentialsRequest
	GetExtend() *string
	SetUserClientIp(v string) *SubmitDomainSpecialBizCredentialsRequest
	GetUserClientIp() *string
}

type SubmitDomainSpecialBizCredentialsRequest struct {
	// The business ID.
	//
	// example:
	//
	// 219
	BizId *int64 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The certificate information.
	Credentials *string `json:"Credentials,omitempty" xml:"Credentials,omitempty"`
	// The extended information.
	//
	// example:
	//
	// {\\"addTransferLock\\":true}
	Extend *string `json:"Extend,omitempty" xml:"Extend,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s SubmitDomainSpecialBizCredentialsRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitDomainSpecialBizCredentialsRequest) GoString() string {
	return s.String()
}

func (s *SubmitDomainSpecialBizCredentialsRequest) GetBizId() *int64 {
	return s.BizId
}

func (s *SubmitDomainSpecialBizCredentialsRequest) GetCredentials() *string {
	return s.Credentials
}

func (s *SubmitDomainSpecialBizCredentialsRequest) GetExtend() *string {
	return s.Extend
}

func (s *SubmitDomainSpecialBizCredentialsRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *SubmitDomainSpecialBizCredentialsRequest) SetBizId(v int64) *SubmitDomainSpecialBizCredentialsRequest {
	s.BizId = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsRequest) SetCredentials(v string) *SubmitDomainSpecialBizCredentialsRequest {
	s.Credentials = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsRequest) SetExtend(v string) *SubmitDomainSpecialBizCredentialsRequest {
	s.Extend = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsRequest) SetUserClientIp(v string) *SubmitDomainSpecialBizCredentialsRequest {
	s.UserClientIp = &v
	return s
}

func (s *SubmitDomainSpecialBizCredentialsRequest) Validate() error {
	return dara.Validate(s)
}
