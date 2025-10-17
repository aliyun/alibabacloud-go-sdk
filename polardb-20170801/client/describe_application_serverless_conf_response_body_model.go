// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApplicationServerlessConfResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribeApplicationServerlessConfResponseBody
	GetApplicationId() *string
	SetRequestId(v string) *DescribeApplicationServerlessConfResponseBody
	GetRequestId() *string
	SetServerlessConfItems(v []*DescribeApplicationServerlessConfResponseBodyServerlessConfItems) *DescribeApplicationServerlessConfResponseBody
	GetServerlessConfItems() []*DescribeApplicationServerlessConfResponseBodyServerlessConfItems
}

type DescribeApplicationServerlessConfResponseBody struct {
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// CDB3258F-B5DE-43C4-8935-CBA0CA******
	RequestId           *string                                                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ServerlessConfItems []*DescribeApplicationServerlessConfResponseBodyServerlessConfItems `json:"ServerlessConfItems,omitempty" xml:"ServerlessConfItems,omitempty" type:"Repeated"`
}

func (s DescribeApplicationServerlessConfResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationServerlessConfResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApplicationServerlessConfResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribeApplicationServerlessConfResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApplicationServerlessConfResponseBody) GetServerlessConfItems() []*DescribeApplicationServerlessConfResponseBodyServerlessConfItems {
	return s.ServerlessConfItems
}

func (s *DescribeApplicationServerlessConfResponseBody) SetApplicationId(v string) *DescribeApplicationServerlessConfResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *DescribeApplicationServerlessConfResponseBody) SetRequestId(v string) *DescribeApplicationServerlessConfResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApplicationServerlessConfResponseBody) SetServerlessConfItems(v []*DescribeApplicationServerlessConfResponseBodyServerlessConfItems) *DescribeApplicationServerlessConfResponseBody {
	s.ServerlessConfItems = v
	return s
}

func (s *DescribeApplicationServerlessConfResponseBody) Validate() error {
	if s.ServerlessConfItems != nil {
		for _, item := range s.ServerlessConfItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeApplicationServerlessConfResponseBodyServerlessConfItems struct {
	// example:
	//
	// gateway
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// example:
	//
	// 16
	ScaleMax *string `json:"ScaleMax,omitempty" xml:"ScaleMax,omitempty"`
	// example:
	//
	// 1
	ScaleMin *string `json:"ScaleMin,omitempty" xml:"ScaleMin,omitempty"`
}

func (s DescribeApplicationServerlessConfResponseBodyServerlessConfItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeApplicationServerlessConfResponseBodyServerlessConfItems) GoString() string {
	return s.String()
}

func (s *DescribeApplicationServerlessConfResponseBodyServerlessConfItems) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeApplicationServerlessConfResponseBodyServerlessConfItems) GetScaleMax() *string {
	return s.ScaleMax
}

func (s *DescribeApplicationServerlessConfResponseBodyServerlessConfItems) GetScaleMin() *string {
	return s.ScaleMin
}

func (s *DescribeApplicationServerlessConfResponseBodyServerlessConfItems) SetComponentType(v string) *DescribeApplicationServerlessConfResponseBodyServerlessConfItems {
	s.ComponentType = &v
	return s
}

func (s *DescribeApplicationServerlessConfResponseBodyServerlessConfItems) SetScaleMax(v string) *DescribeApplicationServerlessConfResponseBodyServerlessConfItems {
	s.ScaleMax = &v
	return s
}

func (s *DescribeApplicationServerlessConfResponseBodyServerlessConfItems) SetScaleMin(v string) *DescribeApplicationServerlessConfResponseBodyServerlessConfItems {
	s.ScaleMin = &v
	return s
}

func (s *DescribeApplicationServerlessConfResponseBodyServerlessConfItems) Validate() error {
	return dara.Validate(s)
}
