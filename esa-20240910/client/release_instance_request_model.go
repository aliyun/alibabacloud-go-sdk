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
	// The ID of the security instance. This must be a DDoS security instance ID in the format of esa-ddos-*. You can call the ListDDoSInstances operation to obtain the ID. Site instance IDs in the format of esa-site-	- are not supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// esa-ddos-2sdj**3s
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
