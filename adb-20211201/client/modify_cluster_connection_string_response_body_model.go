// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterConnectionStringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClusterConnectionStringResponseBody
	GetRequestId() *string
}

type ModifyClusterConnectionStringResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 370D09FD-442A-5225-AAD3-7362CAE39177
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterConnectionStringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterConnectionStringResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterConnectionStringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterConnectionStringResponseBody) SetRequestId(v string) *ModifyClusterConnectionStringResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterConnectionStringResponseBody) Validate() error {
	return dara.Validate(s)
}
