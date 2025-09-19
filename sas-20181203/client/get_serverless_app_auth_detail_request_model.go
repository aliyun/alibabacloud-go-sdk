// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServerlessAppAuthDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *GetServerlessAppAuthDetailRequest
	GetAppId() *string
	SetAppRegionId(v string) *GetServerlessAppAuthDetailRequest
	GetAppRegionId() *string
	SetMachineType(v string) *GetServerlessAppAuthDetailRequest
	GetMachineType() *string
	SetVendorType(v string) *GetServerlessAppAuthDetailRequest
	GetVendorType() *string
}

type GetServerlessAppAuthDetailRequest struct {
	// SAE application ID.
	//
	// > Obtain through the [ListMachineApps](~~ListMachineApps~~) interface.
	//
	// example:
	//
	// 3de9f2ac-f***769226df
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// Application region ID.
	//
	// example:
	//
	// cn-hangzhou
	AppRegionId *string `json:"AppRegionId,omitempty" xml:"AppRegionId,omitempty"`
	// Server type:
	//
	// - **RunD**
	//
	// - **ECI**
	//
	// example:
	//
	// RunD
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Cloud product:
	//
	// - **ASK**
	//
	// - **SAE**
	//
	// - **ACS**
	//
	// example:
	//
	// SAE
	VendorType *string `json:"VendorType,omitempty" xml:"VendorType,omitempty"`
}

func (s GetServerlessAppAuthDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServerlessAppAuthDetailRequest) GoString() string {
	return s.String()
}

func (s *GetServerlessAppAuthDetailRequest) GetAppId() *string {
	return s.AppId
}

func (s *GetServerlessAppAuthDetailRequest) GetAppRegionId() *string {
	return s.AppRegionId
}

func (s *GetServerlessAppAuthDetailRequest) GetMachineType() *string {
	return s.MachineType
}

func (s *GetServerlessAppAuthDetailRequest) GetVendorType() *string {
	return s.VendorType
}

func (s *GetServerlessAppAuthDetailRequest) SetAppId(v string) *GetServerlessAppAuthDetailRequest {
	s.AppId = &v
	return s
}

func (s *GetServerlessAppAuthDetailRequest) SetAppRegionId(v string) *GetServerlessAppAuthDetailRequest {
	s.AppRegionId = &v
	return s
}

func (s *GetServerlessAppAuthDetailRequest) SetMachineType(v string) *GetServerlessAppAuthDetailRequest {
	s.MachineType = &v
	return s
}

func (s *GetServerlessAppAuthDetailRequest) SetVendorType(v string) *GetServerlessAppAuthDetailRequest {
	s.VendorType = &v
	return s
}

func (s *GetServerlessAppAuthDetailRequest) Validate() error {
	return dara.Validate(s)
}
