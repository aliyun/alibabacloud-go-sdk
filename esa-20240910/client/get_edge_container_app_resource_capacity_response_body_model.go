// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEdgeContainerAppResourceCapacityResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*GetEdgeContainerAppResourceCapacityResponseBodyRegions) *GetEdgeContainerAppResourceCapacityResponseBody
	GetRegions() []*GetEdgeContainerAppResourceCapacityResponseBodyRegions
	SetRequestId(v string) *GetEdgeContainerAppResourceCapacityResponseBody
	GetRequestId() *string
}

type GetEdgeContainerAppResourceCapacityResponseBody struct {
	// The queried region.
	Regions []*GetEdgeContainerAppResourceCapacityResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 50423A7F-A83D-1E24-B80E-86DD25790759
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEdgeContainerAppResourceCapacityResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceCapacityResponseBody) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceCapacityResponseBody) GetRegions() []*GetEdgeContainerAppResourceCapacityResponseBodyRegions {
	return s.Regions
}

func (s *GetEdgeContainerAppResourceCapacityResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEdgeContainerAppResourceCapacityResponseBody) SetRegions(v []*GetEdgeContainerAppResourceCapacityResponseBodyRegions) *GetEdgeContainerAppResourceCapacityResponseBody {
	s.Regions = v
	return s
}

func (s *GetEdgeContainerAppResourceCapacityResponseBody) SetRequestId(v string) *GetEdgeContainerAppResourceCapacityResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEdgeContainerAppResourceCapacityResponseBody) Validate() error {
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

type GetEdgeContainerAppResourceCapacityResponseBodyRegions struct {
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
	// The number of container replicas that can be deployed.
	//
	// example:
	//
	// 16
	Replicas *int32 `json:"Replicas,omitempty" xml:"Replicas,omitempty"`
}

func (s GetEdgeContainerAppResourceCapacityResponseBodyRegions) String() string {
	return dara.Prettify(s)
}

func (s GetEdgeContainerAppResourceCapacityResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *GetEdgeContainerAppResourceCapacityResponseBodyRegions) GetIsp() *string {
	return s.Isp
}

func (s *GetEdgeContainerAppResourceCapacityResponseBodyRegions) GetRegion() *string {
	return s.Region
}

func (s *GetEdgeContainerAppResourceCapacityResponseBodyRegions) GetReplicas() *int32 {
	return s.Replicas
}

func (s *GetEdgeContainerAppResourceCapacityResponseBodyRegions) SetIsp(v string) *GetEdgeContainerAppResourceCapacityResponseBodyRegions {
	s.Isp = &v
	return s
}

func (s *GetEdgeContainerAppResourceCapacityResponseBodyRegions) SetRegion(v string) *GetEdgeContainerAppResourceCapacityResponseBodyRegions {
	s.Region = &v
	return s
}

func (s *GetEdgeContainerAppResourceCapacityResponseBodyRegions) SetReplicas(v int32) *GetEdgeContainerAppResourceCapacityResponseBodyRegions {
	s.Replicas = &v
	return s
}

func (s *GetEdgeContainerAppResourceCapacityResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}
