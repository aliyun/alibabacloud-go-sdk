// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentlessRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegionList(v []*string) *ListAgentlessRegionResponseBody
	GetRegionList() []*string
	SetRequestId(v string) *ListAgentlessRegionResponseBody
	GetRequestId() *string
	SetVendorRegionList(v []*ListAgentlessRegionResponseBodyVendorRegionList) *ListAgentlessRegionResponseBody
	GetVendorRegionList() []*ListAgentlessRegionResponseBodyVendorRegionList
}

type ListAgentlessRegionResponseBody struct {
	// The information about the regions.
	RegionList []*string `json:"RegionList,omitempty" xml:"RegionList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 7E0618A9-D5EF-4220-9471-C42B5E92****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the regions.
	VendorRegionList []*ListAgentlessRegionResponseBodyVendorRegionList `json:"VendorRegionList,omitempty" xml:"VendorRegionList,omitempty" type:"Repeated"`
}

func (s ListAgentlessRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRegionResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentlessRegionResponseBody) GetRegionList() []*string {
	return s.RegionList
}

func (s *ListAgentlessRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentlessRegionResponseBody) GetVendorRegionList() []*ListAgentlessRegionResponseBodyVendorRegionList {
	return s.VendorRegionList
}

func (s *ListAgentlessRegionResponseBody) SetRegionList(v []*string) *ListAgentlessRegionResponseBody {
	s.RegionList = v
	return s
}

func (s *ListAgentlessRegionResponseBody) SetRequestId(v string) *ListAgentlessRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentlessRegionResponseBody) SetVendorRegionList(v []*ListAgentlessRegionResponseBodyVendorRegionList) *ListAgentlessRegionResponseBody {
	s.VendorRegionList = v
	return s
}

func (s *ListAgentlessRegionResponseBody) Validate() error {
	if s.VendorRegionList != nil {
		for _, item := range s.VendorRegionList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAgentlessRegionResponseBodyVendorRegionList struct {
	// The region ID of the instance.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the server. Valid values:
	//
	// - **0**: Alibaba Cloud
	//
	// - **3**: Tencent Cloud
	//
	// - **4**: Huawei Cloud
	//
	// - **5**: Azure
	//
	// - **7**: AWS
	//
	// example:
	//
	// 0
	Vendor *int32 `json:"Vendor,omitempty" xml:"Vendor,omitempty"`
}

func (s ListAgentlessRegionResponseBodyVendorRegionList) String() string {
	return dara.Prettify(s)
}

func (s ListAgentlessRegionResponseBodyVendorRegionList) GoString() string {
	return s.String()
}

func (s *ListAgentlessRegionResponseBodyVendorRegionList) GetRegionId() *string {
	return s.RegionId
}

func (s *ListAgentlessRegionResponseBodyVendorRegionList) GetVendor() *int32 {
	return s.Vendor
}

func (s *ListAgentlessRegionResponseBodyVendorRegionList) SetRegionId(v string) *ListAgentlessRegionResponseBodyVendorRegionList {
	s.RegionId = &v
	return s
}

func (s *ListAgentlessRegionResponseBodyVendorRegionList) SetVendor(v int32) *ListAgentlessRegionResponseBodyVendorRegionList {
	s.Vendor = &v
	return s
}

func (s *ListAgentlessRegionResponseBodyVendorRegionList) Validate() error {
	return dara.Validate(s)
}
