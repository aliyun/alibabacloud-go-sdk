// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySoarStrategySubscribeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifySoarStrategySubscribeResponseBody
	GetRequestId() *string
}

type ModifySoarStrategySubscribeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8BF56229-7FF5-51ED-B958-00B8573E77CD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifySoarStrategySubscribeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySoarStrategySubscribeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySoarStrategySubscribeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySoarStrategySubscribeResponseBody) SetRequestId(v string) *ModifySoarStrategySubscribeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySoarStrategySubscribeResponseBody) Validate() error {
	return dara.Validate(s)
}
