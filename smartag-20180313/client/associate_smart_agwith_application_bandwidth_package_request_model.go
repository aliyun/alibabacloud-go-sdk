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
	// The instance ID of the bandwidth plan for application acceleration.
	//
	// This parameter is required.
	//
	// example:
	//
	// abwp-7963l7iqnquyj3****
	ApplicationBandwidthPackageId *string `json:"ApplicationBandwidthPackageId,omitempty" xml:"ApplicationBandwidthPackageId,omitempty"`
	// The list of Smart Access Gateway instances to be associated.
	//
	// This parameter is required.
	AssociateConfigs []*AssociateSmartAGWithApplicationBandwidthPackageRequestAssociateConfigs `json:"AssociateConfigs,omitempty" xml:"AssociateConfigs,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the association request. The dry run includes permission and instance status checks. Valid values:
	//
	// - **false*	- (default): Sends a normal request and associates the bandwidth plan for application acceleration with the Smart Access Gateway instance after the request passes the check.
	//
	// - **true**: Sends a check request and performs only the check without associating the bandwidth plan for application acceleration with the Smart Access Gateway instance. The check items include whether required parameters are specified and whether the request format is valid. If the request fails the check, an error is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region to which the bandwidth plan for application acceleration belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
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
	// The peak bandwidth for application acceleration of the Smart Access Gateway instance. Unit: Mbit/s.
	//
	// > The peak bandwidth for application acceleration of each Smart Access Gateway instance cannot exceed the peak bandwidth of the associated bandwidth plan for application acceleration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
	// The Smart Access Gateway instance ID.
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
