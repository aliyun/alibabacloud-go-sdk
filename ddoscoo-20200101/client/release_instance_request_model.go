// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReleaseInstanceRequest
	GetInstanceId() *string
}

type ReleaseInstanceRequest struct {
	// The ID of the instance to be released.
	//
	// > Only expired instances can be released. You can call [DescribeInstances](https://help.aliyun.com/document_detail/91478.html) to query the IDs and expiration status of all Anti-DDoS Proxy instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// ddoscoo-cn-mp91j1ao****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseInstanceRequest) SetInstanceId(v string) *ReleaseInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseInstanceRequest) Validate() error {
	return dara.Validate(s)
}
