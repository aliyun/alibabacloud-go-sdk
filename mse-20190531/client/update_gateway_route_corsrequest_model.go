// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteCORSRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteCORSRequest
	GetAcceptLanguage() *string
	SetCorsJSON(v *UpdateGatewayRouteCORSRequestCorsJSON) *UpdateGatewayRouteCORSRequest
	GetCorsJSON() *UpdateGatewayRouteCORSRequestCorsJSON
	SetGatewayId(v int64) *UpdateGatewayRouteCORSRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayRouteCORSRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayRouteCORSRequest
	GetId() *int64
}

type UpdateGatewayRouteCORSRequest struct {
	// The language of the response. In compliance with [RFC 7231](https://tools.ietf.org/html/rfc7231), the backend service must return a response based on the language used by the user.
	//
	// 	- No default value.
	//
	// 	- zh-CN: Chinese
	//
	// 	- en-US: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The information about the CORS policy.
	CorsJSON *UpdateGatewayRouteCORSRequestCorsJSON `json:"CorsJSON,omitempty" xml:"CorsJSON,omitempty" type:"Struct"`
	// The ID of the gateway.
	//
	// example:
	//
	// 85
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-f70a6ddf2f0941a2bb997b2d16028f37
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the associated record.
	//
	// example:
	//
	// 55
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s UpdateGatewayRouteCORSRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteCORSRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteCORSRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteCORSRequest) GetCorsJSON() *UpdateGatewayRouteCORSRequestCorsJSON {
	return s.CorsJSON
}

func (s *UpdateGatewayRouteCORSRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteCORSRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteCORSRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteCORSRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteCORSRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteCORSRequest) SetCorsJSON(v *UpdateGatewayRouteCORSRequestCorsJSON) *UpdateGatewayRouteCORSRequest {
	s.CorsJSON = v
	return s
}

func (s *UpdateGatewayRouteCORSRequest) SetGatewayId(v int64) *UpdateGatewayRouteCORSRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteCORSRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteCORSRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteCORSRequest) SetId(v int64) *UpdateGatewayRouteCORSRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteCORSRequest) Validate() error {
	if s.CorsJSON != nil {
		if err := s.CorsJSON.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGatewayRouteCORSRequestCorsJSON struct {
	// The credentials allowed.
	//
	// example:
	//
	// true
	AllowCredentials *bool `json:"AllowCredentials,omitempty" xml:"AllowCredentials,omitempty"`
	// The request headers allowed.
	//
	// example:
	//
	// content-type
	AllowHeaders *string `json:"AllowHeaders,omitempty" xml:"AllowHeaders,omitempty"`
	// The HTTP methods allowed.
	//
	// example:
	//
	// GET,POST
	AllowMethods *string `json:"AllowMethods,omitempty" xml:"AllowMethods,omitempty"`
	// The origins from which access is allowed.
	//
	// example:
	//
	// https://api.aliyun-inc.com/
	AllowOrigins *string `json:"AllowOrigins,omitempty" xml:"AllowOrigins,omitempty"`
	// The response headers allowed.
	//
	// example:
	//
	// *
	ExposeHeaders *string `json:"ExposeHeaders,omitempty" xml:"ExposeHeaders,omitempty"`
	// The status of the policy.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The unit of time.
	//
	// example:
	//
	// h
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	// The value of time.
	//
	// example:
	//
	// 24
	UnitNum *int64 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s UpdateGatewayRouteCORSRequestCorsJSON) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteCORSRequestCorsJSON) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) GetAllowCredentials() *bool {
	return s.AllowCredentials
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) GetAllowHeaders() *string {
	return s.AllowHeaders
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) GetAllowMethods() *string {
	return s.AllowMethods
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) GetAllowOrigins() *string {
	return s.AllowOrigins
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) GetExposeHeaders() *string {
	return s.ExposeHeaders
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) GetStatus() *string {
	return s.Status
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) GetUnitNum() *int64 {
	return s.UnitNum
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) SetAllowCredentials(v bool) *UpdateGatewayRouteCORSRequestCorsJSON {
	s.AllowCredentials = &v
	return s
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) SetAllowHeaders(v string) *UpdateGatewayRouteCORSRequestCorsJSON {
	s.AllowHeaders = &v
	return s
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) SetAllowMethods(v string) *UpdateGatewayRouteCORSRequestCorsJSON {
	s.AllowMethods = &v
	return s
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) SetAllowOrigins(v string) *UpdateGatewayRouteCORSRequestCorsJSON {
	s.AllowOrigins = &v
	return s
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) SetExposeHeaders(v string) *UpdateGatewayRouteCORSRequestCorsJSON {
	s.ExposeHeaders = &v
	return s
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) SetStatus(v string) *UpdateGatewayRouteCORSRequestCorsJSON {
	s.Status = &v
	return s
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) SetTimeUnit(v string) *UpdateGatewayRouteCORSRequestCorsJSON {
	s.TimeUnit = &v
	return s
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) SetUnitNum(v int64) *UpdateGatewayRouteCORSRequestCorsJSON {
	s.UnitNum = &v
	return s
}

func (s *UpdateGatewayRouteCORSRequestCorsJSON) Validate() error {
	return dara.Validate(s)
}
