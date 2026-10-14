// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveInstanceSDGRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *RemoveInstanceSDGRequest
	GetInstanceIds() []*string
}

type RemoveInstanceSDGRequest struct {
	// The list of instance IDs. The value is a JSON array. A maximum of 100 IDs can be specified.
	//
	// This parameter is required.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s RemoveInstanceSDGRequest) String() string {
	return dara.Prettify(s)
}

func (s RemoveInstanceSDGRequest) GoString() string {
	return s.String()
}

func (s *RemoveInstanceSDGRequest) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *RemoveInstanceSDGRequest) SetInstanceIds(v []*string) *RemoveInstanceSDGRequest {
	s.InstanceIds = v
	return s
}

func (s *RemoveInstanceSDGRequest) Validate() error {
	return dara.Validate(s)
}
