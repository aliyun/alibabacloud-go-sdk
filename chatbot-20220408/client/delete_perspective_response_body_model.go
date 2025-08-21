// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePerspectiveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePerspectiveResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeletePerspectiveResponseBody
	GetResult() *bool
}

type DeletePerspectiveResponseBody struct {
	// example:
	//
	// FC384CE1-8D42-1900-84E1-F33F990F2B5E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeletePerspectiveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePerspectiveResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePerspectiveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePerspectiveResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeletePerspectiveResponseBody) SetRequestId(v string) *DeletePerspectiveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePerspectiveResponseBody) SetResult(v bool) *DeletePerspectiveResponseBody {
	s.Result = &v
	return s
}

func (s *DeletePerspectiveResponseBody) Validate() error {
	return dara.Validate(s)
}
