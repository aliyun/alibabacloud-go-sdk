// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsDeliveryConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDeliveryConfigs(v []*DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs) *DescribeSlsDeliveryConfigResponseBody
	GetDeliveryConfigs() []*DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs
	SetRequestId(v string) *DescribeSlsDeliveryConfigResponseBody
	GetRequestId() *string
}

type DescribeSlsDeliveryConfigResponseBody struct {
	// The list of delivery configurations that match the filter conditions. If no configurations match the conditions, an empty list is returned.
	DeliveryConfigs []*DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs `json:"DeliveryConfigs,omitempty" xml:"DeliveryConfigs,omitempty" type:"Repeated"`
	// The request ID used to locate the current call.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSlsDeliveryConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsDeliveryConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSlsDeliveryConfigResponseBody) GetDeliveryConfigs() []*DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs {
	return s.DeliveryConfigs
}

func (s *DescribeSlsDeliveryConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSlsDeliveryConfigResponseBody) SetDeliveryConfigs(v []*DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs) *DescribeSlsDeliveryConfigResponseBody {
	s.DeliveryConfigs = v
	return s
}

func (s *DescribeSlsDeliveryConfigResponseBody) SetRequestId(v string) *DescribeSlsDeliveryConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSlsDeliveryConfigResponseBody) Validate() error {
	if s.DeliveryConfigs != nil {
		for _, item := range s.DeliveryConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs struct {
	// The list of Logstore delivery configurations under the project.
	LogStores []*DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores `json:"LogStores,omitempty" xml:"LogStores,omitempty" type:"Repeated"`
	// The name of the destination SLS project.
	//
	// example:
	//
	// wuying-audit-cn-hangzhou
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the region where the destination SLS project is located.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
}

func (s DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs) GoString() string {
	return s.String()
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs) GetLogStores() []*DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores {
	return s.LogStores
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs) GetProjectName() *string {
	return s.ProjectName
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs) GetSlsRegionId() *string {
	return s.SlsRegionId
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs) SetLogStores(v []*DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores) *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs {
	s.LogStores = v
	return s
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs) SetProjectName(v string) *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs {
	s.ProjectName = &v
	return s
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs) SetSlsRegionId(v string) *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs {
	s.SlsRegionId = &v
	return s
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigs) Validate() error {
	if s.LogStores != nil {
		for _, item := range s.LogStores {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores struct {
	// The product delivery scopes of the Logstore and the delivery status of each scope.
	DeliveryScopes []*DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes `json:"DeliveryScopes,omitempty" xml:"DeliveryScopes,omitempty" type:"Repeated"`
	// The name of the destination SLS Logstore.
	//
	// example:
	//
	// wuying-session-audit
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The data retention period configured for the destination Logstore, in days.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores) GoString() string {
	return s.String()
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores) GetDeliveryScopes() []*DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes {
	return s.DeliveryScopes
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores) GetTtl() *int32 {
	return s.Ttl
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores) SetDeliveryScopes(v []*DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes) *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores {
	s.DeliveryScopes = v
	return s
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores) SetLogStoreName(v string) *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores {
	s.LogStoreName = &v
	return s
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores) SetTtl(v int32) *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores {
	s.Ttl = &v
	return s
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStores) Validate() error {
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

type DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes struct {
	// The product type of the delivery scope. For example, `WuyingServer` indicates Enterprise Edition workstations.
	//
	// example:
	//
	// WuyingServer
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The current status of the delivery configuration. Valid values:
	//
	// - `Stopped`: Delivery is stopped.
	//
	// - `Enabled`: Delivery is enabled.
	//
	// - `AbnormalLogstoreDeleted`: The destination Logstore is deleted, and delivery is abnormal.
	//
	// - `AbnormalProjectDeleted`: The destination project is deleted, and delivery is abnormal.
	//
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes) GoString() string {
	return s.String()
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes) GetStatus() *string {
	return s.Status
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes) SetProductType(v string) *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes {
	s.ProductType = &v
	return s
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes) SetStatus(v string) *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes {
	s.Status = &v
	return s
}

func (s *DescribeSlsDeliveryConfigResponseBodyDeliveryConfigsLogStoresDeliveryScopes) Validate() error {
	return dara.Validate(s)
}
