// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteParamRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteParamRequest
	GetInstanceId() *string
}

type DeleteParamRequest struct {
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteParamRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteParamRequest) GoString() string {
	return s.String()
}

func (s *DeleteParamRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteParamRequest) SetInstanceId(v string) *DeleteParamRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteParamRequest) Validate() error {
	return dara.Validate(s)
}
