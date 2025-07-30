// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIsKeepAliveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIsKeepAlive(v bool) *IsKeepAliveResponseBody
	GetIsKeepAlive() *bool
	SetOfficeSiteId(v string) *IsKeepAliveResponseBody
	GetOfficeSiteId() *string
	SetRequestId(v string) *IsKeepAliveResponseBody
	GetRequestId() *string
	SetTenantId(v string) *IsKeepAliveResponseBody
	GetTenantId() *string
}

type IsKeepAliveResponseBody struct {
	// example:
	//
	// True
	IsKeepAlive *bool `json:"IsKeepAlive,omitempty" xml:"IsKeepAlive,omitempty"`
	// example:
	//
	// cn-hangzhou+dir-885351****
	OfficeSiteId *string `json:"OfficeSiteId,omitempty" xml:"OfficeSiteId,omitempty"`
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 141631846826****
	TenantId *string `json:"TenantId,omitempty" xml:"TenantId,omitempty"`
}

func (s IsKeepAliveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s IsKeepAliveResponseBody) GoString() string {
	return s.String()
}

func (s *IsKeepAliveResponseBody) GetIsKeepAlive() *bool {
	return s.IsKeepAlive
}

func (s *IsKeepAliveResponseBody) GetOfficeSiteId() *string {
	return s.OfficeSiteId
}

func (s *IsKeepAliveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *IsKeepAliveResponseBody) GetTenantId() *string {
	return s.TenantId
}

func (s *IsKeepAliveResponseBody) SetIsKeepAlive(v bool) *IsKeepAliveResponseBody {
	s.IsKeepAlive = &v
	return s
}

func (s *IsKeepAliveResponseBody) SetOfficeSiteId(v string) *IsKeepAliveResponseBody {
	s.OfficeSiteId = &v
	return s
}

func (s *IsKeepAliveResponseBody) SetRequestId(v string) *IsKeepAliveResponseBody {
	s.RequestId = &v
	return s
}

func (s *IsKeepAliveResponseBody) SetTenantId(v string) *IsKeepAliveResponseBody {
	s.TenantId = &v
	return s
}

func (s *IsKeepAliveResponseBody) Validate() error {
	return dara.Validate(s)
}
