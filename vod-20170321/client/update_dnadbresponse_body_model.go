// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDNADBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDNADBResponseBody
	GetRequestId() *string
}

type UpdateDNADBResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDNADBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDNADBResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDNADBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDNADBResponseBody) SetRequestId(v string) *UpdateDNADBResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDNADBResponseBody) Validate() error {
	return dara.Validate(s)
}
