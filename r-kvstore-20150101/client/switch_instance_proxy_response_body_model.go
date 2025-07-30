// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSwitchInstanceProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SwitchInstanceProxyResponseBody
	GetRequestId() *string
}

type SwitchInstanceProxyResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5D622714-AEDD-4609-9167-F5DDD3D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SwitchInstanceProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SwitchInstanceProxyResponseBody) GoString() string {
	return s.String()
}

func (s *SwitchInstanceProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SwitchInstanceProxyResponseBody) SetRequestId(v string) *SwitchInstanceProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *SwitchInstanceProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
