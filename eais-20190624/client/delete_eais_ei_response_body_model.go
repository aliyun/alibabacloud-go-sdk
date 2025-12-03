// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEaisEiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteEaisEiResponseBody
	GetRequestId() *string
}

type DeleteEaisEiResponseBody struct {
	// example:
	//
	// F23AEEC7-4D98-4657-A104-02692701****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEaisEiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEaisEiResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEaisEiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEaisEiResponseBody) SetRequestId(v string) *DeleteEaisEiResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEaisEiResponseBody) Validate() error {
	return dara.Validate(s)
}
