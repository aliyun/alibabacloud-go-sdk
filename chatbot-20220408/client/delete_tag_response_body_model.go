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
	SetSuccess(v bool) *DeleteTagResponseBody
	GetSuccess() *bool
}

type DeleteTagResponseBody struct {
	// example:
	//
	// xxxXxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *DeleteTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteTagResponseBody) SetRequestId(v string) *DeleteTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTagResponseBody) SetSuccess(v bool) *DeleteTagResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteTagResponseBody) Validate() error {
	return dara.Validate(s)
}
