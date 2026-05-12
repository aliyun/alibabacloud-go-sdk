// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeliverToUserSlsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryScopes(v []*DeliverToUserSlsRequestDeliveryScopes) *DeliverToUserSlsRequest
	GetDeliveryScopes() []*DeliverToUserSlsRequestDeliveryScopes
	SetExistedProjectName(v string) *DeliverToUserSlsRequest
	GetExistedProjectName() *string
	SetLogStoreName(v string) *DeliverToUserSlsRequest
	GetLogStoreName() *string
	SetProjectName(v string) *DeliverToUserSlsRequest
	GetProjectName() *string
	SetSlsRegionId(v string) *DeliverToUserSlsRequest
	GetSlsRegionId() *string
	SetTtl(v int32) *DeliverToUserSlsRequest
	GetTtl() *int32
}

type DeliverToUserSlsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// [{"productType":"China_China"}]
	DeliveryScopes []*DeliverToUserSlsRequestDeliveryScopes `json:"DeliveryScopes,omitempty" xml:"DeliveryScopes,omitempty" type:"Repeated"`
	// example:
	//
	// elastic-desktop-xxx
	ExistedProjectName *string `json:"ExistedProjectName,omitempty" xml:"ExistedProjectName,omitempty"`
	// example:
	//
	// elastic_desktop_xxx
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// example:
	//
	// elastic-desktop-xxx
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s DeliverToUserSlsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeliverToUserSlsRequest) GoString() string {
	return s.String()
}

func (s *DeliverToUserSlsRequest) GetDeliveryScopes() []*DeliverToUserSlsRequestDeliveryScopes {
	return s.DeliveryScopes
}

func (s *DeliverToUserSlsRequest) GetExistedProjectName() *string {
	return s.ExistedProjectName
}

func (s *DeliverToUserSlsRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DeliverToUserSlsRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *DeliverToUserSlsRequest) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *DeliverToUserSlsRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *DeliverToUserSlsRequest) SetDeliveryScopes(v []*DeliverToUserSlsRequestDeliveryScopes) *DeliverToUserSlsRequest {
	s.DeliveryScopes = v
	return s
}

func (s *DeliverToUserSlsRequest) SetExistedProjectName(v string) *DeliverToUserSlsRequest {
	s.ExistedProjectName = &v
	return s
}

func (s *DeliverToUserSlsRequest) SetLogStoreName(v string) *DeliverToUserSlsRequest {
	s.LogStoreName = &v
	return s
}

func (s *DeliverToUserSlsRequest) SetProjectName(v string) *DeliverToUserSlsRequest {
	s.ProjectName = &v
	return s
}

func (s *DeliverToUserSlsRequest) SetSlsRegionId(v string) *DeliverToUserSlsRequest {
	s.SlsRegionId = &v
	return s
}

func (s *DeliverToUserSlsRequest) SetTtl(v int32) *DeliverToUserSlsRequest {
	s.Ttl = &v
	return s
}

func (s *DeliverToUserSlsRequest) Validate() error {
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

type DeliverToUserSlsRequestDeliveryScopes struct {
	// This parameter is required.
	//
	// example:
	//
	// CloudBrowser
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s DeliverToUserSlsRequestDeliveryScopes) String() string {
	return dara.Prettify(s)
}

func (s DeliverToUserSlsRequestDeliveryScopes) GoString() string {
	return s.String()
}

func (s *DeliverToUserSlsRequestDeliveryScopes) GetProductType() *string {
	return s.ProductType
}

func (s *DeliverToUserSlsRequestDeliveryScopes) SetProductType(v string) *DeliverToUserSlsRequestDeliveryScopes {
	s.ProductType = &v
	return s
}

func (s *DeliverToUserSlsRequestDeliveryScopes) Validate() error {
	return dara.Validate(s)
}
