// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetModelProviderRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *GetModelProviderRequest
	GetId() *string
	SetInstanceId(v string) *GetModelProviderRequest
	GetInstanceId() *string
}

type GetModelProviderRequest struct {
	// This parameter is required.
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// AgentTeams
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetModelProviderRequest) String() string {
	return dara.Prettify(s)
}

func (s GetModelProviderRequest) GoString() string {
	return s.String()
}

func (s *GetModelProviderRequest) GetId() *string {
	return s.Id
}

func (s *GetModelProviderRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetModelProviderRequest) SetId(v string) *GetModelProviderRequest {
	s.Id = &v
	return s
}

func (s *GetModelProviderRequest) SetInstanceId(v string) *GetModelProviderRequest {
	s.InstanceId = &v
	return s
}

func (s *GetModelProviderRequest) Validate() error {
	return dara.Validate(s)
}
