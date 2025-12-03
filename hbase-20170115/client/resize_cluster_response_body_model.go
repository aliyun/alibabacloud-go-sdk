// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResizeClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ResizeClusterResponseBody
	GetRequestId() *string
}

type ResizeClusterResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ResizeClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ResizeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ResizeClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ResizeClusterResponseBody) SetRequestId(v string) *ResizeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ResizeClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
