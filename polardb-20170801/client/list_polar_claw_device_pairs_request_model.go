// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPolarClawDevicePairsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *ListPolarClawDevicePairsRequest
	GetApplicationId() *string
}

type ListPolarClawDevicePairsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
}

func (s ListPolarClawDevicePairsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPolarClawDevicePairsRequest) GoString() string {
	return s.String()
}

func (s *ListPolarClawDevicePairsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *ListPolarClawDevicePairsRequest) SetApplicationId(v string) *ListPolarClawDevicePairsRequest {
	s.ApplicationId = &v
	return s
}

func (s *ListPolarClawDevicePairsRequest) Validate() error {
	return dara.Validate(s)
}
