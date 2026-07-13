// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateWorkerBootstrapTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateWorkerBootstrapTokenRequest
	GetInstanceId() *string
	SetName(v string) *CreateWorkerBootstrapTokenRequest
	GetName() *string
	SetNetworkType(v string) *CreateWorkerBootstrapTokenRequest
	GetNetworkType() *string
}

type CreateWorkerBootstrapTokenRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
}

func (s CreateWorkerBootstrapTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateWorkerBootstrapTokenRequest) GoString() string {
	return s.String()
}

func (s *CreateWorkerBootstrapTokenRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateWorkerBootstrapTokenRequest) GetName() *string {
	return s.Name
}

func (s *CreateWorkerBootstrapTokenRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateWorkerBootstrapTokenRequest) SetInstanceId(v string) *CreateWorkerBootstrapTokenRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateWorkerBootstrapTokenRequest) SetName(v string) *CreateWorkerBootstrapTokenRequest {
	s.Name = &v
	return s
}

func (s *CreateWorkerBootstrapTokenRequest) SetNetworkType(v string) *CreateWorkerBootstrapTokenRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateWorkerBootstrapTokenRequest) Validate() error {
	return dara.Validate(s)
}
