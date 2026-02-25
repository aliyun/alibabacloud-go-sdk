// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableSSLConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DisableSSLConnectionRequest
	GetInstanceId() *string
}

type DisableSSLConnectionRequest struct {
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DisableSSLConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableSSLConnectionRequest) GoString() string {
	return s.String()
}

func (s *DisableSSLConnectionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DisableSSLConnectionRequest) SetInstanceId(v string) *DisableSSLConnectionRequest {
	s.InstanceId = &v
	return s
}

func (s *DisableSSLConnectionRequest) Validate() error {
	return dara.Validate(s)
}
