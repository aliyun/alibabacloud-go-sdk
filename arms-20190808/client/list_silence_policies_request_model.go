// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSilencePoliciesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsDetail(v bool) *ListSilencePoliciesRequest
	GetIsDetail() *bool
	SetName(v string) *ListSilencePoliciesRequest
	GetName() *string
	SetPage(v int64) *ListSilencePoliciesRequest
	GetPage() *int64
	SetRegionId(v string) *ListSilencePoliciesRequest
	GetRegionId() *string
	SetSize(v int64) *ListSilencePoliciesRequest
	GetSize() *int64
}

type ListSilencePoliciesRequest struct {
	// Specifies whether to query the details of a silence policy. Valid values:
	//
	// 	- `true`: Details of the silence policy are queried.
	//
	// 	- `false`: Details about notification policies are not queried.
	//
	// example:
	//
	// true
	IsDetail *bool `json:"IsDetail,omitempty" xml:"IsDetail,omitempty"`
	// The name of the silence policy.
	//
	// example:
	//
	// silencepolicy_test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The number of the page to return.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	Page *int64 `json:"Page,omitempty" xml:"Page,omitempty"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of entries to return on each page.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ListSilencePoliciesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSilencePoliciesRequest) GoString() string {
	return s.String()
}

func (s *ListSilencePoliciesRequest) GetIsDetail() *bool {
	return s.IsDetail
}

func (s *ListSilencePoliciesRequest) GetName() *string {
	return s.Name
}

func (s *ListSilencePoliciesRequest) GetPage() *int64 {
	return s.Page
}

func (s *ListSilencePoliciesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListSilencePoliciesRequest) GetSize() *int64 {
	return s.Size
}

func (s *ListSilencePoliciesRequest) SetIsDetail(v bool) *ListSilencePoliciesRequest {
	s.IsDetail = &v
	return s
}

func (s *ListSilencePoliciesRequest) SetName(v string) *ListSilencePoliciesRequest {
	s.Name = &v
	return s
}

func (s *ListSilencePoliciesRequest) SetPage(v int64) *ListSilencePoliciesRequest {
	s.Page = &v
	return s
}

func (s *ListSilencePoliciesRequest) SetRegionId(v string) *ListSilencePoliciesRequest {
	s.RegionId = &v
	return s
}

func (s *ListSilencePoliciesRequest) SetSize(v int64) *ListSilencePoliciesRequest {
	s.Size = &v
	return s
}

func (s *ListSilencePoliciesRequest) Validate() error {
	return dara.Validate(s)
}
