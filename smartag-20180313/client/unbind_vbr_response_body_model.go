// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindVbrResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindVbrResponseBody
	GetRequestId() *string
}

type UnbindVbrResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 61C33444-D8C5-4018-A06C-BA8C8812BEF6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindVbrResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindVbrResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindVbrResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindVbrResponseBody) SetRequestId(v string) *UnbindVbrResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindVbrResponseBody) Validate() error {
	return dara.Validate(s)
}
