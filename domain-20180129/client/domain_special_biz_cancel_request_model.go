// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomainSpecialBizCancelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v int64) *DomainSpecialBizCancelRequest
	GetBizId() *int64
	SetUserClientIp(v string) *DomainSpecialBizCancelRequest
	GetUserClientIp() *string
}

type DomainSpecialBizCancelRequest struct {
	// The business ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3666
	BizId *int64 `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// The IP address of the client.
	//
	// example:
	//
	// 127.0.0.1
	UserClientIp *string `json:"UserClientIp,omitempty" xml:"UserClientIp,omitempty"`
}

func (s DomainSpecialBizCancelRequest) String() string {
	return dara.Prettify(s)
}

func (s DomainSpecialBizCancelRequest) GoString() string {
	return s.String()
}

func (s *DomainSpecialBizCancelRequest) GetBizId() *int64 {
	return s.BizId
}

func (s *DomainSpecialBizCancelRequest) GetUserClientIp() *string {
	return s.UserClientIp
}

func (s *DomainSpecialBizCancelRequest) SetBizId(v int64) *DomainSpecialBizCancelRequest {
	s.BizId = &v
	return s
}

func (s *DomainSpecialBizCancelRequest) SetUserClientIp(v string) *DomainSpecialBizCancelRequest {
	s.UserClientIp = &v
	return s
}

func (s *DomainSpecialBizCancelRequest) Validate() error {
	return dara.Validate(s)
}
