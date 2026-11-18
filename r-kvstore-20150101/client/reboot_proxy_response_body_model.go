// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootProxyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootProxyResponseBody
	GetRequestId() *string
}

type RebootProxyResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 561AFBF1-BE20-44DB-9BD1-6988B53E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RebootProxyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RebootProxyResponseBody) GoString() string {
	return s.String()
}

func (s *RebootProxyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootProxyResponseBody) SetRequestId(v string) *RebootProxyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootProxyResponseBody) Validate() error {
	return dara.Validate(s)
}
