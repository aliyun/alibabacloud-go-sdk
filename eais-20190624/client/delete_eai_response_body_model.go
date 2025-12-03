// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEaiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEaiResponseBody
	GetRequestId() *string
}

type DeleteEaiResponseBody struct {
	// example:
	//
	// F23AEEC7-4D98-4657-A104-0269270*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEaiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEaiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEaiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEaiResponseBody) SetRequestId(v string) *DeleteEaiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEaiResponseBody) Validate() error {
	return dara.Validate(s)
}
