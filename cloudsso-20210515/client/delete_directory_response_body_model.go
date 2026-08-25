// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDirectoryResponseBody
	GetRequestId() *string
}

type DeleteDirectoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B182C041-8C64-5F2F-A07B-FC67FAF89CF9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDirectoryResponseBody) SetRequestId(v string) *DeleteDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
