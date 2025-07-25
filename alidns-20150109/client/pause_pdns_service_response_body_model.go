// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPausePdnsServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PausePdnsServiceResponseBody
	GetRequestId() *string
}

type PausePdnsServiceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PausePdnsServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PausePdnsServiceResponseBody) GoString() string {
	return s.String()
}

func (s *PausePdnsServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PausePdnsServiceResponseBody) SetRequestId(v string) *PausePdnsServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PausePdnsServiceResponseBody) Validate() error {
	return dara.Validate(s)
}
