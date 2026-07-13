// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteModelRequest
	GetId() *string
	SetInstanceId(v string) *DeleteModelRequest
	GetInstanceId() *string
	SetProviderId(v string) *DeleteModelRequest
	GetProviderId() *string
}

type DeleteModelRequest struct {
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	ProviderId *string `json:"ProviderId,omitempty" xml:"ProviderId,omitempty"`
}

func (s DeleteModelRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelRequest) GetId() *string {
	return s.Id
}

func (s *DeleteModelRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteModelRequest) GetProviderId() *string {
	return s.ProviderId
}

func (s *DeleteModelRequest) SetId(v string) *DeleteModelRequest {
	s.Id = &v
	return s
}

func (s *DeleteModelRequest) SetInstanceId(v string) *DeleteModelRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteModelRequest) SetProviderId(v string) *DeleteModelRequest {
	s.ProviderId = &v
	return s
}

func (s *DeleteModelRequest) Validate() error {
	return dara.Validate(s)
}
