// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIPv6Request interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *UpdateIPv6Request
	GetEnable() *string
	SetRegion(v string) *UpdateIPv6Request
	GetRegion() *string
	SetSiteId(v int64) *UpdateIPv6Request
	GetSiteId() *int64
}

type UpdateIPv6Request struct {
	// Whether to enable IPv6. Valid values:
	//
	// - **on**: Enables IPv6.
	//
	// - **off**: Disables IPv6.
	//
	// This parameter is required.
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The region in which to enable IPv6. The default value is x.x.
	//
	// - x.x: global
	//
	// - cn.cn: Chinese mainland
	//
	// example:
	//
	// x.x
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The site ID. You can get this ID by calling [ListSites](https://help.aliyun.com/document_detail/2850189.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// 5407498413****
	SiteId *int64 `json:"SiteId,omitempty" xml:"SiteId,omitempty"`
}

func (s UpdateIPv6Request) String() string {
	return dara.Prettify(s)
}

func (s UpdateIPv6Request) GoString() string {
	return s.String()
}

func (s *UpdateIPv6Request) GetEnable() *string {
	return s.Enable
}

func (s *UpdateIPv6Request) GetRegion() *string {
	return s.Region
}

func (s *UpdateIPv6Request) GetSiteId() *int64 {
	return s.SiteId
}

func (s *UpdateIPv6Request) SetEnable(v string) *UpdateIPv6Request {
	s.Enable = &v
	return s
}

func (s *UpdateIPv6Request) SetRegion(v string) *UpdateIPv6Request {
	s.Region = &v
	return s
}

func (s *UpdateIPv6Request) SetSiteId(v int64) *UpdateIPv6Request {
	s.SiteId = &v
	return s
}

func (s *UpdateIPv6Request) Validate() error {
	return dara.Validate(s)
}
