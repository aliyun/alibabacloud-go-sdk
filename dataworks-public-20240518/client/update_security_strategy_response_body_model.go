// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSecurityStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateSecurityStrategyResponseBody
	GetRequestId() *string
}

type UpdateSecurityStrategyResponseBody struct {
	// example:
	//
	// 0bc5df3a17***903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateSecurityStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateSecurityStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateSecurityStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateSecurityStrategyResponseBody) SetRequestId(v string) *UpdateSecurityStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateSecurityStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
