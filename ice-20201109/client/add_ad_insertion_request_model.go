// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAdInsertionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdMarkerPassthrough(v string) *AddAdInsertionRequest
	GetAdMarkerPassthrough() *string
	SetAdsUrl(v string) *AddAdInsertionRequest
	GetAdsUrl() *string
	SetCdnAdSegmentUrlPrefix(v string) *AddAdInsertionRequest
	GetCdnAdSegmentUrlPrefix() *string
	SetCdnContentSegmentUrlPrefix(v string) *AddAdInsertionRequest
	GetCdnContentSegmentUrlPrefix() *string
	SetClientToken(v string) *AddAdInsertionRequest
	GetClientToken() *string
	SetConfigAliases(v string) *AddAdInsertionRequest
	GetConfigAliases() *string
	SetContentUrlPrefix(v string) *AddAdInsertionRequest
	GetContentUrlPrefix() *string
	SetName(v string) *AddAdInsertionRequest
	GetName() *string
	SetPersonalizationThreshold(v int32) *AddAdInsertionRequest
	GetPersonalizationThreshold() *int32
	SetSlateAdUrl(v string) *AddAdInsertionRequest
	GetSlateAdUrl() *string
}

type AddAdInsertionRequest struct {
	// Specifies whether to enable ad marker passthrough. Default value: OFF.
	//
	// Valid values:
	//
	// 	- OFF: Disable.
	//
	// 	- ON: Enable.
	//
	// example:
	//
	// ON
	AdMarkerPassthrough *string `json:"AdMarkerPassthrough,omitempty" xml:"AdMarkerPassthrough,omitempty"`
	// The request URL of the ad decision server (ADS). HTTP and HTTPS are supported. The maximum length is 2,048 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://ads.com/ad1?param1=[palyer_params.p1]
	AdsUrl *string `json:"AdsUrl,omitempty" xml:"AdsUrl,omitempty"`
	// The CDN prefix for ad segments. HTTP and HTTPS are supported. The maximum length is 512 characters.
	//
	// example:
	//
	// http://cdn.com/
	CdnAdSegmentUrlPrefix *string `json:"CdnAdSegmentUrlPrefix,omitempty" xml:"CdnAdSegmentUrlPrefix,omitempty"`
	// The CDN prefix for content segments. HTTP and HTTPS are supported. The maximum length is 512 characters.
	//
	// example:
	//
	// http://cdn.com/
	CdnContentSegmentUrlPrefix *string `json:"CdnContentSegmentUrlPrefix,omitempty" xml:"CdnContentSegmentUrlPrefix,omitempty"`
	// The idempotency key that is used to avoid repeated submission. The value can be up to 200 characters in length.
	//
	// example:
	//
	// ****0311a423d11a5f7dee713535****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// A JSON string that specifies the player parameter variables and aliases. You can add up to 20 player_params.{name} entries. The name field can be up to 150 characters in length. Each player parameter can include up to 50 key-value pairs. A key can be up to 150 characters long, and a value can be up to 500 characters. Example: { "player_params.{name}": { "{key}": "{value}" } }
	//
	// example:
	//
	// { "player_params.p1": { "1": "abc" } }
	ConfigAliases *string `json:"ConfigAliases,omitempty" xml:"ConfigAliases,omitempty"`
	// The URL prefix for the source content. HTTP and HTTPS are supported. The maximum length is 512 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://source.com/
	ContentUrlPrefix *string `json:"ContentUrlPrefix,omitempty" xml:"ContentUrlPrefix,omitempty"`
	// The name of the configuration. The name must be unique and can be up to 128 characters in length. Letters, digits, underscores (_), and hyphens (-) are supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// my_ad
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Specifies the maximum duration of underfilled time allowed in an ad break. Unit: seconds. Default value: 8 seconds.
	//
	// example:
	//
	// 5
	PersonalizationThreshold *int32 `json:"PersonalizationThreshold,omitempty" xml:"PersonalizationThreshold,omitempty"`
	// The HTTP or HTTPS URL of the slate ad. Only MP4 format is supported. The maximum length is 2,048 characters.
	//
	// example:
	//
	// http://storage.com/slate1.mp4
	SlateAdUrl *string `json:"SlateAdUrl,omitempty" xml:"SlateAdUrl,omitempty"`
}

func (s AddAdInsertionRequest) String() string {
	return dara.Prettify(s)
}

func (s AddAdInsertionRequest) GoString() string {
	return s.String()
}

func (s *AddAdInsertionRequest) GetAdMarkerPassthrough() *string {
	return s.AdMarkerPassthrough
}

func (s *AddAdInsertionRequest) GetAdsUrl() *string {
	return s.AdsUrl
}

func (s *AddAdInsertionRequest) GetCdnAdSegmentUrlPrefix() *string {
	return s.CdnAdSegmentUrlPrefix
}

func (s *AddAdInsertionRequest) GetCdnContentSegmentUrlPrefix() *string {
	return s.CdnContentSegmentUrlPrefix
}

func (s *AddAdInsertionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AddAdInsertionRequest) GetConfigAliases() *string {
	return s.ConfigAliases
}

func (s *AddAdInsertionRequest) GetContentUrlPrefix() *string {
	return s.ContentUrlPrefix
}

func (s *AddAdInsertionRequest) GetName() *string {
	return s.Name
}

func (s *AddAdInsertionRequest) GetPersonalizationThreshold() *int32 {
	return s.PersonalizationThreshold
}

func (s *AddAdInsertionRequest) GetSlateAdUrl() *string {
	return s.SlateAdUrl
}

func (s *AddAdInsertionRequest) SetAdMarkerPassthrough(v string) *AddAdInsertionRequest {
	s.AdMarkerPassthrough = &v
	return s
}

func (s *AddAdInsertionRequest) SetAdsUrl(v string) *AddAdInsertionRequest {
	s.AdsUrl = &v
	return s
}

func (s *AddAdInsertionRequest) SetCdnAdSegmentUrlPrefix(v string) *AddAdInsertionRequest {
	s.CdnAdSegmentUrlPrefix = &v
	return s
}

func (s *AddAdInsertionRequest) SetCdnContentSegmentUrlPrefix(v string) *AddAdInsertionRequest {
	s.CdnContentSegmentUrlPrefix = &v
	return s
}

func (s *AddAdInsertionRequest) SetClientToken(v string) *AddAdInsertionRequest {
	s.ClientToken = &v
	return s
}

func (s *AddAdInsertionRequest) SetConfigAliases(v string) *AddAdInsertionRequest {
	s.ConfigAliases = &v
	return s
}

func (s *AddAdInsertionRequest) SetContentUrlPrefix(v string) *AddAdInsertionRequest {
	s.ContentUrlPrefix = &v
	return s
}

func (s *AddAdInsertionRequest) SetName(v string) *AddAdInsertionRequest {
	s.Name = &v
	return s
}

func (s *AddAdInsertionRequest) SetPersonalizationThreshold(v int32) *AddAdInsertionRequest {
	s.PersonalizationThreshold = &v
	return s
}

func (s *AddAdInsertionRequest) SetSlateAdUrl(v string) *AddAdInsertionRequest {
	s.SlateAdUrl = &v
	return s
}

func (s *AddAdInsertionRequest) Validate() error {
	return dara.Validate(s)
}
