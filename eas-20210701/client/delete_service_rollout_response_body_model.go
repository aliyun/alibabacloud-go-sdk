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
	// **Message*	-
	//
	// 	- Type: `string`
	//
	// 	- Description: The response message. Describes the result of the request processing.
	//
	// 	- Example: `Rollout deleted successfully`
	//
	// example:
	//
	// Rollout deleted successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// **RequestId*	-
	//
	// 	- Type: `string`
	//
	// 	- Description: The request ID. A unique identifier returned for each request, used for troubleshooting and tracing.
	//
	// 	- Example: `40325405-579C-4D82-9B4F-8A7C6D5E4F3A`
	//
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
