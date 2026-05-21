// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTempFileTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateTempFileTaskRequest
	GetInstanceId() *string
}

type CreateTempFileTaskRequest struct {
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateTempFileTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTempFileTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTempFileTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateTempFileTaskRequest) SetInstanceId(v string) *CreateTempFileTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateTempFileTaskRequest) Validate() error {
	return dara.Validate(s)
}
