// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLifecycleHookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLifecycleHookResponseBody
	GetRequestId() *string
}

type DeleteLifecycleHookResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLifecycleHookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLifecycleHookResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLifecycleHookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLifecycleHookResponseBody) SetRequestId(v string) *DeleteLifecycleHookResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLifecycleHookResponseBody) Validate() error {
	return dara.Validate(s)
}
