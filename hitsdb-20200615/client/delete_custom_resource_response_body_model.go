// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteCustomResourceResponseBody
	GetRequestId() *string
}

type DeleteCustomResourceResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteCustomResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomResourceResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteCustomResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteCustomResourceResponseBody) SetRequestId(v string) *DeleteCustomResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteCustomResourceResponseBody) Validate() error {
	return dara.Validate(s)
}
