// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetConsumerAuthorizationRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApiInfo(v *HttpApiApiInfo) *GetConsumerAuthorizationRuleResponseBody
	GetApiInfo() *HttpApiApiInfo
	SetConsumerAuthorizationRuleId(v string) *GetConsumerAuthorizationRuleResponseBody
	GetConsumerAuthorizationRuleId() *string
	SetConsumerId(v string) *GetConsumerAuthorizationRuleResponseBody
	GetConsumerId() *string
	SetCreateTimestamp(v int64) *GetConsumerAuthorizationRuleResponseBody
	GetCreateTimestamp() *int64
	SetDeployStatus(v string) *GetConsumerAuthorizationRuleResponseBody
	GetDeployStatus() *string
	SetEnvironmentInfo(v *EnvironmentInfo) *GetConsumerAuthorizationRuleResponseBody
	GetEnvironmentInfo() *EnvironmentInfo
	SetExpireMode(v string) *GetConsumerAuthorizationRuleResponseBody
	GetExpireMode() *string
	SetExpireStatus(v string) *GetConsumerAuthorizationRuleResponseBody
	GetExpireStatus() *string
	SetExpireTimestamp(v int64) *GetConsumerAuthorizationRuleResponseBody
	GetExpireTimestamp() *int64
	SetGatewayInfo(v *GatewayInfo) *GetConsumerAuthorizationRuleResponseBody
	GetGatewayInfo() *GatewayInfo
	SetRequestId(v string) *GetConsumerAuthorizationRuleResponseBody
	GetRequestId() *string
	SetResourceType(v string) *GetConsumerAuthorizationRuleResponseBody
	GetResourceType() *string
	SetUpdateTimestamp(v int64) *GetConsumerAuthorizationRuleResponseBody
	GetUpdateTimestamp() *int64
}

type GetConsumerAuthorizationRuleResponseBody struct {
	ApiInfo *HttpApiApiInfo `json:"apiInfo,omitempty" xml:"apiInfo,omitempty"`
	// example:
	//
	// car-ctgdn8em1hko5krqq03g
	ConsumerAuthorizationRuleId *string `json:"consumerAuthorizationRuleId,omitempty" xml:"consumerAuthorizationRuleId,omitempty"`
	// example:
	//
	// cs-ctgdn2um1hkossul8gvg
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// example:
	//
	// 1750852089975
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// ""
	DeployStatus    *string          `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	EnvironmentInfo *EnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty"`
	// example:
	//
	// LongTerm
	ExpireMode *string `json:"expireMode,omitempty" xml:"expireMode,omitempty"`
	// example:
	//
	// true
	ExpireStatus *string `json:"expireStatus,omitempty" xml:"expireStatus,omitempty"`
	// example:
	//
	// 1750852089975
	ExpireTimestamp *int64       `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
	GatewayInfo     *GatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 53102737-1E4E-5A8B-8E0A-4184B0959B84
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// API
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// 1750852089975
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s GetConsumerAuthorizationRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetConsumerAuthorizationRuleResponseBody) GoString() string {
	return s.String()
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetApiInfo() *HttpApiApiInfo {
	return s.ApiInfo
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetConsumerAuthorizationRuleId() *string {
	return s.ConsumerAuthorizationRuleId
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetEnvironmentInfo() *EnvironmentInfo {
	return s.EnvironmentInfo
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetExpireMode() *string {
	return s.ExpireMode
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetExpireStatus() *string {
	return s.ExpireStatus
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetGatewayInfo() *GatewayInfo {
	return s.GatewayInfo
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetResourceType() *string {
	return s.ResourceType
}

func (s *GetConsumerAuthorizationRuleResponseBody) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetApiInfo(v *HttpApiApiInfo) *GetConsumerAuthorizationRuleResponseBody {
	s.ApiInfo = v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetConsumerAuthorizationRuleId(v string) *GetConsumerAuthorizationRuleResponseBody {
	s.ConsumerAuthorizationRuleId = &v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetConsumerId(v string) *GetConsumerAuthorizationRuleResponseBody {
	s.ConsumerId = &v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetCreateTimestamp(v int64) *GetConsumerAuthorizationRuleResponseBody {
	s.CreateTimestamp = &v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetDeployStatus(v string) *GetConsumerAuthorizationRuleResponseBody {
	s.DeployStatus = &v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetEnvironmentInfo(v *EnvironmentInfo) *GetConsumerAuthorizationRuleResponseBody {
	s.EnvironmentInfo = v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetExpireMode(v string) *GetConsumerAuthorizationRuleResponseBody {
	s.ExpireMode = &v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetExpireStatus(v string) *GetConsumerAuthorizationRuleResponseBody {
	s.ExpireStatus = &v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetExpireTimestamp(v int64) *GetConsumerAuthorizationRuleResponseBody {
	s.ExpireTimestamp = &v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetGatewayInfo(v *GatewayInfo) *GetConsumerAuthorizationRuleResponseBody {
	s.GatewayInfo = v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetRequestId(v string) *GetConsumerAuthorizationRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetResourceType(v string) *GetConsumerAuthorizationRuleResponseBody {
	s.ResourceType = &v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) SetUpdateTimestamp(v int64) *GetConsumerAuthorizationRuleResponseBody {
	s.UpdateTimestamp = &v
	return s
}

func (s *GetConsumerAuthorizationRuleResponseBody) Validate() error {
	if s.ApiInfo != nil {
		if err := s.ApiInfo.Validate(); err != nil {
			return err
		}
	}
	if s.EnvironmentInfo != nil {
		if err := s.EnvironmentInfo.Validate(); err != nil {
			return err
		}
	}
	if s.GatewayInfo != nil {
		if err := s.GatewayInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}
