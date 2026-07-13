// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceEndpointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListServiceEndpointsResponseBody
	GetCode() *string
	SetItems(v []*ListServiceEndpointsResponseBodyItems) *ListServiceEndpointsResponseBody
	GetItems() []*ListServiceEndpointsResponseBodyItems
	SetMaxResults(v int32) *ListServiceEndpointsResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListServiceEndpointsResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListServiceEndpointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListServiceEndpointsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListServiceEndpointsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListServiceEndpointsResponseBody
	GetTotalCount() *int32
}

type ListServiceEndpointsResponseBody struct {
	// example:
	//
	// SUCCESS
	Code  *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Items []*ListServiceEndpointsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// example:
	//
	// 1
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// success
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// xxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListServiceEndpointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListServiceEndpointsResponseBody) GetItems() []*ListServiceEndpointsResponseBodyItems {
	return s.Items
}

func (s *ListServiceEndpointsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListServiceEndpointsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListServiceEndpointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListServiceEndpointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceEndpointsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListServiceEndpointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListServiceEndpointsResponseBody) SetCode(v string) *ListServiceEndpointsResponseBody {
	s.Code = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetItems(v []*ListServiceEndpointsResponseBodyItems) *ListServiceEndpointsResponseBody {
	s.Items = v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetMaxResults(v int32) *ListServiceEndpointsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetMessage(v string) *ListServiceEndpointsResponseBody {
	s.Message = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetNextToken(v string) *ListServiceEndpointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetRequestId(v string) *ListServiceEndpointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetSuccess(v bool) *ListServiceEndpointsResponseBody {
	s.Success = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) SetTotalCount(v int32) *ListServiceEndpointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListServiceEndpointsResponseBody) Validate() error {
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

type ListServiceEndpointsResponseBodyItems struct {
	// example:
	//
	// cert-xxx
	CertIdentifier *string `json:"CertIdentifier,omitempty" xml:"CertIdentifier,omitempty"`
	// example:
	//
	// MATRIX
	Component  *string `json:"Component,omitempty" xml:"Component,omitempty"`
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// matrix.example.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// CUSTOM
	DomainType     *string                                              `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	EndpointConfig *ListServiceEndpointsResponseBodyItemsEndpointConfig `json:"EndpointConfig,omitempty" xml:"EndpointConfig,omitempty" type:"Struct"`
	EndpointId     *string                                              `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// example:
	//
	// matrix-service
	EndpointName *string `json:"EndpointName,omitempty" xml:"EndpointName,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// INTERNET
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// example:
	//
	// CONFIGURED
	Status     *string `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListServiceEndpointsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBodyItems) GetCertIdentifier() *string {
	return s.CertIdentifier
}

func (s *ListServiceEndpointsResponseBodyItems) GetComponent() *string {
	return s.Component
}

func (s *ListServiceEndpointsResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListServiceEndpointsResponseBodyItems) GetDomain() *string {
	return s.Domain
}

func (s *ListServiceEndpointsResponseBodyItems) GetDomainType() *string {
	return s.DomainType
}

func (s *ListServiceEndpointsResponseBodyItems) GetEndpointConfig() *ListServiceEndpointsResponseBodyItemsEndpointConfig {
	return s.EndpointConfig
}

func (s *ListServiceEndpointsResponseBodyItems) GetEndpointId() *string {
	return s.EndpointId
}

func (s *ListServiceEndpointsResponseBodyItems) GetEndpointName() *string {
	return s.EndpointName
}

func (s *ListServiceEndpointsResponseBodyItems) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListServiceEndpointsResponseBodyItems) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListServiceEndpointsResponseBodyItems) GetStatus() *string {
	return s.Status
}

func (s *ListServiceEndpointsResponseBodyItems) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListServiceEndpointsResponseBodyItems) SetCertIdentifier(v string) *ListServiceEndpointsResponseBodyItems {
	s.CertIdentifier = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetComponent(v string) *ListServiceEndpointsResponseBodyItems {
	s.Component = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetCreateTime(v string) *ListServiceEndpointsResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetDomain(v string) *ListServiceEndpointsResponseBodyItems {
	s.Domain = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetDomainType(v string) *ListServiceEndpointsResponseBodyItems {
	s.DomainType = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetEndpointConfig(v *ListServiceEndpointsResponseBodyItemsEndpointConfig) *ListServiceEndpointsResponseBodyItems {
	s.EndpointConfig = v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetEndpointId(v string) *ListServiceEndpointsResponseBodyItems {
	s.EndpointId = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetEndpointName(v string) *ListServiceEndpointsResponseBodyItems {
	s.EndpointName = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetInstanceId(v string) *ListServiceEndpointsResponseBodyItems {
	s.InstanceId = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetNetworkType(v string) *ListServiceEndpointsResponseBodyItems {
	s.NetworkType = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetStatus(v string) *ListServiceEndpointsResponseBodyItems {
	s.Status = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) SetUpdateTime(v string) *ListServiceEndpointsResponseBodyItems {
	s.UpdateTime = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItems) Validate() error {
	if s.EndpointConfig != nil {
		if err := s.EndpointConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServiceEndpointsResponseBodyItemsEndpointConfig struct {
	Auth *ListServiceEndpointsResponseBodyItemsEndpointConfigAuth `json:"Auth,omitempty" xml:"Auth,omitempty" type:"Struct"`
}

func (s ListServiceEndpointsResponseBodyItemsEndpointConfig) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBodyItemsEndpointConfig) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBodyItemsEndpointConfig) GetAuth() *ListServiceEndpointsResponseBodyItemsEndpointConfigAuth {
	return s.Auth
}

func (s *ListServiceEndpointsResponseBodyItemsEndpointConfig) SetAuth(v *ListServiceEndpointsResponseBodyItemsEndpointConfigAuth) *ListServiceEndpointsResponseBodyItemsEndpointConfig {
	s.Auth = v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsEndpointConfig) Validate() error {
	if s.Auth != nil {
		if err := s.Auth.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListServiceEndpointsResponseBodyItemsEndpointConfigAuth struct {
	ApiKey     *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
	ApiKeyName *string `json:"ApiKeyName,omitempty" xml:"ApiKeyName,omitempty"`
}

func (s ListServiceEndpointsResponseBodyItemsEndpointConfigAuth) String() string {
	return dara.Prettify(s)
}

func (s ListServiceEndpointsResponseBodyItemsEndpointConfigAuth) GoString() string {
	return s.String()
}

func (s *ListServiceEndpointsResponseBodyItemsEndpointConfigAuth) GetApiKey() *string {
	return s.ApiKey
}

func (s *ListServiceEndpointsResponseBodyItemsEndpointConfigAuth) GetApiKeyName() *string {
	return s.ApiKeyName
}

func (s *ListServiceEndpointsResponseBodyItemsEndpointConfigAuth) SetApiKey(v string) *ListServiceEndpointsResponseBodyItemsEndpointConfigAuth {
	s.ApiKey = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsEndpointConfigAuth) SetApiKeyName(v string) *ListServiceEndpointsResponseBodyItemsEndpointConfigAuth {
	s.ApiKeyName = &v
	return s
}

func (s *ListServiceEndpointsResponseBodyItemsEndpointConfigAuth) Validate() error {
	return dara.Validate(s)
}
