// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEipResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEipResponseBody
	GetRequestId() *string
}

type DeleteEipResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9341CDC2-D6AC-5992-86C8-D5774CFCC708
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEipResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEipResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEipResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEipResponseBody) SetRequestId(v string) *DeleteEipResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEipResponseBody) Validate() error {
	return dara.Validate(s)
}
