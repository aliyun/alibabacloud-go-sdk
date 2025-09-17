// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetServiceRequest
	GetInstanceId() *string
}

type GetServiceRequest struct {
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRequest) GoString() string {
	return s.String()
}

func (s *GetServiceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetServiceRequest) SetInstanceId(v string) *GetServiceRequest {
	s.InstanceId = &v
	return s
}

func (s *GetServiceRequest) Validate() error {
	return dara.Validate(s)
}
