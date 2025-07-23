// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitializeClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *InitializeClusterResponseBody
	GetRequestId() *string
}

type InitializeClusterResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049366F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s InitializeClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s InitializeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *InitializeClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *InitializeClusterResponseBody) SetRequestId(v string) *InitializeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *InitializeClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
