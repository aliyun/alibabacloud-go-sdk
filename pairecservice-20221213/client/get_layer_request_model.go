// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLayerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetLayerRequest
	GetInstanceId() *string
}

type GetLayerRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pairec-cn-abcdefg1234
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetLayerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetLayerRequest) GoString() string {
	return s.String()
}

func (s *GetLayerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetLayerRequest) SetInstanceId(v string) *GetLayerRequest {
	s.InstanceId = &v
	return s
}

func (s *GetLayerRequest) Validate() error {
	return dara.Validate(s)
}
