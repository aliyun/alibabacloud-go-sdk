// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetYikeCallbackConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetYikeCallbackConfigResponseBody
	GetRequestId() *string
}

type SetYikeCallbackConfigResponseBody struct {
	// RequestId
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetYikeCallbackConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetYikeCallbackConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetYikeCallbackConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetYikeCallbackConfigResponseBody) SetRequestId(v string) *SetYikeCallbackConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetYikeCallbackConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
