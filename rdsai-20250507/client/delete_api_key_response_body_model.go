// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteApiKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *DeleteApiKeyResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteApiKeyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteApiKeyResponseBody
	GetSuccess() *bool
}

type DeleteApiKeyResponseBody struct {
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteApiKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteApiKeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteApiKeyResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteApiKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteApiKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteApiKeyResponseBody) SetMessage(v string) *DeleteApiKeyResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteApiKeyResponseBody) SetRequestId(v string) *DeleteApiKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteApiKeyResponseBody) SetSuccess(v bool) *DeleteApiKeyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteApiKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
