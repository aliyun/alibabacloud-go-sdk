// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDingTalkOnlineTestRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DingTalkOnlineTestRequest
	GetId() *int64
}

type DingTalkOnlineTestRequest struct {
	// The ID of the DingTalk notification configuration.
	//
	// > You can call the [DescribeDingTalk](~~DescribeDingTalk~~) operation to obtain the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2373
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DingTalkOnlineTestRequest) String() string {
	return dara.Prettify(s)
}

func (s DingTalkOnlineTestRequest) GoString() string {
	return s.String()
}

func (s *DingTalkOnlineTestRequest) GetId() *int64 {
	return s.Id
}

func (s *DingTalkOnlineTestRequest) SetId(v int64) *DingTalkOnlineTestRequest {
	s.Id = &v
	return s
}

func (s *DingTalkOnlineTestRequest) Validate() error {
	return dara.Validate(s)
}
