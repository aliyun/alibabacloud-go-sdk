// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHybridProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateHybridProxyResponseBody
	GetRequestId() *string
}

type UpdateHybridProxyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// C0DD0616-0DA0-5450-B89E-F30D49E63D6B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateHybridProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateHybridProxyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateHybridProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateHybridProxyResponseBody) SetRequestId(v string) *UpdateHybridProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateHybridProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
