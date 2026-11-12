// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceRolloutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *UpdateServiceRolloutResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateServiceRolloutResponseBody
	GetRequestId() *string
}

type UpdateServiceRolloutResponseBody struct {
	// The returned message. Describes the request processing result. Valid values:
	//
	// - "Rollout updated successfully": The rollout policy is updated.
	//
	// - "Rollout paused successfully": The rollout is paused.
	//
	// - "Rollout resumed successfully": The rollout is resumed.
	//
	// example:
	//
	// Rollout updated successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID. The unique identifier returned for each request, used for troubleshooting and tracking.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateServiceRolloutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRolloutResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateServiceRolloutResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateServiceRolloutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateServiceRolloutResponseBody) SetMessage(v string) *UpdateServiceRolloutResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateServiceRolloutResponseBody) SetRequestId(v string) *UpdateServiceRolloutResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateServiceRolloutResponseBody) Validate() error {
	return dara.Validate(s)
}
