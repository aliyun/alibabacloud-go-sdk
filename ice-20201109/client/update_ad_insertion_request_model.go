// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdInsertionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdMarkerPassthrough(v string) *UpdateAdInsertionRequest
	GetAdMarkerPassthrough() *string
	SetAdsUrl(v string) *UpdateAdInsertionRequest
	GetAdsUrl() *string
	SetCdnAdSegmentUrlPrefix(v string) *UpdateAdInsertionRequest
	GetCdnAdSegmentUrlPrefix() *string
	SetCdnContentSegmentUrlPrefix(v string) *UpdateAdInsertionRequest
	GetCdnContentSegmentUrlPrefix() *string
	SetConfigAliases(v string) *UpdateAdInsertionRequest
	GetConfigAliases() *string
	SetContentUrlPrefix(v string) *UpdateAdInsertionRequest
	GetContentUrlPrefix() *string
	SetName(v string) *UpdateAdInsertionRequest
	GetName() *string
	SetPersonalizationThreshold(v int32) *UpdateAdInsertionRequest
	GetPersonalizationThreshold() *int32
	SetSlateAdUrl(v string) *UpdateAdInsertionRequest
	GetSlateAdUrl() *string
}

type UpdateAdInsertionRequest struct {
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
	// A JSON string that specifies the player parameter variables and aliases. Format: { "player_params.{name}": { "{key}": "{value}" } }. You can add up to 20 player_params.{name} entries. The name field can be up to 150 characters in length. Each player parameter can include up to 50 key-value pairs. A key can be up to 150 characters long, and a value can be up to 500 characters.
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
	// The configuration name, which cannot be modified.
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

func (s UpdateAdInsertionRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdInsertionRequest) GoString() string {
	return s.String()
}

func (s *UpdateAdInsertionRequest) GetAdMarkerPassthrough() *string {
	return s.AdMarkerPassthrough
}

func (s *UpdateAdInsertionRequest) GetAdsUrl() *string {
	return s.AdsUrl
}

func (s *UpdateAdInsertionRequest) GetCdnAdSegmentUrlPrefix() *string {
	return s.CdnAdSegmentUrlPrefix
}

func (s *UpdateAdInsertionRequest) GetCdnContentSegmentUrlPrefix() *string {
	return s.CdnContentSegmentUrlPrefix
}

func (s *UpdateAdInsertionRequest) GetConfigAliases() *string {
	return s.ConfigAliases
}

func (s *UpdateAdInsertionRequest) GetContentUrlPrefix() *string {
	return s.ContentUrlPrefix
}

func (s *UpdateAdInsertionRequest) GetName() *string {
	return s.Name
}

func (s *UpdateAdInsertionRequest) GetPersonalizationThreshold() *int32 {
	return s.PersonalizationThreshold
}

func (s *UpdateAdInsertionRequest) GetSlateAdUrl() *string {
	return s.SlateAdUrl
}

func (s *UpdateAdInsertionRequest) SetAdMarkerPassthrough(v string) *UpdateAdInsertionRequest {
	s.AdMarkerPassthrough = &v
	return s
}

func (s *UpdateAdInsertionRequest) SetAdsUrl(v string) *UpdateAdInsertionRequest {
	s.AdsUrl = &v
	return s
}

func (s *UpdateAdInsertionRequest) SetCdnAdSegmentUrlPrefix(v string) *UpdateAdInsertionRequest {
	s.CdnAdSegmentUrlPrefix = &v
	return s
}

func (s *UpdateAdInsertionRequest) SetCdnContentSegmentUrlPrefix(v string) *UpdateAdInsertionRequest {
	s.CdnContentSegmentUrlPrefix = &v
	return s
}

func (s *UpdateAdInsertionRequest) SetConfigAliases(v string) *UpdateAdInsertionRequest {
	s.ConfigAliases = &v
	return s
}

func (s *UpdateAdInsertionRequest) SetContentUrlPrefix(v string) *UpdateAdInsertionRequest {
	s.ContentUrlPrefix = &v
	return s
}

func (s *UpdateAdInsertionRequest) SetName(v string) *UpdateAdInsertionRequest {
	s.Name = &v
	return s
}

func (s *UpdateAdInsertionRequest) SetPersonalizationThreshold(v int32) *UpdateAdInsertionRequest {
	s.PersonalizationThreshold = &v
	return s
}

func (s *UpdateAdInsertionRequest) SetSlateAdUrl(v string) *UpdateAdInsertionRequest {
	s.SlateAdUrl = &v
	return s
}

func (s *UpdateAdInsertionRequest) Validate() error {
	return dara.Validate(s)
}
