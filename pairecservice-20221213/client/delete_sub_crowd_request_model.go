// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSubCrowdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteSubCrowdRequest
	GetInstanceId() *string
}

type DeleteSubCrowdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteSubCrowdRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSubCrowdRequest) GoString() string {
	return s.String()
}

func (s *DeleteSubCrowdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteSubCrowdRequest) SetInstanceId(v string) *DeleteSubCrowdRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteSubCrowdRequest) Validate() error {
	return dara.Validate(s)
}
