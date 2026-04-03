// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteChunksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteChunksResponseBody
	GetRequestId() *string
}

type DeleteChunksResponseBody struct {
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteChunksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteChunksResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteChunksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteChunksResponseBody) SetRequestId(v string) *DeleteChunksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteChunksResponseBody) Validate() error {
	return dara.Validate(s)
}
