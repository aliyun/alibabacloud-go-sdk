// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteStorageResponseBody
	GetRequestId() *string
}

type DeleteStorageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteStorageResponseBody) SetRequestId(v string) *DeleteStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
