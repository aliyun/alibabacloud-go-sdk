// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDNADBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDNADBResponseBody
	GetRequestId() *string
}

type DeleteDNADBResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDNADBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDNADBResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDNADBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDNADBResponseBody) SetRequestId(v string) *DeleteDNADBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDNADBResponseBody) Validate() error {
	return dara.Validate(s)
}
