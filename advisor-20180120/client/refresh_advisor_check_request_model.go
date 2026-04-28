// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRefreshAdvisorCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAssumeAliyunId(v int64) *RefreshAdvisorCheckRequest
	GetAssumeAliyunId() *int64
	SetCheckId(v string) *RefreshAdvisorCheckRequest
	GetCheckId() *string
	SetCheckPlanId(v int64) *RefreshAdvisorCheckRequest
	GetCheckPlanId() *int64
	SetLanguage(v string) *RefreshAdvisorCheckRequest
	GetLanguage() *string
	SetProduct(v string) *RefreshAdvisorCheckRequest
	GetProduct() *string
	SetResourceDimensionList(v []*RefreshAdvisorCheckRequestResourceDimensionList) *RefreshAdvisorCheckRequest
	GetResourceDimensionList() []*RefreshAdvisorCheckRequestResourceDimensionList
	SetResourceId(v string) *RefreshAdvisorCheckRequest
	GetResourceId() *string
	SetToken(v string) *RefreshAdvisorCheckRequest
	GetToken() *string
}

type RefreshAdvisorCheckRequest struct {
	AssumeAliyunId *int64 `json:"AssumeAliyunId,omitempty" xml:"AssumeAliyunId,omitempty"`
	// example:
	//
	// EcsHighCpuUtilization
	CheckId     *string `json:"CheckId,omitempty" xml:"CheckId,omitempty"`
	CheckPlanId *int64  `json:"CheckPlanId,omitempty" xml:"CheckPlanId,omitempty"`
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// example:
	//
	// ecs
	Product               *string                                            `json:"Product,omitempty" xml:"Product,omitempty"`
	ResourceDimensionList []*RefreshAdvisorCheckRequestResourceDimensionList `json:"ResourceDimensionList,omitempty" xml:"ResourceDimensionList,omitempty" type:"Repeated"`
	// example:
	//
	// i-bp67acfmxazb4p****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	Token      *string `json:"Token,omitempty" xml:"Token,omitempty"`
}

func (s RefreshAdvisorCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorCheckRequest) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckRequest) GetAssumeAliyunId() *int64 {
	return s.AssumeAliyunId
}

func (s *RefreshAdvisorCheckRequest) GetCheckId() *string {
	return s.CheckId
}

func (s *RefreshAdvisorCheckRequest) GetCheckPlanId() *int64 {
	return s.CheckPlanId
}

func (s *RefreshAdvisorCheckRequest) GetLanguage() *string {
	return s.Language
}

func (s *RefreshAdvisorCheckRequest) GetProduct() *string {
	return s.Product
}

func (s *RefreshAdvisorCheckRequest) GetResourceDimensionList() []*RefreshAdvisorCheckRequestResourceDimensionList {
	return s.ResourceDimensionList
}

func (s *RefreshAdvisorCheckRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *RefreshAdvisorCheckRequest) GetToken() *string {
	return s.Token
}

func (s *RefreshAdvisorCheckRequest) SetAssumeAliyunId(v int64) *RefreshAdvisorCheckRequest {
	s.AssumeAliyunId = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetCheckId(v string) *RefreshAdvisorCheckRequest {
	s.CheckId = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetCheckPlanId(v int64) *RefreshAdvisorCheckRequest {
	s.CheckPlanId = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetLanguage(v string) *RefreshAdvisorCheckRequest {
	s.Language = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetProduct(v string) *RefreshAdvisorCheckRequest {
	s.Product = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetResourceDimensionList(v []*RefreshAdvisorCheckRequestResourceDimensionList) *RefreshAdvisorCheckRequest {
	s.ResourceDimensionList = v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetResourceId(v string) *RefreshAdvisorCheckRequest {
	s.ResourceId = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) SetToken(v string) *RefreshAdvisorCheckRequest {
	s.Token = &v
	return s
}

func (s *RefreshAdvisorCheckRequest) Validate() error {
	if s.ResourceDimensionList != nil {
		for _, item := range s.ResourceDimensionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RefreshAdvisorCheckRequestResourceDimensionList struct {
	Cost         *bool   `json:"Cost,omitempty" xml:"Cost,omitempty"`
	Performance  *bool   `json:"Performance,omitempty" xml:"Performance,omitempty"`
	Product      *string `json:"Product,omitempty" xml:"Product,omitempty"`
	ProductName  *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	Reliablility *bool   `json:"Reliablility,omitempty" xml:"Reliablility,omitempty"`
	Security     *bool   `json:"Security,omitempty" xml:"Security,omitempty"`
	Service      *bool   `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s RefreshAdvisorCheckRequestResourceDimensionList) String() string {
	return dara.Prettify(s)
}

func (s RefreshAdvisorCheckRequestResourceDimensionList) GoString() string {
	return s.String()
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) GetCost() *bool {
	return s.Cost
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) GetPerformance() *bool {
	return s.Performance
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) GetProduct() *string {
	return s.Product
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) GetProductName() *string {
	return s.ProductName
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) GetReliablility() *bool {
	return s.Reliablility
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) GetSecurity() *bool {
	return s.Security
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) GetService() *bool {
	return s.Service
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetCost(v bool) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.Cost = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetPerformance(v bool) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.Performance = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetProduct(v string) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.Product = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetProductName(v string) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.ProductName = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetReliablility(v bool) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.Reliablility = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetSecurity(v bool) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.Security = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) SetService(v bool) *RefreshAdvisorCheckRequestResourceDimensionList {
	s.Service = &v
	return s
}

func (s *RefreshAdvisorCheckRequestResourceDimensionList) Validate() error {
	return dara.Validate(s)
}
