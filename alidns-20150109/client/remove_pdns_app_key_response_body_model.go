// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemovePdnsAppKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RemovePdnsAppKeyResponseBody
	GetRequestId() *string
}

type RemovePdnsAppKeyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RemovePdnsAppKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemovePdnsAppKeyResponseBody) GoString() string {
	return s.String()
}

func (s *RemovePdnsAppKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemovePdnsAppKeyResponseBody) SetRequestId(v string) *RemovePdnsAppKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemovePdnsAppKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
