// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarClawChannelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *DescribePolarClawChannelsRequest
	GetApplicationId() *string
	SetChannelList(v []*string) *DescribePolarClawChannelsRequest
	GetChannelList() []*string
}

type DescribePolarClawChannelsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// pa-**************
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// example:
	//
	// feishu,telegram
	ChannelList []*string `json:"ChannelList,omitempty" xml:"ChannelList,omitempty" type:"Repeated"`
}

func (s DescribePolarClawChannelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarClawChannelsRequest) GoString() string {
	return s.String()
}

func (s *DescribePolarClawChannelsRequest) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *DescribePolarClawChannelsRequest) GetChannelList() []*string {
	return s.ChannelList
}

func (s *DescribePolarClawChannelsRequest) SetApplicationId(v string) *DescribePolarClawChannelsRequest {
	s.ApplicationId = &v
	return s
}

func (s *DescribePolarClawChannelsRequest) SetChannelList(v []*string) *DescribePolarClawChannelsRequest {
	s.ChannelList = v
	return s
}

func (s *DescribePolarClawChannelsRequest) Validate() error {
	return dara.Validate(s)
}
