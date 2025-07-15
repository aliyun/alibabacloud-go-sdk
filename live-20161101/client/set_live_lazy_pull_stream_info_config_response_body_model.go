// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetLiveLazyPullStreamInfoConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetLiveLazyPullStreamInfoConfigResponseBody
	GetRequestId() *string
}

type SetLiveLazyPullStreamInfoConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetLiveLazyPullStreamInfoConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetLiveLazyPullStreamInfoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetLiveLazyPullStreamInfoConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetLiveLazyPullStreamInfoConfigResponseBody) SetRequestId(v string) *SetLiveLazyPullStreamInfoConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetLiveLazyPullStreamInfoConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
