// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDbProxyInstanceSslResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDbProxyInstanceSslResponseBody
	GetRequestId() *string
}

type ModifyDbProxyInstanceSslResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// BF46A62B-3717-4397-9338-36BB95C898B3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDbProxyInstanceSslResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDbProxyInstanceSslResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDbProxyInstanceSslResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDbProxyInstanceSslResponseBody) SetRequestId(v string) *ModifyDbProxyInstanceSslResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDbProxyInstanceSslResponseBody) Validate() error {
	return dara.Validate(s)
}
