// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDomainWebSocketStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetDomainWebSocketStatusResponseBody
	GetRequestId() *string
}

type SetDomainWebSocketStatusResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 61A16D46-EC04-5288-8A18-811B0F536CC2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SetDomainWebSocketStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetDomainWebSocketStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetDomainWebSocketStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetDomainWebSocketStatusResponseBody) SetRequestId(v string) *SetDomainWebSocketStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetDomainWebSocketStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
