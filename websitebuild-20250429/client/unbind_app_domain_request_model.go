// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAppDomainRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *UnbindAppDomainRequest
	GetBizId() *string
	SetDomainName(v string) *UnbindAppDomainRequest
	GetDomainName() *string
}

type UnbindAppDomainRequest struct {
	// example:
	//
	// WD20250908140837000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// conway.cn
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
}

func (s UnbindAppDomainRequest) String() string {
	return dara.Prettify(s)
}

func (s UnbindAppDomainRequest) GoString() string {
	return s.String()
}

func (s *UnbindAppDomainRequest) GetBizId() *string {
	return s.BizId
}

func (s *UnbindAppDomainRequest) GetDomainName() *string {
	return s.DomainName
}

func (s *UnbindAppDomainRequest) SetBizId(v string) *UnbindAppDomainRequest {
	s.BizId = &v
	return s
}

func (s *UnbindAppDomainRequest) SetDomainName(v string) *UnbindAppDomainRequest {
	s.DomainName = &v
	return s
}

func (s *UnbindAppDomainRequest) Validate() error {
	return dara.Validate(s)
}
