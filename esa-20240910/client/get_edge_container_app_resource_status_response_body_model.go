// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppResourceStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*GetEdgeContainerAppResourceStatusResponseBodyRegions) *GetEdgeContainerAppResourceStatusResponseBody
	GetRegions() []*GetEdgeContainerAppResourceStatusResponseBodyRegions
	SetRequestId(v string) *GetEdgeContainerAppResourceStatusResponseBody
	GetRequestId() *string
}

type GetEdgeContainerAppResourceStatusResponseBody struct {
	Regions []*GetEdgeContainerAppResourceStatusResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// example:
	//
	// 15C66C7B-671A-4297-9187-2C4477247A74
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEdgeContainerAppResourceStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceStatusResponseBody) GetRegions() []*GetEdgeContainerAppResourceStatusResponseBodyRegions {
	return s.Regions
}

func (s *GetEdgeContainerAppResourceStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerAppResourceStatusResponseBody) SetRegions(v []*GetEdgeContainerAppResourceStatusResponseBodyRegions) *GetEdgeContainerAppResourceStatusResponseBody {
	s.Regions = v
	return s
}

func (s *GetEdgeContainerAppResourceStatusResponseBody) SetRequestId(v string) *GetEdgeContainerAppResourceStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerAppResourceStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEdgeContainerAppResourceStatusResponseBodyRegions struct {
	IsOffline *bool `json:"IsOffline,omitempty" xml:"IsOffline,omitempty"`
	IsStaging *bool `json:"IsStaging,omitempty" xml:"IsStaging,omitempty"`
	// example:
	//
	// unicom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// example:
	//
	// 1
	Ready *int32 `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// example:
	//
	// huadong
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// example:
	//
	// 1
	Total *int32 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s GetEdgeContainerAppResourceStatusResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceStatusResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) GetIsOffline() *bool {
	return s.IsOffline
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) GetIsStaging() *bool {
	return s.IsStaging
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) GetIsp() *string {
	return s.Isp
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) GetReady() *int32 {
	return s.Ready
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) GetRegion() *string {
	return s.Region
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) GetTotal() *int32 {
	return s.Total
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) SetIsOffline(v bool) *GetEdgeContainerAppResourceStatusResponseBodyRegions {
	s.IsOffline = &v
	return s
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) SetIsStaging(v bool) *GetEdgeContainerAppResourceStatusResponseBodyRegions {
	s.IsStaging = &v
	return s
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) SetIsp(v string) *GetEdgeContainerAppResourceStatusResponseBodyRegions {
	s.Isp = &v
	return s
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) SetReady(v int32) *GetEdgeContainerAppResourceStatusResponseBodyRegions {
	s.Ready = &v
	return s
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) SetRegion(v string) *GetEdgeContainerAppResourceStatusResponseBodyRegions {
	s.Region = &v
	return s
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) SetTotal(v int32) *GetEdgeContainerAppResourceStatusResponseBodyRegions {
	s.Total = &v
	return s
}

func (s *GetEdgeContainerAppResourceStatusResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
