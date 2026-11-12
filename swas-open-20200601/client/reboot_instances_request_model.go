// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RebootInstancesRequest
	GetClientToken() *string
	SetForceReboot(v bool) *RebootInstancesRequest
	GetForceReboot() *bool
	SetInstanceIds(v string) *RebootInstancesRequest
	GetInstanceIds() *string
	SetRegionId(v string) *RebootInstancesRequest
	GetRegionId() *string
}

type RebootInstancesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to forcibly restart the servers. Valid values:
	//
	// - true: forcibly restarts the servers. This operation is equivalent to the typical power-off operation. Cache data that is not written to storage devices on the server will be lost.
	//
	// - false: normally restarts the instance.
	//
	// Default value: false
	//
	// example:
	//
	// false
	ForceReboot *bool `json:"ForceReboot,omitempty" xml:"ForceReboot,omitempty"`
	// The IDs of the simple application servers. The value can be a JSON array that consists of up to 100 simple application server IDs. Separate multiple server IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// ["2ad1ae67295445f598017499dc****", "2ad1ae67295445f598017123dc****"]
	InstanceIds *string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty"`
	// The region ID of the simple application servers. You can call the [ListRegions](https://help.aliyun.com/document_detail/189315.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootInstancesRequest) GoString() string {
	return s.String()
}

func (s *RebootInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RebootInstancesRequest) GetForceReboot() *bool {
	return s.ForceReboot
}

func (s *RebootInstancesRequest) GetInstanceIds() *string {
	return s.InstanceIds
}

func (s *RebootInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebootInstancesRequest) SetClientToken(v string) *RebootInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *RebootInstancesRequest) SetForceReboot(v bool) *RebootInstancesRequest {
	s.ForceReboot = &v
	return s
}

func (s *RebootInstancesRequest) SetInstanceIds(v string) *RebootInstancesRequest {
	s.InstanceIds = &v
	return s
}

func (s *RebootInstancesRequest) SetRegionId(v string) *RebootInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RebootInstancesRequest) Validate() error {
	return dara.Validate(s)
}
