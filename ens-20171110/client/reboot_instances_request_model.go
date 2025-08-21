// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *RebootInstancesRequest
	GetInstanceIds() []*string
}

type RebootInstancesRequest struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s RebootInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootInstancesRequest) GoString() string {
	return s.String()
}

func (s *RebootInstancesRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RebootInstancesRequest) SetInstanceIds(v []*string) *RebootInstancesRequest {
	s.InstanceIds = v
	return s
}

func (s *RebootInstancesRequest) Validate() error {
	return dara.Validate(s)
}
