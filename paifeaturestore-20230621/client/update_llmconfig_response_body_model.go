// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLLMConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLLMConfigResponseBody
	GetRequestId() *string
}

type UpdateLLMConfigResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLLMConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLLMConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLLMConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLLMConfigResponseBody) SetRequestId(v string) *UpdateLLMConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLLMConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
