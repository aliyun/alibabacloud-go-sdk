// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssociateSmartAGWithApplicationBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationBandwidthPackageId(v string) *AssociateSmartAGWithApplicationBandwidthPackageRequest
	GetApplicationBandwidthPackageId() *string
	SetAssociateConfigs(v []*AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs) *AssociateSmartAGWithApplicationBandwidthPackageRequest
	GetAssociateConfigs() []*AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs
	SetClientToken(v string) *AssociateSmartAGWithApplicationBandwidthPackageRequest
	GetClientToken() *string
	SetDryRun(v bool) *AssociateSmartAGWithApplicationBandwidthPackageRequest
	GetDryRun() *bool
	SetRegionId(v string) *AssociateSmartAGWithApplicationBandwidthPackageRequest
	GetRegionId() *string
}

type AssociateSmartAGWithApplicationBandwidthPackageRequest struct {
	// The ID of the bandwidth plan for application acceleration.
	//
	// This parameter is required.
	//
	// example:
	//
	// abwp-7963l7iqnquyj3****
	ApplicationBandwidthPackageId *string `json:"ApplicationBandwidthPackageId,omitempty" xml:"ApplicationBandwidthPackageId,omitempty"`
	// The configuration of application acceleration.
	//
	// This parameter is required.
	AssociateConfigs []*AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs `json:"AssociateConfigs,omitempty" xml:"AssociateConfigs,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not set this parameter, ClientToken is set to the value of RequestId. The value of RequestId for each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to precheck the request. Check items include permissions and the status of the specified cloud resources. Valid values:
	//
	// 	- **false*	- (default): sends the request. After the request passes the check, the bandwidth plan for application acceleration is associated with the SAG instance.
	//
	// 	- **true**: prechecks the request but does not associate the bandwidth plan for application acceleration with the SAG instance. If you use this value, the system checks the required parameters and the request syntax. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the bandwidth plan for application acceleration.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s AssociateSmartAGWithApplicationBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s AssociateSmartAGWithApplicationBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequest) GetApplicationBandwidthPackageId() *string {
	return s.ApplicationBandwidthPackageId
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequest) GetAssociateConfigs() []*AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs {
	return s.AssociateConfigs
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequest) SetApplicationBandwidthPackageId(v string) *AssociateSmartAGWithApplicationBandwidthPackageRequest {
	s.ApplicationBandwidthPackageId = &v
	return s
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequest) SetAssociateConfigs(v []*AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs) *AssociateSmartAGWithApplicationBandwidthPackageRequest {
	s.AssociateConfigs = v
	return s
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequest) SetClientToken(v string) *AssociateSmartAGWithApplicationBandwidthPackageRequest {
	s.ClientToken = &v
	return s
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequest) SetDryRun(v bool) *AssociateSmartAGWithApplicationBandwidthPackageRequest {
	s.DryRun = &v
	return s
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequest) SetRegionId(v string) *AssociateSmartAGWithApplicationBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequest) Validate() error {
	if s.AssociateConfigs != nil {
		for _, item := range s.AssociateConfigs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs struct {
	// The maximum bandwidth value for application acceleration. Unit: Mbit/s.
	//
	// >  The maximum bandwidth value of each SAG instance cannot exceed that of the associated bandwidth plan for application acceleration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-gf5serujbhrn2j****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs) String() string {
	return dara.Prettify(s)
}

func (s AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs) GoString() string {
	return s.String()
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs) SetBandwidth(v int32) *AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs {
	s.Bandwidth = &v
	return s
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs) SetSmartAGId(v string) *AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs {
	s.SmartAGId = &v
	return s
}

func (s *AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs) Validate() error {
	return dara.Validate(s)
}
