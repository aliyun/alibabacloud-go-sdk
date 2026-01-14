// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListListenerCertificatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListListenerCertificatesRequest
	GetAcceleratorId() *string
	SetListenerId(v string) *ListListenerCertificatesRequest
	GetListenerId() *string
	SetMaxResults(v int32) *ListListenerCertificatesRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListListenerCertificatesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListListenerCertificatesRequest
	GetRegionId() *string
	SetRole(v string) *ListListenerCertificatesRequest
	GetRole() *string
}

type ListListenerCertificatesRequest struct {
	// The ID of the GA instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The ID of the listener.
	//
	// This parameter is required.
	//
	// example:
	//
	// lsr-bp1bpn0kn908w4nbw****
	ListenerId *string `json:"ListenerId,omitempty" xml:"ListenerId,omitempty"`
	// The maximum number of entries to return. Valid values: **1*	- to **50**. Default value: **20**.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. Valid values:
	//
	// 	- You do not need to specify this parameter for the first request.
	//
	// 	- You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the region where the GA instance is deployed. Set the value to **cn-hangzhou**.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the certificate. Valid values:
	//
	// 	- **default**
	//
	// 	- **additional**
	//
	// If you do not specify this parameter, default and additional certificates are returned by default.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// default
	Role *string `json:"Role,omitempty" xml:"Role,omitempty"`
}

func (s ListListenerCertificatesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListListenerCertificatesRequest) GoString() string {
	return s.String()
}

func (s *ListListenerCertificatesRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListListenerCertificatesRequest) GetListenerId() *string {
	return s.ListenerId
}

func (s *ListListenerCertificatesRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListListenerCertificatesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListListenerCertificatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListListenerCertificatesRequest) GetRole() *string {
	return s.Role
}

func (s *ListListenerCertificatesRequest) SetAcceleratorId(v string) *ListListenerCertificatesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetListenerId(v string) *ListListenerCertificatesRequest {
	s.ListenerId = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetMaxResults(v int32) *ListListenerCertificatesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetNextToken(v string) *ListListenerCertificatesRequest {
	s.NextToken = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetRegionId(v string) *ListListenerCertificatesRequest {
	s.RegionId = &v
	return s
}

func (s *ListListenerCertificatesRequest) SetRole(v string) *ListListenerCertificatesRequest {
	s.Role = &v
	return s
}

func (s *ListListenerCertificatesRequest) Validate() error {
	return dara.Validate(s)
}
