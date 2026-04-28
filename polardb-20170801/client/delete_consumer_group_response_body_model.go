// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteConsumerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteConsumerGroupResponseBody
	GetRequestId() *string
}

type DeleteConsumerGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// E56531A4-E552-40BA-9C58-137B80******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteConsumerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteConsumerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteConsumerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteConsumerGroupResponseBody) SetRequestId(v string) *DeleteConsumerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteConsumerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
