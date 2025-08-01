// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteRetryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteRetryRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *UpdateGatewayRouteRetryRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayRouteRetryRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayRouteRetryRequest
	GetId() *int64
	SetRetryJSON(v *UpdateGatewayRouteRetryRequestRetryJSON) *UpdateGatewayRouteRetryRequest
	GetRetryJSON() *UpdateGatewayRouteRetryRequestRetryJSON
}

type UpdateGatewayRouteRetryRequest struct {
	// The language of the response. Valid values:
	//
	// 	- zh: Chinese
	//
	// 	- en: English
	//
	// example:
	//
	// zh
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
	// The ID of the gateway.
	//
	// example:
	//
	// 501
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-3f97e2989c344f35ab3fd62b19f1d10a
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the associated record.
	//
	// example:
	//
	// 508
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The information about the retry policy.
	RetryJSON *UpdateGatewayRouteRetryRequestRetryJSON `json:"RetryJSON,omitempty" xml:"RetryJSON,omitempty" type:"Struct"`
}

func (s UpdateGatewayRouteRetryRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRetryRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRetryRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteRetryRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteRetryRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteRetryRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteRetryRequest) GetRetryJSON() *UpdateGatewayRouteRetryRequestRetryJSON {
	return s.RetryJSON
}

func (s *UpdateGatewayRouteRetryRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteRetryRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteRetryRequest) SetGatewayId(v int64) *UpdateGatewayRouteRetryRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteRetryRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteRetryRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteRetryRequest) SetId(v int64) *UpdateGatewayRouteRetryRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteRetryRequest) SetRetryJSON(v *UpdateGatewayRouteRetryRequestRetryJSON) *UpdateGatewayRouteRetryRequest {
	s.RetryJSON = v
	return s
}

func (s *UpdateGatewayRouteRetryRequest) Validate() error {
	return dara.Validate(s)
}

type UpdateGatewayRouteRetryRequestRetryJSON struct {
	// The number of retries.
	//
	// example:
	//
	// 2
	Attempts *int32 `json:"Attempts,omitempty" xml:"Attempts,omitempty"`
	// The HTTP status codes.
	HttpCodes []*string `json:"HttpCodes,omitempty" xml:"HttpCodes,omitempty" type:"Repeated"`
	// The retry conditions.
	RetryOn []*string `json:"RetryOn,omitempty" xml:"RetryOn,omitempty" type:"Repeated"`
	// The status of the policy.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateGatewayRouteRetryRequestRetryJSON) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteRetryRequestRetryJSON) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteRetryRequestRetryJSON) GetAttempts() *int32 {
	return s.Attempts
}

func (s *UpdateGatewayRouteRetryRequestRetryJSON) GetHttpCodes() []*string {
	return s.HttpCodes
}

func (s *UpdateGatewayRouteRetryRequestRetryJSON) GetRetryOn() []*string {
	return s.RetryOn
}

func (s *UpdateGatewayRouteRetryRequestRetryJSON) GetStatus() *string {
	return s.Status
}

func (s *UpdateGatewayRouteRetryRequestRetryJSON) SetAttempts(v int32) *UpdateGatewayRouteRetryRequestRetryJSON {
	s.Attempts = &v
	return s
}

func (s *UpdateGatewayRouteRetryRequestRetryJSON) SetHttpCodes(v []*string) *UpdateGatewayRouteRetryRequestRetryJSON {
	s.HttpCodes = v
	return s
}

func (s *UpdateGatewayRouteRetryRequestRetryJSON) SetRetryOn(v []*string) *UpdateGatewayRouteRetryRequestRetryJSON {
	s.RetryOn = v
	return s
}

func (s *UpdateGatewayRouteRetryRequestRetryJSON) SetStatus(v string) *UpdateGatewayRouteRetryRequestRetryJSON {
	s.Status = &v
	return s
}

func (s *UpdateGatewayRouteRetryRequestRetryJSON) Validate() error {
	return dara.Validate(s)
}
