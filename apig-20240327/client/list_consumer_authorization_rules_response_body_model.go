// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConsumerAuthorizationRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListConsumerAuthorizationRulesResponseBody
	GetCode() *string
	SetData(v *ListConsumerAuthorizationRulesResponseBodyData) *ListConsumerAuthorizationRulesResponseBody
	GetData() *ListConsumerAuthorizationRulesResponseBodyData
	SetMessage(v string) *ListConsumerAuthorizationRulesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListConsumerAuthorizationRulesResponseBody
	GetRequestId() *string
}

type ListConsumerAuthorizationRulesResponseBody struct {
	// example:
	//
	// Ok
	Code *string                                         `json:"code,omitempty" xml:"code,omitempty"`
	Data *ListConsumerAuthorizationRulesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 294382D9-EE60-5735-A4CD-F2AC2840423D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListConsumerAuthorizationRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerAuthorizationRulesResponseBody) GoString() string {
	return s.String()
}

func (s *ListConsumerAuthorizationRulesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListConsumerAuthorizationRulesResponseBody) GetData() *ListConsumerAuthorizationRulesResponseBodyData {
	return s.Data
}

func (s *ListConsumerAuthorizationRulesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListConsumerAuthorizationRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConsumerAuthorizationRulesResponseBody) SetCode(v string) *ListConsumerAuthorizationRulesResponseBody {
	s.Code = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBody) SetData(v *ListConsumerAuthorizationRulesResponseBodyData) *ListConsumerAuthorizationRulesResponseBody {
	s.Data = v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBody) SetMessage(v string) *ListConsumerAuthorizationRulesResponseBody {
	s.Message = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBody) SetRequestId(v string) *ListConsumerAuthorizationRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConsumerAuthorizationRulesResponseBodyData struct {
	Items []*ListConsumerAuthorizationRulesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"pageNumber,omitempty" xml:"pageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 9
	TotalSize *string `json:"totalSize,omitempty" xml:"totalSize,omitempty"`
}

func (s ListConsumerAuthorizationRulesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerAuthorizationRulesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListConsumerAuthorizationRulesResponseBodyData) GetItems() []*ListConsumerAuthorizationRulesResponseBodyDataItems {
	return s.Items
}

func (s *ListConsumerAuthorizationRulesResponseBodyData) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListConsumerAuthorizationRulesResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListConsumerAuthorizationRulesResponseBodyData) GetTotalSize() *string {
	return s.TotalSize
}

func (s *ListConsumerAuthorizationRulesResponseBodyData) SetItems(v []*ListConsumerAuthorizationRulesResponseBodyDataItems) *ListConsumerAuthorizationRulesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyData) SetPageNumber(v int32) *ListConsumerAuthorizationRulesResponseBodyData {
	s.PageNumber = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyData) SetPageSize(v int32) *ListConsumerAuthorizationRulesResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyData) SetTotalSize(v string) *ListConsumerAuthorizationRulesResponseBodyData {
	s.TotalSize = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConsumerAuthorizationRulesResponseBodyDataItems struct {
	ApiInfo *HttpApiApiInfo `json:"apiInfo,omitempty" xml:"apiInfo,omitempty"`
	// example:
	//
	// car-csgeka5lhtggrjcprok0
	ConsumerAuthorizationRuleId *string `json:"consumerAuthorizationRuleId,omitempty" xml:"consumerAuthorizationRuleId,omitempty"`
	// example:
	//
	// cs-csheiftlhtgmp0j0hp4g
	ConsumerId *string `json:"consumerId,omitempty" xml:"consumerId,omitempty"`
	// example:
	//
	// 1719386834548
	CreateTimestamp *int64 `json:"createTimestamp,omitempty" xml:"createTimestamp,omitempty"`
	// example:
	//
	// {}
	DeployStatus    *string          `json:"deployStatus,omitempty" xml:"deployStatus,omitempty"`
	EnvironmentInfo *EnvironmentInfo `json:"environmentInfo,omitempty" xml:"environmentInfo,omitempty"`
	// example:
	//
	// ShortTerm
	ExpireMode *string `json:"expireMode,omitempty" xml:"expireMode,omitempty"`
	// example:
	//
	// InEffect
	ExpireStatus *string `json:"expireStatus,omitempty" xml:"expireStatus,omitempty"`
	// example:
	//
	// 172086834548
	ExpireTimestamp *int64       `json:"expireTimestamp,omitempty" xml:"expireTimestamp,omitempty"`
	GatewayInfo     *GatewayInfo `json:"gatewayInfo,omitempty" xml:"gatewayInfo,omitempty"`
	// example:
	//
	// 2351944
	ResourceId *string `json:"resourceId,omitempty" xml:"resourceId,omitempty"`
	// example:
	//
	// HttpApiRoute
	ResourceType *string `json:"resourceType,omitempty" xml:"resourceType,omitempty"`
	// example:
	//
	// 1721116090326
	UpdateTimestamp *int64 `json:"updateTimestamp,omitempty" xml:"updateTimestamp,omitempty"`
}

func (s ListConsumerAuthorizationRulesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListConsumerAuthorizationRulesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetApiInfo() *HttpApiApiInfo {
	return s.ApiInfo
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetConsumerAuthorizationRuleId() *string {
	return s.ConsumerAuthorizationRuleId
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetConsumerId() *string {
	return s.ConsumerId
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetCreateTimestamp() *int64 {
	return s.CreateTimestamp
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetDeployStatus() *string {
	return s.DeployStatus
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetEnvironmentInfo() *EnvironmentInfo {
	return s.EnvironmentInfo
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetExpireMode() *string {
	return s.ExpireMode
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetExpireStatus() *string {
	return s.ExpireStatus
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetExpireTimestamp() *int64 {
	return s.ExpireTimestamp
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetGatewayInfo() *GatewayInfo {
	return s.GatewayInfo
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) GetUpdateTimestamp() *int64 {
	return s.UpdateTimestamp
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetApiInfo(v *HttpApiApiInfo) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.ApiInfo = v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetConsumerAuthorizationRuleId(v string) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.ConsumerAuthorizationRuleId = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetConsumerId(v string) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.ConsumerId = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetCreateTimestamp(v int64) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.CreateTimestamp = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetDeployStatus(v string) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.DeployStatus = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetEnvironmentInfo(v *EnvironmentInfo) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.EnvironmentInfo = v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetExpireMode(v string) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.ExpireMode = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetExpireStatus(v string) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.ExpireStatus = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetExpireTimestamp(v int64) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.ExpireTimestamp = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetGatewayInfo(v *GatewayInfo) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.GatewayInfo = v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetResourceId(v string) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.ResourceId = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetResourceType(v string) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.ResourceType = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) SetUpdateTimestamp(v int64) *ListConsumerAuthorizationRulesResponseBodyDataItems {
	s.UpdateTimestamp = &v
	return s
}

func (s *ListConsumerAuthorizationRulesResponseBodyDataItems) Validate() error {
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
