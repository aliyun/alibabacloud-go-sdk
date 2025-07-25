// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIpamScopeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIpamScopeResponseBody
	GetRequestId() *string
}

type DeleteIpamScopeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 9F8315CB-560E-5F1E-B069-6E44B440CAF8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteIpamScopeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIpamScopeResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIpamScopeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIpamScopeResponseBody) SetRequestId(v string) *DeleteIpamScopeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIpamScopeResponseBody) Validate() error {
	return dara.Validate(s)
}
