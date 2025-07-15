// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLivePullStreamInfoConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLivePullStreamInfoConfigResponseBody
	GetRequestId() *string
}

type DeleteLivePullStreamInfoConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLivePullStreamInfoConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLivePullStreamInfoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLivePullStreamInfoConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLivePullStreamInfoConfigResponseBody) SetRequestId(v string) *DeleteLivePullStreamInfoConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLivePullStreamInfoConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
