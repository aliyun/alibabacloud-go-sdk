// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseSubscriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReleaseSubscriptionResponseBody
	GetRequestId() *string
}

type ReleaseSubscriptionResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReleaseSubscriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReleaseSubscriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ReleaseSubscriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReleaseSubscriptionResponseBody) SetRequestId(v string) *ReleaseSubscriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReleaseSubscriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
