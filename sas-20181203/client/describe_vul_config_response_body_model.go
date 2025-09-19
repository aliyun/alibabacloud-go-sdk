// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVulConfigResponseBody
	GetRequestId() *string
	SetTargetConfigs(v []*DescribeVulConfigResponseBodyTargetConfigs) *DescribeVulConfigResponseBody
	GetTargetConfigs() []*DescribeVulConfigResponseBodyTargetConfigs
	SetTotalCount(v int32) *DescribeVulConfigResponseBody
	GetTotalCount() *int32
}

type DescribeVulConfigResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 24A20733-10A0-4AF6-BE6B-E3322413BB68
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the configurations of vulnerability management.
	TargetConfigs []*DescribeVulConfigResponseBodyTargetConfigs `json:"TargetConfigs,omitempty" xml:"TargetConfigs,omitempty" type:"Repeated"`
	// The total number of configurations.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVulConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulConfigResponseBody) GetTargetConfigs() []*DescribeVulConfigResponseBodyTargetConfigs {
	return s.TargetConfigs
}

func (s *DescribeVulConfigResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVulConfigResponseBody) SetRequestId(v string) *DescribeVulConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulConfigResponseBody) SetTargetConfigs(v []*DescribeVulConfigResponseBodyTargetConfigs) *DescribeVulConfigResponseBody {
	s.TargetConfigs = v
	return s
}

func (s *DescribeVulConfigResponseBody) SetTotalCount(v int32) *DescribeVulConfigResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVulConfigResponseBodyTargetConfigs struct {
	// The configuration of vulnerability scan.
	//
	// > Valid values when you set the Type parameter to **cve**, **sys**, **cms**, **app**, **emg**, or **yum**:
	//
	// 	- **on**: enabled
	//
	// 	- **off**: disabled
	//
	// Valid values when you set the Type parameter to **scanMode**:
	//
	// 	- **real**: displays easily exploitable vulnerability.
	//
	// 	- **all**: displays all vulnerabilities.
	//
	// When you set the Type parameter to **imageVulClean**, the value of this parameter indicates the vulnerability retention period in days.
	//
	// example:
	//
	// 90
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Indicates whether the vulnerability management feature is enabled for all servers. Valid values:
	//
	// 	- **off**: disabled
	//
	// 	- **on**: enabled
	//
	// example:
	//
	// on
	OverAllConfig *string `json:"OverAllConfig,omitempty" xml:"OverAllConfig,omitempty"`
	// The type of configuration. Valid values:
	//
	// 	- **cve**: Linux software vulnerability.
	//
	// 	- **sys**: Windows system vulnerability.
	//
	// 	- **cms**: Web-CMS vulnerability.
	//
	// 	- **app**: application vulnerability that is detected by using web scanner.
	//
	// 	- **emg**: urgent vulnerability.
	//
	// 	- **scanMode**: displays easily exploitable vulnerability.
	//
	// 	- **imageVulClean**: vulnerability retention duration.
	//
	// 	- **yum**: preferentially uses YUM or APT sources of Alibaba Cloud to fix vulnerabilities.
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVulConfigResponseBodyTargetConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulConfigResponseBodyTargetConfigs) GoString() string {
	return s.String()
}

func (s *DescribeVulConfigResponseBodyTargetConfigs) GetConfig() *string {
	return s.Config
}

func (s *DescribeVulConfigResponseBodyTargetConfigs) GetOverAllConfig() *string {
	return s.OverAllConfig
}

func (s *DescribeVulConfigResponseBodyTargetConfigs) GetType() *string {
	return s.Type
}

func (s *DescribeVulConfigResponseBodyTargetConfigs) SetConfig(v string) *DescribeVulConfigResponseBodyTargetConfigs {
	s.Config = &v
	return s
}

func (s *DescribeVulConfigResponseBodyTargetConfigs) SetOverAllConfig(v string) *DescribeVulConfigResponseBodyTargetConfigs {
	s.OverAllConfig = &v
	return s
}

func (s *DescribeVulConfigResponseBodyTargetConfigs) SetType(v string) *DescribeVulConfigResponseBodyTargetConfigs {
	s.Type = &v
	return s
}

func (s *DescribeVulConfigResponseBodyTargetConfigs) Validate() error {
	return dara.Validate(s)
}
