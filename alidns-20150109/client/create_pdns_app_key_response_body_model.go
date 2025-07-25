// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePdnsAppKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreatePdnsAppKeyResponseBody
	GetRequestId() *string
}

type CreatePdnsAppKeyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePdnsAppKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePdnsAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePdnsAppKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePdnsAppKeyResponseBody) SetRequestId(v string) *CreatePdnsAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePdnsAppKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
