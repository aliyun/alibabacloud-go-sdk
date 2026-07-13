// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerBootstrapOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetWorkerBootstrapOptionsRequest
	GetInstanceId() *string
	SetName(v string) *GetWorkerBootstrapOptionsRequest
	GetName() *string
}

type GetWorkerBootstrapOptionsRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetWorkerBootstrapOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerBootstrapOptionsRequest) GoString() string {
	return s.String()
}

func (s *GetWorkerBootstrapOptionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkerBootstrapOptionsRequest) GetName() *string {
	return s.Name
}

func (s *GetWorkerBootstrapOptionsRequest) SetInstanceId(v string) *GetWorkerBootstrapOptionsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetWorkerBootstrapOptionsRequest) SetName(v string) *GetWorkerBootstrapOptionsRequest {
	s.Name = &v
	return s
}

func (s *GetWorkerBootstrapOptionsRequest) Validate() error {
	return dara.Validate(s)
}
