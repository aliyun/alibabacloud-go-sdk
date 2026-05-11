// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceRolloutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteServiceRolloutResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteServiceRolloutResponseBody
	GetRequestId() *string
}

type DeleteServiceRolloutResponseBody struct {
	// example:
	//
	// Rollout deleted successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteServiceRolloutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceRolloutResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteServiceRolloutResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteServiceRolloutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteServiceRolloutResponseBody) SetMessage(v string) *DeleteServiceRolloutResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteServiceRolloutResponseBody) SetRequestId(v string) *DeleteServiceRolloutResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteServiceRolloutResponseBody) Validate() error {
	return dara.Validate(s)
}
