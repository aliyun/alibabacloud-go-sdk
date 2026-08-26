// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrustedOriginsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListTrustedOriginsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *ListTrustedOriginsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListTrustedOriginsRequest
	GetNextToken() *string
	SetOrigin(v string) *ListTrustedOriginsRequest
	GetOrigin() *string
	SetStatus(v string) *ListTrustedOriginsRequest
	GetStatus() *string
	SetTrustOriginName(v string) *ListTrustedOriginsRequest
	GetTrustOriginName() *string
	SetTrustedOriginScene(v []*string) *ListTrustedOriginsRequest
	GetTrustedOriginScene() []*string
}

type ListTrustedOriginsRequest struct {
	// The ID of the IDaaS EIAM instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_example
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The number of entries per page. Default value: 20. Maximum value: 100. If you set this parameter to 0, the default value is used.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The NextToken returned by the previous call.
	//
	// example:
	//
	// NT_example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Filters by the exact normalized origin.
	//
	// example:
	//
	// https://console.qoder.com
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// Filters by exact status. Valid values: Enabled or Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Filters by exact name.
	//
	// example:
	//
	// Qoder Production Console
	TrustOriginName *string `json:"TrustOriginName,omitempty" xml:"TrustOriginName,omitempty"`
	// Filters by exact trusted origin scene. You can specify at most one value.
	//
	// example:
	//
	// iframe_embed
	TrustedOriginScene []*string `json:"TrustedOriginScene,omitempty" xml:"TrustedOriginScene,omitempty" type:"Repeated"`
}

func (s ListTrustedOriginsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTrustedOriginsRequest) GoString() string {
	return s.String()
}

func (s *ListTrustedOriginsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTrustedOriginsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListTrustedOriginsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTrustedOriginsRequest) GetOrigin() *string {
	return s.Origin
}

func (s *ListTrustedOriginsRequest) GetStatus() *string {
	return s.Status
}

func (s *ListTrustedOriginsRequest) GetTrustOriginName() *string {
	return s.TrustOriginName
}

func (s *ListTrustedOriginsRequest) GetTrustedOriginScene() []*string {
	return s.TrustedOriginScene
}

func (s *ListTrustedOriginsRequest) SetInstanceId(v string) *ListTrustedOriginsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTrustedOriginsRequest) SetMaxResults(v int32) *ListTrustedOriginsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListTrustedOriginsRequest) SetNextToken(v string) *ListTrustedOriginsRequest {
	s.NextToken = &v
	return s
}

func (s *ListTrustedOriginsRequest) SetOrigin(v string) *ListTrustedOriginsRequest {
	s.Origin = &v
	return s
}

func (s *ListTrustedOriginsRequest) SetStatus(v string) *ListTrustedOriginsRequest {
	s.Status = &v
	return s
}

func (s *ListTrustedOriginsRequest) SetTrustOriginName(v string) *ListTrustedOriginsRequest {
	s.TrustOriginName = &v
	return s
}

func (s *ListTrustedOriginsRequest) SetTrustedOriginScene(v []*string) *ListTrustedOriginsRequest {
	s.TrustedOriginScene = v
	return s
}

func (s *ListTrustedOriginsRequest) Validate() error {
	return dara.Validate(s)
}
