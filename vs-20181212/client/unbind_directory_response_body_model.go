// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindDirectoryResponseBody
	GetRequestId() *string
}

type UnbindDirectoryResponseBody struct {
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindDirectoryResponseBody) SetRequestId(v string) *UnbindDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
