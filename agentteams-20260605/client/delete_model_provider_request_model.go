// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteModelProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DeleteModelProviderRequest
	GetId() *string
	SetInstanceId(v string) *DeleteModelProviderRequest
	GetInstanceId() *string
}

type DeleteModelProviderRequest struct {
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DeleteModelProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteModelProviderRequest) GoString() string {
	return s.String()
}

func (s *DeleteModelProviderRequest) GetId() *string {
	return s.Id
}

func (s *DeleteModelProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteModelProviderRequest) SetId(v string) *DeleteModelProviderRequest {
	s.Id = &v
	return s
}

func (s *DeleteModelProviderRequest) SetInstanceId(v string) *DeleteModelProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteModelProviderRequest) Validate() error {
	return dara.Validate(s)
}
