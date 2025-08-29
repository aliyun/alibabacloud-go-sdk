// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCrowdsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListCrowdsRequest
	GetInstanceId() *string
}

type ListCrowdsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ListCrowdsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListCrowdsRequest) GoString() string {
	return s.String()
}

func (s *ListCrowdsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListCrowdsRequest) SetInstanceId(v string) *ListCrowdsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListCrowdsRequest) Validate() error {
	return dara.Validate(s)
}
