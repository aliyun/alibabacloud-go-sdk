// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UpdateModelRequest
	GetClientToken() *string
	SetDescription(v string) *UpdateModelRequest
	GetDescription() *string
	SetId(v string) *UpdateModelRequest
	GetId() *string
	SetInstanceId(v string) *UpdateModelRequest
	GetInstanceId() *string
}

type UpdateModelRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s UpdateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateModelRequest) GoString() string {
	return s.String()
}

func (s *UpdateModelRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UpdateModelRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateModelRequest) GetId() *string {
	return s.Id
}

func (s *UpdateModelRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateModelRequest) SetClientToken(v string) *UpdateModelRequest {
	s.ClientToken = &v
	return s
}

func (s *UpdateModelRequest) SetDescription(v string) *UpdateModelRequest {
	s.Description = &v
	return s
}

func (s *UpdateModelRequest) SetId(v string) *UpdateModelRequest {
	s.Id = &v
	return s
}

func (s *UpdateModelRequest) SetInstanceId(v string) *UpdateModelRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateModelRequest) Validate() error {
	return dara.Validate(s)
}
