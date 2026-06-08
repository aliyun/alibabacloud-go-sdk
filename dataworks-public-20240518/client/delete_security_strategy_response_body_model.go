// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSecurityStrategyResponseBody
	GetRequestId() *string
}

type DeleteSecurityStrategyResponseBody struct {
	// example:
	//
	// 0bc5df3a17****903790e8e8a
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSecurityStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSecurityStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSecurityStrategyResponseBody) SetRequestId(v string) *DeleteSecurityStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
