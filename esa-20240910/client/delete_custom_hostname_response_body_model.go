// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomHostnameResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomHostnameResponseBody
	GetRequestId() *string
}

type DeleteCustomHostnameResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomHostnameResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomHostnameResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomHostnameResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomHostnameResponseBody) SetRequestId(v string) *DeleteCustomHostnameResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomHostnameResponseBody) Validate() error {
	return dara.Validate(s)
}
