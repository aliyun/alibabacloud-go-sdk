// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootRCInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *RebootRCInstanceRequest
	GetDryRun() *bool
	SetForceStop(v bool) *RebootRCInstanceRequest
	GetForceStop() *bool
	SetInstanceId(v string) *RebootRCInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *RebootRCInstanceRequest
	GetRegionId() *string
}

type RebootRCInstanceRequest struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- **true**: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and insufficient inventory errors. If the request passes the dry run, the DryRunOperation error code is returned. Otherwise, an error message is returned.
	//
	// 	- **false**: performs a dry run and performs the actual request. If the request passes the dry run, the instance is restarted.
	//
	// Default value: false
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully stop the instance before you restart the instance Valid values:
	//
	// 	- **true**: forcefully stops the instance. This operation is equivalent to the typical power-off operation. Cache data that is not written to storage devices on the instance is lost.
	//
	// 	- **false*	- (default): normally stops the instance.
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// rc-m5sc1271fv344a1r****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootRCInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootRCInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootRCInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RebootRCInstanceRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *RebootRCInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RebootRCInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebootRCInstanceRequest) SetDryRun(v bool) *RebootRCInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *RebootRCInstanceRequest) SetForceStop(v bool) *RebootRCInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *RebootRCInstanceRequest) SetInstanceId(v string) *RebootRCInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootRCInstanceRequest) SetRegionId(v string) *RebootRCInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RebootRCInstanceRequest) Validate() error {
	return dara.Validate(s)
}
