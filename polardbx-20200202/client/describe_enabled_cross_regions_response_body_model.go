// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnabledCrossRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeEnabledCrossRegionsResponseBodyData) *DescribeEnabledCrossRegionsResponseBody
	GetData() []*DescribeEnabledCrossRegionsResponseBodyData
	SetMessage(v string) *DescribeEnabledCrossRegionsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEnabledCrossRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEnabledCrossRegionsResponseBody
	GetSuccess() *bool
}

type DescribeEnabledCrossRegionsResponseBody struct {
	Data []*DescribeEnabledCrossRegionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3WE34
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEnabledCrossRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnabledCrossRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEnabledCrossRegionsResponseBody) GetData() []*DescribeEnabledCrossRegionsResponseBodyData {
	return s.Data
}

func (s *DescribeEnabledCrossRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEnabledCrossRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEnabledCrossRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEnabledCrossRegionsResponseBody) SetData(v []*DescribeEnabledCrossRegionsResponseBodyData) *DescribeEnabledCrossRegionsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEnabledCrossRegionsResponseBody) SetMessage(v string) *DescribeEnabledCrossRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEnabledCrossRegionsResponseBody) SetRequestId(v string) *DescribeEnabledCrossRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEnabledCrossRegionsResponseBody) SetSuccess(v bool) *DescribeEnabledCrossRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEnabledCrossRegionsResponseBody) Validate() error {
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

type DescribeEnabledCrossRegionsResponseBodyData struct {
	Regions []*string `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
}

func (s DescribeEnabledCrossRegionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnabledCrossRegionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEnabledCrossRegionsResponseBodyData) GetRegions() []*string {
	return s.Regions
}

func (s *DescribeEnabledCrossRegionsResponseBodyData) SetRegions(v []*string) *DescribeEnabledCrossRegionsResponseBodyData {
	s.Regions = v
	return s
}

func (s *DescribeEnabledCrossRegionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
