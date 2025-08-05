// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseExpiredInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ReleaseExpiredInstanceRequest
	GetInstanceId() *string
}

type ReleaseExpiredInstanceRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// cfw-y3gpqq705****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ReleaseExpiredInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ReleaseExpiredInstanceRequest) GoString() string {
	return s.String()
}

func (s *ReleaseExpiredInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ReleaseExpiredInstanceRequest) SetInstanceId(v string) *ReleaseExpiredInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ReleaseExpiredInstanceRequest) Validate() error {
	return dara.Validate(s)
}
