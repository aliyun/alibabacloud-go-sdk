// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServerlessAuthSummaryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppRegionId(v string) *GetServerlessAuthSummaryRequest
	GetAppRegionId() *string
	SetMachineType(v string) *GetServerlessAuthSummaryRequest
	GetMachineType() *string
	SetVendorType(v string) *GetServerlessAuthSummaryRequest
	GetVendorType() *string
}

type GetServerlessAuthSummaryRequest struct {
	// The region ID of the application.
	//
	// example:
	//
	// cn-hangzhou
	AppRegionId *string `json:"AppRegionId,omitempty" xml:"AppRegionId,omitempty"`
	// The server type. Valid values:
	//
	// - **RunD**
	//
	// - **ECI**.
	//
	// example:
	//
	// RunD
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The cloud service to which the resource belongs. Valid values:
	//
	// - **ASK**
	//
	// - **SAE**
	//
	// - **ACS**.
	//
	// example:
	//
	// SAE
	VendorType *string `json:"VendorType,omitempty" xml:"VendorType,omitempty"`
}

func (s GetServerlessAuthSummaryRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServerlessAuthSummaryRequest) GoString() string {
	return s.String()
}

func (s *GetServerlessAuthSummaryRequest) GetAppRegionId() *string {
	return s.AppRegionId
}

func (s *GetServerlessAuthSummaryRequest) GetMachineType() *string {
	return s.MachineType
}

func (s *GetServerlessAuthSummaryRequest) GetVendorType() *string {
	return s.VendorType
}

func (s *GetServerlessAuthSummaryRequest) SetAppRegionId(v string) *GetServerlessAuthSummaryRequest {
	s.AppRegionId = &v
	return s
}

func (s *GetServerlessAuthSummaryRequest) SetMachineType(v string) *GetServerlessAuthSummaryRequest {
	s.MachineType = &v
	return s
}

func (s *GetServerlessAuthSummaryRequest) SetVendorType(v string) *GetServerlessAuthSummaryRequest {
	s.VendorType = &v
	return s
}

func (s *GetServerlessAuthSummaryRequest) Validate() error {
	return dara.Validate(s)
}
