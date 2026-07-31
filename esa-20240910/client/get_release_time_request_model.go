// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetReleaseTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetReleaseTimeRequest
	GetInstanceId() *string
}

type GetReleaseTimeRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-ddos-2sdj**3s
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetReleaseTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetReleaseTimeRequest) GoString() string {
	return s.String()
}

func (s *GetReleaseTimeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetReleaseTimeRequest) SetInstanceId(v string) *GetReleaseTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *GetReleaseTimeRequest) Validate() error {
	return dara.Validate(s)
}
