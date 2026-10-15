// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryChannelPublishRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDraftId(v string) *RetryChannelPublishRequest
	GetDraftId() *string
}

type RetryChannelPublishRequest struct {
	// The channel draft ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// draft-789
	DraftId *string `json:"DraftId,omitempty" xml:"DraftId,omitempty"`
}

func (s RetryChannelPublishRequest) String() string {
	return dara.Prettify(s)
}

func (s RetryChannelPublishRequest) GoString() string {
	return s.String()
}

func (s *RetryChannelPublishRequest) GetDraftId() *string {
	return s.DraftId
}

func (s *RetryChannelPublishRequest) SetDraftId(v string) *RetryChannelPublishRequest {
	s.DraftId = &v
	return s
}

func (s *RetryChannelPublishRequest) Validate() error {
	return dara.Validate(s)
}
