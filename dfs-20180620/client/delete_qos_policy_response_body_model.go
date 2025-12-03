// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessages(v string) *DeleteQosPolicyResponseBody
	GetErrorMessages() *string
	SetRequestId(v string) *DeleteQosPolicyResponseBody
	GetRequestId() *string
}

type DeleteQosPolicyResponseBody struct {
	ErrorMessages *string `json:"ErrorMessages,omitempty" xml:"ErrorMessages,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQosPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQosPolicyResponseBody) GetErrorMessages() *string {
	return s.ErrorMessages
}

func (s *DeleteQosPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQosPolicyResponseBody) SetErrorMessages(v string) *DeleteQosPolicyResponseBody {
	s.ErrorMessages = &v
	return s
}

func (s *DeleteQosPolicyResponseBody) SetRequestId(v string) *DeleteQosPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQosPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
