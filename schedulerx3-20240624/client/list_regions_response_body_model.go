// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *ListRegionsResponseBody
	GetCode() *int32
	SetMessage(v string) *ListRegionsResponseBody
	GetMessage() *string
	SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody
	GetRegions() []*ListRegionsResponseBodyRegions
	SetRequestId(v string) *ListRegionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListRegionsResponseBody
	GetSuccess() *bool
}

type ListRegionsResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// Parameter check error
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// -
	Regions []*ListRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// AFD5B166-4A7D-50DF-91BF-EFAFD41F7335
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListRegionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *ListRegionsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListRegionsResponseBody) GetRegions() []*ListRegionsResponseBodyRegions {
	return s.Regions
}

func (s *ListRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRegionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListRegionsResponseBody) SetCode(v int32) *ListRegionsResponseBody {
	s.Code = &v
	return s
}

func (s *ListRegionsResponseBody) SetMessage(v string) *ListRegionsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRegionsResponseBody) SetRegions(v []*ListRegionsResponseBodyRegions) *ListRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *ListRegionsResponseBody) SetRequestId(v string) *ListRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegionsResponseBody) SetSuccess(v bool) *ListRegionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListRegionsResponseBody) Validate() error {
	if s.Regions != nil {
		for _, item := range s.Regions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListRegionsResponseBodyRegions struct {
	// example:
	//
	// China (Hangzhou)
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// endpoint
	//
	// example:
	//
	// schedulerx3.cn-hangzhou.aliyuncs.com
	RegionEndpoint *string `json:"RegionEndpoint,omitempty" xml:"RegionEndpoint,omitempty"`
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListRegionsResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s ListRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *ListRegionsResponseBodyRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *ListRegionsResponseBodyRegions) GetRegionEndpoint() *string {
	return s.RegionEndpoint
}

func (s *ListRegionsResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *ListRegionsResponseBodyRegions) SetLocalName(v string) *ListRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionEndpoint(v string) *ListRegionsResponseBodyRegions {
	s.RegionEndpoint = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) SetRegionId(v string) *ListRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *ListRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
