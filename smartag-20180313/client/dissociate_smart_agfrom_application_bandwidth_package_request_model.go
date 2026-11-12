// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDissociateSmartAGFromApplicationBandwidthPackageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationBandwidthPackageId(v string) *DissociateSmartAGFromApplicationBandwidthPackageRequest
	GetApplicationBandwidthPackageId() *string
	SetClientToken(v string) *DissociateSmartAGFromApplicationBandwidthPackageRequest
	GetClientToken() *string
	SetDryRun(v bool) *DissociateSmartAGFromApplicationBandwidthPackageRequest
	GetDryRun() *bool
	SetRegionId(v string) *DissociateSmartAGFromApplicationBandwidthPackageRequest
	GetRegionId() *string
	SetSmartAGId(v string) *DissociateSmartAGFromApplicationBandwidthPackageRequest
	GetSmartAGId() *string
	SetSmartAGIdList(v []*string) *DissociateSmartAGFromApplicationBandwidthPackageRequest
	GetSmartAGIdList() []*string
}

type DissociateSmartAGFromApplicationBandwidthPackageRequest struct {
	// The instance ID of the bandwidth plan for application acceleration.
	//
	// This parameter is required.
	//
	// example:
	//
	// abwp-7963l7iqnquyj3****
	ApplicationBandwidthPackageId *string `json:"ApplicationBandwidthPackageId,omitempty" xml:"ApplicationBandwidthPackageId,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, and you must make sure that the value is unique among different requests. The client token can contain only ASCII characters.
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the client token. The RequestId of each API request may be different.
	//
	// example:
	//
	// 02fb3da4-130e-11e9-8e44-001****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run on the disassociation request. The dry run includes checks on permissions and instance status. Valid values:
	//
	// - **false*	- (default): Sends a normal request. If the request passes the check, the bandwidth plan for application acceleration is disassociated from the Smart Access Gateway instance.
	//
	// - **true**: Sends a dry run request. Only the check is performed, and the bandwidth plan for application acceleration is not disassociated from the Smart Access Gateway instance. Check items include whether required parameters are specified and the request format. If the check fails, the corresponding error is returned. If the check passes, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The region ID of the bandwidth plan for application acceleration.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query region IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The instance ID of the Smart Access Gateway.
	//
	// example:
	//
	// sag-w11hk4bqxpakem****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The list of Smart Access Gateway instance IDs.
	//
	// You can enter up to 100 Smart Access Gateway instance IDs at a time.
	SmartAGIdList []*string `json:"SmartAGIdList,omitempty" xml:"SmartAGIdList,omitempty" type:"Repeated"`
}

func (s DissociateSmartAGFromApplicationBandwidthPackageRequest) String() string {
	return dara.Prettify(s)
}

func (s DissociateSmartAGFromApplicationBandwidthPackageRequest) GoString() string {
	return s.String()
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) GetApplicationBandwidthPackageId() *string {
	return s.ApplicationBandwidthPackageId
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) GetSmartAGIdList() []*string {
	return s.SmartAGIdList
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) SetApplicationBandwidthPackageId(v string) *DissociateSmartAGFromApplicationBandwidthPackageRequest {
	s.ApplicationBandwidthPackageId = &v
	return s
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) SetClientToken(v string) *DissociateSmartAGFromApplicationBandwidthPackageRequest {
	s.ClientToken = &v
	return s
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) SetDryRun(v bool) *DissociateSmartAGFromApplicationBandwidthPackageRequest {
	s.DryRun = &v
	return s
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) SetRegionId(v string) *DissociateSmartAGFromApplicationBandwidthPackageRequest {
	s.RegionId = &v
	return s
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) SetSmartAGId(v string) *DissociateSmartAGFromApplicationBandwidthPackageRequest {
	s.SmartAGId = &v
	return s
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) SetSmartAGIdList(v []*string) *DissociateSmartAGFromApplicationBandwidthPackageRequest {
	s.SmartAGIdList = v
	return s
}

func (s *DissociateSmartAGFromApplicationBandwidthPackageRequest) Validate() error {
	return dara.Validate(s)
}
