// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceUsageRequest
	GetInstanceId() *string
}

type GetInstanceUsageRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cri-xkx6vujuhay0****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceUsageRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceUsageRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceUsageRequest) SetInstanceId(v string) *GetInstanceUsageRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceUsageRequest) Validate() error {
	return dara.Validate(s)
}
