// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSolutionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSolutionResponseBody
	GetRequestId() *string
}

type DeleteSolutionResponseBody struct {
	// example:
	//
	// F79E7305-5314-5069-A701-9591AD051902
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSolutionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSolutionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSolutionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSolutionResponseBody) SetRequestId(v string) *DeleteSolutionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSolutionResponseBody) Validate() error {
	return dara.Validate(s)
}
