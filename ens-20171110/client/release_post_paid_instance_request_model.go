// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePostPaidInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReleasePostPaidInstanceRequest
	GetInstanceId() *string
}

type ReleasePostPaidInstanceRequest struct {
	// The ID of the instance to be deleted. You can specify only one instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-instance****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleasePostPaidInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleasePostPaidInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleasePostPaidInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleasePostPaidInstanceRequest) SetInstanceId(v string) *ReleasePostPaidInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleasePostPaidInstanceRequest) Validate() error {
	return dara.Validate(s)
}
