// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawChannelsShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawChannelsShrinkRequest
	GetApplicationId() *string
	SetChannelListShrink(v string) *DescribePolarClawChannelsShrinkRequest
	GetChannelListShrink() *string
}

type DescribePolarClawChannelsShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// feishu,telegram
	ChannelListShrink *string `json:"ChannelList,omitempty" xml:"ChannelList,omitempty"`
}

func (s DescribePolarClawChannelsShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawChannelsShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawChannelsShrinkRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawChannelsShrinkRequest) GetChannelListShrink() *string {
	return s.ChannelListShrink
}

func (s *DescribePolarClawChannelsShrinkRequest) SetApplicationId(v string) *DescribePolarClawChannelsShrinkRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawChannelsShrinkRequest) SetChannelListShrink(v string) *DescribePolarClawChannelsShrinkRequest {
	s.ChannelListShrink = &v
	return s
}

func (s *DescribePolarClawChannelsShrinkRequest) Validate() error {
	return dara.Validate(s)
}
