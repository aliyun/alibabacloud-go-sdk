// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainSpecialBizDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v int64) *QueryDomainSpecialBizDetailRequest
	GetBizId() *int64
	SetUserClientIp(v string) *QueryDomainSpecialBizDetailRequest
	GetUserClientIp() *string
}

type QueryDomainSpecialBizDetailRequest struct {
	// The business ID.
	//
	// example:
	//
	// 123
	BizId *int64 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainSpecialBizDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizDetailRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizDetailRequest) GetBizId() *int64 {
	return s.BizId
}

func (s *QueryDomainSpecialBizDetailRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDomainSpecialBizDetailRequest) SetBizId(v int64) *QueryDomainSpecialBizDetailRequest {
	s.BizId = &v
	return s
}

func (s *QueryDomainSpecialBizDetailRequest) SetUserClientIp(v string) *QueryDomainSpecialBizDetailRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDomainSpecialBizDetailRequest) Validate() error {
	return dara.Validate(s)
}
