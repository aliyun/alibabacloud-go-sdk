// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMqConsoleLinkResult interface {
	dara.Model
	String() string
	GoString() string
	SetGroupLinkUrl(v string) *MqConsoleLinkResult
	GetGroupLinkUrl() *string
	SetRequestId(v string) *MqConsoleLinkResult
	GetRequestId() *string
	SetTopicLinkUrl(v string) *MqConsoleLinkResult
	GetTopicLinkUrl() *string
}

type MqConsoleLinkResult struct {
	GroupLinkUrl *string `json:"groupLinkUrl,omitempty" xml:"groupLinkUrl,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TopicLinkUrl *string `json:"topicLinkUrl,omitempty" xml:"topicLinkUrl,omitempty"`
}

func (s MqConsoleLinkResult) String() string {
	return dara.Prettify(s)
}

func (s MqConsoleLinkResult) GoString() string {
	return s.String()
}

func (s *MqConsoleLinkResult) GetGroupLinkUrl() *string {
	return s.GroupLinkUrl
}

func (s *MqConsoleLinkResult) GetRequestId() *string {
	return s.RequestId
}

func (s *MqConsoleLinkResult) GetTopicLinkUrl() *string {
	return s.TopicLinkUrl
}

func (s *MqConsoleLinkResult) SetGroupLinkUrl(v string) *MqConsoleLinkResult {
	s.GroupLinkUrl = &v
	return s
}

func (s *MqConsoleLinkResult) SetRequestId(v string) *MqConsoleLinkResult {
	s.RequestId = &v
	return s
}

func (s *MqConsoleLinkResult) SetTopicLinkUrl(v string) *MqConsoleLinkResult {
	s.TopicLinkUrl = &v
	return s
}

func (s *MqConsoleLinkResult) Validate() error {
	return dara.Validate(s)
}
