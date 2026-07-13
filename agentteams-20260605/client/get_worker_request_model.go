// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetWorkerRequest
	GetInstanceId() *string
	SetName(v string) *GetWorkerRequest
	GetName() *string
}

type GetWorkerRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetWorkerRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerRequest) GoString() string {
	return s.String()
}

func (s *GetWorkerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkerRequest) GetName() *string {
	return s.Name
}

func (s *GetWorkerRequest) SetInstanceId(v string) *GetWorkerRequest {
	s.InstanceId = &v
	return s
}

func (s *GetWorkerRequest) SetName(v string) *GetWorkerRequest {
	s.Name = &v
	return s
}

func (s *GetWorkerRequest) Validate() error {
	return dara.Validate(s)
}
