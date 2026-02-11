// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopWuyingServerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *StopWuyingServerRequest
	GetForce() *bool
	SetProductType(v string) *StopWuyingServerRequest
	GetProductType() *string
	SetWuyingServerIdList(v []*string) *StopWuyingServerRequest
	GetWuyingServerIdList() []*string
}

type StopWuyingServerRequest struct {
	// Force restart.
	//
	// Valid values:
	//
	// 	- True.
	//
	// 	- False
	//
	// example:
	//
	// True
	Force       *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// The list of workstation IDs.
	WuyingServerIdList []*string `json:"WuyingServerIdList,omitempty" xml:"WuyingServerIdList,omitempty" type:"Repeated"`
}

func (s StopWuyingServerRequest) String() string {
	return dara.Prettify(s)
}

func (s StopWuyingServerRequest) GoString() string {
	return s.String()
}

func (s *StopWuyingServerRequest) GetForce() *bool {
	return s.Force
}

func (s *StopWuyingServerRequest) GetProductType() *string {
	return s.ProductType
}

func (s *StopWuyingServerRequest) GetWuyingServerIdList() []*string {
	return s.WuyingServerIdList
}

func (s *StopWuyingServerRequest) SetForce(v bool) *StopWuyingServerRequest {
	s.Force = &v
	return s
}

func (s *StopWuyingServerRequest) SetProductType(v string) *StopWuyingServerRequest {
	s.ProductType = &v
	return s
}

func (s *StopWuyingServerRequest) SetWuyingServerIdList(v []*string) *StopWuyingServerRequest {
	s.WuyingServerIdList = v
	return s
}

func (s *StopWuyingServerRequest) Validate() error {
	return dara.Validate(s)
}
