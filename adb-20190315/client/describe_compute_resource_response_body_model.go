// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComputeResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetComputeResource(v []*DescribeComputeResourceResponseBodyComputeResource) *DescribeComputeResourceResponseBody
	GetComputeResource() []*DescribeComputeResourceResponseBodyComputeResource
	SetRequestId(v string) *DescribeComputeResourceResponseBody
	GetRequestId() *string
}

type DescribeComputeResourceResponseBody struct {
	// The queried specifications of computing resources.
	ComputeResource []*DescribeComputeResourceResponseBodyComputeResource `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEAW
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeComputeResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeComputeResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceResponseBody) GetComputeResource() []*DescribeComputeResourceResponseBodyComputeResource {
	return s.ComputeResource
}

func (s *DescribeComputeResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeComputeResourceResponseBody) SetComputeResource(v []*DescribeComputeResourceResponseBodyComputeResource) *DescribeComputeResourceResponseBody {
	s.ComputeResource = v
	return s
}

func (s *DescribeComputeResourceResponseBody) SetRequestId(v string) *DescribeComputeResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeComputeResourceResponseBody) Validate() error {
	if s.ComputeResource != nil {
		for _, item := range s.ComputeResource {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeComputeResourceResponseBodyComputeResource struct {
	// The specifications of computing resources displayed in the console.
	//
	// example:
	//
	// 8 Core 32 GB
	DisplayValue *string `json:"DisplayValue,omitempty" xml:"DisplayValue,omitempty"`
	// The actual specifications of computing resources.
	//
	// example:
	//
	// 8 Core 32 GB
	RealValue *string `json:"RealValue,omitempty" xml:"RealValue,omitempty"`
}

func (s DescribeComputeResourceResponseBodyComputeResource) String() string {
	return dara.Prettify(s)
}

func (s DescribeComputeResourceResponseBodyComputeResource) GoString() string {
	return s.String()
}

func (s *DescribeComputeResourceResponseBodyComputeResource) GetDisplayValue() *string {
	return s.DisplayValue
}

func (s *DescribeComputeResourceResponseBodyComputeResource) GetRealValue() *string {
	return s.RealValue
}

func (s *DescribeComputeResourceResponseBodyComputeResource) SetDisplayValue(v string) *DescribeComputeResourceResponseBodyComputeResource {
	s.DisplayValue = &v
	return s
}

func (s *DescribeComputeResourceResponseBodyComputeResource) SetRealValue(v string) *DescribeComputeResourceResponseBodyComputeResource {
	s.RealValue = &v
	return s
}

func (s *DescribeComputeResourceResponseBodyComputeResource) Validate() error {
	return dara.Validate(s)
}
