// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKMSServiceKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateKMSServiceKeyResponseBody
	GetRequestId() *string
}

type CreateKMSServiceKeyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateKMSServiceKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateKMSServiceKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateKMSServiceKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateKMSServiceKeyResponseBody) SetRequestId(v string) *CreateKMSServiceKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateKMSServiceKeyResponseBody) Validate() error {
	return dara.Validate(s)
}
