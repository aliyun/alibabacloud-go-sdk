// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSanityCheckTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateSanityCheckTaskRequest
	GetInstanceId() *string
}

type CreateSanityCheckTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// dsw-730xxxxxxxxxx
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s CreateSanityCheckTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSanityCheckTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSanityCheckTaskRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateSanityCheckTaskRequest) SetInstanceId(v string) *CreateSanityCheckTaskRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateSanityCheckTaskRequest) Validate() error {
	return dara.Validate(s)
}
