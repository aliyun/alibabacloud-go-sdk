// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDecreaseNodeGroupParam interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupId(v string) *DecreaseNodeGroupParam
	GetNodeGroupId() *string
	SetReleaseInstanceIds(v []*string) *DecreaseNodeGroupParam
	GetReleaseInstanceIds() []*string
}

type DecreaseNodeGroupParam struct {
	// This parameter is required.
	//
	// example:
	//
	// G-21E39B11837E****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// This parameter is required.
	ReleaseInstanceIds []*string `json:"ReleaseInstanceIds,omitempty" xml:"ReleaseInstanceIds,omitempty" type:"Repeated"`
}

func (s DecreaseNodeGroupParam) String() string {
	return dara.Prettify(s)
}

func (s DecreaseNodeGroupParam) GoString() string {
	return s.String()
}

func (s *DecreaseNodeGroupParam) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DecreaseNodeGroupParam) GetReleaseInstanceIds() []*string {
	return s.ReleaseInstanceIds
}

func (s *DecreaseNodeGroupParam) SetNodeGroupId(v string) *DecreaseNodeGroupParam {
	s.NodeGroupId = &v
	return s
}

func (s *DecreaseNodeGroupParam) SetReleaseInstanceIds(v []*string) *DecreaseNodeGroupParam {
	s.ReleaseInstanceIds = v
	return s
}

func (s *DecreaseNodeGroupParam) Validate() error {
	return dara.Validate(s)
}
