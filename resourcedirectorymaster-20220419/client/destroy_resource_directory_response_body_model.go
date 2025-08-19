// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDestroyResourceDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DestroyResourceDirectoryResponseBody
	GetRequestId() *string
}

type DestroyResourceDirectoryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 9B34724D-54B0-4A51-B34D-4512372FE1BE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DestroyResourceDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DestroyResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DestroyResourceDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DestroyResourceDirectoryResponseBody) SetRequestId(v string) *DestroyResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DestroyResourceDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
