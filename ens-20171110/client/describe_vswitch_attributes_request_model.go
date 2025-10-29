// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVSwitchId(v string) *DescribeVSwitchAttributesRequest
	GetVSwitchId() *string
}

type DescribeVSwitchAttributesRequest struct {
	// The ID of the VSwitch.
	//
	// This parameter is required.
	//
	// example:
	//
	// vsw-5****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s DescribeVSwitchAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchAttributesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeVSwitchAttributesRequest) SetVSwitchId(v string) *DescribeVSwitchAttributesRequest {
	s.VSwitchId = &v
	return s
}

func (s *DescribeVSwitchAttributesRequest) Validate() error {
	return dara.Validate(s)
}
