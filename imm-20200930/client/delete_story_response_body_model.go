// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStoryResponseBody
	GetRequestId() *string
}

type DeleteStoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1B3D5E0A-D8B8-4DA0-8127-ED32C851****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStoryResponseBody) SetRequestId(v string) *DeleteStoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStoryResponseBody) Validate() error {
	return dara.Validate(s)
}
