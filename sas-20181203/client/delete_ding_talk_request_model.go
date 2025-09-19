// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDingTalkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *DeleteDingTalkRequest
	GetIds() *string
}

type DeleteDingTalkRequest struct {
	// The ID of the notification from the DingTalk chatbot. Separate multiple IDs with commas (,).
	//
	// >  You can call the [DescribeDingTalk](~~DescribeDingTalk~~) operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2170,256
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
}

func (s DeleteDingTalkRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDingTalkRequest) GoString() string {
	return s.String()
}

func (s *DeleteDingTalkRequest) GetIds() *string {
	return s.Ids
}

func (s *DeleteDingTalkRequest) SetIds(v string) *DeleteDingTalkRequest {
	s.Ids = &v
	return s
}

func (s *DeleteDingTalkRequest) Validate() error {
	return dara.Validate(s)
}
