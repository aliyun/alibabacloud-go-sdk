// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIdentityProvidersForNetworkAccessEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIdentityProvidersForNetworkAccessEndpoint(v []*ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint) *ListIdentityProvidersForNetworkAccessEndpointResponseBody
	GetIdentityProvidersForNetworkAccessEndpoint() []*ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint
	SetNextToken(v string) *ListIdentityProvidersForNetworkAccessEndpointResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListIdentityProvidersForNetworkAccessEndpointResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListIdentityProvidersForNetworkAccessEndpointResponseBody
	GetTotalCount() *int64
}

type ListIdentityProvidersForNetworkAccessEndpointResponseBody struct {
	IdentityProvidersForNetworkAccessEndpoint []*ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint `json:"IdentityProvidersForNetworkAccessEndpoint,omitempty" xml:"IdentityProvidersForNetworkAccessEndpoint,omitempty" type:"Repeated"`
	// 本次调用返回的查询凭证（Token）值，用于下一次翻页查询。
	//
	// example:
	//
	// NTxxxexample
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// example:
	//
	// 0441BD79-92F3-53AA-8657-F8CE4A2B912A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListIdentityProvidersForNetworkAccessEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersForNetworkAccessEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBody) GetIdentityProvidersForNetworkAccessEndpoint() []*ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint {
	return s.IdentityProvidersForNetworkAccessEndpoint
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBody) SetIdentityProvidersForNetworkAccessEndpoint(v []*ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint) *ListIdentityProvidersForNetworkAccessEndpointResponseBody {
	s.IdentityProvidersForNetworkAccessEndpoint = v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBody) SetNextToken(v string) *ListIdentityProvidersForNetworkAccessEndpointResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBody) SetRequestId(v string) *ListIdentityProvidersForNetworkAccessEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBody) SetTotalCount(v int64) *ListIdentityProvidersForNetworkAccessEndpointResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBody) Validate() error {
	if s.IdentityProvidersForNetworkAccessEndpoint != nil {
		for _, item := range s.IdentityProvidersForNetworkAccessEndpoint {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint struct {
	// IdP的ID。
	//
	// example:
	//
	// idp_nbq7i4ylodmm64iy6t5muxxxxx
	IdentityProviderId *string `json:"IdentityProviderId,omitempty" xml:"IdentityProviderId,omitempty"`
	// IdP名称。
	//
	// example:
	//
	// OIDC Provider
	IdentityProviderName *string `json:"IdentityProviderName,omitempty" xml:"IdentityProviderName,omitempty"`
	// IDaaS EIAM 实例ID
	//
	// example:
	//
	// idaas_elk5evwagodqlmwpfehasxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint) GoString() string {
	return s.String()
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint) GetIdentityProviderId() *string {
	return s.IdentityProviderId
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint) GetIdentityProviderName() *string {
	return s.IdentityProviderName
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint) SetIdentityProviderId(v string) *ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint {
	s.IdentityProviderId = &v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint) SetIdentityProviderName(v string) *ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint {
	s.IdentityProviderName = &v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint) SetInstanceId(v string) *ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint {
	s.InstanceId = &v
	return s
}

func (s *ListIdentityProvidersForNetworkAccessEndpointResponseBodyIdentityProvidersForNetworkAccessEndpoint) Validate() error {
	return dara.Validate(s)
}
