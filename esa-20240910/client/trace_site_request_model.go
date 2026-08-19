// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTraceSiteRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *TraceSiteRequestBody) *TraceSiteRequest
	GetBody() *TraceSiteRequestBody
	SetContext(v *TraceSiteRequestContext) *TraceSiteRequest
	GetContext() *TraceSiteRequestContext
	SetCookies(v []*TraceSiteRequestCookies) *TraceSiteRequest
	GetCookies() []*TraceSiteRequestCookies
	SetHeaders(v []*TraceSiteRequestHeaders) *TraceSiteRequest
	GetHeaders() []*TraceSiteRequestHeaders
	SetMethod(v string) *TraceSiteRequest
	GetMethod() *string
	SetProtocol(v string) *TraceSiteRequest
	GetProtocol() *string
	SetUrl(v string) *TraceSiteRequest
	GetUrl() *string
}

type TraceSiteRequest struct {
	// The HTTP request body.
	//
	// example:
	//
	// {"PlainText":"bc58c54211db"}
	Body *TraceSiteRequestBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	// The environment context. This parameter is optional.
	Context *TraceSiteRequestContext `json:"Context,omitempty" xml:"Context,omitempty" type:"Struct"`
	// The cookie parameters.
	//
	// example:
	//
	// []
	Cookies []*TraceSiteRequestCookies `json:"Cookies,omitempty" xml:"Cookies,omitempty" type:"Repeated"`
	// The request headers.
	//
	// example:
	//
	// []
	Headers []*TraceSiteRequestHeaders `json:"Headers,omitempty" xml:"Headers,omitempty" type:"Repeated"`
	// The HTTP method.
	//
	// example:
	//
	// GET
	Method *string `json:"Method,omitempty" xml:"Method,omitempty"`
	// The HTTP protocol.
	//
	// example:
	//
	// HTTP/1.1
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The URL of the request.
	//
	// This parameter is required.
	//
	// example:
	//
	// http://example.com/test
	Url *string `json:"Url,omitempty" xml:"Url,omitempty"`
}

func (s TraceSiteRequest) String() string {
	return dara.Prettify(s)
}

func (s TraceSiteRequest) GoString() string {
	return s.String()
}

func (s *TraceSiteRequest) GetBody() *TraceSiteRequestBody {
	return s.Body
}

func (s *TraceSiteRequest) GetContext() *TraceSiteRequestContext {
	return s.Context
}

func (s *TraceSiteRequest) GetCookies() []*TraceSiteRequestCookies {
	return s.Cookies
}

func (s *TraceSiteRequest) GetHeaders() []*TraceSiteRequestHeaders {
	return s.Headers
}

func (s *TraceSiteRequest) GetMethod() *string {
	return s.Method
}

func (s *TraceSiteRequest) GetProtocol() *string {
	return s.Protocol
}

func (s *TraceSiteRequest) GetUrl() *string {
	return s.Url
}

func (s *TraceSiteRequest) SetBody(v *TraceSiteRequestBody) *TraceSiteRequest {
	s.Body = v
	return s
}

func (s *TraceSiteRequest) SetContext(v *TraceSiteRequestContext) *TraceSiteRequest {
	s.Context = v
	return s
}

func (s *TraceSiteRequest) SetCookies(v []*TraceSiteRequestCookies) *TraceSiteRequest {
	s.Cookies = v
	return s
}

func (s *TraceSiteRequest) SetHeaders(v []*TraceSiteRequestHeaders) *TraceSiteRequest {
	s.Headers = v
	return s
}

func (s *TraceSiteRequest) SetMethod(v string) *TraceSiteRequest {
	s.Method = &v
	return s
}

func (s *TraceSiteRequest) SetProtocol(v string) *TraceSiteRequest {
	s.Protocol = &v
	return s
}

func (s *TraceSiteRequest) SetUrl(v string) *TraceSiteRequest {
	s.Url = &v
	return s
}

func (s *TraceSiteRequest) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.Context != nil {
		if err := s.Context.Validate(); err != nil {
			return err
		}
	}
	if s.Cookies != nil {
		for _, item := range s.Cookies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Headers != nil {
		for _, item := range s.Headers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type TraceSiteRequestBody struct {
	// The content in JSON format. If both JSON format content and plain text content are specified, the JSON format content takes precedence.
	//
	// example:
	//
	// {"request_id","3f809c32"}
	Json interface{} `json:"Json,omitempty" xml:"Json,omitempty"`
	// The plain text content.
	//
	// example:
	//
	// bc58c54211db
	PlainText *string `json:"PlainText,omitempty" xml:"PlainText,omitempty"`
}

func (s TraceSiteRequestBody) String() string {
	return dara.Prettify(s)
}

func (s TraceSiteRequestBody) GoString() string {
	return s.String()
}

func (s *TraceSiteRequestBody) GetJson() interface{} {
	return s.Json
}

func (s *TraceSiteRequestBody) GetPlainText() *string {
	return s.PlainText
}

func (s *TraceSiteRequestBody) SetJson(v interface{}) *TraceSiteRequestBody {
	s.Json = v
	return s
}

func (s *TraceSiteRequestBody) SetPlainText(v string) *TraceSiteRequestBody {
	s.PlainText = &v
	return s
}

func (s *TraceSiteRequestBody) Validate() error {
	return dara.Validate(s)
}

type TraceSiteRequestContext struct {
	// The simulated geolocation information.
	GeoLocation *TraceSiteRequestContextGeoLocation `json:"GeoLocation,omitempty" xml:"GeoLocation,omitempty" type:"Struct"`
	// Specifies whether to skip the security challenge test.
	//
	// example:
	//
	// true
	SkipChallenge *bool `json:"SkipChallenge,omitempty" xml:"SkipChallenge,omitempty"`
}

func (s TraceSiteRequestContext) String() string {
	return dara.Prettify(s)
}

func (s TraceSiteRequestContext) GoString() string {
	return s.String()
}

func (s *TraceSiteRequestContext) GetGeoLocation() *TraceSiteRequestContextGeoLocation {
	return s.GeoLocation
}

func (s *TraceSiteRequestContext) GetSkipChallenge() *bool {
	return s.SkipChallenge
}

func (s *TraceSiteRequestContext) SetGeoLocation(v *TraceSiteRequestContextGeoLocation) *TraceSiteRequestContext {
	s.GeoLocation = v
	return s
}

func (s *TraceSiteRequestContext) SetSkipChallenge(v bool) *TraceSiteRequestContext {
	s.SkipChallenge = &v
	return s
}

func (s *TraceSiteRequestContext) Validate() error {
	if s.GeoLocation != nil {
		if err := s.GeoLocation.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type TraceSiteRequestContextGeoLocation struct {
	// The country/region code.
	//
	// example:
	//
	// CN
	CountryCode *string `json:"CountryCode,omitempty" xml:"CountryCode,omitempty"`
	// The Internet service provider (ISP) code. This parameter is valid only when the country or region is the Chinese mainland.
	//
	// example:
	//
	// 100025
	IspCode *string `json:"IspCode,omitempty" xml:"IspCode,omitempty"`
	// The region or province code. This parameter is valid only when the country or region is the Chinese mainland.
	//
	// example:
	//
	// CN-BJ
	RegionCode *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
}

func (s TraceSiteRequestContextGeoLocation) String() string {
	return dara.Prettify(s)
}

func (s TraceSiteRequestContextGeoLocation) GoString() string {
	return s.String()
}

func (s *TraceSiteRequestContextGeoLocation) GetCountryCode() *string {
	return s.CountryCode
}

func (s *TraceSiteRequestContextGeoLocation) GetIspCode() *string {
	return s.IspCode
}

func (s *TraceSiteRequestContextGeoLocation) GetRegionCode() *string {
	return s.RegionCode
}

func (s *TraceSiteRequestContextGeoLocation) SetCountryCode(v string) *TraceSiteRequestContextGeoLocation {
	s.CountryCode = &v
	return s
}

func (s *TraceSiteRequestContextGeoLocation) SetIspCode(v string) *TraceSiteRequestContextGeoLocation {
	s.IspCode = &v
	return s
}

func (s *TraceSiteRequestContextGeoLocation) SetRegionCode(v string) *TraceSiteRequestContextGeoLocation {
	s.RegionCode = &v
	return s
}

func (s *TraceSiteRequestContextGeoLocation) Validate() error {
	return dara.Validate(s)
}

type TraceSiteRequestCookies struct {
	// The cookie name.
	//
	// example:
	//
	// sessionId
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The cookie value.
	//
	// example:
	//
	// f9ca1f7d-15bb-4c60-ad99-71b8e3e4985b
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TraceSiteRequestCookies) String() string {
	return dara.Prettify(s)
}

func (s TraceSiteRequestCookies) GoString() string {
	return s.String()
}

func (s *TraceSiteRequestCookies) GetName() *string {
	return s.Name
}

func (s *TraceSiteRequestCookies) GetValue() *string {
	return s.Value
}

func (s *TraceSiteRequestCookies) SetName(v string) *TraceSiteRequestCookies {
	s.Name = &v
	return s
}

func (s *TraceSiteRequestCookies) SetValue(v string) *TraceSiteRequestCookies {
	s.Value = &v
	return s
}

func (s *TraceSiteRequestCookies) Validate() error {
	return dara.Validate(s)
}

type TraceSiteRequestHeaders struct {
	// The HTTP request header name.
	//
	// example:
	//
	// User-Agent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The HTTP request header value.
	//
	// example:
	//
	// trace-test
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TraceSiteRequestHeaders) String() string {
	return dara.Prettify(s)
}

func (s TraceSiteRequestHeaders) GoString() string {
	return s.String()
}

func (s *TraceSiteRequestHeaders) GetName() *string {
	return s.Name
}

func (s *TraceSiteRequestHeaders) GetValue() *string {
	return s.Value
}

func (s *TraceSiteRequestHeaders) SetName(v string) *TraceSiteRequestHeaders {
	s.Name = &v
	return s
}

func (s *TraceSiteRequestHeaders) SetValue(v string) *TraceSiteRequestHeaders {
	s.Value = &v
	return s
}

func (s *TraceSiteRequestHeaders) Validate() error {
	return dara.Validate(s)
}
