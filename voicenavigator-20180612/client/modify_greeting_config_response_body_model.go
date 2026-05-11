// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGreetingConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyGreetingConfigResponseBody
	GetRequestId() *string
}

type ModifyGreetingConfigResponseBody struct {
	// example:
	//
	// 14C39896-AE6D-4643-9C9A-E0566B2C2DDD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGreetingConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGreetingConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGreetingConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGreetingConfigResponseBody) SetRequestId(v string) *ModifyGreetingConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGreetingConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
