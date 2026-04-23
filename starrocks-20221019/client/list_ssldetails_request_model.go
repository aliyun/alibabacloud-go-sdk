// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSSLDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListSSLDetailsRequest
	GetInstanceId() *string
}

type ListSSLDetailsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// c-b25e21e24388****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListSSLDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSSLDetailsRequest) GoString() string {
	return s.String()
}

func (s *ListSSLDetailsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSSLDetailsRequest) SetInstanceId(v string) *ListSSLDetailsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSSLDetailsRequest) Validate() error {
	return dara.Validate(s)
}
