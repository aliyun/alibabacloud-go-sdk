// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWebCacheConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainCacheConfigs(v []*DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) *DescribeWebCacheConfigsResponseBody
	GetDomainCacheConfigs() []*DescribeWebCacheConfigsResponseBodyDomainCacheConfigs
	SetRequestId(v string) *DescribeWebCacheConfigsResponseBody
	GetRequestId() *string
}

type DescribeWebCacheConfigsResponseBody struct {
	// An array that consists of Static Page Caching configurations.
	DomainCacheConfigs []*DescribeWebCacheConfigsResponseBodyDomainCacheConfigs `json:"DomainCacheConfigs,omitempty" xml:"DomainCacheConfigs,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 6623EA1F-30FB-5BC8-BEC9-74D55F6F08F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeWebCacheConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCacheConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWebCacheConfigsResponseBody) GetDomainCacheConfigs() []*DescribeWebCacheConfigsResponseBodyDomainCacheConfigs {
	return s.DomainCacheConfigs
}

func (s *DescribeWebCacheConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWebCacheConfigsResponseBody) SetDomainCacheConfigs(v []*DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) *DescribeWebCacheConfigsResponseBody {
	s.DomainCacheConfigs = v
	return s
}

func (s *DescribeWebCacheConfigsResponseBody) SetRequestId(v string) *DescribeWebCacheConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBody) Validate() error {
	if s.DomainCacheConfigs != nil {
		for _, item := range s.DomainCacheConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebCacheConfigsResponseBodyDomainCacheConfigs struct {
	// An array that consists of custom caching rules.
	CustomRules []*DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules `json:"CustomRules,omitempty" xml:"CustomRules,omitempty" type:"Repeated"`
	// The domain name of the website.
	//
	// example:
	//
	// www.aliyundoc.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// The status of the Static Page Caching policy. Valid values:
	//
	// 	- **1**: enabled
	//
	// 	- **0**: disabled
	//
	// example:
	//
	// 1
	Enable *int32 `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The cache mode. Valid values:
	//
	// 	- **standard**: The standard cache mode is used.
	//
	// 	- **aggressive**: The enhanced cache mode is used.
	//
	// 	- **bypass**: No data is cached.
	//
	// example:
	//
	// bypass
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) GoString() string {
	return s.String()
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) GetCustomRules() []*DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules {
	return s.CustomRules
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) GetEnable() *int32 {
	return s.Enable
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) GetMode() *string {
	return s.Mode
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) SetCustomRules(v []*DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs {
	s.CustomRules = v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) SetDomain(v string) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs {
	s.Domain = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) SetEnable(v int32) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs {
	s.Enable = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) SetMode(v string) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs {
	s.Mode = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigs) Validate() error {
	if s.CustomRules != nil {
		for _, item := range s.CustomRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules struct {
	// The expiration time of the page cache. Unit: seconds.
	//
	// example:
	//
	// 86400
	CacheTtl *int64 `json:"CacheTtl,omitempty" xml:"CacheTtl,omitempty"`
	// The cache mode. Valid values:
	//
	// 	- **standard**: The standard cache mode is used.
	//
	// 	- **aggressive**: The enhanced cache mode is used.
	//
	// 	- **bypass**: No data is cached.
	//
	// example:
	//
	// standard
	Mode *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
	// The name of the rule.
	//
	// example:
	//
	// c1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The path to the cached page.
	//
	// example:
	//
	// /blog/
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) String() string {
	return dara.Prettify(s)
}

func (s DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) GoString() string {
	return s.String()
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) GetCacheTtl() *int64 {
	return s.CacheTtl
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) GetMode() *string {
	return s.Mode
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) GetName() *string {
	return s.Name
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) GetUri() *string {
	return s.Uri
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) SetCacheTtl(v int64) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules {
	s.CacheTtl = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) SetMode(v string) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules {
	s.Mode = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) SetName(v string) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules {
	s.Name = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) SetUri(v string) *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules {
	s.Uri = &v
	return s
}

func (s *DescribeWebCacheConfigsResponseBodyDomainCacheConfigsCustomRules) Validate() error {
	return dara.Validate(s)
}
