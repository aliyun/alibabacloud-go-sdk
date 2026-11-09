// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteGroupIdResponseBody
	GetRequestId() *string
}

type DeleteGroupIdResponseBody struct {
	// The common parameter. The ID of each request is unique.
	//
	// example:
	//
	// 0621DDD7-F0E9-4D35-8900-518116D6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteGroupIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupIdResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteGroupIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteGroupIdResponseBody) SetRequestId(v string) *DeleteGroupIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteGroupIdResponseBody) Validate() error {
	return dara.Validate(s)
}
