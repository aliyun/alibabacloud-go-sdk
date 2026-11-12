// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StopInstancesRequest
	GetClientToken() *string
	SetForceStop(v bool) *StopInstancesRequest
	GetForceStop() *bool
	SetInstanceIds(v string) *StopInstancesRequest
	GetInstanceIds() *string
	SetRegionId(v string) *StopInstancesRequest
	GetRegionId() *string
}

type StopInstancesRequest struct {
	// Ensures the idempotence of the request. You can generate the parameter value on your client and ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to forcibly stop the instance.
	//
	// - **true**: forcibly stops the instance.
	//
	// - **false*	- (default): normally stops the instance.
	//
	// example:
	//
	// true
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The instance IDs of the simple application servers. The value can be a JSON array that consists of up to 100 instance IDs. Separate multiple instance IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application server instance. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the supported Alibaba Cloud regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s StopInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstancesRequest) GoString() string {
	return s.String()
}

func (s *StopInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopInstancesRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *StopInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *StopInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopInstancesRequest) SetClientToken(v string) *StopInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *StopInstancesRequest) SetForceStop(v bool) *StopInstancesRequest {
	s.ForceStop = &v
	return s
}

func (s *StopInstancesRequest) SetInstanceIds(v string) *StopInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *StopInstancesRequest) SetRegionId(v string) *StopInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *StopInstancesRequest) Validate() error {
	return dara.Validate(s)
}
