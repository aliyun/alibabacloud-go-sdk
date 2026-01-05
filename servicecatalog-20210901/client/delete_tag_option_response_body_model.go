// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagOptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTagOptionResponseBody
	GetRequestId() *string
}

type DeleteTagOptionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0FEEF92D-4052-5202-87D0-3D8EC16F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagOptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagOptionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagOptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTagOptionResponseBody) SetRequestId(v string) *DeleteTagOptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTagOptionResponseBody) Validate() error {
	return dara.Validate(s)
}
