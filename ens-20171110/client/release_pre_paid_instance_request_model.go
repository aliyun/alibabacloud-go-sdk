// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleasePrePaidInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReleasePrePaidInstanceRequest
	GetInstanceId() *string
}

type ReleasePrePaidInstanceRequest struct {
	// The ID of the instance to be deleted. You can specify only one instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-instance****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleasePrePaidInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleasePrePaidInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleasePrePaidInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleasePrePaidInstanceRequest) SetInstanceId(v string) *ReleasePrePaidInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleasePrePaidInstanceRequest) Validate() error {
	return dara.Validate(s)
}
