// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLLMConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLLMConfigResponseBody
	GetRequestId() *string
}

type DeleteLLMConfigResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLLMConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLLMConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLLMConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLLMConfigResponseBody) SetRequestId(v string) *DeleteLLMConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLLMConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
