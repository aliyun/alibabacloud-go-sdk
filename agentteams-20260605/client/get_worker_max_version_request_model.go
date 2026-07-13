// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkerMaxVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetWorkerMaxVersionRequest
	GetInstanceId() *string
}

type GetWorkerMaxVersionRequest struct {
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetWorkerMaxVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s GetWorkerMaxVersionRequest) GoString() string {
	return s.String()
}

func (s *GetWorkerMaxVersionRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetWorkerMaxVersionRequest) SetInstanceId(v string) *GetWorkerMaxVersionRequest {
	s.InstanceId = &v
	return s
}

func (s *GetWorkerMaxVersionRequest) Validate() error {
	return dara.Validate(s)
}
