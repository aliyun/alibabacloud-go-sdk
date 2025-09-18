// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCommandResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteCommandResponseBody
	GetSuccess() *bool
}

type DeleteCommandResponseBody struct {
	// example:
	//
	// xxxx-xx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteCommandResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCommandResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCommandResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteCommandResponseBody) SetRequestId(v string) *DeleteCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCommandResponseBody) SetSuccess(v bool) *DeleteCommandResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteCommandResponseBody) Validate() error {
	return dara.Validate(s)
}
