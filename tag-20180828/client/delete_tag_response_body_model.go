// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTagResponseBody
	GetRequestId() *string
}

type DeleteTagResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6E27F22C-EDA3-132E-A53F-77DE3BC2343D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTagResponseBody) SetRequestId(v string) *DeleteTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTagResponseBody) Validate() error {
	return dara.Validate(s)
}
