// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachDdosToAcceleratorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceleratorId(v string) *AttachDdosToAcceleratorRequest
	GetAcceleratorId() *string
	SetDdosConfigList(v []*AttachDdosToAcceleratorRequestDdosConfigList) *AttachDdosToAcceleratorRequest
	GetDdosConfigList() []*AttachDdosToAcceleratorRequestDdosConfigList
	SetDdosId(v string) *AttachDdosToAcceleratorRequest
	GetDdosId() *string
	SetDdosRegionId(v string) *AttachDdosToAcceleratorRequest
	GetDdosRegionId() *string
	SetDryRun(v bool) *AttachDdosToAcceleratorRequest
	GetDryRun() *bool
	SetRegionId(v string) *AttachDdosToAcceleratorRequest
	GetRegionId() *string
}

type AttachDdosToAcceleratorRequest struct {
	// The ID of the Global Accelerator (GA) instance with which you want to associate the Anti-DDoS Pro or Anti-DDoS Premium instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// ga-bp1odcab8tmno0hdq****
	AcceleratorId *string `json:"AcceleratorId,omitempty" xml:"AcceleratorId,omitempty"`
	// The list of Anti-DDoS Pro or Anti-DDoS Premium instances to associate with the Global Accelerator (GA) instance.
	DdosConfigList []*AttachDdosToAcceleratorRequestDdosConfigList `json:"DdosConfigList,omitempty" xml:"DdosConfigList,omitempty" type:"Repeated"`
	// Deprecated
	//
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance to associate with the Global Accelerator (GA) instance.
	//
	// example:
	//
	// ddoscoo-cn-zz11vq7j****
	DdosId *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	// Deprecated
	//
	// The region of the Anti-DDoS Pro or Anti-DDoS Premium instance. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// - **true**: performs a dry run without actually associating the instances. The system checks the required parameters, request syntax, and business limits. If the check fails, the corresponding error is returned. If the check passes, an HTTP 2xx status code is returned.
	//
	// - **false*	- (default): sends the request. After the request passes the check, an HTTP 2xx status code is returned and the operation is performed.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the Global Accelerator (GA) instance. Set the value to **cn-hangzhou**.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AttachDdosToAcceleratorRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachDdosToAcceleratorRequest) GoString() string {
	return s.String()
}

func (s *AttachDdosToAcceleratorRequest) GetAcceleratorId() *string {
	return s.AcceleratorId
}

func (s *AttachDdosToAcceleratorRequest) GetDdosConfigList() []*AttachDdosToAcceleratorRequestDdosConfigList {
	return s.DdosConfigList
}

func (s *AttachDdosToAcceleratorRequest) GetDdosId() *string {
	return s.DdosId
}

func (s *AttachDdosToAcceleratorRequest) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *AttachDdosToAcceleratorRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AttachDdosToAcceleratorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AttachDdosToAcceleratorRequest) SetAcceleratorId(v string) *AttachDdosToAcceleratorRequest {
	s.AcceleratorId = &v
	return s
}

func (s *AttachDdosToAcceleratorRequest) SetDdosConfigList(v []*AttachDdosToAcceleratorRequestDdosConfigList) *AttachDdosToAcceleratorRequest {
	s.DdosConfigList = v
	return s
}

func (s *AttachDdosToAcceleratorRequest) SetDdosId(v string) *AttachDdosToAcceleratorRequest {
	s.DdosId = &v
	return s
}

func (s *AttachDdosToAcceleratorRequest) SetDdosRegionId(v string) *AttachDdosToAcceleratorRequest {
	s.DdosRegionId = &v
	return s
}

func (s *AttachDdosToAcceleratorRequest) SetDryRun(v bool) *AttachDdosToAcceleratorRequest {
	s.DryRun = &v
	return s
}

func (s *AttachDdosToAcceleratorRequest) SetRegionId(v string) *AttachDdosToAcceleratorRequest {
	s.RegionId = &v
	return s
}

func (s *AttachDdosToAcceleratorRequest) Validate() error {
	if s.DdosConfigList != nil {
		for _, item := range s.DdosConfigList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AttachDdosToAcceleratorRequestDdosConfigList struct {
	// The ID of the Anti-DDoS Pro or Anti-DDoS Premium instance to associate with the Global Accelerator (GA) instance.
	//
	// example:
	//
	// ddoscoo-cn-zz11vq7j****
	DdosId *string `json:"DdosId,omitempty" xml:"DdosId,omitempty"`
	// The region of the Anti-DDoS Pro or Anti-DDoS Premium instance. Valid values:
	//
	// - **cn-hangzhou**: the Chinese mainland.
	//
	// - **ap-southeast-1**: outside the Chinese mainland.
	//
	// example:
	//
	// cn-hangzhou
	DdosRegionId *string `json:"DdosRegionId,omitempty" xml:"DdosRegionId,omitempty"`
}

func (s AttachDdosToAcceleratorRequestDdosConfigList) String() string {
	return dara.Prettify(s)
}

func (s AttachDdosToAcceleratorRequestDdosConfigList) GoString() string {
	return s.String()
}

func (s *AttachDdosToAcceleratorRequestDdosConfigList) GetDdosId() *string {
	return s.DdosId
}

func (s *AttachDdosToAcceleratorRequestDdosConfigList) GetDdosRegionId() *string {
	return s.DdosRegionId
}

func (s *AttachDdosToAcceleratorRequestDdosConfigList) SetDdosId(v string) *AttachDdosToAcceleratorRequestDdosConfigList {
	s.DdosId = &v
	return s
}

func (s *AttachDdosToAcceleratorRequestDdosConfigList) SetDdosRegionId(v string) *AttachDdosToAcceleratorRequestDdosConfigList {
	s.DdosRegionId = &v
	return s
}

func (s *AttachDdosToAcceleratorRequestDdosConfigList) Validate() error {
	return dara.Validate(s)
}
