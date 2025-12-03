// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHbaseSlbServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateHbaseSlbServerResponseBody
	GetRequestId() *string
}

type CreateHbaseSlbServerResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHbaseSlbServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHbaseSlbServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHbaseSlbServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHbaseSlbServerResponseBody) SetRequestId(v string) *CreateHbaseSlbServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHbaseSlbServerResponseBody) Validate() error {
	return dara.Validate(s)
}
