// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnsSaleControlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateEnsSaleControlResponseBody
	GetRequestId() *string
}

type CreateEnsSaleControlResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateEnsSaleControlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateEnsSaleControlResponseBody) GoString() string {
	return s.String()
}

func (s *CreateEnsSaleControlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateEnsSaleControlResponseBody) SetRequestId(v string) *CreateEnsSaleControlResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateEnsSaleControlResponseBody) Validate() error {
	return dara.Validate(s)
}
