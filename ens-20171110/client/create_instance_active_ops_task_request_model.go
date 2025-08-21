// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceActiveOpsTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *CreateInstanceActiveOpsTaskRequest
	GetInstanceIds() []*string
}

type CreateInstanceActiveOpsTaskRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s CreateInstanceActiveOpsTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceActiveOpsTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceActiveOpsTaskRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateInstanceActiveOpsTaskRequest) SetInstanceIds(v []*string) *CreateInstanceActiveOpsTaskRequest {
	s.InstanceIds = v
	return s
}

func (s *CreateInstanceActiveOpsTaskRequest) Validate() error {
	return dara.Validate(s)
}
