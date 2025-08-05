// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePostInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReleasePostInstanceRequest
	GetInstanceId() *string
}

type ReleasePostInstanceRequest struct {
	// The ID of the Cloud Firewall instance.
	//
	// example:
	//
	// cfw_elasticity_public_cn-xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleasePostInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleasePostInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleasePostInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleasePostInstanceRequest) SetInstanceId(v string) *ReleasePostInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleasePostInstanceRequest) Validate() error {
	return dara.Validate(s)
}
