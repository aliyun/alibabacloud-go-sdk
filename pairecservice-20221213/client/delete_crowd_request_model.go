// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCrowdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteCrowdRequest
	GetInstanceId() *string
}

type DeleteCrowdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteCrowdRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCrowdRequest) GoString() string {
	return s.String()
}

func (s *DeleteCrowdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCrowdRequest) SetInstanceId(v string) *DeleteCrowdRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCrowdRequest) Validate() error {
	return dara.Validate(s)
}
