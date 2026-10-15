// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDistributionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArticleId(v string) *CreateDistributionRequest
	GetArticleId() *string
	SetChannels(v []*string) *CreateDistributionRequest
	GetChannels() []*string
}

type CreateDistributionRequest struct {
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
	Channels []*string `json:"Channels,omitempty" xml:"Channels,omitempty" type:"Repeated"`
}

func (s CreateDistributionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDistributionRequest) GoString() string {
	return s.String()
}

func (s *CreateDistributionRequest) GetArticleId() *string {
	return s.ArticleId
}

func (s *CreateDistributionRequest) GetChannels() []*string {
	return s.Channels
}

func (s *CreateDistributionRequest) SetArticleId(v string) *CreateDistributionRequest {
	s.ArticleId = &v
	return s
}

func (s *CreateDistributionRequest) SetChannels(v []*string) *CreateDistributionRequest {
	s.Channels = v
	return s
}

func (s *CreateDistributionRequest) Validate() error {
	return dara.Validate(s)
}
