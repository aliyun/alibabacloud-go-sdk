// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindAegisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UnbindAegisResponseBody
	GetRequestId() *string
}

type UnbindAegisResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 825F5526-2A17-4279-857F-F790E9590171
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindAegisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindAegisResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindAegisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindAegisResponseBody) SetRequestId(v string) *UnbindAegisResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindAegisResponseBody) Validate() error {
	return dara.Validate(s)
}
