// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDistributionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArticleId(v string) *CreateDistributionShrinkRequest
	GetArticleId() *string
	SetChannelsShrink(v string) *CreateDistributionShrinkRequest
	GetChannelsShrink() *string
}

type CreateDistributionShrinkRequest struct {
	// The article ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 28733
	ArticleId *string `json:"ArticleId,omitempty" xml:"ArticleId,omitempty"`
	// The list of selected channels.
	//
	// This parameter is required.
	ChannelsShrink *string `json:"Channels,omitempty" xml:"Channels,omitempty"`
}

func (s CreateDistributionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDistributionShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDistributionShrinkRequest) GetArticleId() *string {
	return s.ArticleId
}

func (s *CreateDistributionShrinkRequest) GetChannelsShrink() *string {
	return s.ChannelsShrink
}

func (s *CreateDistributionShrinkRequest) SetArticleId(v string) *CreateDistributionShrinkRequest {
	s.ArticleId = &v
	return s
}

func (s *CreateDistributionShrinkRequest) SetChannelsShrink(v string) *CreateDistributionShrinkRequest {
	s.ChannelsShrink = &v
	return s
}

func (s *CreateDistributionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
