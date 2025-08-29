// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSubCrowdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetSubCrowdRequest
	GetInstanceId() *string
}

type GetSubCrowdRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetSubCrowdRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSubCrowdRequest) GoString() string {
	return s.String()
}

func (s *GetSubCrowdRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetSubCrowdRequest) SetInstanceId(v string) *GetSubCrowdRequest {
	s.InstanceId = &v
	return s
}

func (s *GetSubCrowdRequest) Validate() error {
	return dara.Validate(s)
}
