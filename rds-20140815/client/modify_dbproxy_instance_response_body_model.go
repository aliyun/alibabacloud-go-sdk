// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBProxyInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBProxyInstanceResponseBody
	GetRequestId() *string
}

type ModifyDBProxyInstanceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBProxyInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBProxyInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBProxyInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBProxyInstanceResponseBody) SetRequestId(v string) *ModifyDBProxyInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBProxyInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
