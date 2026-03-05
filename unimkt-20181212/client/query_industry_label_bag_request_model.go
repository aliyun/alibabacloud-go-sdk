// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIndustryLabelBagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetChannelId(v string) *QueryIndustryLabelBagRequest
	GetChannelId() *string
}

type QueryIndustryLabelBagRequest struct {
	// This parameter is required.
	ChannelId *string `json:"ChannelId,omitempty" xml:"ChannelId,omitempty"`
}

func (s QueryIndustryLabelBagRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryIndustryLabelBagRequest) GoString() string {
	return s.String()
}

func (s *QueryIndustryLabelBagRequest) GetChannelId() *string {
	return s.ChannelId
}

func (s *QueryIndustryLabelBagRequest) SetChannelId(v string) *QueryIndustryLabelBagRequest {
	s.ChannelId = &v
	return s
}

func (s *QueryIndustryLabelBagRequest) Validate() error {
	return dara.Validate(s)
}
