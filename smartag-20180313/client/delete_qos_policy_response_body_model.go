// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQosPolicyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteQosPolicyResponseBody
	GetRequestId() *string
}

type DeleteQosPolicyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 551CD836-9E46-4F2C-A167-B4363180A647
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteQosPolicyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteQosPolicyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteQosPolicyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteQosPolicyResponseBody) SetRequestId(v string) *DeleteQosPolicyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteQosPolicyResponseBody) Validate() error {
	return dara.Validate(s)
}
