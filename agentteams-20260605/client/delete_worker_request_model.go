// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWorkerRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteWorkerRequest
	GetInstanceId() *string
	SetName(v string) *DeleteWorkerRequest
	GetName() *string
}

type DeleteWorkerRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DeleteWorkerRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteWorkerRequest) GoString() string {
	return s.String()
}

func (s *DeleteWorkerRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteWorkerRequest) GetName() *string {
	return s.Name
}

func (s *DeleteWorkerRequest) SetInstanceId(v string) *DeleteWorkerRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteWorkerRequest) SetName(v string) *DeleteWorkerRequest {
	s.Name = &v
	return s
}

func (s *DeleteWorkerRequest) Validate() error {
	return dara.Validate(s)
}
