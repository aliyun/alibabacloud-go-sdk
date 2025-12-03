// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHbaseSlbServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteHbaseSlbServerResponseBody
	GetRequestId() *string
}

type DeleteHbaseSlbServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteHbaseSlbServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteHbaseSlbServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteHbaseSlbServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteHbaseSlbServerResponseBody) SetRequestId(v string) *DeleteHbaseSlbServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteHbaseSlbServerResponseBody) Validate() error {
	return dara.Validate(s)
}
