// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVulTargetConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeVulTargetConfigResponseBody
	GetRequestId() *string
	SetTargetConfigs(v []*DescribeVulTargetConfigResponseBodyTargetConfigs) *DescribeVulTargetConfigResponseBody
	GetTargetConfigs() []*DescribeVulTargetConfigResponseBodyTargetConfigs
	SetTotalCount(v int32) *DescribeVulTargetConfigResponseBody
	GetTotalCount() *int32
}

type DescribeVulTargetConfigResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9FBC6E47-7508-58C9-9E76-528E118CB1CC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// An array that consists of the configurations.
	TargetConfigs []*DescribeVulTargetConfigResponseBodyTargetConfigs `json:"TargetConfigs,omitempty" xml:"TargetConfigs,omitempty" type:"Repeated"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeVulTargetConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulTargetConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVulTargetConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVulTargetConfigResponseBody) GetTargetConfigs() []*DescribeVulTargetConfigResponseBodyTargetConfigs {
	return s.TargetConfigs
}

func (s *DescribeVulTargetConfigResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVulTargetConfigResponseBody) SetRequestId(v string) *DescribeVulTargetConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVulTargetConfigResponseBody) SetTargetConfigs(v []*DescribeVulTargetConfigResponseBodyTargetConfigs) *DescribeVulTargetConfigResponseBody {
	s.TargetConfigs = v
	return s
}

func (s *DescribeVulTargetConfigResponseBody) SetTotalCount(v int32) *DescribeVulTargetConfigResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVulTargetConfigResponseBody) Validate() error {
	if s.TargetConfigs != nil {
		for _, item := range s.TargetConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVulTargetConfigResponseBodyTargetConfigs struct {
	// Indicates whether the vulnerability scan feature is enabled for the server.
	//
	// 	- **off**: disabled
	//
	// 	- **on**: enabled
	//
	// example:
	//
	// on
	Config *string `json:"Config,omitempty" xml:"Config,omitempty"`
	// Indicates whether the vulnerability scan feature is enabled for all servers. Valid values:
	//
	// 	- **off**: disabled
	//
	// 	- **on**: enabled
	//
	// example:
	//
	// on
	OverAllConfig *string `json:"OverAllConfig,omitempty" xml:"OverAllConfig,omitempty"`
	// The type of the vulnerability. Valid values:
	//
	// 	- **cve**: Linux software vulnerability
	//
	// 	- **sys**: Windows system vulnerability
	//
	// 	- **cms**: Web-CMS vulnerability
	//
	// 	- **emg**: urgent vulnerability
	//
	// example:
	//
	// cve
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeVulTargetConfigResponseBodyTargetConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVulTargetConfigResponseBodyTargetConfigs) GoString() string {
	return s.String()
}

func (s *DescribeVulTargetConfigResponseBodyTargetConfigs) GetConfig() *string {
	return s.Config
}

func (s *DescribeVulTargetConfigResponseBodyTargetConfigs) GetOverAllConfig() *string {
	return s.OverAllConfig
}

func (s *DescribeVulTargetConfigResponseBodyTargetConfigs) GetType() *string {
	return s.Type
}

func (s *DescribeVulTargetConfigResponseBodyTargetConfigs) SetConfig(v string) *DescribeVulTargetConfigResponseBodyTargetConfigs {
	s.Config = &v
	return s
}

func (s *DescribeVulTargetConfigResponseBodyTargetConfigs) SetOverAllConfig(v string) *DescribeVulTargetConfigResponseBodyTargetConfigs {
	s.OverAllConfig = &v
	return s
}

func (s *DescribeVulTargetConfigResponseBodyTargetConfigs) SetType(v string) *DescribeVulTargetConfigResponseBodyTargetConfigs {
	s.Type = &v
	return s
}

func (s *DescribeVulTargetConfigResponseBodyTargetConfigs) Validate() error {
	return dara.Validate(s)
}
