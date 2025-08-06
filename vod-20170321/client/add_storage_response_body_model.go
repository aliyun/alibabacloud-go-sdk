// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddStorageResponseBody
	GetRequestId() *string
}

type AddStorageResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddStorageResponseBody) GoString() string {
	return s.String()
}

func (s *AddStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddStorageResponseBody) SetRequestId(v string) *AddStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddStorageResponseBody) Validate() error {
	return dara.Validate(s)
}
