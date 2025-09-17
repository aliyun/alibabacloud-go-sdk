// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSourceServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSourceServerResponseBody
	GetRequestId() *string
}

type DeleteSourceServerResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSourceServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSourceServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSourceServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSourceServerResponseBody) SetRequestId(v string) *DeleteSourceServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSourceServerResponseBody) Validate() error {
	return dara.Validate(s)
}
