// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComponentIndexResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteComponentIndexResponseBody
	GetRequestId() *string
	SetResult(v bool) *DeleteComponentIndexResponseBody
	GetResult() *bool
}

type DeleteComponentIndexResponseBody struct {
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DCC47D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s DeleteComponentIndexResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteComponentIndexResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteComponentIndexResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteComponentIndexResponseBody) GetResult() *bool {
	return s.Result
}

func (s *DeleteComponentIndexResponseBody) SetRequestId(v string) *DeleteComponentIndexResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteComponentIndexResponseBody) SetResult(v bool) *DeleteComponentIndexResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteComponentIndexResponseBody) Validate() error {
	return dara.Validate(s)
}
