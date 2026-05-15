// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMcuLvsIpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetMcuLvsIpRequest
	GetInstanceId() *string
}

type GetMcuLvsIpRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc_xp_pre***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetMcuLvsIpRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMcuLvsIpRequest) GoString() string {
	return s.String()
}

func (s *GetMcuLvsIpRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetMcuLvsIpRequest) SetInstanceId(v string) *GetMcuLvsIpRequest {
	s.InstanceId = &v
	return s
}

func (s *GetMcuLvsIpRequest) Validate() error {
	return dara.Validate(s)
}
