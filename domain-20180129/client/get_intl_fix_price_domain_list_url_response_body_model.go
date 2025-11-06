// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIntlFixPriceDomainListUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v *GetIntlFixPriceDomainListUrlResponseBodyModule) *GetIntlFixPriceDomainListUrlResponseBody
	GetModule() *GetIntlFixPriceDomainListUrlResponseBodyModule
	SetRequestId(v string) *GetIntlFixPriceDomainListUrlResponseBody
	GetRequestId() *string
}

type GetIntlFixPriceDomainListUrlResponseBody struct {
	// The returned data.
	Module *GetIntlFixPriceDomainListUrlResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// BF014B60-C708-4253-B5F2-3F9B493F398B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIntlFixPriceDomainListUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIntlFixPriceDomainListUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetIntlFixPriceDomainListUrlResponseBody) GetModule() *GetIntlFixPriceDomainListUrlResponseBodyModule {
	return s.Module
}

func (s *GetIntlFixPriceDomainListUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIntlFixPriceDomainListUrlResponseBody) SetModule(v *GetIntlFixPriceDomainListUrlResponseBodyModule) *GetIntlFixPriceDomainListUrlResponseBody {
	s.Module = v
	return s
}

func (s *GetIntlFixPriceDomainListUrlResponseBody) SetRequestId(v string) *GetIntlFixPriceDomainListUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIntlFixPriceDomainListUrlResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetIntlFixPriceDomainListUrlResponseBodyModule struct {
	// The URL for downloading the list that contains available fixed-price domain names at the international site (alibabacloud.com).
	//
	// example:
	//
	// http://intl-fixed-price.oss-cn-zhangjiakou.aliyuncs.com/aliyun_intl_fixed_price_domain_20240827.gz?Expires=1724830838&OSSAccessKeyId=LTAI5tPMAybR4gfSEjdf****&Signature=tb0SPs6tKb9gLKyQ5ibpQnBUuT****
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
}

func (s GetIntlFixPriceDomainListUrlResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s GetIntlFixPriceDomainListUrlResponseBodyModule) GoString() string {
	return s.String()
}

func (s *GetIntlFixPriceDomainListUrlResponseBodyModule) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetIntlFixPriceDomainListUrlResponseBodyModule) SetDownloadUrl(v string) *GetIntlFixPriceDomainListUrlResponseBodyModule {
	s.DownloadUrl = &v
	return s
}

func (s *GetIntlFixPriceDomainListUrlResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
