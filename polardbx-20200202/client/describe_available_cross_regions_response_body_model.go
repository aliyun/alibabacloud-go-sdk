// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableCrossRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeAvailableCrossRegionsResponseBodyData) *DescribeAvailableCrossRegionsResponseBody
	GetData() []*DescribeAvailableCrossRegionsResponseBodyData
	SetMessage(v string) *DescribeAvailableCrossRegionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAvailableCrossRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAvailableCrossRegionsResponseBody
	GetSuccess() *bool
}

type DescribeAvailableCrossRegionsResponseBody struct {
	Data []*DescribeAvailableCrossRegionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeAvailableCrossRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableCrossRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCrossRegionsResponseBody) GetData() []*DescribeAvailableCrossRegionsResponseBodyData {
	return s.Data
}

func (s *DescribeAvailableCrossRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAvailableCrossRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAvailableCrossRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAvailableCrossRegionsResponseBody) SetData(v []*DescribeAvailableCrossRegionsResponseBodyData) *DescribeAvailableCrossRegionsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAvailableCrossRegionsResponseBody) SetMessage(v string) *DescribeAvailableCrossRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAvailableCrossRegionsResponseBody) SetRequestId(v string) *DescribeAvailableCrossRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAvailableCrossRegionsResponseBody) SetSuccess(v bool) *DescribeAvailableCrossRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAvailableCrossRegionsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAvailableCrossRegionsResponseBodyData struct {
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeAvailableCrossRegionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableCrossRegionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAvailableCrossRegionsResponseBodyData) GetRegions() []*string {
	return s.Regions
}

func (s *DescribeAvailableCrossRegionsResponseBodyData) SetRegions(v []*string) *DescribeAvailableCrossRegionsResponseBodyData {
	s.Regions = v
	return s
}

func (s *DescribeAvailableCrossRegionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
