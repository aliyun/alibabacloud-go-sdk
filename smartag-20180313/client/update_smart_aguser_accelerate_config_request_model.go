// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSmartAGUserAccelerateConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidth(v int32) *UpdateSmartAGUserAccelerateConfigRequest
	GetBandwidth() *int32
	SetClientToken(v string) *UpdateSmartAGUserAccelerateConfigRequest
	GetClientToken() *string
	SetDryRun(v bool) *UpdateSmartAGUserAccelerateConfigRequest
	GetDryRun() *bool
	SetRegionId(v string) *UpdateSmartAGUserAccelerateConfigRequest
	GetRegionId() *string
	SetSmartAGId(v string) *UpdateSmartAGUserAccelerateConfigRequest
	GetSmartAGId() *string
	SetUserName(v string) *UpdateSmartAGUserAccelerateConfigRequest
	GetUserName() *string
}

type UpdateSmartAGUserAccelerateConfigRequest struct {
	// The maximum bandwidth value for the client. Unit: Kbit/s.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Bandwidth *int32 `json:"Bandwidth,omitempty" xml:"Bandwidth,omitempty"`
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
	// Specifies whether to perform a precheck to check information such as the permissions and instance status. Valid values:
	//
	// 	- **false*	- (default): sends the request. After the request passes the check, the maximum bandwidth value of the client is modified.
	//
	// 	- **true**: prechecks the request but does not modify the maximum bandwidth value of the client. If you use this value, the system checks the required parameters and the request syntax. If the request fails the precheck, an error message is returned. If the request passes the precheck, the `DryRunOperation` error code is returned.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the region where the SAG app instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the SAG app instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-w11hk4bqxpakem****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The username of the client account.
	//
	// This parameter is required.
	//
	// example:
	//
	// username@example.com
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s UpdateSmartAGUserAccelerateConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateSmartAGUserAccelerateConfigRequest) GoString() string {
	return s.String()
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) GetBandwidth() *int32 {
	return s.Bandwidth
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) GetUserName() *string {
	return s.UserName
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) SetBandwidth(v int32) *UpdateSmartAGUserAccelerateConfigRequest {
	s.Bandwidth = &v
	return s
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) SetClientToken(v string) *UpdateSmartAGUserAccelerateConfigRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) SetDryRun(v bool) *UpdateSmartAGUserAccelerateConfigRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) SetRegionId(v string) *UpdateSmartAGUserAccelerateConfigRequest {
	s.RegionId = &v
	return s
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) SetSmartAGId(v string) *UpdateSmartAGUserAccelerateConfigRequest {
	s.SmartAGId = &v
	return s
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) SetUserName(v string) *UpdateSmartAGUserAccelerateConfigRequest {
	s.UserName = &v
	return s
}

func (s *UpdateSmartAGUserAccelerateConfigRequest) Validate() error {
	return dara.Validate(s)
}
