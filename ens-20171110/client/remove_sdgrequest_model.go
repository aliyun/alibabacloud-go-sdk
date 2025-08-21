// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *RemoveSDGRequest
	GetInstanceIds() []*string
}

type RemoveSDGRequest struct {
	// IDs of Android in Container (AIC) instances.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s RemoveSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveSDGRequest) GoString() string {
	return s.String()
}

func (s *RemoveSDGRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RemoveSDGRequest) SetInstanceIds(v []*string) *RemoveSDGRequest {
	s.InstanceIds = v
	return s
}

func (s *RemoveSDGRequest) Validate() error {
	return dara.Validate(s)
}
