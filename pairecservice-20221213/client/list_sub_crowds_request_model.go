// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSubCrowdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListSubCrowdsRequest
	GetInstanceId() *string
}

type ListSubCrowdsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListSubCrowdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSubCrowdsRequest) GoString() string {
	return s.String()
}

func (s *ListSubCrowdsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListSubCrowdsRequest) SetInstanceId(v string) *ListSubCrowdsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListSubCrowdsRequest) Validate() error {
	return dara.Validate(s)
}
