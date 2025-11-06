// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainSpecialBizInfoByDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizType(v string) *QueryDomainSpecialBizInfoByDomainRequest
	GetBizType() *string
	SetDomainName(v string) *QueryDomainSpecialBizInfoByDomainRequest
	GetDomainName() *string
	SetUserClientIp(v string) *QueryDomainSpecialBizInfoByDomainRequest
	GetUserClientIp() *string
}

type QueryDomainSpecialBizInfoByDomainRequest struct {
	// The business type.
	//
	// This parameter is required.
	//
	// example:
	//
	// govRegister
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// The domain name.
	//
	// This parameter is required.
	//
	// example:
	//
	// test003.cn
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s QueryDomainSpecialBizInfoByDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainSpecialBizInfoByDomainRequest) GoString() string {
	return s.String()
}

func (s *QueryDomainSpecialBizInfoByDomainRequest) GetBizType() *string {
	return s.BizType
}

func (s *QueryDomainSpecialBizInfoByDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *QueryDomainSpecialBizInfoByDomainRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *QueryDomainSpecialBizInfoByDomainRequest) SetBizType(v string) *QueryDomainSpecialBizInfoByDomainRequest {
	s.BizType = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainRequest) SetDomainName(v string) *QueryDomainSpecialBizInfoByDomainRequest {
	s.DomainName = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainRequest) SetUserClientIp(v string) *QueryDomainSpecialBizInfoByDomainRequest {
	s.UserClientIp = &v
	return s
}

func (s *QueryDomainSpecialBizInfoByDomainRequest) Validate() error {
	return dara.Validate(s)
}
