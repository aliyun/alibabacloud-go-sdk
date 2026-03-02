// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePasskeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePasskeyResponseBody
	GetRequestId() *string
}

type DeletePasskeyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePasskeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePasskeyResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePasskeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePasskeyResponseBody) SetRequestId(v string) *DeletePasskeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePasskeyResponseBody) Validate() error {
	return dara.Validate(s)
}
