// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLivePullStreamInfoConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddLivePullStreamInfoConfigResponseBody
	GetRequestId() *string
}

type AddLivePullStreamInfoConfigResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddLivePullStreamInfoConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddLivePullStreamInfoConfigResponseBody) GoString() string {
	return s.String()
}

func (s *AddLivePullStreamInfoConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddLivePullStreamInfoConfigResponseBody) SetRequestId(v string) *AddLivePullStreamInfoConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddLivePullStreamInfoConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
