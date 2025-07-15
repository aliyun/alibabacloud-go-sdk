// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLivePullStreamInfoConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateLivePullStreamInfoConfigResponseBody
	GetRequestId() *string
}

type UpdateLivePullStreamInfoConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateLivePullStreamInfoConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLivePullStreamInfoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLivePullStreamInfoConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLivePullStreamInfoConfigResponseBody) SetRequestId(v string) *UpdateLivePullStreamInfoConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLivePullStreamInfoConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
