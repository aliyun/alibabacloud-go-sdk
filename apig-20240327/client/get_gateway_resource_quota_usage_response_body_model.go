// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetGatewayResourceQuotaUsageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetGatewayResourceQuotaUsageResponseBody
	GetCode() *string
	SetData(v *GetGatewayResourceQuotaUsageResponseBodyData) *GetGatewayResourceQuotaUsageResponseBody
	GetData() *GetGatewayResourceQuotaUsageResponseBodyData
	SetMessage(v string) *GetGatewayResourceQuotaUsageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetGatewayResourceQuotaUsageResponseBody
	GetRequestId() *string
}

type GetGatewayResourceQuotaUsageResponseBody struct {
	// The response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The quota query result.
	//
	// example:
	//
	// {"gatewayId":"gw-d5be0s9q5z6f1234567g","observedAt":"2026-09-17T08:00:00Z","items":[{"quotaKey":"Route","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":12,"limit":800},{"quotaKey":"ConsumerAuthorizationRule","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":20,"limit":1000},{"quotaKey":"McpServer","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":false},{"quotaKey":"Domain","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":2,"limit":100},{"quotaKey":"Service","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":5,"limit":800},{"quotaKey":"ConsumerQuotaRule","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":0,"limit":1000},{"quotaKey":"K8sServiceSource","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":1,"limit":3},{"quotaKey":"InstalledPlugin","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":2,"limit":5},{"quotaKey":"CustomPlugin","usedScope":"ACCOUNT_REGION","limitScope":"ACCOUNT_GATEWAY_TYPE","applicable":false}]}
	Data *GetGatewayResourceQuotaUsageResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// The specified gateway does not exist.
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D0A6A1A0-8793-4C10-AB1A-2C03E770A912
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetGatewayResourceQuotaUsageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResourceQuotaUsageResponseBody) GoString() string {
	return s.String()
}

func (s *GetGatewayResourceQuotaUsageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetGatewayResourceQuotaUsageResponseBody) GetData() *GetGatewayResourceQuotaUsageResponseBodyData {
	return s.Data
}

func (s *GetGatewayResourceQuotaUsageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetGatewayResourceQuotaUsageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetGatewayResourceQuotaUsageResponseBody) SetCode(v string) *GetGatewayResourceQuotaUsageResponseBody {
	s.Code = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBody) SetData(v *GetGatewayResourceQuotaUsageResponseBodyData) *GetGatewayResourceQuotaUsageResponseBody {
	s.Data = v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBody) SetMessage(v string) *GetGatewayResourceQuotaUsageResponseBody {
	s.Message = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBody) SetRequestId(v string) *GetGatewayResourceQuotaUsageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetGatewayResourceQuotaUsageResponseBodyData struct {
	// The gateway ID.
	//
	// example:
	//
	// gw-d5be0s9q5z6f1234567g
	GatewayId *string `json:"gatewayId,omitempty" xml:"gatewayId,omitempty"`
	// The resource quota items.
	//
	// example:
	//
	// [{"quotaKey":"Route","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":12,"limit":800},{"quotaKey":"ConsumerAuthorizationRule","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":20,"limit":1000},{"quotaKey":"McpServer","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":false},{"quotaKey":"Domain","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":2,"limit":100},{"quotaKey":"Service","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":5,"limit":800},{"quotaKey":"ConsumerQuotaRule","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":0,"limit":1000},{"quotaKey":"K8sServiceSource","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":1,"limit":3},{"quotaKey":"InstalledPlugin","usedScope":"GATEWAY","limitScope":"GATEWAY","applicable":true,"used":2,"limit":5},{"quotaKey":"CustomPlugin","usedScope":"ACCOUNT_REGION","limitScope":"ACCOUNT_GATEWAY_TYPE","applicable":false}]
	Items []*GetGatewayResourceQuotaUsageResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The time when the statistics collection was completed.
	//
	// example:
	//
	// 2026-09-17T08:00:00Z
	ObservedAt *string `json:"observedAt,omitempty" xml:"observedAt,omitempty"`
}

func (s GetGatewayResourceQuotaUsageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResourceQuotaUsageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) GetGatewayId() *string {
	return s.GatewayId
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) GetItems() []*GetGatewayResourceQuotaUsageResponseBodyDataItems {
	return s.Items
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) GetObservedAt() *string {
	return s.ObservedAt
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) SetGatewayId(v string) *GetGatewayResourceQuotaUsageResponseBodyData {
	s.GatewayId = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) SetItems(v []*GetGatewayResourceQuotaUsageResponseBodyDataItems) *GetGatewayResourceQuotaUsageResponseBodyData {
	s.Items = v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) SetObservedAt(v string) *GetGatewayResourceQuotaUsageResponseBodyData {
	s.ObservedAt = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyData) Validate() error {
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

type GetGatewayResourceQuotaUsageResponseBodyDataItems struct {
	// Indicates whether quota numerical values are provided.
	//
	// example:
	//
	// true
	Applicable *bool `json:"applicable,omitempty" xml:"applicable,omitempty"`
	// The effective limit.
	//
	// example:
	//
	// 800
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The scope to which the limit applies.
	//
	// example:
	//
	// GATEWAY
	LimitScope *string `json:"limitScope,omitempty" xml:"limitScope,omitempty"`
	// The quota metric.
	//
	// example:
	//
	// Route
	QuotaKey *string `json:"quotaKey,omitempty" xml:"quotaKey,omitempty"`
	// The current usage.
	//
	// example:
	//
	// 12
	Used *int64 `json:"used,omitempty" xml:"used,omitempty"`
	// The scope for usage statistics.
	//
	// example:
	//
	// GATEWAY
	UsedScope *string `json:"usedScope,omitempty" xml:"usedScope,omitempty"`
}

func (s GetGatewayResourceQuotaUsageResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s GetGatewayResourceQuotaUsageResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) GetApplicable() *bool {
	return s.Applicable
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) GetLimit() *int64 {
	return s.Limit
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) GetLimitScope() *string {
	return s.LimitScope
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) GetQuotaKey() *string {
	return s.QuotaKey
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) GetUsed() *int64 {
	return s.Used
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) GetUsedScope() *string {
	return s.UsedScope
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) SetApplicable(v bool) *GetGatewayResourceQuotaUsageResponseBodyDataItems {
	s.Applicable = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) SetLimit(v int64) *GetGatewayResourceQuotaUsageResponseBodyDataItems {
	s.Limit = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) SetLimitScope(v string) *GetGatewayResourceQuotaUsageResponseBodyDataItems {
	s.LimitScope = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) SetQuotaKey(v string) *GetGatewayResourceQuotaUsageResponseBodyDataItems {
	s.QuotaKey = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) SetUsed(v int64) *GetGatewayResourceQuotaUsageResponseBodyDataItems {
	s.Used = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) SetUsedScope(v string) *GetGatewayResourceQuotaUsageResponseBodyDataItems {
	s.UsedScope = &v
	return s
}

func (s *GetGatewayResourceQuotaUsageResponseBodyDataItems) Validate() error {
	return dara.Validate(s)
}
