// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIspTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *ListIspTypesRequest
	GetAcceleratorId() *string
	SetAcceleratorType(v string) *ListIspTypesRequest
	GetAcceleratorType() *string
	SetBusinessRegionId(v string) *ListIspTypesRequest
	GetBusinessRegionId() *string
}

type ListIspTypesRequest struct {
	// The ID of the GA instance that you want to query.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The type of the Global Accelerator (GA) instance to be queried. Valid values:
	//
	// 	- **basic**: basic GA instance
	//
	// 	- **standard**: standard GA instance
	//
	// example:
	//
	// basic
	AcceleratorType *string `json:"AcceleratorType,omitempty" xml:"AcceleratorType,omitempty"`
	// The ID of the acceleration region to be queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	BusinessRegionId *string `json:"BusinessRegionId,omitempty" xml:"BusinessRegionId,omitempty"`
}

func (s ListIspTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListIspTypesRequest) GoString() string {
	return s.String()
}

func (s *ListIspTypesRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *ListIspTypesRequest) GetAcceleratorType() *string {
	return s.AcceleratorType
}

func (s *ListIspTypesRequest) GetBusinessRegionId() *string {
	return s.BusinessRegionId
}

func (s *ListIspTypesRequest) SetAcceleratorId(v string) *ListIspTypesRequest {
	s.AcceleratorId = &v
	return s
}

func (s *ListIspTypesRequest) SetAcceleratorType(v string) *ListIspTypesRequest {
	s.AcceleratorType = &v
	return s
}

func (s *ListIspTypesRequest) SetBusinessRegionId(v string) *ListIspTypesRequest {
	s.BusinessRegionId = &v
	return s
}

func (s *ListIspTypesRequest) Validate() error {
	return dara.Validate(s)
}
