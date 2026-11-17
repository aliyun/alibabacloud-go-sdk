// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWuyingServerSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAfterStatus(v string) *ModifyWuyingServerSpecRequest
	GetAfterStatus() *string
	SetProductType(v string) *ModifyWuyingServerSpecRequest
	GetProductType() *string
	SetPromotionId(v string) *ModifyWuyingServerSpecRequest
	GetPromotionId() *string
	SetServerInstanceType(v string) *ModifyWuyingServerSpecRequest
	GetServerInstanceType() *string
	SetWuyingServerId(v string) *ModifyWuyingServerSpecRequest
	GetWuyingServerId() *string
}

type ModifyWuyingServerSpecRequest struct {
	// The status after the specification change is complete.
	//
	// - `stopped`: Shuts down the workstation after the specification change is complete.
	//
	// This option applies to Enterprise Edition or container workstations that support this setting.
	//
	// example:
	//
	// stopped
	AfterStatus *string `json:"AfterStatus,omitempty" xml:"AfterStatus,omitempty"`
	// The product type of the target workstation. It must match the actual type of the resource. For Enterprise Edition workstations, use `WuyingServer`. If this parameter is omitted, `WuyingServer` is used by default.
	//
	// example:
	//
	// WuyingServer
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The ID of the promotion. It is used to identify a promotional event or coupon.
	//
	// example:
	//
	// 17440009****
	PromotionId *string `json:"PromotionId,omitempty" xml:"PromotionId,omitempty"`
	// The ID of the target workstation specification. This parameter is required when you call the operation. Select a specification that is available in the region where the workstation is located and compatible with existing resources.
	//
	// **Limits:*	- Conversion between CPU specifications and GPU specifications is not supported.
	//
	// example:
	//
	// eds.proworkstation_flagship_elite_ne.96c384g.192g4x
	ServerInstanceType *string `json:"ServerInstanceType,omitempty" xml:"ServerInstanceType,omitempty"`
	// The ID of the workstation whose specifications you want to change. This parameter is required.
	//
	// example:
	//
	// ws-0bw2f11****dial
	WuyingServerId *string `json:"WuyingServerId,omitempty" xml:"WuyingServerId,omitempty"`
}

func (s ModifyWuyingServerSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWuyingServerSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyWuyingServerSpecRequest) GetAfterStatus() *string {
	return s.AfterStatus
}

func (s *ModifyWuyingServerSpecRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ModifyWuyingServerSpecRequest) GetPromotionId() *string {
	return s.PromotionId
}

func (s *ModifyWuyingServerSpecRequest) GetServerInstanceType() *string {
	return s.ServerInstanceType
}

func (s *ModifyWuyingServerSpecRequest) GetWuyingServerId() *string {
	return s.WuyingServerId
}

func (s *ModifyWuyingServerSpecRequest) SetAfterStatus(v string) *ModifyWuyingServerSpecRequest {
	s.AfterStatus = &v
	return s
}

func (s *ModifyWuyingServerSpecRequest) SetProductType(v string) *ModifyWuyingServerSpecRequest {
	s.ProductType = &v
	return s
}

func (s *ModifyWuyingServerSpecRequest) SetPromotionId(v string) *ModifyWuyingServerSpecRequest {
	s.PromotionId = &v
	return s
}

func (s *ModifyWuyingServerSpecRequest) SetServerInstanceType(v string) *ModifyWuyingServerSpecRequest {
	s.ServerInstanceType = &v
	return s
}

func (s *ModifyWuyingServerSpecRequest) SetWuyingServerId(v string) *ModifyWuyingServerSpecRequest {
	s.WuyingServerId = &v
	return s
}

func (s *ModifyWuyingServerSpecRequest) Validate() error {
	return dara.Validate(s)
}
