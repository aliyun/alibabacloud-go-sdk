// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteValidateFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteValidateFileResponseBody
	GetRequestId() *string
}

type DeleteValidateFileResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteValidateFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteValidateFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteValidateFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteValidateFileResponseBody) SetRequestId(v string) *DeleteValidateFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteValidateFileResponseBody) Validate() error {
	return dara.Validate(s)
}
