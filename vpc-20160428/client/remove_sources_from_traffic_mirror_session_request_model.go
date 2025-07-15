// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSourcesFromTrafficMirrorSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RemoveSourcesFromTrafficMirrorSessionRequest
	GetClientToken() *string
	SetDryRun(v bool) *RemoveSourcesFromTrafficMirrorSessionRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *RemoveSourcesFromTrafficMirrorSessionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RemoveSourcesFromTrafficMirrorSessionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RemoveSourcesFromTrafficMirrorSessionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RemoveSourcesFromTrafficMirrorSessionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RemoveSourcesFromTrafficMirrorSessionRequest
	GetResourceOwnerId() *int64
	SetTrafficMirrorSessionId(v string) *RemoveSourcesFromTrafficMirrorSessionRequest
	GetTrafficMirrorSessionId() *string
	SetTrafficMirrorSourceIds(v []*string) *RemoveSourcesFromTrafficMirrorSessionRequest
	GetTrafficMirrorSourceIds() []*string
}

type RemoveSourcesFromTrafficMirrorSessionRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the value, but you must make sure that the value is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.
	//
	// >  If you do not set this parameter, the system uses **RequestId*	- as **ClientToken**. **RequestId*	- may be different for each API request.
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe63****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to check the request without performing the operation. Valid values:
	//
	// 	- **true**: checks the request without performing the operation. The system checks the required parameters, request format, and limits. If the request fails the check, an error message is returned. If the request passes the check, the `DryRunOperation` error code is returned.
	//
	// 	- **false*	- (default): sends the request. After the request passes the check, the operation is performed.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region to which the traffic mirror session belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list. For more information about regions that support traffic mirror, see [Overview of traffic mirror](https://help.aliyun.com/document_detail/207513.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hongkong
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the traffic mirror session from which you want to delete a traffic mirror source.
	//
	// This parameter is required.
	//
	// example:
	//
	// tms-j6cla50buc44ap8tu****
	TrafficMirrorSessionId *string `json:"TrafficMirrorSessionId,omitempty" xml:"TrafficMirrorSessionId,omitempty"`
	// The ID of the traffic mirror source to be deleted. Maximum value of N: 10.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-j6c8znm5l1yt4sox****
	TrafficMirrorSourceIds []*string `json:"TrafficMirrorSourceIds,omitempty" xml:"TrafficMirrorSourceIds,omitempty" type:"Repeated"`
}

func (s RemoveSourcesFromTrafficMirrorSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveSourcesFromTrafficMirrorSessionRequest) GoString() string {
	return s.String()
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) GetTrafficMirrorSessionId() *string {
	return s.TrafficMirrorSessionId
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) GetTrafficMirrorSourceIds() []*string {
	return s.TrafficMirrorSourceIds
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) SetClientToken(v string) *RemoveSourcesFromTrafficMirrorSessionRequest {
	s.ClientToken = &v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) SetDryRun(v bool) *RemoveSourcesFromTrafficMirrorSessionRequest {
	s.DryRun = &v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) SetOwnerAccount(v string) *RemoveSourcesFromTrafficMirrorSessionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) SetOwnerId(v int64) *RemoveSourcesFromTrafficMirrorSessionRequest {
	s.OwnerId = &v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) SetRegionId(v string) *RemoveSourcesFromTrafficMirrorSessionRequest {
	s.RegionId = &v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) SetResourceOwnerAccount(v string) *RemoveSourcesFromTrafficMirrorSessionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) SetResourceOwnerId(v int64) *RemoveSourcesFromTrafficMirrorSessionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) SetTrafficMirrorSessionId(v string) *RemoveSourcesFromTrafficMirrorSessionRequest {
	s.TrafficMirrorSessionId = &v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) SetTrafficMirrorSourceIds(v []*string) *RemoveSourcesFromTrafficMirrorSessionRequest {
	s.TrafficMirrorSourceIds = v
	return s
}

func (s *RemoveSourcesFromTrafficMirrorSessionRequest) Validate() error {
	return dara.Validate(s)
}
