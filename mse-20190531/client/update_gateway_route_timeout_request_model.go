// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGatewayRouteTimeoutRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *UpdateGatewayRouteTimeoutRequest
	GetAcceptLanguage() *string
	SetGatewayId(v int64) *UpdateGatewayRouteTimeoutRequest
	GetGatewayId() *int64
	SetGatewayUniqueId(v string) *UpdateGatewayRouteTimeoutRequest
	GetGatewayUniqueId() *string
	SetId(v int64) *UpdateGatewayRouteTimeoutRequest
	GetId() *int64
	SetTimeoutJSON(v *UpdateGatewayRouteTimeoutRequestTimeoutJSON) *UpdateGatewayRouteTimeoutRequest
	GetTimeoutJSON() *UpdateGatewayRouteTimeoutRequestTimeoutJSON
}

type UpdateGatewayRouteTimeoutRequest struct {
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
	// 85
	GatewayId *int64 `json:"GatewayId,omitempty" xml:"GatewayId,omitempty"`
	// The unique ID of the gateway.
	//
	// example:
	//
	// gw-533290d279c1405f9628c64f7c8272ee
	GatewayUniqueId *string `json:"GatewayUniqueId,omitempty" xml:"GatewayUniqueId,omitempty"`
	// The ID of the associated record.
	//
	// example:
	//
	// 567
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The timeout period.
	TimeoutJSON *UpdateGatewayRouteTimeoutRequestTimeoutJSON `json:"TimeoutJSON,omitempty" xml:"TimeoutJSON,omitempty" type:"Struct"`
}

func (s UpdateGatewayRouteTimeoutRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteTimeoutRequest) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteTimeoutRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *UpdateGatewayRouteTimeoutRequest) GetGatewayId() *int64 {
	return s.GatewayId
}

func (s *UpdateGatewayRouteTimeoutRequest) GetGatewayUniqueId() *string {
	return s.GatewayUniqueId
}

func (s *UpdateGatewayRouteTimeoutRequest) GetId() *int64 {
	return s.Id
}

func (s *UpdateGatewayRouteTimeoutRequest) GetTimeoutJSON() *UpdateGatewayRouteTimeoutRequestTimeoutJSON {
	return s.TimeoutJSON
}

func (s *UpdateGatewayRouteTimeoutRequest) SetAcceptLanguage(v string) *UpdateGatewayRouteTimeoutRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutRequest) SetGatewayId(v int64) *UpdateGatewayRouteTimeoutRequest {
	s.GatewayId = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutRequest) SetGatewayUniqueId(v string) *UpdateGatewayRouteTimeoutRequest {
	s.GatewayUniqueId = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutRequest) SetId(v int64) *UpdateGatewayRouteTimeoutRequest {
	s.Id = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutRequest) SetTimeoutJSON(v *UpdateGatewayRouteTimeoutRequestTimeoutJSON) *UpdateGatewayRouteTimeoutRequest {
	s.TimeoutJSON = v
	return s
}

func (s *UpdateGatewayRouteTimeoutRequest) Validate() error {
	if s.TimeoutJSON != nil {
		if err := s.TimeoutJSON.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateGatewayRouteTimeoutRequestTimeoutJSON struct {
	// The status of the policy.
	//
	// example:
	//
	// off
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The unit of time. A value of s indicates seconds.
	//
	// example:
	//
	// s
	TimeUnit *string `json:"TimeUnit,omitempty" xml:"TimeUnit,omitempty"`
	// The value of the timeout period.
	//
	// example:
	//
	// 1
	UnitNum *int32 `json:"UnitNum,omitempty" xml:"UnitNum,omitempty"`
}

func (s UpdateGatewayRouteTimeoutRequestTimeoutJSON) String() string {
	return dara.Prettify(s)
}

func (s UpdateGatewayRouteTimeoutRequestTimeoutJSON) GoString() string {
	return s.String()
}

func (s *UpdateGatewayRouteTimeoutRequestTimeoutJSON) GetStatus() *string {
	return s.Status
}

func (s *UpdateGatewayRouteTimeoutRequestTimeoutJSON) GetTimeUnit() *string {
	return s.TimeUnit
}

func (s *UpdateGatewayRouteTimeoutRequestTimeoutJSON) GetUnitNum() *int32 {
	return s.UnitNum
}

func (s *UpdateGatewayRouteTimeoutRequestTimeoutJSON) SetStatus(v string) *UpdateGatewayRouteTimeoutRequestTimeoutJSON {
	s.Status = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutRequestTimeoutJSON) SetTimeUnit(v string) *UpdateGatewayRouteTimeoutRequestTimeoutJSON {
	s.TimeUnit = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutRequestTimeoutJSON) SetUnitNum(v int32) *UpdateGatewayRouteTimeoutRequestTimeoutJSON {
	s.UnitNum = &v
	return s
}

func (s *UpdateGatewayRouteTimeoutRequestTimeoutJSON) Validate() error {
	return dara.Validate(s)
}
