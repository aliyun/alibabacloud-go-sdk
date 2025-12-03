// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEaiAllResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEaiAllResponseBody
	GetRequestId() *string
}

type DeleteEaiAllResponseBody struct {
	// example:
	//
	// AD4EA714-A35B-4710-BF92-8275BCDDB69F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEaiAllResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEaiAllResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEaiAllResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEaiAllResponseBody) SetRequestId(v string) *DeleteEaiAllResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEaiAllResponseBody) Validate() error {
	return dara.Validate(s)
}
