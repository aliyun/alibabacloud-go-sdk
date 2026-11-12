// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RebootInstanceRequest
	GetClientToken() *string
	SetInstanceId(v string) *RebootInstanceRequest
	GetInstanceId() *string
	SetRegionId(v string) *RebootInstanceRequest
	GetRegionId() *string
}

type RebootInstanceRequest struct {
	// Ensures the idempotence of the request. Generate a parameter value from your client to ensure that the value is unique across different requests. The ClientToken parameter supports only ASCII characters and can be up to 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The instance ID of the simple application server.
	//
	// This parameter is required.
	//
	// example:
	//
	// ace0706b2ac4454d984295a94213****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s RebootInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RebootInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RebootInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RebootInstanceRequest) SetClientToken(v string) *RebootInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *RebootInstanceRequest) SetInstanceId(v string) *RebootInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootInstanceRequest) SetRegionId(v string) *RebootInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *RebootInstanceRequest) Validate() error {
	return dara.Validate(s)
}
