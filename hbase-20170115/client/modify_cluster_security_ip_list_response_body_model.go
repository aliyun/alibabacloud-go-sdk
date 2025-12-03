// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterSecurityIpListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClusterSecurityIpListResponseBody
	GetRequestId() *string
}

type ModifyClusterSecurityIpListResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterSecurityIpListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterSecurityIpListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterSecurityIpListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterSecurityIpListResponseBody) SetRequestId(v string) *ModifyClusterSecurityIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterSecurityIpListResponseBody) Validate() error {
	return dara.Validate(s)
}
