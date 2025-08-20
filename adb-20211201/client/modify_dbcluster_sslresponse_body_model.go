// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBClusterSSLResponseBody
	GetRequestId() *string
}

type ModifyDBClusterSSLResponseBody struct {
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBClusterSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterSSLResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBClusterSSLResponseBody) SetRequestId(v string) *ModifyDBClusterSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBClusterSSLResponseBody) Validate() error {
	return dara.Validate(s)
}
