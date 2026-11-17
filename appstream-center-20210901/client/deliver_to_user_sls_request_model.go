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
	// The list of product scopes for which delivery needs to be configured. Specify at least one item. You cannot create duplicate enabled delivery configurations for the same account and product type.
	//
	// This parameter is required.
	//
	// example:
	//
	// [{"ProductType":"CloudBrowser"}]
	DeliveryScopes []*DeliverToUserSlsRequestDeliveryScopes `json:"DeliveryScopes,omitempty" xml:"DeliveryScopes,omitempty" type:"Repeated"`
	// The name of an existing SLS project. The name must start with `elastic-desktop-` and the project must already exist in the specified SLS region of the current account.
	//
	// Specify either this parameter or `ProjectName`.
	//
	// example:
	//
	// elastic-desktop-audit
	ExistedProjectName *string `json:"ExistedProjectName,omitempty" xml:"ExistedProjectName,omitempty"`
	// The name of the destination Logstore. The name must start with `elastic_desktop_`. If the Logstore does not exist, it is created. If the Logstore already exists, it is reused.
	//
	// If this parameter is not specified, the part after `elastic-desktop-` in the selected project name is used as a suffix to generate a name that starts with `elastic_desktop_`. For example, if the project name is `elastic-desktop-audit`, the default Logstore name is `elastic_desktop_audit`.
	//
	// example:
	//
	// elastic_desktop_audit
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// The name of the SLS project to be created or reused. The name must start with `elastic-desktop-`.
	//
	// Specify either this parameter or `ExistedProjectName`. If the project does not exist, it is created. If the project already exists, it is reused.
	//
	// example:
	//
	// elastic-desktop-audit
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The ID of the region where the destination SLS project is located. This parameter is required.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	SlsRegionId *string `json:"SlsRegionId,omitempty" xml:"SlsRegionId,omitempty"`
	// The data retention period of the new Logstore. Unit: days.
	//
	// **Valid values:*	- 1 to 3650.
	//
	// **Default value:*	- 30.
	//
	// If an existing Logstore is reused, this operation does not update its data retention period.
	//
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
	// The product type to which the logs belong. Specify the product type for which you want to enable log delivery. For example, `CloudBrowser` indicates Cloud Browser.
	//
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
