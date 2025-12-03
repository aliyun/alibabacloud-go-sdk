// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUIProxyAccountPasswordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUIProxyAccountPasswordResponseBody
	GetRequestId() *string
}

type ModifyUIProxyAccountPasswordResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUIProxyAccountPasswordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUIProxyAccountPasswordResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUIProxyAccountPasswordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUIProxyAccountPasswordResponseBody) SetRequestId(v string) *ModifyUIProxyAccountPasswordResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUIProxyAccountPasswordResponseBody) Validate() error {
	return dara.Validate(s)
}
