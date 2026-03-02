// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessKeyInRecycleBinResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAccessKeyInRecycleBinResponseBody
	GetRequestId() *string
}

type DeleteAccessKeyInRecycleBinResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAccessKeyInRecycleBinResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessKeyInRecycleBinResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAccessKeyInRecycleBinResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAccessKeyInRecycleBinResponseBody) SetRequestId(v string) *DeleteAccessKeyInRecycleBinResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAccessKeyInRecycleBinResponseBody) Validate() error {
	return dara.Validate(s)
}
