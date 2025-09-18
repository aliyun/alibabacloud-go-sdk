// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceDetailRequest
	GetInstanceId() *string
}

type GetInstanceDetailRequest struct {
	// The ID of the instance.
	//
	// example:
	//
	// c-123xxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetInstanceDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceDetailRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceDetailRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceDetailRequest) SetInstanceId(v string) *GetInstanceDetailRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceDetailRequest) Validate() error {
	return dara.Validate(s)
}
