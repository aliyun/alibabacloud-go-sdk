// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceActiveOpsGroup interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *InstanceActiveOpsGroup
	GetInstanceIds() []*string
}

type InstanceActiveOpsGroup struct {
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
}

func (s InstanceActiveOpsGroup) String() string {
	return dara.Prettify(s)
}

func (s InstanceActiveOpsGroup) GoString() string {
	return s.String()
}

func (s *InstanceActiveOpsGroup) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *InstanceActiveOpsGroup) SetInstanceIds(v []*string) *InstanceActiveOpsGroup {
	s.InstanceIds = v
	return s
}

func (s *InstanceActiveOpsGroup) Validate() error {
	return dara.Validate(s)
}
