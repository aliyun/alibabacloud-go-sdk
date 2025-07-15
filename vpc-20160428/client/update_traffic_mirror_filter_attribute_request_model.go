// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrafficMirrorFilterAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateTrafficMirrorFilterAttributeRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateTrafficMirrorFilterAttributeRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *UpdateTrafficMirrorFilterAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UpdateTrafficMirrorFilterAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UpdateTrafficMirrorFilterAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UpdateTrafficMirrorFilterAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UpdateTrafficMirrorFilterAttributeRequest
	GetResourceOwnerId() *int64
	SetTrafficMirrorFilterDescription(v string) *UpdateTrafficMirrorFilterAttributeRequest
	GetTrafficMirrorFilterDescription() *string
	SetTrafficMirrorFilterId(v string) *UpdateTrafficMirrorFilterAttributeRequest
	GetTrafficMirrorFilterId() *string
	SetTrafficMirrorFilterName(v string) *UpdateTrafficMirrorFilterAttributeRequest
	GetTrafficMirrorFilterName() *string
}

type UpdateTrafficMirrorFilterAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that it is unique among different requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// 	- **true**: checks the request without performing the operation. The system checks the required parameters, request syntax, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. If the request passes the check, a 2xx HTTP status code is returned and the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the mirrored traffic belongs.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list. For more information about regions that support traffic mirror, see [Overview of traffic mirror](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The description of the filter.
	//
	// The description must be 1 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is a new filter.
	TrafficMirrorFilterDescription *string `json:"TrafficMirrorFilterDescription,omitempty" xml:"TrafficMirrorFilterDescription,omitempty"`
	// The ID of the filter.
	//
	// This parameter is required.
	//
	// example:
	//
	// tmf-j6cmls82xnc86vtpe****
	TrafficMirrorFilterId *string `json:"TrafficMirrorFilterId,omitempty" xml:"TrafficMirrorFilterId,omitempty"`
	// The name of the filter.
	//
	// The name must be 1 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// test
	TrafficMirrorFilterName *string `json:"TrafficMirrorFilterName,omitempty" xml:"TrafficMirrorFilterName,omitempty"`
}

func (s UpdateTrafficMirrorFilterAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrafficMirrorFilterAttributeRequest) GoString() string {
	return s.String()
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) GetTrafficMirrorFilterDescription() *string {
	return s.TrafficMirrorFilterDescription
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) GetTrafficMirrorFilterId() *string {
	return s.TrafficMirrorFilterId
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) GetTrafficMirrorFilterName() *string {
	return s.TrafficMirrorFilterName
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) SetClientToken(v string) *UpdateTrafficMirrorFilterAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) SetDryRun(v bool) *UpdateTrafficMirrorFilterAttributeRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) SetOwnerAccount(v string) *UpdateTrafficMirrorFilterAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) SetOwnerId(v int64) *UpdateTrafficMirrorFilterAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) SetRegionId(v string) *UpdateTrafficMirrorFilterAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) SetResourceOwnerAccount(v string) *UpdateTrafficMirrorFilterAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) SetResourceOwnerId(v int64) *UpdateTrafficMirrorFilterAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) SetTrafficMirrorFilterDescription(v string) *UpdateTrafficMirrorFilterAttributeRequest {
	s.TrafficMirrorFilterDescription = &v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) SetTrafficMirrorFilterId(v string) *UpdateTrafficMirrorFilterAttributeRequest {
	s.TrafficMirrorFilterId = &v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) SetTrafficMirrorFilterName(v string) *UpdateTrafficMirrorFilterAttributeRequest {
	s.TrafficMirrorFilterName = &v
	return s
}

func (s *UpdateTrafficMirrorFilterAttributeRequest) Validate() error {
	return dara.Validate(s)
}
