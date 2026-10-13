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
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// idaas_example
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of results per page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// NT_example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The origin filter.
	//
	// example:
	//
	// https://console.qoder.com
	Origin *string `json:"Origin,omitempty" xml:"Origin,omitempty"`
	// The status filter. Valid values:
	//
	// - Enabled: Enabled.
	//
	// - Disabled: Disabled.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The name filter.
	//
	// example:
	//
	// Qoder Production Console
	TrustOriginName *string `json:"TrustOriginName,omitempty" xml:"TrustOriginName,omitempty"`
	// The trusted origin scene filter.
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
