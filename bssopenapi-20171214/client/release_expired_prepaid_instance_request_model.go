// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseExpiredPrepaidInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v string) *ReleaseExpiredPrepaidInstanceRequest
	GetInstanceIds() *string
	SetProductCode(v string) *ReleaseExpiredPrepaidInstanceRequest
	GetProductCode() *string
	SetProductType(v string) *ReleaseExpiredPrepaidInstanceRequest
	GetProductType() *string
}

type ReleaseExpiredPrepaidInstanceRequest struct {
	// The instance IDs. Separate multiple IDs with commas. You can specify up to 20 IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// rm-1,rm-2
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The product code. You can call the **QueryProductList*	- operation to query the product code.
	//
	// This parameter is required.
	//
	// example:
	//
	// rds
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The product type.
	//
	// example:
	//
	// rds
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
}

func (s ReleaseExpiredPrepaidInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseExpiredPrepaidInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseExpiredPrepaidInstanceRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *ReleaseExpiredPrepaidInstanceRequest) GetProductCode() *string {
	return s.ProductCode
}

func (s *ReleaseExpiredPrepaidInstanceRequest) GetProductType() *string {
	return s.ProductType
}

func (s *ReleaseExpiredPrepaidInstanceRequest) SetInstanceIds(v string) *ReleaseExpiredPrepaidInstanceRequest {
	s.InstanceIds = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceRequest) SetProductCode(v string) *ReleaseExpiredPrepaidInstanceRequest {
	s.ProductCode = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceRequest) SetProductType(v string) *ReleaseExpiredPrepaidInstanceRequest {
	s.ProductType = &v
	return s
}

func (s *ReleaseExpiredPrepaidInstanceRequest) Validate() error {
	return dara.Validate(s)
}
