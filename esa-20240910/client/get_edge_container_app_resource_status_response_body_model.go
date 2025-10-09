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
	// Queries the regions of deployment.
	Regions []*GetEdgeContainerAppResourceStatusResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
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
	// Whether smooth offline is being used.
	//
	// example:
	//
	// false
	IsOffline *bool `json:"IsOffline,omitempty" xml:"IsOffline,omitempty"`
	// Whether it is a test environment.
	//
	// example:
	//
	// false
	IsStaging *bool `json:"IsStaging,omitempty" xml:"IsStaging,omitempty"`
	// Supported ISPs are as follows. The parameter is left empty for regions outside the Chinese mainland. ISP:
	//
	// 	- China Mobile: cmcc
	//
	// 	- China Telecom: chinanet
	//
	// 	- China Unicom: unicom
	//
	// example:
	//
	// unicom
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The number of ready replicas.
	//
	// example:
	//
	// 1
	Ready *int32 `json:"Ready,omitempty" xml:"Ready,omitempty"`
	// Chinese mainland:
	//
	// 	- East China: huadong
	//
	// 	- South China: huanan
	//
	// 	- Central China: huazhong
	//
	// 	- North China: huabei
	//
	// 	- Northwest China: xibei
	//
	// 	- Southwest China: xinan
	//
	// 	- Northeast China: dongbei
	//
	// Special Administrative Regions and Overseas:
	//
	// 	- Taiwan, China: tw
	//
	// 	- Macau, China: mo
	//
	// 	- Hong Kong, China: hk
	//
	// 	- Japan: jp
	//
	// 	- United States: us
	//
	// 	- Thailand: th
	//
	// 	- Korea: kr
	//
	// 	- Russia: ru
	//
	// 	- Singapore: sg
	//
	// 	- France: fr
	//
	// 	- Spain: es
	//
	// 	- Italy: it
	//
	// 	- Sweden: se
	//
	// 	- UAE: ae
	//
	// 	- Indonesia: id
	//
	// 	- Chile: cl
	//
	// 	- Philippines: ph
	//
	// 	- Malaysia: my
	//
	// 	- Vietnam: vn
	//
	// 	- Argentina: AR
	//
	// 	- Australia: au
	//
	// 	- Brazil: br
	//
	// 	- Colombia: co
	//
	// 	- Germany: de
	//
	// 	- UK: GB
	//
	// 	- Peru: pe
	//
	// 	- Saudi Arabia: sa
	//
	// 	- Netherlands: nl
	//
	// 	- South Africa: za
	//
	// example:
	//
	// huadong
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The number of replicas that are deployed.
	//
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
