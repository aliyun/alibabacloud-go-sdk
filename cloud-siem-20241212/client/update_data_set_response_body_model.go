// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataSetResponseBody
	GetRequestId() *string
}

type UpdateDataSetResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDataSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataSetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataSetResponseBody) SetRequestId(v string) *UpdateDataSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataSetResponseBody) Validate() error {
	return dara.Validate(s)
}
