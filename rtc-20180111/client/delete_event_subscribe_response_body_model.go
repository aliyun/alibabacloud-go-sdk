// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventSubscribeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEventSubscribeResponseBody
	GetRequestId() *string
}

type DeleteEventSubscribeResponseBody struct {
	// example:
	//
	// 760bad53276431c499e30dc36f6b26be
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEventSubscribeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventSubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventSubscribeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventSubscribeResponseBody) SetRequestId(v string) *DeleteEventSubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventSubscribeResponseBody) Validate() error {
	return dara.Validate(s)
}
