// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupplierInformationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetSupplierInformationRequest
	GetRegionId() *string
}

type GetSupplierInformationRequest struct {
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetSupplierInformationRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSupplierInformationRequest) GoString() string {
	return s.String()
}

func (s *GetSupplierInformationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetSupplierInformationRequest) SetRegionId(v string) *GetSupplierInformationRequest {
	s.RegionId = &v
	return s
}

func (s *GetSupplierInformationRequest) Validate() error {
	return dara.Validate(s)
}
