// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVideoProcessingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVideoProcessingResponseBody
	GetRequestId() *string
}

type DeleteVideoProcessingResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6abd807e-ed2a-44de-ac54-ac38a62472e6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVideoProcessingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVideoProcessingResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVideoProcessingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVideoProcessingResponseBody) SetRequestId(v string) *DeleteVideoProcessingResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVideoProcessingResponseBody) Validate() error {
	return dara.Validate(s)
}
