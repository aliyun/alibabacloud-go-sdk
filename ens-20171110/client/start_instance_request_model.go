// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartInstanceRequest
	GetInstanceId() *string
}

type StartInstanceRequest struct {
	// The ID of the instance. You can start only one instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-instance****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s StartInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceRequest) GoString() string {
	return s.String()
}

func (s *StartInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartInstanceRequest) SetInstanceId(v string) *StartInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceRequest) Validate() error {
	return dara.Validate(s)
}
