// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSiteFunctionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigs(v *ListSiteFunctionsResponseBodyConfigs) *ListSiteFunctionsResponseBody
	GetConfigs() *ListSiteFunctionsResponseBodyConfigs
	SetPageNumber(v int32) *ListSiteFunctionsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListSiteFunctionsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListSiteFunctionsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListSiteFunctionsResponseBody
	GetTotalCount() *int32
	SetTotalPage(v int32) *ListSiteFunctionsResponseBody
	GetTotalPage() *int32
}

type ListSiteFunctionsResponseBody struct {
	// The response body configurations.
	Configs *ListSiteFunctionsResponseBodyConfigs `json:"Configs,omitempty" xml:"Configs,omitempty" type:"Struct"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The page size.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 16
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The total number of pages.
	//
	// example:
	//
	// 2
	TotalPage *int32 `json:"TotalPage,omitempty" xml:"TotalPage,omitempty"`
}

func (s ListSiteFunctionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBody) GetConfigs() *ListSiteFunctionsResponseBodyConfigs {
	return s.Configs
}

func (s *ListSiteFunctionsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListSiteFunctionsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListSiteFunctionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSiteFunctionsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListSiteFunctionsResponseBody) GetTotalPage() *int32 {
	return s.TotalPage
}

func (s *ListSiteFunctionsResponseBody) SetConfigs(v *ListSiteFunctionsResponseBodyConfigs) *ListSiteFunctionsResponseBody {
	s.Configs = v
	return s
}

func (s *ListSiteFunctionsResponseBody) SetPageNumber(v int32) *ListSiteFunctionsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListSiteFunctionsResponseBody) SetPageSize(v int32) *ListSiteFunctionsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListSiteFunctionsResponseBody) SetRequestId(v string) *ListSiteFunctionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSiteFunctionsResponseBody) SetTotalCount(v int32) *ListSiteFunctionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListSiteFunctionsResponseBody) SetTotalPage(v int32) *ListSiteFunctionsResponseBody {
	s.TotalPage = &v
	return s
}

func (s *ListSiteFunctionsResponseBody) Validate() error {
	if s.Configs != nil {
		if err := s.Configs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListSiteFunctionsResponseBodyConfigs struct {
	// The cache reserve configuration.
	CacheReserve []*ListSiteFunctionsResponseBodyConfigsCacheReserve `json:"CacheReserve,omitempty" xml:"CacheReserve,omitempty" type:"Repeated"`
	// The cache rules.
	CacheRules []*ListSiteFunctionsResponseBodyConfigsCacheRules `json:"CacheRules,omitempty" xml:"CacheRules,omitempty" type:"Repeated"`
	// The cache tag configuration. When using the purge-by-cache-tag feature, specifies the CacheTag name carried in the origin server response.
	CacheTags []*ListSiteFunctionsResponseBodyConfigsCacheTags `json:"CacheTags,omitempty" xml:"CacheTags,omitempty" type:"Repeated"`
	// The CNAME flattening configuration.
	CnameFlattening []*ListSiteFunctionsResponseBodyConfigsCnameFlattening `json:"CnameFlattening,omitempty" xml:"CnameFlattening,omitempty" type:"Repeated"`
	// The compression rules.
	CompressionRules []*ListSiteFunctionsResponseBodyConfigsCompressionRules `json:"CompressionRules,omitempty" xml:"CompressionRules,omitempty" type:"Repeated"`
	// The Chinese mainland network optimization configuration.
	CrossBorderOptimization []*ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization `json:"CrossBorderOptimization,omitempty" xml:"CrossBorderOptimization,omitempty" type:"Repeated"`
	// The custom response code configurations.
	CustomResponseCode []*ListSiteFunctionsResponseBodyConfigsCustomResponseCode `json:"CustomResponseCode,omitempty" xml:"CustomResponseCode,omitempty" type:"Repeated"`
	// The development mode configuration.
	DevelopmentMode []*ListSiteFunctionsResponseBodyConfigsDevelopmentMode `json:"DevelopmentMode,omitempty" xml:"DevelopmentMode,omitempty" type:"Repeated"`
	// The error code redirect rules.
	ErrorPagesRedirects []*ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects `json:"ErrorPagesRedirects,omitempty" xml:"ErrorPagesRedirects,omitempty" type:"Repeated"`
	// The inbound request header modification rules.
	HttpIncomingRequestHeaderModificationRules []*ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules `json:"HttpIncomingRequestHeaderModificationRules,omitempty" xml:"HttpIncomingRequestHeaderModificationRules,omitempty" type:"Repeated"`
	// The rules for modifying inbound response headers.
	HttpIncomingResponseHeaderModificationRules []*ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules `json:"HttpIncomingResponseHeaderModificationRules,omitempty" xml:"HttpIncomingResponseHeaderModificationRules,omitempty" type:"Repeated"`
	// The request header modification rules.
	HttpRequestHeaderModificationRules []*ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules `json:"HttpRequestHeaderModificationRules,omitempty" xml:"HttpRequestHeaderModificationRules,omitempty" type:"Repeated"`
	// The response header modification rules.
	HttpResponseHeaderModificationRules []*ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules `json:"HttpResponseHeaderModificationRules,omitempty" xml:"HttpResponseHeaderModificationRules,omitempty" type:"Repeated"`
	// The HTTPS application configuration.
	HttpsApplicationConfiguration []*ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration `json:"HttpsApplicationConfiguration,omitempty" xml:"HttpsApplicationConfiguration,omitempty" type:"Repeated"`
	// The HTTPS basic configuration.
	HttpsBasicConfiguration []*ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration `json:"HttpsBasicConfiguration,omitempty" xml:"HttpsBasicConfiguration,omitempty" type:"Repeated"`
	// The image transformation configuration.
	ImageTransform []*ListSiteFunctionsResponseBodyConfigsImageTransform `json:"ImageTransform,omitempty" xml:"ImageTransform,omitempty" type:"Repeated"`
	// The IPv6 configuration.
	Ipv6 []*ListSiteFunctionsResponseBodyConfigsIpv6 `json:"Ipv6,omitempty" xml:"Ipv6,omitempty" type:"Repeated"`
	// The managed transforms.
	ManagedTransforms []*ListSiteFunctionsResponseBodyConfigsManagedTransforms `json:"ManagedTransforms,omitempty" xml:"ManagedTransforms,omitempty" type:"Repeated"`
	MarkdownForAgent  []*ListSiteFunctionsResponseBodyConfigsMarkdownForAgent  `json:"MarkdownForAgent,omitempty" xml:"MarkdownForAgent,omitempty" type:"Repeated"`
	// The network optimization configuration.
	NetworkOptimization []*ListSiteFunctionsResponseBodyConfigsNetworkOptimization `json:"NetworkOptimization,omitempty" xml:"NetworkOptimization,omitempty" type:"Repeated"`
	// The back-to-origin rules.
	OriginRules []*ListSiteFunctionsResponseBodyConfigsOriginRules `json:"OriginRules,omitempty" xml:"OriginRules,omitempty" type:"Repeated"`
	// The redirect rules.
	RedirectRules []*ListSiteFunctionsResponseBodyConfigsRedirectRules `json:"RedirectRules,omitempty" xml:"RedirectRules,omitempty" type:"Repeated"`
	// The URL rewrite rules.
	RewriteUrlRules []*ListSiteFunctionsResponseBodyConfigsRewriteUrlRules `json:"RewriteUrlRules,omitempty" xml:"RewriteUrlRules,omitempty" type:"Repeated"`
	// The search engine crawler bypass configuration.
	SeoBypass []*ListSiteFunctionsResponseBodyConfigsSeoBypass `json:"SeoBypass,omitempty" xml:"SeoBypass,omitempty" type:"Repeated"`
	// Site name exclusive. When enabled, other accounts cannot create sites or subsites with the same name as the current site.
	SiteNameExclusive []*ListSiteFunctionsResponseBodyConfigsSiteNameExclusive `json:"SiteNameExclusive,omitempty" xml:"SiteNameExclusive,omitempty" type:"Repeated"`
	// Site acceleration pause. Temporarily pauses the proxy acceleration feature for the entire site. When enabled, all DNS records directly return record values to the client.
	SitePause []*ListSiteFunctionsResponseBodyConfigsSitePause `json:"SitePause,omitempty" xml:"SitePause,omitempty" type:"Repeated"`
	// The tiered cache configuration.
	TieredCache []*ListSiteFunctionsResponseBodyConfigsTieredCache `json:"TieredCache,omitempty" xml:"TieredCache,omitempty" type:"Repeated"`
	// The video processing configurations.
	VideoProcessing []*ListSiteFunctionsResponseBodyConfigsVideoProcessing `json:"VideoProcessing,omitempty" xml:"VideoProcessing,omitempty" type:"Repeated"`
}

func (s ListSiteFunctionsResponseBodyConfigs) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigs) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetCacheReserve() []*ListSiteFunctionsResponseBodyConfigsCacheReserve {
	return s.CacheReserve
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetCacheRules() []*ListSiteFunctionsResponseBodyConfigsCacheRules {
	return s.CacheRules
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetCacheTags() []*ListSiteFunctionsResponseBodyConfigsCacheTags {
	return s.CacheTags
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetCnameFlattening() []*ListSiteFunctionsResponseBodyConfigsCnameFlattening {
	return s.CnameFlattening
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetCompressionRules() []*ListSiteFunctionsResponseBodyConfigsCompressionRules {
	return s.CompressionRules
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetCrossBorderOptimization() []*ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization {
	return s.CrossBorderOptimization
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetCustomResponseCode() []*ListSiteFunctionsResponseBodyConfigsCustomResponseCode {
	return s.CustomResponseCode
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetDevelopmentMode() []*ListSiteFunctionsResponseBodyConfigsDevelopmentMode {
	return s.DevelopmentMode
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetErrorPagesRedirects() []*ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects {
	return s.ErrorPagesRedirects
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetHttpIncomingRequestHeaderModificationRules() []*ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules {
	return s.HttpIncomingRequestHeaderModificationRules
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetHttpIncomingResponseHeaderModificationRules() []*ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules {
	return s.HttpIncomingResponseHeaderModificationRules
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetHttpRequestHeaderModificationRules() []*ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules {
	return s.HttpRequestHeaderModificationRules
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetHttpResponseHeaderModificationRules() []*ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules {
	return s.HttpResponseHeaderModificationRules
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetHttpsApplicationConfiguration() []*ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	return s.HttpsApplicationConfiguration
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetHttpsBasicConfiguration() []*ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	return s.HttpsBasicConfiguration
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetImageTransform() []*ListSiteFunctionsResponseBodyConfigsImageTransform {
	return s.ImageTransform
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetIpv6() []*ListSiteFunctionsResponseBodyConfigsIpv6 {
	return s.Ipv6
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetManagedTransforms() []*ListSiteFunctionsResponseBodyConfigsManagedTransforms {
	return s.ManagedTransforms
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetMarkdownForAgent() []*ListSiteFunctionsResponseBodyConfigsMarkdownForAgent {
	return s.MarkdownForAgent
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetNetworkOptimization() []*ListSiteFunctionsResponseBodyConfigsNetworkOptimization {
	return s.NetworkOptimization
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetOriginRules() []*ListSiteFunctionsResponseBodyConfigsOriginRules {
	return s.OriginRules
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetRedirectRules() []*ListSiteFunctionsResponseBodyConfigsRedirectRules {
	return s.RedirectRules
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetRewriteUrlRules() []*ListSiteFunctionsResponseBodyConfigsRewriteUrlRules {
	return s.RewriteUrlRules
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetSeoBypass() []*ListSiteFunctionsResponseBodyConfigsSeoBypass {
	return s.SeoBypass
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetSiteNameExclusive() []*ListSiteFunctionsResponseBodyConfigsSiteNameExclusive {
	return s.SiteNameExclusive
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetSitePause() []*ListSiteFunctionsResponseBodyConfigsSitePause {
	return s.SitePause
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetTieredCache() []*ListSiteFunctionsResponseBodyConfigsTieredCache {
	return s.TieredCache
}

func (s *ListSiteFunctionsResponseBodyConfigs) GetVideoProcessing() []*ListSiteFunctionsResponseBodyConfigsVideoProcessing {
	return s.VideoProcessing
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetCacheReserve(v []*ListSiteFunctionsResponseBodyConfigsCacheReserve) *ListSiteFunctionsResponseBodyConfigs {
	s.CacheReserve = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetCacheRules(v []*ListSiteFunctionsResponseBodyConfigsCacheRules) *ListSiteFunctionsResponseBodyConfigs {
	s.CacheRules = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetCacheTags(v []*ListSiteFunctionsResponseBodyConfigsCacheTags) *ListSiteFunctionsResponseBodyConfigs {
	s.CacheTags = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetCnameFlattening(v []*ListSiteFunctionsResponseBodyConfigsCnameFlattening) *ListSiteFunctionsResponseBodyConfigs {
	s.CnameFlattening = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetCompressionRules(v []*ListSiteFunctionsResponseBodyConfigsCompressionRules) *ListSiteFunctionsResponseBodyConfigs {
	s.CompressionRules = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetCrossBorderOptimization(v []*ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization) *ListSiteFunctionsResponseBodyConfigs {
	s.CrossBorderOptimization = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetCustomResponseCode(v []*ListSiteFunctionsResponseBodyConfigsCustomResponseCode) *ListSiteFunctionsResponseBodyConfigs {
	s.CustomResponseCode = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetDevelopmentMode(v []*ListSiteFunctionsResponseBodyConfigsDevelopmentMode) *ListSiteFunctionsResponseBodyConfigs {
	s.DevelopmentMode = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetErrorPagesRedirects(v []*ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) *ListSiteFunctionsResponseBodyConfigs {
	s.ErrorPagesRedirects = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetHttpIncomingRequestHeaderModificationRules(v []*ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) *ListSiteFunctionsResponseBodyConfigs {
	s.HttpIncomingRequestHeaderModificationRules = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetHttpIncomingResponseHeaderModificationRules(v []*ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) *ListSiteFunctionsResponseBodyConfigs {
	s.HttpIncomingResponseHeaderModificationRules = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetHttpRequestHeaderModificationRules(v []*ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) *ListSiteFunctionsResponseBodyConfigs {
	s.HttpRequestHeaderModificationRules = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetHttpResponseHeaderModificationRules(v []*ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) *ListSiteFunctionsResponseBodyConfigs {
	s.HttpResponseHeaderModificationRules = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetHttpsApplicationConfiguration(v []*ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) *ListSiteFunctionsResponseBodyConfigs {
	s.HttpsApplicationConfiguration = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetHttpsBasicConfiguration(v []*ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) *ListSiteFunctionsResponseBodyConfigs {
	s.HttpsBasicConfiguration = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetImageTransform(v []*ListSiteFunctionsResponseBodyConfigsImageTransform) *ListSiteFunctionsResponseBodyConfigs {
	s.ImageTransform = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetIpv6(v []*ListSiteFunctionsResponseBodyConfigsIpv6) *ListSiteFunctionsResponseBodyConfigs {
	s.Ipv6 = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetManagedTransforms(v []*ListSiteFunctionsResponseBodyConfigsManagedTransforms) *ListSiteFunctionsResponseBodyConfigs {
	s.ManagedTransforms = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetMarkdownForAgent(v []*ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) *ListSiteFunctionsResponseBodyConfigs {
	s.MarkdownForAgent = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetNetworkOptimization(v []*ListSiteFunctionsResponseBodyConfigsNetworkOptimization) *ListSiteFunctionsResponseBodyConfigs {
	s.NetworkOptimization = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetOriginRules(v []*ListSiteFunctionsResponseBodyConfigsOriginRules) *ListSiteFunctionsResponseBodyConfigs {
	s.OriginRules = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetRedirectRules(v []*ListSiteFunctionsResponseBodyConfigsRedirectRules) *ListSiteFunctionsResponseBodyConfigs {
	s.RedirectRules = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetRewriteUrlRules(v []*ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) *ListSiteFunctionsResponseBodyConfigs {
	s.RewriteUrlRules = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetSeoBypass(v []*ListSiteFunctionsResponseBodyConfigsSeoBypass) *ListSiteFunctionsResponseBodyConfigs {
	s.SeoBypass = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetSiteNameExclusive(v []*ListSiteFunctionsResponseBodyConfigsSiteNameExclusive) *ListSiteFunctionsResponseBodyConfigs {
	s.SiteNameExclusive = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetSitePause(v []*ListSiteFunctionsResponseBodyConfigsSitePause) *ListSiteFunctionsResponseBodyConfigs {
	s.SitePause = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetTieredCache(v []*ListSiteFunctionsResponseBodyConfigsTieredCache) *ListSiteFunctionsResponseBodyConfigs {
	s.TieredCache = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) SetVideoProcessing(v []*ListSiteFunctionsResponseBodyConfigsVideoProcessing) *ListSiteFunctionsResponseBodyConfigs {
	s.VideoProcessing = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigs) Validate() error {
	if s.CacheReserve != nil {
		for _, item := range s.CacheReserve {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CacheRules != nil {
		for _, item := range s.CacheRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CacheTags != nil {
		for _, item := range s.CacheTags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CnameFlattening != nil {
		for _, item := range s.CnameFlattening {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CompressionRules != nil {
		for _, item := range s.CompressionRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CrossBorderOptimization != nil {
		for _, item := range s.CrossBorderOptimization {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CustomResponseCode != nil {
		for _, item := range s.CustomResponseCode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DevelopmentMode != nil {
		for _, item := range s.DevelopmentMode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ErrorPagesRedirects != nil {
		for _, item := range s.ErrorPagesRedirects {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HttpIncomingRequestHeaderModificationRules != nil {
		for _, item := range s.HttpIncomingRequestHeaderModificationRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HttpIncomingResponseHeaderModificationRules != nil {
		for _, item := range s.HttpIncomingResponseHeaderModificationRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HttpRequestHeaderModificationRules != nil {
		for _, item := range s.HttpRequestHeaderModificationRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HttpResponseHeaderModificationRules != nil {
		for _, item := range s.HttpResponseHeaderModificationRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HttpsApplicationConfiguration != nil {
		for _, item := range s.HttpsApplicationConfiguration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.HttpsBasicConfiguration != nil {
		for _, item := range s.HttpsBasicConfiguration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageTransform != nil {
		for _, item := range s.ImageTransform {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Ipv6 != nil {
		for _, item := range s.Ipv6 {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ManagedTransforms != nil {
		for _, item := range s.ManagedTransforms {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.MarkdownForAgent != nil {
		for _, item := range s.MarkdownForAgent {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NetworkOptimization != nil {
		for _, item := range s.NetworkOptimization {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OriginRules != nil {
		for _, item := range s.OriginRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RedirectRules != nil {
		for _, item := range s.RedirectRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RewriteUrlRules != nil {
		for _, item := range s.RewriteUrlRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SeoBypass != nil {
		for _, item := range s.SeoBypass {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SiteNameExclusive != nil {
		for _, item := range s.SiteNameExclusive {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SitePause != nil {
		for _, item := range s.SitePause {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.TieredCache != nil {
		for _, item := range s.TieredCache {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoProcessing != nil {
		for _, item := range s.VideoProcessing {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSiteFunctionsResponseBodyConfigsCacheReserve struct {
	// The configuration ID.
	//
	// example:
	//
	// 392382988376064
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to enable cache reserve. This feature is disabled by default. Valid values:
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The cache reserve instance ID.
	//
	// example:
	//
	// cr_hk_123456789
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsCacheReserve) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsCacheReserve) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheReserve) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheReserve) GetEnable() *string {
	return s.Enable
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheReserve) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheReserve) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsCacheReserve {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheReserve) SetEnable(v string) *ListSiteFunctionsResponseBodyConfigsCacheReserve {
	s.Enable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheReserve) SetInstanceId(v string) *ListSiteFunctionsResponseBodyConfigsCacheReserve {
	s.InstanceId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheReserve) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsCacheRules struct {
	// The ports on which caching is enabled. Valid values: 8880, 2052, 2082, 2086, 2095, 2053, 2083, 2087, and 2096.
	//
	// example:
	//
	// 8880
	AdditionalCacheablePorts *string `json:"AdditionalCacheablePorts,omitempty" xml:"AdditionalCacheablePorts,omitempty"`
	// The browser cache mode. Valid values:
	//
	// example:
	//
	// follow_origin
	BrowserCacheMode *string `json:"BrowserCacheMode,omitempty" xml:"BrowserCacheMode,omitempty"`
	// The browser cache expiration time, in seconds.
	//
	// example:
	//
	// 300
	BrowserCacheTtl *string `json:"BrowserCacheTtl,omitempty" xml:"BrowserCacheTtl,omitempty"`
	// The bypass cache mode. Valid values:
	//
	// - cache_all: all requests are cached.
	//
	// - bypass_all: all requests bypass the cache.
	//
	// example:
	//
	// cache_all
	BypassCache *string `json:"BypassCache,omitempty" xml:"BypassCache,omitempty"`
	// Specifies whether cache deception armor is enabled. This feature protects against web cache deception attacks by caching only content that passes validation. Valid values:
	//
	// example:
	//
	// on
	CacheDeceptionArmor *string `json:"CacheDeceptionArmor,omitempty" xml:"CacheDeceptionArmor,omitempty"`
	// The cache reserve eligibility. Controls whether user requests bypass cache reserve nodes during back-to-origin. Valid values:
	//
	// - bypass_cache_reserve: requests bypass cache reserve.
	//
	// - eligible_for_cache_reserve: eligible for cache reserve.
	//
	// example:
	//
	// bypass_cache_reserve
	CacheReserveEligibility *string `json:"CacheReserveEligibility,omitempty" xml:"CacheReserveEligibility,omitempty"`
	// The cookie names to check for presence when generating cache keys. If a cookie exists, its name (case-insensitive) is added to the cache key. Multiple cookie names are separated by spaces.
	//
	// example:
	//
	// cookiename
	CheckPresenceCookie *string `json:"CheckPresenceCookie,omitempty" xml:"CheckPresenceCookie,omitempty"`
	// The header names to check for presence when generating cache keys. If a header exists, its name (case-insensitive) is added to the cache key. Multiple header names are separated by spaces.
	//
	// example:
	//
	// headername
	CheckPresenceHeader *string `json:"CheckPresenceHeader,omitempty" xml:"CheckPresenceHeader,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The edge cache mode. Valid values:
	//
	// - follow_origin: follows the origin cache policy if one exists. Otherwise, uses the default cache policy.
	//
	// - no_cache: does not cache.
	//
	// - override_origin: overrides the origin cache policy.
	//
	// - follow_origin_bypass: follows the origin cache policy if one exists. Otherwise, does not cache.
	//
	// example:
	//
	// follow_origin
	EdgeCacheMode *string `json:"EdgeCacheMode,omitempty" xml:"EdgeCacheMode,omitempty"`
	// The edge node cache expiration time, in seconds.
	//
	// example:
	//
	// 300
	EdgeCacheTtl *string `json:"EdgeCacheTtl,omitempty" xml:"EdgeCacheTtl,omitempty"`
	// The status code cache expiration time, in seconds.
	//
	// example:
	//
	// 300
	EdgeStatusCodeCacheTtl *string `json:"EdgeStatusCodeCacheTtl,omitempty" xml:"EdgeStatusCodeCacheTtl,omitempty"`
	// The specified cookie names and their values to include when generating cache keys. Multiple values are separated by spaces.
	//
	// example:
	//
	// cookie_exapmle
	IncludeCookie *string `json:"IncludeCookie,omitempty" xml:"IncludeCookie,omitempty"`
	// The specified header names and their values to include when generating cache keys. Multiple values are separated by spaces.
	//
	// example:
	//
	// example
	IncludeHeader *string `json:"IncludeHeader,omitempty" xml:"IncludeHeader,omitempty"`
	// The cache key processing mode.
	//
	// example:
	//
	// ignore
	PostBodyCacheKey *string `json:"PostBodyCacheKey,omitempty" xml:"PostBodyCacheKey,omitempty"`
	// The body size limit, in KB. Supports body sizes from 1 to 8 KB. If the value is empty, the default value of 8 KB takes effect.
	//
	// example:
	//
	// 1
	PostBodySizeLimit *string `json:"PostBodySizeLimit,omitempty" xml:"PostBodySizeLimit,omitempty"`
	// Specifies whether POST caching is enabled.
	//
	// example:
	//
	// on
	PostCache *string `json:"PostCache,omitempty" xml:"PostCache,omitempty"`
	// The query strings to retain or remove. Multiple values are separated by spaces.
	//
	// example:
	//
	// example
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// The processing mode for query strings when generating cache keys. Valid values:
	//
	// example:
	//
	// reserve_all
	QueryStringMode *string `json:"QueryStringMode,omitempty" xml:"QueryStringMode,omitempty"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Specifies whether to serve stale cache. When enabled, edge nodes can respond to user requests with cached expired content when the origin server is unavailable. Valid values:
	//
	// example:
	//
	// on
	ServeStale *string `json:"ServeStale,omitempty" xml:"ServeStale,omitempty"`
	// Specifies whether to sort query strings for caching. Valid values:
	//
	// example:
	//
	// on
	SortQueryStringForCache *string `json:"SortQueryStringForCache,omitempty" xml:"SortQueryStringForCache,omitempty"`
	// Specifies whether to include the type of the client when generating cache keys. Valid values:
	//
	// example:
	//
	// on
	UserDeviceType *string `json:"UserDeviceType,omitempty" xml:"UserDeviceType,omitempty"`
	// Specifies whether to include the client geographic location when generating cache keys. Valid values:
	//
	// example:
	//
	// on
	UserGeo *string `json:"UserGeo,omitempty" xml:"UserGeo,omitempty"`
	// Specifies whether to include the client language type when generating cache keys. Valid values:
	//
	// example:
	//
	// on
	UserLanguage *string `json:"UserLanguage,omitempty" xml:"UserLanguage,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsCacheRules) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsCacheRules) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetAdditionalCacheablePorts() *string {
	return s.AdditionalCacheablePorts
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetBrowserCacheMode() *string {
	return s.BrowserCacheMode
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetBrowserCacheTtl() *string {
	return s.BrowserCacheTtl
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetBypassCache() *string {
	return s.BypassCache
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetCacheDeceptionArmor() *string {
	return s.CacheDeceptionArmor
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetCacheReserveEligibility() *string {
	return s.CacheReserveEligibility
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetCheckPresenceCookie() *string {
	return s.CheckPresenceCookie
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetCheckPresenceHeader() *string {
	return s.CheckPresenceHeader
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetEdgeCacheMode() *string {
	return s.EdgeCacheMode
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetEdgeCacheTtl() *string {
	return s.EdgeCacheTtl
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetEdgeStatusCodeCacheTtl() *string {
	return s.EdgeStatusCodeCacheTtl
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetIncludeCookie() *string {
	return s.IncludeCookie
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetIncludeHeader() *string {
	return s.IncludeHeader
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetPostBodyCacheKey() *string {
	return s.PostBodyCacheKey
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetPostBodySizeLimit() *string {
	return s.PostBodySizeLimit
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetPostCache() *string {
	return s.PostCache
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetQueryString() *string {
	return s.QueryString
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetQueryStringMode() *string {
	return s.QueryStringMode
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetServeStale() *string {
	return s.ServeStale
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetSortQueryStringForCache() *string {
	return s.SortQueryStringForCache
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetUserDeviceType() *string {
	return s.UserDeviceType
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetUserGeo() *string {
	return s.UserGeo
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) GetUserLanguage() *string {
	return s.UserLanguage
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetAdditionalCacheablePorts(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.AdditionalCacheablePorts = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetBrowserCacheMode(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.BrowserCacheMode = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetBrowserCacheTtl(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.BrowserCacheTtl = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetBypassCache(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.BypassCache = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetCacheDeceptionArmor(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.CacheDeceptionArmor = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetCacheReserveEligibility(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.CacheReserveEligibility = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetCheckPresenceCookie(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.CheckPresenceCookie = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetCheckPresenceHeader(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.CheckPresenceHeader = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetEdgeCacheMode(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.EdgeCacheMode = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetEdgeCacheTtl(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.EdgeCacheTtl = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetEdgeStatusCodeCacheTtl(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.EdgeStatusCodeCacheTtl = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetIncludeCookie(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.IncludeCookie = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetIncludeHeader(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.IncludeHeader = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetPostBodyCacheKey(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.PostBodyCacheKey = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetPostBodySizeLimit(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.PostBodySizeLimit = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetPostCache(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.PostCache = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetQueryString(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.QueryString = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetQueryStringMode(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.QueryStringMode = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetServeStale(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.ServeStale = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetSortQueryStringForCache(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.SortQueryStringForCache = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetUserDeviceType(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.UserDeviceType = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetUserGeo(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.UserGeo = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) SetUserLanguage(v string) *ListSiteFunctionsResponseBodyConfigsCacheRules {
	s.UserLanguage = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheRules) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsCacheTags struct {
	// Specifies whether to ignore case. Valid values:
	//
	// example:
	//
	// on
	CaseInsensitive *string `json:"CaseInsensitive,omitempty" xml:"CaseInsensitive,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The custom CacheTag name.
	//
	// example:
	//
	// example
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsCacheTags) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsCacheTags) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheTags) GetCaseInsensitive() *string {
	return s.CaseInsensitive
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheTags) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheTags) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheTags) GetTagName() *string {
	return s.TagName
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheTags) SetCaseInsensitive(v string) *ListSiteFunctionsResponseBodyConfigsCacheTags {
	s.CaseInsensitive = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheTags) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsCacheTags {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheTags) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsCacheTags {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheTags) SetTagName(v string) *ListSiteFunctionsResponseBodyConfigsCacheTags {
	s.TagName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCacheTags) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsCnameFlattening struct {
	// The configuration ID.
	//
	// example:
	//
	// 245523334529026
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The flattening mode. Valid values:
	//
	// - flatten_all: flattens all records.
	//
	// - flatten_at_root: flattens only the root domain. The root domain is flattened by default.
	//
	// example:
	//
	// flatten_all
	FlattenMode *string `json:"FlattenMode,omitempty" xml:"FlattenMode,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 0
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsCnameFlattening) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsCnameFlattening) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsCnameFlattening) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsCnameFlattening) GetFlattenMode() *string {
	return s.FlattenMode
}

func (s *ListSiteFunctionsResponseBodyConfigsCnameFlattening) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsCnameFlattening) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsCnameFlattening {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCnameFlattening) SetFlattenMode(v string) *ListSiteFunctionsResponseBodyConfigsCnameFlattening {
	s.FlattenMode = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCnameFlattening) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsCnameFlattening {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCnameFlattening) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsCompressionRules struct {
	// The Brotli compression setting. Valid values:
	//
	// example:
	//
	// on
	Brotli *string `json:"Brotli,omitempty" xml:"Brotli,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The Gzip compression setting. Valid values:
	//
	// example:
	//
	// on
	Gzip *string `json:"Gzip,omitempty" xml:"Gzip,omitempty"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The Zstd compression setting. Valid values:
	//
	// example:
	//
	// on
	Zstd *string `json:"Zstd,omitempty" xml:"Zstd,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsCompressionRules) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsCompressionRules) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) GetBrotli() *string {
	return s.Brotli
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) GetGzip() *string {
	return s.Gzip
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) GetZstd() *string {
	return s.Zstd
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) SetBrotli(v string) *ListSiteFunctionsResponseBodyConfigsCompressionRules {
	s.Brotli = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsCompressionRules {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) SetGzip(v string) *ListSiteFunctionsResponseBodyConfigsCompressionRules {
	s.Gzip = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsCompressionRules {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsCompressionRules {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsCompressionRules {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsCompressionRules {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) SetZstd(v string) *ListSiteFunctionsResponseBodyConfigsCompressionRules {
	s.Zstd = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCompressionRules) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization struct {
	// The configuration ID.
	//
	// example:
	//
	// 245523334529026
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to enable Chinese mainland network access optimization. Disabled by default. Valid values:
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization) GetEnable() *string {
	return s.Enable
}

func (s *ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization) SetEnable(v string) *ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization {
	s.Enable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCrossBorderOptimization) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsCustomResponseCode struct {
	// The configuration ID.
	//
	// example:
	//
	// 457325144242176
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The response page.
	//
	// example:
	//
	// 1
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// The response code.
	//
	// example:
	//
	// 200
	ReturnCode *string `json:"ReturnCode,omitempty" xml:"ReturnCode,omitempty"`
	// The rule content. A conditional expression is used to match user requests. You do not need to set this parameter when you add a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specified requests: Set the value to a custom expression, for example: (http.host eq \\"video.example.com\\")
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. You do not need to set this parameter when adding a global configuration. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution priority of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsCustomResponseCode) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsCustomResponseCode) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) GetPageId() *string {
	return s.PageId
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) GetReturnCode() *string {
	return s.ReturnCode
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsCustomResponseCode {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) SetPageId(v string) *ListSiteFunctionsResponseBodyConfigsCustomResponseCode {
	s.PageId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) SetReturnCode(v string) *ListSiteFunctionsResponseBodyConfigsCustomResponseCode {
	s.ReturnCode = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsCustomResponseCode {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsCustomResponseCode {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsCustomResponseCode {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsCustomResponseCode {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsCustomResponseCode) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsDevelopmentMode struct {
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The feature switch. Disabled by default. Valid values:
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsDevelopmentMode) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsDevelopmentMode) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsDevelopmentMode) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsDevelopmentMode) GetEnable() *string {
	return s.Enable
}

func (s *ListSiteFunctionsResponseBodyConfigsDevelopmentMode) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsDevelopmentMode) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsDevelopmentMode {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsDevelopmentMode) SetEnable(v string) *ListSiteFunctionsResponseBodyConfigsDevelopmentMode {
	s.Enable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsDevelopmentMode) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsDevelopmentMode {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsDevelopmentMode) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects struct {
	// The configuration ID.
	//
	// example:
	//
	// 473117342636032
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The error code redirect configuration.
	ErrorPagesRedirect []*ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect `json:"ErrorPagesRedirect,omitempty" xml:"ErrorPagesRedirect,omitempty" type:"Repeated"`
	// The rule content. A conditional expression is used to match user requests. You do not need to set this parameter when you add a global configuration. Two scenarios are supported:
	//
	// - Match all incoming requests: Set the value to true.
	//
	// - Match specified requests: Set the value to a custom expression, such as (http.host eq \\"video.example.com\\").
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether to enable the rule. You do not need to set this parameter when adding a global configuration. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name. You do not need to set this parameter when adding a global configuration.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution priority of the rule. A smaller value indicates a higher priority.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) GetErrorPagesRedirect() []*ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect {
	return s.ErrorPagesRedirect
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) SetErrorPagesRedirect(v []*ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect) *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects {
	s.ErrorPagesRedirect = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirects) Validate() error {
	if s.ErrorPagesRedirect != nil {
		for _, item := range s.ErrorPagesRedirect {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect struct {
	// The response status code used by the node when returning the redirect address to the client. Valid values:
	//
	// example:
	//
	// 400
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The target URL to which the request is redirected.
	//
	// example:
	//
	// http://example.com/test
	TargetURL *string `json:"TargetURL,omitempty" xml:"TargetURL,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect) GetTargetURL() *string {
	return s.TargetURL
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect) SetStatusCode(v string) *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect {
	s.StatusCode = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect) SetTargetURL(v string) *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect {
	s.TargetURL = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsErrorPagesRedirectsErrorPagesRedirect) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules struct {
	// The configuration ID.
	//
	// example:
	//
	// 430893999331328
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The request header modifications. Supports add, delete, and modify operations.
	RequestHeaderModification []*ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	// The rule content.
	//
	// example:
	//
	// http.host eq "videoo.example.com"
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) GetRequestHeaderModification() []*ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification {
	return s.RequestHeaderModification
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) SetRequestHeaderModification(v []*ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification) *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules {
	s.RequestHeaderModification = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRules) Validate() error {
	if s.RequestHeaderModification != nil {
		for _, item := range s.RequestHeaderModification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification struct {
	// The request header name.
	//
	// example:
	//
	// headername
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation type. Valid values:
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The request header value.
	//
	// example:
	//
	// headervalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification) GetName() *string {
	return s.Name
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification) GetValue() *string {
	return s.Value
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification) SetName(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification {
	s.Name = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification) SetOperation(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification {
	s.Operation = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification) SetValue(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification {
	s.Value = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingRequestHeaderModificationRulesRequestHeaderModification) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules struct {
	// The configuration ID.
	//
	// example:
	//
	// 430893999331328
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The response header modifications. Supports add, delete, and modify operations.
	ResponseHeaderModification []*ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty" type:"Repeated"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) GetResponseHeaderModification() []*ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification {
	return s.ResponseHeaderModification
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) SetResponseHeaderModification(v []*ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification) *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules {
	s.ResponseHeaderModification = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRules) Validate() error {
	if s.ResponseHeaderModification != nil {
		for _, item := range s.ResponseHeaderModification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification struct {
	// The response header name.
	//
	// example:
	//
	// headername
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation type. Valid values:
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The response header value.
	//
	// example:
	//
	// headervalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification) GetName() *string {
	return s.Name
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification) GetValue() *string {
	return s.Value
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification) SetName(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification {
	s.Name = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification) SetOperation(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification {
	s.Operation = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification) SetValue(v string) *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification {
	s.Value = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpIncomingResponseHeaderModificationRulesResponseHeaderModification) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules struct {
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The request header modifications. Supports add, delete, and modify operations.
	//
	// example:
	//
	// [{"operation":"add","name":"header_example_add","value":"value_exapme_add"},{"operation":"del","name":"header_example_delete","value":"value_exapme_delete"},{"operation":"modify","name":"header_example_update","value":"value_exapme_example"}]
	RequestHeaderModification []*ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification `json:"RequestHeaderModification,omitempty" xml:"RequestHeaderModification,omitempty" type:"Repeated"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) GetRequestHeaderModification() []*ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification {
	return s.RequestHeaderModification
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) SetRequestHeaderModification(v []*ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification) *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules {
	s.RequestHeaderModification = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRules) Validate() error {
	if s.RequestHeaderModification != nil {
		for _, item := range s.RequestHeaderModification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification struct {
	// The request header name.
	//
	// example:
	//
	// headername
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation type. Valid values:
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The request header value.
	//
	// example:
	//
	// headervalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification) GetName() *string {
	return s.Name
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification) GetValue() *string {
	return s.Value
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification) SetName(v string) *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification {
	s.Name = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification) SetOperation(v string) *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification {
	s.Operation = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification) SetValue(v string) *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification {
	s.Value = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpRequestHeaderModificationRulesRequestHeaderModification) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules struct {
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The response header modifications. Supports add, delete, and modify operations.
	//
	// example:
	//
	// [{"operation":"add","name":"header_example_add","value":"value_exapme_add"},{"operation":"del","name":"header_example_delete","value":"value_exapme_delete"},{"operation":"modify","name":"header_example_update","value":"value_exapme_example"}]
	ResponseHeaderModification []*ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification `json:"ResponseHeaderModification,omitempty" xml:"ResponseHeaderModification,omitempty" type:"Repeated"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) GetResponseHeaderModification() []*ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification {
	return s.ResponseHeaderModification
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) SetResponseHeaderModification(v []*ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification) *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules {
	s.ResponseHeaderModification = v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRules) Validate() error {
	if s.ResponseHeaderModification != nil {
		for _, item := range s.ResponseHeaderModification {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification struct {
	// The response header name.
	//
	// example:
	//
	// headername
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The operation type. Valid values:
	//
	// example:
	//
	// add
	Operation *string `json:"Operation,omitempty" xml:"Operation,omitempty"`
	// The response header value.
	//
	// example:
	//
	// headervalue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification) GetName() *string {
	return s.Name
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification) GetOperation() *string {
	return s.Operation
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification) GetValue() *string {
	return s.Value
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification) SetName(v string) *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification {
	s.Name = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification) SetOperation(v string) *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification {
	s.Operation = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification) SetValue(v string) *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification {
	s.Value = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpResponseHeaderModificationRulesResponseHeaderModification) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration struct {
	// The Alt-Svc feature switch. Disabled by default. Valid values:
	//
	// example:
	//
	// on
	AltSvc *string `json:"AltSvc,omitempty" xml:"AltSvc,omitempty"`
	// Specifies whether the Alt-Svc header includes the clear parameter. Disabled by default. Valid values:
	//
	// example:
	//
	// on
	AltSvcClear *string `json:"AltSvcClear,omitempty" xml:"AltSvcClear,omitempty"`
	// The Alt-Svc validity period, in seconds. Default value: 86400.
	//
	// example:
	//
	// 86400
	AltSvcMa *string `json:"AltSvcMa,omitempty" xml:"AltSvcMa,omitempty"`
	// Specifies whether the Alt-Svc header includes the persist parameter. Disabled by default. Valid values:
	//
	// example:
	//
	// on
	AltSvcPersist *string `json:"AltSvcPersist,omitempty" xml:"AltSvcPersist,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 391240445274112
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to enable HSTS. Disabled by default. Valid values:
	//
	// example:
	//
	// on
	Hsts *string `json:"Hsts,omitempty" xml:"Hsts,omitempty"`
	// Specifies whether to include subdomains in HSTS. Disabled by default. Valid values:
	//
	// example:
	//
	// on
	HstsIncludeSubdomains *string `json:"HstsIncludeSubdomains,omitempty" xml:"HstsIncludeSubdomains,omitempty"`
	// The HSTS expiration time, in seconds.
	//
	// example:
	//
	// 3600
	HstsMaxAge *string `json:"HstsMaxAge,omitempty" xml:"HstsMaxAge,omitempty"`
	// Specifies whether to enable HSTS preload. Disabled by default. Valid values:
	//
	// example:
	//
	// on
	HstsPreload *string `json:"HstsPreload,omitempty" xml:"HstsPreload,omitempty"`
	// Specifies whether to enable forced HTTPS. Disabled by default. Valid values:
	//
	// example:
	//
	// on
	HttpsForce *string `json:"HttpsForce,omitempty" xml:"HttpsForce,omitempty"`
	// The HTTP status code for forced HTTPS redirect. Valid values:
	//
	// example:
	//
	// 301
	HttpsForceCode *string `json:"HttpsForceCode,omitempty" xml:"HttpsForceCode,omitempty"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetAltSvc() *string {
	return s.AltSvc
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetAltSvcClear() *string {
	return s.AltSvcClear
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetAltSvcMa() *string {
	return s.AltSvcMa
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetAltSvcPersist() *string {
	return s.AltSvcPersist
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetHsts() *string {
	return s.Hsts
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetHstsIncludeSubdomains() *string {
	return s.HstsIncludeSubdomains
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetHstsMaxAge() *string {
	return s.HstsMaxAge
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetHstsPreload() *string {
	return s.HstsPreload
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetHttpsForce() *string {
	return s.HttpsForce
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetHttpsForceCode() *string {
	return s.HttpsForceCode
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetAltSvc(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.AltSvc = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetAltSvcClear(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.AltSvcClear = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetAltSvcMa(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.AltSvcMa = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetAltSvcPersist(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.AltSvcPersist = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetHsts(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.Hsts = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetHstsIncludeSubdomains(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.HstsIncludeSubdomains = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetHstsMaxAge(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.HstsMaxAge = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetHstsPreload(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.HstsPreload = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetHttpsForce(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.HttpsForce = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetHttpsForceCode(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.HttpsForceCode = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsApplicationConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration struct {
	// The custom cipher suite. Specifies the specific encryption algorithms selected when CiphersuiteGroup is set to custom.
	//
	// example:
	//
	// TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256
	Ciphersuite *string `json:"Ciphersuite,omitempty" xml:"Ciphersuite,omitempty"`
	// The cipher suite group. By default, all cipher suites are enabled. Valid values:
	//
	// - all: all cipher suites.
	//
	// - strict: strong cipher suites.
	//
	// - custom: custom cipher suites.
	//
	// example:
	//
	// all
	CiphersuiteGroup *string `json:"CiphersuiteGroup,omitempty" xml:"CiphersuiteGroup,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 391380266602496
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to enable HTTP/2. Enabled by default. Valid values:
	//
	// example:
	//
	// on
	Http2 *string `json:"Http2,omitempty" xml:"Http2,omitempty"`
	// Specifies whether to enable HTTP/3. Enabled by default. Valid values:
	//
	// example:
	//
	// on
	Http3 *string `json:"Http3,omitempty" xml:"Http3,omitempty"`
	// Specifies whether to enable HTTPS. Enabled by default. Valid values:
	//
	// example:
	//
	// on
	Https *string `json:"Https,omitempty" xml:"Https,omitempty"`
	// Specifies whether to enable OCSP stapling. Disabled by default. Valid values:
	//
	// example:
	//
	// on
	OcspStapling *string `json:"OcspStapling,omitempty" xml:"OcspStapling,omitempty"`
	// The matching rule content.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Specifies whether to enable TLS 1.0. Disabled by default. Valid values:
	//
	// example:
	//
	// on
	Tls10 *string `json:"Tls10,omitempty" xml:"Tls10,omitempty"`
	// Specifies whether to enable TLS 1.1. Enabled by default. Valid values:
	//
	// example:
	//
	// on
	Tls11 *string `json:"Tls11,omitempty" xml:"Tls11,omitempty"`
	// Specifies whether to enable TLS 1.2. Enabled by default. Valid values:
	//
	// example:
	//
	// on
	Tls12 *string `json:"Tls12,omitempty" xml:"Tls12,omitempty"`
	// Specifies whether to enable TLS 1.3. Enabled by default. Valid values:
	//
	// example:
	//
	// on
	Tls13 *string `json:"Tls13,omitempty" xml:"Tls13,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetCiphersuite() *string {
	return s.Ciphersuite
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetCiphersuiteGroup() *string {
	return s.CiphersuiteGroup
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetHttp2() *string {
	return s.Http2
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetHttp3() *string {
	return s.Http3
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetHttps() *string {
	return s.Https
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetOcspStapling() *string {
	return s.OcspStapling
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetTls10() *string {
	return s.Tls10
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetTls11() *string {
	return s.Tls11
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetTls12() *string {
	return s.Tls12
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) GetTls13() *string {
	return s.Tls13
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetCiphersuite(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.Ciphersuite = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetCiphersuiteGroup(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.CiphersuiteGroup = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetHttp2(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.Http2 = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetHttp3(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.Http3 = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetHttps(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.Https = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetOcspStapling(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.OcspStapling = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetTls10(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.Tls10 = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetTls11(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.Tls11 = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetTls12(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.Tls12 = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) SetTls13(v string) *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration {
	s.Tls13 = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsHttpsBasicConfiguration) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsImageTransform struct {
	// The adaptive AVIF setting.
	//
	// example:
	//
	// on
	AutoAvif *string `json:"AutoAvif,omitempty" xml:"AutoAvif,omitempty"`
	// The adaptive WebP setting.
	//
	// example:
	//
	// on
	AutoWebp *string `json:"AutoWebp,omitempty" xml:"AutoWebp,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to enable image transformation. This feature is disabled by default. Valid values:
	//
	// example:
	//
	// off
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsImageTransform) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsImageTransform) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) GetAutoAvif() *string {
	return s.AutoAvif
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) GetAutoWebp() *string {
	return s.AutoWebp
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) GetEnable() *string {
	return s.Enable
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) SetAutoAvif(v string) *ListSiteFunctionsResponseBodyConfigsImageTransform {
	s.AutoAvif = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) SetAutoWebp(v string) *ListSiteFunctionsResponseBodyConfigsImageTransform {
	s.AutoWebp = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsImageTransform {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) SetEnable(v string) *ListSiteFunctionsResponseBodyConfigsImageTransform {
	s.Enable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsImageTransform {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsImageTransform {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsImageTransform {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsImageTransform {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsImageTransform) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsIpv6 struct {
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to enable IPv6. Enabled by default. Valid values:
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsIpv6) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsIpv6) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsIpv6) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsIpv6) GetEnable() *string {
	return s.Enable
}

func (s *ListSiteFunctionsResponseBodyConfigsIpv6) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsIpv6) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsIpv6 {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsIpv6) SetEnable(v string) *ListSiteFunctionsResponseBodyConfigsIpv6 {
	s.Enable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsIpv6) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsIpv6 {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsIpv6) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsManagedTransforms struct {
	// Specifies whether to add visitor geolocation headers. Valid values:
	//
	// example:
	//
	// on
	AddClientGeolocationHeaders *string `json:"AddClientGeolocationHeaders,omitempty" xml:"AddClientGeolocationHeaders,omitempty"`
	// Adds the "ali-real-client-ip" header that contains the originating IP address of the client. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// example:
	//
	// on
	AddRealClientIpHeader *string `json:"AddRealClientIpHeader,omitempty" xml:"AddRealClientIpHeader,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsManagedTransforms) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsManagedTransforms) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsManagedTransforms) GetAddClientGeolocationHeaders() *string {
	return s.AddClientGeolocationHeaders
}

func (s *ListSiteFunctionsResponseBodyConfigsManagedTransforms) GetAddRealClientIpHeader() *string {
	return s.AddRealClientIpHeader
}

func (s *ListSiteFunctionsResponseBodyConfigsManagedTransforms) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsManagedTransforms) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsManagedTransforms) SetAddClientGeolocationHeaders(v string) *ListSiteFunctionsResponseBodyConfigsManagedTransforms {
	s.AddClientGeolocationHeaders = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsManagedTransforms) SetAddRealClientIpHeader(v string) *ListSiteFunctionsResponseBodyConfigsManagedTransforms {
	s.AddRealClientIpHeader = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsManagedTransforms) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsManagedTransforms {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsManagedTransforms) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsManagedTransforms {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsManagedTransforms) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsMarkdownForAgent struct {
	ConfigId   *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	Enable     *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Rule       *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	RuleName   *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	Sequence   *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) GetEnable() *string {
	return s.Enable
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) SetEnable(v string) *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent {
	s.Enable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsMarkdownForAgent) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsNetworkOptimization struct {
	// The configuration ID.
	//
	// example:
	//
	// 395901755670528
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to enable gRPC. This feature is disabled by default. Valid values:
	//
	// example:
	//
	// on
	Grpc *string `json:"Grpc,omitempty" xml:"Grpc,omitempty"`
	// Specifies whether to enable HTTP/2 back-to-origin. This feature is disabled by default. Valid values:
	//
	// example:
	//
	// on
	Http2Origin *string `json:"Http2Origin,omitempty" xml:"Http2Origin,omitempty"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// The rule switch. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Specifies whether to enable the smart routing service. This feature is disabled by default. Valid values:
	//
	// example:
	//
	// on
	SmartRouting *string `json:"SmartRouting,omitempty" xml:"SmartRouting,omitempty"`
	// The maximum upload file size. Unit: MB. Valid values: 100 to 500.
	//
	// example:
	//
	// 300
	UploadMaxFilesize *string `json:"UploadMaxFilesize,omitempty" xml:"UploadMaxFilesize,omitempty"`
	// Specifies whether to enable WebSocket. This feature is enabled by default. Valid values:
	//
	// example:
	//
	// on
	Websocket *string `json:"Websocket,omitempty" xml:"Websocket,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsNetworkOptimization) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsNetworkOptimization) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) GetGrpc() *string {
	return s.Grpc
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) GetHttp2Origin() *string {
	return s.Http2Origin
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) GetSmartRouting() *string {
	return s.SmartRouting
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) GetUploadMaxFilesize() *string {
	return s.UploadMaxFilesize
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) GetWebsocket() *string {
	return s.Websocket
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsNetworkOptimization {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) SetGrpc(v string) *ListSiteFunctionsResponseBodyConfigsNetworkOptimization {
	s.Grpc = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) SetHttp2Origin(v string) *ListSiteFunctionsResponseBodyConfigsNetworkOptimization {
	s.Http2Origin = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsNetworkOptimization {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsNetworkOptimization {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsNetworkOptimization {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsNetworkOptimization {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) SetSmartRouting(v string) *ListSiteFunctionsResponseBodyConfigsNetworkOptimization {
	s.SmartRouting = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) SetUploadMaxFilesize(v string) *ListSiteFunctionsResponseBodyConfigsNetworkOptimization {
	s.UploadMaxFilesize = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) SetWebsocket(v string) *ListSiteFunctionsResponseBodyConfigsNetworkOptimization {
	s.Websocket = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsNetworkOptimization) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsOriginRules struct {
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The rewritten DNS resolution record for back-to-origin requests.
	//
	// example:
	//
	// test.example.com
	DnsRecord *string `json:"DnsRecord,omitempty" xml:"DnsRecord,omitempty"`
	// The HOST header carried in the back-to-origin request.
	//
	// example:
	//
	// origin.example.com
	OriginHost *string `json:"OriginHost,omitempty" xml:"OriginHost,omitempty"`
	// The origin server port used when fetching content over HTTP.
	//
	// example:
	//
	// 8080
	OriginHttpPort *string `json:"OriginHttpPort,omitempty" xml:"OriginHttpPort,omitempty"`
	// The origin server port used when fetching content over HTTPS.
	//
	// example:
	//
	// 4433
	OriginHttpsPort *string `json:"OriginHttpsPort,omitempty" xml:"OriginHttpsPort,omitempty"`
	// Specifies whether to enable mTLS for back-to-origin requests. Valid values:
	//
	// example:
	//
	// on
	OriginMtls *string `json:"OriginMtls,omitempty" xml:"OriginMtls,omitempty"`
	// The origin read timeout, in seconds.
	//
	// example:
	//
	// 300
	OriginReadTimeout *string `json:"OriginReadTimeout,omitempty" xml:"OriginReadTimeout,omitempty"`
	// The protocol used for back-to-origin requests. Valid values:
	//
	// example:
	//
	// http
	OriginScheme *string `json:"OriginScheme,omitempty" xml:"OriginScheme,omitempty"`
	// The SNI carried in the back-to-origin request.
	//
	// example:
	//
	// origin.example.com
	OriginSni *string `json:"OriginSni,omitempty" xml:"OriginSni,omitempty"`
	// Specifies whether to enable origin server certificate verification. Valid values:
	//
	// example:
	//
	// on
	OriginVerify *string `json:"OriginVerify,omitempty" xml:"OriginVerify,omitempty"`
	// Uses range-based origin fetch to download files. Valid values:
	//
	// - on: enabled.
	//
	// - off: disabled.
	//
	// - force: forced.
	//
	// example:
	//
	// on
	Range *string `json:"Range,omitempty" xml:"Range,omitempty"`
	// The range chunk size. Valid values:
	//
	// example:
	//
	// 512KB
	RangeChunkSize *string `json:"RangeChunkSize,omitempty" xml:"RangeChunkSize,omitempty"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsOriginRules) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsOriginRules) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetDnsRecord() *string {
	return s.DnsRecord
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetOriginHost() *string {
	return s.OriginHost
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetOriginHttpPort() *string {
	return s.OriginHttpPort
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetOriginHttpsPort() *string {
	return s.OriginHttpsPort
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetOriginMtls() *string {
	return s.OriginMtls
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetOriginReadTimeout() *string {
	return s.OriginReadTimeout
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetOriginScheme() *string {
	return s.OriginScheme
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetOriginSni() *string {
	return s.OriginSni
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetOriginVerify() *string {
	return s.OriginVerify
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetRange() *string {
	return s.Range
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetRangeChunkSize() *string {
	return s.RangeChunkSize
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetDnsRecord(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.DnsRecord = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetOriginHost(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.OriginHost = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetOriginHttpPort(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.OriginHttpPort = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetOriginHttpsPort(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.OriginHttpsPort = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetOriginMtls(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.OriginMtls = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetOriginReadTimeout(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.OriginReadTimeout = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetOriginScheme(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.OriginScheme = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetOriginSni(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.OriginSni = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetOriginVerify(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.OriginVerify = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetRange(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.Range = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetRangeChunkSize(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.RangeChunkSize = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsOriginRules {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsOriginRules) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsRedirectRules struct {
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to reserve the query string. Valid values:
	//
	// example:
	//
	// on
	ReserveQueryString *string `json:"ReserveQueryString,omitempty" xml:"ReserveQueryString,omitempty"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The HTTP status code used when the node responds to the client with a redirect address. Valid values:
	//
	// - 301
	//
	// - 302
	//
	// - 303
	//
	// - 307
	//
	// - 308
	//
	// example:
	//
	// 301
	StatusCode *string `json:"StatusCode,omitempty" xml:"StatusCode,omitempty"`
	// The target URL after redirection.
	//
	// example:
	//
	// http://www.exapmle.com/index.html
	TargetUrl *string `json:"TargetUrl,omitempty" xml:"TargetUrl,omitempty"`
	// The redirect type. Valid values:
	//
	// example:
	//
	// static
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsRedirectRules) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsRedirectRules) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) GetReserveQueryString() *string {
	return s.ReserveQueryString
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) GetStatusCode() *string {
	return s.StatusCode
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) GetTargetUrl() *string {
	return s.TargetUrl
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) GetType() *string {
	return s.Type
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsRedirectRules {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) SetReserveQueryString(v string) *ListSiteFunctionsResponseBodyConfigsRedirectRules {
	s.ReserveQueryString = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsRedirectRules {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsRedirectRules {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsRedirectRules {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsRedirectRules {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) SetStatusCode(v string) *ListSiteFunctionsResponseBodyConfigsRedirectRules {
	s.StatusCode = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) SetTargetUrl(v string) *ListSiteFunctionsResponseBodyConfigsRedirectRules {
	s.TargetUrl = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) SetType(v string) *ListSiteFunctionsResponseBodyConfigsRedirectRules {
	s.Type = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRedirectRules) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsRewriteUrlRules struct {
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The query string after rewriting.
	//
	// example:
	//
	// example=123
	QueryString *string `json:"QueryString,omitempty" xml:"QueryString,omitempty"`
	// The query string rewrite type. Valid values:
	//
	// example:
	//
	// static
	RewriteQueryStringType *string `json:"RewriteQueryStringType,omitempty" xml:"RewriteQueryStringType,omitempty"`
	// The path rewrite type. Valid values:
	//
	// example:
	//
	// static
	RewriteUriType *string `json:"RewriteUriType,omitempty" xml:"RewriteUriType,omitempty"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq \\"video.example.com\\")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The target URI after rewriting.
	//
	// example:
	//
	// /image.example.com/index.html
	Uri *string `json:"Uri,omitempty" xml:"Uri,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) GetQueryString() *string {
	return s.QueryString
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) GetRewriteQueryStringType() *string {
	return s.RewriteQueryStringType
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) GetRewriteUriType() *string {
	return s.RewriteUriType
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) GetUri() *string {
	return s.Uri
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) SetQueryString(v string) *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules {
	s.QueryString = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) SetRewriteQueryStringType(v string) *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules {
	s.RewriteQueryStringType = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) SetRewriteUriType(v string) *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules {
	s.RewriteUriType = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) SetUri(v string) *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules {
	s.Uri = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsRewriteUrlRules) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsSeoBypass struct {
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The feature switch. Disabled by default. Valid values:
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsSeoBypass) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsSeoBypass) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsSeoBypass) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsSeoBypass) GetEnable() *string {
	return s.Enable
}

func (s *ListSiteFunctionsResponseBodyConfigsSeoBypass) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsSeoBypass) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsSeoBypass {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsSeoBypass) SetEnable(v string) *ListSiteFunctionsResponseBodyConfigsSeoBypass {
	s.Enable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsSeoBypass) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsSeoBypass {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsSeoBypass) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsSiteNameExclusive struct {
	// The configuration ID.
	//
	// example:
	//
	// 380858020294656
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The feature switch. Valid values:
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 0
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsSiteNameExclusive) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsSiteNameExclusive) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsSiteNameExclusive) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsSiteNameExclusive) GetEnable() *string {
	return s.Enable
}

func (s *ListSiteFunctionsResponseBodyConfigsSiteNameExclusive) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsSiteNameExclusive) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsSiteNameExclusive {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsSiteNameExclusive) SetEnable(v string) *ListSiteFunctionsResponseBodyConfigsSiteNameExclusive {
	s.Enable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsSiteNameExclusive) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsSiteNameExclusive {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsSiteNameExclusive) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsSitePause struct {
	// The configuration ID.
	//
	// example:
	//
	// 302426190190592
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Specifies whether to temporarily pause the proxy acceleration feature for the entire site. When enabled, all DNS records directly return record values to the client. Valid values:
	//
	// - true: Site acceleration is paused.
	//
	// - false: Site acceleration is active.
	//
	// example:
	//
	// false
	Paused *string `json:"Paused,omitempty" xml:"Paused,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsSitePause) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsSitePause) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsSitePause) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsSitePause) GetPaused() *string {
	return s.Paused
}

func (s *ListSiteFunctionsResponseBodyConfigsSitePause) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsSitePause) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsSitePause {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsSitePause) SetPaused(v string) *ListSiteFunctionsResponseBodyConfigsSitePause {
	s.Paused = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsSitePause) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsSitePause {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsSitePause) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsTieredCache struct {
	// The tiered cache architecture mode. Valid values:
	//
	// example:
	//
	// edge_smart
	CacheArchitectureMode *string `json:"CacheArchitectureMode,omitempty" xml:"CacheArchitectureMode,omitempty"`
	// The configuration ID.
	//
	// example:
	//
	// 352816096987136
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsTieredCache) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsTieredCache) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsTieredCache) GetCacheArchitectureMode() *string {
	return s.CacheArchitectureMode
}

func (s *ListSiteFunctionsResponseBodyConfigsTieredCache) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsTieredCache) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsTieredCache) SetCacheArchitectureMode(v string) *ListSiteFunctionsResponseBodyConfigsTieredCache {
	s.CacheArchitectureMode = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsTieredCache) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsTieredCache {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsTieredCache) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsTieredCache {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsTieredCache) Validate() error {
	return dara.Validate(s)
}

type ListSiteFunctionsResponseBodyConfigsVideoProcessing struct {
	// The configuration ID.
	//
	// example:
	//
	// 455153377667072
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The custom FLV end parameter.
	//
	// example:
	//
	// f_end
	FlvSeekEnd *string `json:"FlvSeekEnd,omitempty" xml:"FlvSeekEnd,omitempty"`
	// The custom FLV start parameter.
	//
	// example:
	//
	// f_start
	FlvSeekStart *string `json:"FlvSeekStart,omitempty" xml:"FlvSeekStart,omitempty"`
	// The FLV seeking mode. Valid values:
	//
	// example:
	//
	// by_time
	FlvVideoSeekMode *string `json:"FlvVideoSeekMode,omitempty" xml:"FlvVideoSeekMode,omitempty"`
	// The custom MP4 end parameter.
	//
	// example:
	//
	// m_end
	Mp4SeekEnd *string `json:"Mp4SeekEnd,omitempty" xml:"Mp4SeekEnd,omitempty"`
	// The custom MP4 start parameter.
	//
	// example:
	//
	// m_start
	Mp4SeekStart *string `json:"Mp4SeekStart,omitempty" xml:"Mp4SeekStart,omitempty"`
	// The rule content.
	//
	// example:
	//
	// (http.host eq "video.example.com")
	Rule *string `json:"Rule,omitempty" xml:"Rule,omitempty"`
	// Specifies whether the rule is enabled. Valid values:
	//
	// example:
	//
	// on
	RuleEnable *string `json:"RuleEnable,omitempty" xml:"RuleEnable,omitempty"`
	// The rule name.
	//
	// example:
	//
	// rule_example
	RuleName *string `json:"RuleName,omitempty" xml:"RuleName,omitempty"`
	// The execution order of the rule.
	//
	// example:
	//
	// 1
	Sequence *string `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// Specifies whether to enable the audio seeking feature. Valid values:
	//
	// example:
	//
	// on
	VideoSeekEnable *string `json:"VideoSeekEnable,omitempty" xml:"VideoSeekEnable,omitempty"`
}

func (s ListSiteFunctionsResponseBodyConfigsVideoProcessing) String() string {
	return dara.Prettify(s)
}

func (s ListSiteFunctionsResponseBodyConfigsVideoProcessing) GoString() string {
	return s.String()
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) GetFlvSeekEnd() *string {
	return s.FlvSeekEnd
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) GetFlvSeekStart() *string {
	return s.FlvSeekStart
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) GetFlvVideoSeekMode() *string {
	return s.FlvVideoSeekMode
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) GetMp4SeekEnd() *string {
	return s.Mp4SeekEnd
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) GetMp4SeekStart() *string {
	return s.Mp4SeekStart
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) GetRule() *string {
	return s.Rule
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) GetRuleEnable() *string {
	return s.RuleEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) GetRuleName() *string {
	return s.RuleName
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) GetSequence() *string {
	return s.Sequence
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) GetVideoSeekEnable() *string {
	return s.VideoSeekEnable
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) SetConfigId(v int64) *ListSiteFunctionsResponseBodyConfigsVideoProcessing {
	s.ConfigId = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) SetFlvSeekEnd(v string) *ListSiteFunctionsResponseBodyConfigsVideoProcessing {
	s.FlvSeekEnd = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) SetFlvSeekStart(v string) *ListSiteFunctionsResponseBodyConfigsVideoProcessing {
	s.FlvSeekStart = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) SetFlvVideoSeekMode(v string) *ListSiteFunctionsResponseBodyConfigsVideoProcessing {
	s.FlvVideoSeekMode = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) SetMp4SeekEnd(v string) *ListSiteFunctionsResponseBodyConfigsVideoProcessing {
	s.Mp4SeekEnd = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) SetMp4SeekStart(v string) *ListSiteFunctionsResponseBodyConfigsVideoProcessing {
	s.Mp4SeekStart = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) SetRule(v string) *ListSiteFunctionsResponseBodyConfigsVideoProcessing {
	s.Rule = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) SetRuleEnable(v string) *ListSiteFunctionsResponseBodyConfigsVideoProcessing {
	s.RuleEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) SetRuleName(v string) *ListSiteFunctionsResponseBodyConfigsVideoProcessing {
	s.RuleName = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) SetSequence(v string) *ListSiteFunctionsResponseBodyConfigsVideoProcessing {
	s.Sequence = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) SetVideoSeekEnable(v string) *ListSiteFunctionsResponseBodyConfigsVideoProcessing {
	s.VideoSeekEnable = &v
	return s
}

func (s *ListSiteFunctionsResponseBodyConfigsVideoProcessing) Validate() error {
	return dara.Validate(s)
}
