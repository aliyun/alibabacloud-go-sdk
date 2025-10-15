// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForNetworkAccessEndpointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationsForNetworkAccessEndpoint(v []*ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint) *ListApplicationsForNetworkAccessEndpointResponseBody
	GetApplicationsForNetworkAccessEndpoint() []*ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint
	SetNextToken(v string) *ListApplicationsForNetworkAccessEndpointResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListApplicationsForNetworkAccessEndpointResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *ListApplicationsForNetworkAccessEndpointResponseBody
	GetTotalCount() *int64
}

type ListApplicationsForNetworkAccessEndpointResponseBody struct {
	ApplicationsForNetworkAccessEndpoint []*ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint `json:"ApplicationsForNetworkAccessEndpoint,omitempty" xml:"ApplicationsForNetworkAccessEndpoint,omitempty" type:"Repeated"`
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

func (s ListApplicationsForNetworkAccessEndpointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForNetworkAccessEndpointResponseBody) GoString() string {
	return s.String()
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBody) GetApplicationsForNetworkAccessEndpoint() []*ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint {
	return s.ApplicationsForNetworkAccessEndpoint
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBody) SetApplicationsForNetworkAccessEndpoint(v []*ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint) *ListApplicationsForNetworkAccessEndpointResponseBody {
	s.ApplicationsForNetworkAccessEndpoint = v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBody) SetNextToken(v string) *ListApplicationsForNetworkAccessEndpointResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBody) SetRequestId(v string) *ListApplicationsForNetworkAccessEndpointResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBody) SetTotalCount(v int64) *ListApplicationsForNetworkAccessEndpointResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBody) Validate() error {
	if s.ApplicationsForNetworkAccessEndpoint != nil {
		for _, item := range s.ApplicationsForNetworkAccessEndpoint {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint struct {
	// 应用ID。
	//
	// example:
	//
	// app_m5nzr3kk4njkco2nnc4wjxxxxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// 应用名称。
	ApplicationName *string `json:"ApplicationName,omitempty" xml:"ApplicationName,omitempty"`
	// IDaaS EIAM 实例ID
	//
	// example:
	//
	// idaas_6ed5syotlwdrgmbzn7qn5xxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint) GoString() string {
	return s.String()
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint) GetApplicationName() *string {
	return s.ApplicationName
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint) SetApplicationId(v string) *ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint {
	s.ApplicationId = &v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint) SetApplicationName(v string) *ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint {
	s.ApplicationName = &v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint) SetInstanceId(v string) *ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint {
	s.InstanceId = &v
	return s
}

func (s *ListApplicationsForNetworkAccessEndpointResponseBodyApplicationsForNetworkAccessEndpoint) Validate() error {
	return dara.Validate(s)
}
