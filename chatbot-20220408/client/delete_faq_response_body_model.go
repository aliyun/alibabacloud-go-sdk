// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaqResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteFaqResponseBody
	GetRequestId() *string
}

type DeleteFaqResponseBody struct {
	// example:
	//
	// F79E7305-5314-5069-A701-9591AD051902
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteFaqResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaqResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFaqResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteFaqResponseBody) SetRequestId(v string) *DeleteFaqResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFaqResponseBody) Validate() error {
	return dara.Validate(s)
}
