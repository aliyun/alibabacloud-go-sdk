// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListIspTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIspTypeList(v []*string) *ListIspTypesResponseBody
	GetIspTypeList() []*string
	SetRequestId(v string) *ListIspTypesResponseBody
	GetRequestId() *string
}

type ListIspTypesResponseBody struct {
	// The line types of EIPs in the acceleration region.
	//
	// 	- **BGP*	- (default): BGP (Multi-ISP) lines
	//
	// 	- **BGP_PRO**: BGP (Multi-ISP) Pro lines
	//
	// If you have the permissions to use single-ISP bandwidth, one of the following values may be returned:
	//
	// 	- **ChinaTelecom**: China Telecom (single ISP)
	//
	// 	- **ChinaUnicom**: China Unicom (single ISP)
	//
	// 	- **ChinaMobile**: China Mobile (single ISP)
	//
	// 	- **ChinaTelecom_L2**: China Telecom_L2 (single ISP)
	//
	// 	- **ChinaUnicom_L2**: China Unicom_L2 (single ISP)
	//
	// 	- **ChinaMobile_L2**: China Mobile_L2 (single ISP)
	//
	// > Different acceleration regions support different single-ISP BGP lines.
	IspTypeList []*string `json:"IspTypeList,omitempty" xml:"IspTypeList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// F591955F-5CB5-4CCE-A75D-17CF2085CE22
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListIspTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListIspTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListIspTypesResponseBody) GetIspTypeList() []*string {
	return s.IspTypeList
}

func (s *ListIspTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListIspTypesResponseBody) SetIspTypeList(v []*string) *ListIspTypesResponseBody {
	s.IspTypeList = v
	return s
}

func (s *ListIspTypesResponseBody) SetRequestId(v string) *ListIspTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListIspTypesResponseBody) Validate() error {
	return dara.Validate(s)
}
