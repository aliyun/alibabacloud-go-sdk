// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsDeliveryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryScopes(v []*DescribeSlsDeliveryConfigRequestDeliveryScopes) *DescribeSlsDeliveryConfigRequest
	GetDeliveryScopes() []*DescribeSlsDeliveryConfigRequestDeliveryScopes
	SetLogStoreName(v string) *DescribeSlsDeliveryConfigRequest
	GetLogStoreName() *string
	SetProjectName(v string) *DescribeSlsDeliveryConfigRequest
	GetProjectName() *string
	SetSlsRegionId(v string) *DescribeSlsDeliveryConfigRequest
	GetSlsRegionId() *string
}

type DescribeSlsDeliveryConfigRequest struct {
	// Filters delivery scopes by the product type to which the logs belong. If this parameter is omitted or an empty list is passed, delivery scopes are not filtered by product type.
	//
	// example:
	//
	// [{"ProductType":"WuyingServer"}]
	DeliveryScopes []*DescribeSlsDeliveryConfigRequestDeliveryScopes `json:"DeliveryScopes,omitempty" xml:"DeliveryScopes,omitempty" type:"Repeated"`
	// The name of the SLS Logstore. Exact match is used. If this parameter is omitted, delivery scopes are not filtered by Logstore name.
	//
	// example:
	//
	// wuying-session-audit
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The name of the SLS project. Exact match is used. If this parameter is omitted, delivery scopes are not filtered by project name.
	//
	// example:
	//
	// wuying-audit-cn-hangzhou
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the region where the SLS project is located. Exact match is used. If this parameter is omitted, delivery scopes are not filtered by region.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
}

func (s DescribeSlsDeliveryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsDeliveryConfigRequest) GoString() string {
	return s.String()
}

func (s *DescribeSlsDeliveryConfigRequest) GetDeliveryScopes() []*DescribeSlsDeliveryConfigRequestDeliveryScopes {
	return s.DeliveryScopes
}

func (s *DescribeSlsDeliveryConfigRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeSlsDeliveryConfigRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeSlsDeliveryConfigRequest) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *DescribeSlsDeliveryConfigRequest) SetDeliveryScopes(v []*DescribeSlsDeliveryConfigRequestDeliveryScopes) *DescribeSlsDeliveryConfigRequest {
	s.DeliveryScopes = v
	return s
}

func (s *DescribeSlsDeliveryConfigRequest) SetLogStoreName(v string) *DescribeSlsDeliveryConfigRequest {
	s.LogStoreName = &v
	return s
}

func (s *DescribeSlsDeliveryConfigRequest) SetProjectName(v string) *DescribeSlsDeliveryConfigRequest {
	s.ProjectName = &v
	return s
}

func (s *DescribeSlsDeliveryConfigRequest) SetSlsRegionId(v string) *DescribeSlsDeliveryConfigRequest {
	s.SlsRegionId = &v
	return s
}

func (s *DescribeSlsDeliveryConfigRequest) Validate() error {
	if s.DeliveryScopes != nil {
		for _, item := range s.DeliveryScopes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlsDeliveryConfigRequestDeliveryScopes struct {
	// The query condition for the product type to which the logs belong. If you specify this parameter, delivery scopes are filtered by the specified product type. The value must be consistent with the existing delivery configurations. For example, `WuyingServer` indicates Enterprise Edition workstations.
	//
	// This parameter is required.
	//
	// example:
	//
	// WuyingServer
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DescribeSlsDeliveryConfigRequestDeliveryScopes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsDeliveryConfigRequestDeliveryScopes) GoString() string {
	return s.String()
}

func (s *DescribeSlsDeliveryConfigRequestDeliveryScopes) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeSlsDeliveryConfigRequestDeliveryScopes) SetProductType(v string) *DescribeSlsDeliveryConfigRequestDeliveryScopes {
	s.ProductType = &v
	return s
}

func (s *DescribeSlsDeliveryConfigRequestDeliveryScopes) Validate() error {
	return dara.Validate(s)
}
